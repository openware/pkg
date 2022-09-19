package signer

import (
	"crypto/ecdsa"
)

type SignerInterface interface {
	Sign(digest []byte) (SignatureECDSA, error)
	GetPublicKey() ecdsa.PublicKey
}
