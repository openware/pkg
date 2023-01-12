package fujin

import (
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt"
)

type PassEntries struct {
	adminEmail      string
	domainName      string
	platformVersion string
}

// ParseJwtPass verifies jwtPass with brokerId, that should be a valid ECDSA secp256k1 public key
// and returns email and domain values if error didn't occur
func ParseJwtPass(brokerId, jwtPass string) (*PassEntries, error) {
	pubKeyRaw, err := base64.StdEncoding.DecodeString(brokerId)
	if err != nil {
		return nil, err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(jwtPass, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return pubKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return nil, fmt.Errorf("token email value is invalid: '%v'", claims["email"])
	}

	domain, ok := claims["domain"].(string)
	if !ok || domain == "" {
		return nil, fmt.Errorf("token domain value is invalid: '%v'", claims["domain"])
	}

	version, ok := claims["version"].(string)
	if !ok || version == "" {
		return nil, fmt.Errorf("token version value is invalid: '%v'", claims["version"])
	}

	return &PassEntries{
		adminEmail:      email,
		domainName:      domain,
		platformVersion: version,
	}, nil
}
