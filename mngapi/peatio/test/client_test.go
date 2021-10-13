package test

import (
	"testing"

	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	t.Run("Success creation", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)

		assert.NotNil(t, client)
		assert.Nil(t, err)
	})

	t.Run("JWT issuer unset", func(t *testing.T) {
		client, err := New(URL, "", jwtAlgo, jwtPrivateKey)

		assert.Nil(t, client)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "JWT issuer unset")
	})

	t.Run("Invalid signing algorithm", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, "RS999", jwtPrivateKey)

		assert.Nil(t, client)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Unsupported signing method RS999")
	})

	t.Run("Invalid private key", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, "")

		assert.Nil(t, client)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Invalid Key: Key must be PEM encoded PKCS1 or PKCS8 private key")
	})
}
