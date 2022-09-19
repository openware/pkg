package signer

import (
	"crypto/ecdsa"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
	"hash/crc32"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/cryptobyte"
	cb_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

var (
	secp256k1N, _     = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN    = new(big.Int).Div(secp256k1N, big.NewInt(2))
	oidPublicKeyECDSA = asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}
)

type SignatureECDSA []byte

// publicKeyInfo is an ASN.1 encoded Subject Public Key Info, defined here: https://tools.ietf.org/html/rfc5280#section-4.1.2.7
type publicKeyInfo struct {
	Raw       asn1.RawContent
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

func crc32c(data []byte) int64 {
	return int64(crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli)))
}

// recoverRS recovers R and S from KMS signature
func recoverRS(signature []byte) (r *big.Int, s *big.Int, err error) {
	r, s = &big.Int{}, &big.Int{}
	var inner cryptobyte.String
	input := cryptobyte.String(signature)
	if !input.ReadASN1(&inner, cb_asn1.SEQUENCE) ||
		!input.Empty() ||
		!inner.ReadASN1Integer(r) ||
		!inner.ReadASN1Integer(s) ||
		!inner.Empty() {
		return nil, nil, errors.New("invalid signature")
	}
	// Google may have already encured that the signature is valid, but we
	// can't assume that.
	if s.Cmp(secp256k1halfN) > 0 {
		s = s.Sub(secp256k1N, s)
	}
	return r, s, nil
}

// pemToPubkey PEM public key to ecdsa.PublicKey
func pemToPubkey(publicKey []byte) (*ecdsa.PublicKey, error) {
	var pub publicKeyInfo
	rest, err := asn1.Unmarshal(publicKey, &pub)
	if err != nil || len(rest) > 0 {
		return nil, fmt.Errorf("error unmarshaling public key: %w", err)
	}
	if !pub.Algorithm.Algorithm.Equal(oidPublicKeyECDSA) {
		return nil, errors.New("not a ECDSA public key")
	}

	// Convert to ecdsa.PublicKey
	pk, err := crypto.UnmarshalPubkey(pub.PublicKey.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal secp256k1 curve point: %w", err)
	}

	return pk, nil
}

// verifyDigest checks that the given public key created signature over digest
func verifyDigest(pubKey ecdsa.PublicKey, hash, sig []byte) bool {
	if len(sig) != 65 {
		return false
	} else if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27
	defer func() { sig[64] += 27 }()

	publicKey, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return false
	}

	verified := crypto.VerifySignature(publicKey, hash, sig[:len(sig)-1])
	addr := common.BytesToAddress(crypto.Keccak256(publicKey[1:])[12:])

	return verified && addr == crypto.PubkeyToAddress(pubKey)
}

// recoverAndVerify recovers R and S from signature and verifies digest
func recoverAndVerify(digest, signature []byte, pubKey ecdsa.PublicKey) (SignatureECDSA, error) {
	r, s, err := recoverRS(signature)
	if err != nil {
		return nil, err
	}

	// Reconstruct the eth signature R || S || V
	sign := make([]byte, 65)
	copy(sign[:32], r.Bytes())
	copy(sign[32:64], s.Bytes())
	sign[64] = 0x1b

	if !verifyDigest(pubKey, digest, sign) {
		sign[64]++
		if !verifyDigest(pubKey, digest, sign) {
			return nil, fmt.Errorf("signer: signature failed, unable to determine V")
		}
	}
	return sign, nil
}
