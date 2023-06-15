package jwt

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestAuth_appendClaims(t *testing.T) {
	claims := jwt.MapClaims{
		"iat":         time.Now().Unix(),
		"jti":         strconv.FormatInt(time.Now().Unix(), 10),
		"exp":         time.Now().UTC().Add(time.Hour).Unix(),
		"sub":         "session",
		"iss":         "barong",
		"aud":         [2]string{"peatio", "barong"},
		"state":       "active",
		"referral_id": 1,
	}

	t.Run("merges claims with nil", func(t *testing.T) {
		res := appendClaims(claims, nil)

		if !reflect.DeepEqual(claims, res) {
			t.Errorf("expected: %v actual: %v", claims, res)
		}
	})

	t.Run("merges nil with claims", func(t *testing.T) {
		res := appendClaims(nil, claims)
		if !reflect.DeepEqual(claims, res) {
			t.Errorf("expected: %v actual: %v", claims, res)
		}
	})

	t.Run("adds claim", func(t *testing.T) {
		res := appendClaims(claims, jwt.MapClaims{"custom": "claim"})

		if claims["custom"] != "claim" {
			t.Errorf("expected: %v actual: %v", claims, res)
		}
	})

	t.Run("rewrites claim", func(t *testing.T) {
		res := appendClaims(claims, jwt.MapClaims{"state": "banned"})

		if claims["state"] != "banned" {
			t.Errorf("expected: %v actual: %v", claims, res)
		}
	})
}

func TestAuth_JWT_RSA(t *testing.T) {
	ks, err := LoadOrGenerateKeys("./testdata/rsa-key", "./testdata/rsa-key.pub")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should validate jwt", func(t *testing.T) {
		token, err := ForgeToken("uid", "email", "role", 3, 1, ks.PrivateKey, nil)
		if err != nil {
			t.Fatal(err)
		}

		_, err = ParseAndValidate(token, ks.PublicKey)
		if err != nil {
			t.Fatal(err)
		}
	})
}
func TestAuth_JWT_EdDSA(t *testing.T) {
	ks, err := LoadOrGenerateKeysEdDSA("./testdata/ed25519-key", "./testdata/ed25519-key.pub")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should validate jwt", func(t *testing.T) {
		token, err := ForgeTokenEdDSA("uid", "email", "role", 3, 1, ks.PrivateKey, nil)
		if err != nil {
			t.Fatal(err)
		}

		_, err = ParseAndValidateEdDSA(token, ks.PublicKey)
		if err != nil {
			t.Fatal(err)
		}
	})
}
