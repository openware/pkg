# Cloud KMS Signer

Sign the requested digest with a`secp256k1` key stored in cloud KMS, without holding the private key in memory. 

**The package implements `SignerInterface`:**

    type SignerInterface interface {
	    Sign(digest []byte) (signatureECDSA, error)
	    GetPublicKey() ecdsa.PublicKey
	}
	
	type signatureECDSA []byte


**How to use:**

 1. Create a client of your cloud provider
 2. Create a new signer by passing in your client and `secp256k1` signing key path.
 3. That's all! You are ready to sign your transactions!

**Currently supported KMS providers:** `AWS`, `GCP`.

## Examples

### Create the correct signing key in KMS

**AWS:**

 1. Go to `KMS > Customer managed keys > Create key`
 2. Create a key with the following parameters: `Type: Asymmetric`, `Key usage: Sign and Verify`, `Key spec: ECC_SECG_P256K1`.
 3. Select your key in `Customer managed keys` tab and copy its ARN.

**GCP:**

 1. Go to `Cloud Key Management Service` and create a `KEY RING`.
 2. Select the key ring and create a new key with the following parameters: `Protection level: HSM`, `Purpose: Asymmetric sign`, `Algorithm: Elliptic Curve secp256k1 - SHA256 Digest`.
 4. Select your key in the key ring and do `Actions > Copy resource name`.
 
### Create a client

**AWS:**

    // Parse key ARN to get the region of the key (this step is optional)
    arn, err := arn.Parse(keyARN)
    if err != nil {
	    return err
    }
    // Start a new session with aws access key id and secret key
    sess, err := session.NewSession(&aws.Config{
	    Region: &arn.Region,
	    Credentials: credentials.NewStaticCredentials(accessKeyID, secretKey, ""),
	})
	if err != nil {
		return err
	}
	// Create a new client
	client := kms.New(sess)

*Useful links:*
 - [Passing Credentials to a Session](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/)
 - [Configuring the AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)
  
**GCP:**

    // Create a new client
    client, err := kms.NewKeyManagementClient(context.Background(), option.WithCredentialsJSON([]byte(credentials)))
    if err != nil {
	    return err
	}

*Useful links:*
 - [Passing credentials to a GCP client](https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-go)

### Sign hash

**AWS:**

    signer, err := NewAWSSigner(client, keyARN)
    if err != nil {
	    return err
	}
	signed, err := signer.Sign(hash)
	if err != nil {
		return err
	}

**GCP:**

    signer, err := NewGCPSigner(client, keyPath)
    if err != nil {
	    return err
	}
	signed, err := signer.Sign(hash)
	if err != nil {
		return err
	}


### Sign and send a blockchain transaction

	func sendTransaction(s SignerInterface, amount int64, destAddress, rpcURL string) {
		ctx := context.Background()
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			log.Fatal(err)
		}

		address := crypto.PubkeyToAddress(s.GetPublicKey())
		balance, err := client.BalanceAt(ctx, address, nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("current balance: %v", balance)

		nonce, err := client.PendingNonceAt(ctx, address)
		if err != nil {
			log.Fatal(err)
		}

		value := big.NewInt(amount) // in wei (1 eth = 10{^18} wei)
		gasLimit := uint64(21000)   // in units
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			log.Fatal(err)
		}

		var data []byte
		tx := types.NewTransaction(nonce, common.HexToAddress(destAddress), value, gasLimit, gasPrice, data)

		chainID, err := client.NetworkID(ctx)
		if err != nil {
			log.Fatal(err)
		}

		signer := types.LatestSignerForChainID(chainID)
		hash := signer.Hash(tx)
		res, err := s.Sign(hash[:])
		if err != nil {
			log.Fatal(err)
		}
		if res[64] == 27 || res[64] == 28 {
			res[64] -= 27 // Transform V from Ethereum-legacy to 0/1
		}
		signedTx, err := tx.WithSignature(signer, res)
		if err != nil {
			log.Fatal(err)
		}

		err = client.SendTransaction(ctx, signedTx)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("tx sent: %s", tx.Hash())
	}