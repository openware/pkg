package fujin

import (
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt"
)

// ParseJwtPass verifies jwtPass with brokerId, that should be a valid ECDSA secp256k1 public key
// and returns email and domain values if error didn't occur
func ParseJwtPass(brokerId, jwtPass string) (string, string, error) {
	pubKeyRaw, err := base64.StdEncoding.DecodeString(brokerId)
	if err != nil {
		return "", "", err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return "", "", err
	}

	token, err := jwt.Parse(jwtPass, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return pubKey, nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("token is invalid")
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return "", "", fmt.Errorf("token email value is invalid: '%v'", claims["email"])
	}

	domain, ok := claims["domain"].(string)
	if !ok || domain == "" {
		return "", "", fmt.Errorf("token domain value is invalid: '%v'", claims["domain"])
	}

	return email, domain, nil
}
