package signer

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
)

type AWSSigner struct {
	client    *kms.KMS
	keyARN    string
	publicKey ecdsa.PublicKey
}

// NewAWSSigner creates a new AWS signer with the provided signing key
func NewAWSSigner(client *kms.KMS, keyARN string) (*AWSSigner, error) {
	// Get public key from KMS
	key, err := client.GetPublicKey(&kms.GetPublicKeyInput{KeyId: &keyARN})
	if err != nil {
		return nil, fmt.Errorf("signer: unable to get public key: %w", err)
	}

	// Convert to ecdsa.PublicKey
	pk, err := pemToPubkey(key.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("signer: failed to decode public key: %w", err)
	}

	return &AWSSigner{
		client:    client,
		keyARN:    keyARN,
		publicKey: *pk,
	}, nil
}

// GetPublicKey returns a public key
func (c *AWSSigner) GetPublicKey() ecdsa.PublicKey {
	return c.publicKey
}

// Sign the given digest using a KMS key and return ECDSA signature
func (c *AWSSigner) Sign(digest []byte) (SignatureECDSA, error) {
	// Call the API
	res, err := c.client.Sign(&kms.SignInput{
		KeyId:            aws.String(c.keyARN),
		Message:          digest,
		MessageType:      aws.String(kms.MessageTypeDigest),
		SigningAlgorithm: aws.String(kms.SigningAlgorithmSpecEcdsaSha256),
	})
	if err != nil {
		return nil, err
	}

	return recoverAndVerify(digest, res.Signature, c.publicKey)
}
