package signer

import (
	"context"
	"crypto/ecdsa"
	"encoding/pem"
	"fmt"

	api "cloud.google.com/go/kms/apiv1"

	"google.golang.org/genproto/googleapis/cloud/kms/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type GCPSigner struct {
	client    *api.KeyManagementClient
	keyPath   string
	publicKey ecdsa.PublicKey
}

// NewGCPSigner creates a new GCP signer with the provided signing key
func NewGCPSigner(client *api.KeyManagementClient, keyPath string) (*GCPSigner, error) {
	// Get public key from KMS
	key, err := client.GetPublicKey(context.Background(), &kms.GetPublicKeyRequest{
		Name: keyPath,
	})
	if err != nil {
		return nil, fmt.Errorf("signer: unable to get public key: %w", err)
	}

	// Convert to ecdsa.PublicKey
	pem, _ := pem.Decode([]byte(key.Pem))
	pk, err := pemToPubkey(pem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("signer: failed to decode public key: %w", err)
	}
	return &GCPSigner{
		client:    client,
		keyPath:   keyPath,
		publicKey: *pk,
	}, nil
}

// GetPublicKey returns a public key
func (c *GCPSigner) GetPublicKey() ecdsa.PublicKey {
	return c.publicKey
}

// Sign the given digest using a KMS key and return ECDSA signature
func (c *GCPSigner) Sign(digest []byte) (SignatureECDSA, error) {
	req := &kms.AsymmetricSignRequest{
		Name: c.keyPath,
		Digest: &kms.Digest{
			Digest: &kms.Digest_Sha256{
				Sha256: digest,
			},
		},
		DigestCrc32C: wrapperspb.Int64(crc32c(digest)),
	}

	// Call the API
	res, err := c.client.AsymmetricSign(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if !res.VerifiedDigestCrc32C {
		return nil, fmt.Errorf("signer: request corrupted in-transit")
	}

	if crc32c(res.Signature) != res.SignatureCrc32C.Value {
		return nil, fmt.Errorf("signer: response corrupted in-transit")
	}

	return recoverAndVerify(digest, res.Signature, c.publicKey)
}
