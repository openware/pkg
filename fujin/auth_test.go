package fujin

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestParseJwtPass(t *testing.T) {
	validEcdsaPrivKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	validEcdsaPubKeyRaw := base64.StdEncoding.EncodeToString(crypto.FromECDSAPub(&validEcdsaPrivKey.PublicKey))

	invalidEcdsaPrivKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	invalidEcdsaPubKeyRaw := base64.StdEncoding.EncodeToString(crypto.FromECDSAPub(&invalidEcdsaPrivKey.PublicKey))

	invalidAlgPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Fatal(err)
	}

	publickey := &invalidAlgPrivKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		t.Fatal(err)
	}
	publicKeyBlock := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	invalidAlgPubKeyRaw := base64.StdEncoding.EncodeToString(publicKeyBlock)

	tcs := []struct {
		name         string
		privKey      interface{}
		brokerId     string
		jwtMapClaims jwt.MapClaims
		errorMessage string
	}{
		{"correct", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"email": "some@email.com", "domain": "some.domain.com", "version": "4.0.0", "custody": "mainnet",
		}, ""},
		{"absent email", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"domain": "some.domain.com", "version": "4.0.0", "custody": "mainnet",
		}, "token email value is invalid: '<nil>'"},
		{"absent domain", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"email": "some@email.com", "version": "4.0.0", "custody": "mainnet",
		}, "token domain value is invalid: '<nil>'"},
		{"absent version", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"email": "some@email.com", "domain": "some.domain.com",
		}, ""},
		{"absent custody", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"email": "some@email.com", "domain": "some.domain.com", "version": "4.0.0",
		}, ""},
		{"empty email", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"email": "", "domain": "some.domain.com", "version": "4.0.0", "custody": "mainnet",
		}, "token email value is invalid: ''"},
		{"empty domain", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"email": "some@email.com", "domain": "", "version": "4.0.0", "custody": "mainnet",
		}, "token domain value is invalid: ''"},
		{"empty version", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"email": "some@email.com", "domain": "some.domain.com", "version": "", "custody": "mainnet",
		}, ""},
		{"empty custody", validEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{
			"email": "some@email.com", "domain": "some.domain.com", "version": "4.0.0", "custody": "",
		}, ""},
		{"invalid ECDSA public key", validEcdsaPrivKey, invalidEcdsaPubKeyRaw, jwt.MapClaims{}, "invalid secp256k1 public key"},
		{"invalid ECDSA private key", invalidEcdsaPrivKey, validEcdsaPubKeyRaw, jwt.MapClaims{}, "crypto/ecdsa: verification error"},
		{"invalid algorithm public key", validEcdsaPrivKey, invalidAlgPubKeyRaw, jwt.MapClaims{}, "invalid secp256k1 public key"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			token := jwt.NewWithClaims(jwt.SigningMethodES256, tc.jwtMapClaims)
			tokenString, err := token.SignedString(tc.privKey)
			if err != nil {
				t.Fatal(err)
			}

			passEntries, err := ParseJwtPass(tc.brokerId, tokenString)

			if tc.errorMessage != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.errorMessage, err.Error())
			} else {
				assert.NoError(t, err)

				if tcEmail, ok := tc.jwtMapClaims["email"]; ok {
					assert.Equal(t, tcEmail, passEntries.AdminEmail)
				}

				if tcDomain, ok := tc.jwtMapClaims["domain"]; ok {
					assert.Equal(t, tcDomain, passEntries.DomainName)
				}

				if tcVersion, ok := tc.jwtMapClaims["version"]; ok && tcVersion == "" {
					assert.Equal(t, "none", passEntries.PlatformVersion)
				} else if ok {
					assert.Equal(t, tcVersion, passEntries.PlatformVersion)
				}

				if tcCustody, ok := tc.jwtMapClaims["custody"]; ok && tcCustody == "" {
					assert.Equal(t, "none", passEntries.CustodyMode)
				} else if ok {
					assert.Equal(t, tcCustody, passEntries.CustodyMode)
				}
			}
		})
	}
}

func Test_GenerateAuthCredentials(t *testing.T) {
	validEcdsaPrivKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	brokerId := base64.StdEncoding.EncodeToString(crypto.FromECDSAPub(&validEcdsaPrivKey.PublicKey))
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":   "some@email.com",
		"domain":  "some.domain.com",
		"version": "4.0.0",
		"custody": "mainnet",
	})

	jwtToken, err := token.SignedString(validEcdsaPrivKey)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("broker:%s\n", brokerId)
	fmt.Printf("jwt:%s\n", jwtToken)
}
