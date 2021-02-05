package jwt

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/openware/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthJWT(t *testing.T) {
	ks := &jwt.KeyStore{}
	ks.GenerateKeys()
	token, err := jwt.ForgeToken("ABC0001", "john@barong.io", "author", 12, ks.PrivateKey, jwtgo.MapClaims{})
	require.NoError(t, err)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.Header = http.Header{
		"Authorization": []string{"Bearer " + token},
	}
	c := &gin.Context{
		Request: req,
	}
	// c.Header("Authorization", "Bearer "+token)

	AuthJWT(ks)(c)
	assert.Equal(t, []string(nil), c.Errors.Errors())
	uid, ok := c.Get("uid")
	assert.True(t, ok)
	role, ok := c.Get("role")
	assert.True(t, ok)
	level, ok := c.Get("level")
	assert.True(t, ok)
	email, ok := c.Get("email")
	assert.True(t, ok)
	assert.Equal(t, "ABC0001", uid)
	assert.Equal(t, "author", role)
	assert.Equal(t, json.Number("12"), level)
	assert.Equal(t, "john@barong.io", email)
}
