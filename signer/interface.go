package signer

import (
	"crypto/ecdsa"
)

type SignerInterface interface {
	Sign(digest []byte) (signatureECDSA, error)
	GetPublicKey() ecdsa.PublicKey
}
