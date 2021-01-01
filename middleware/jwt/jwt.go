package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/openware/pkg/jwt"
)

const bearerPrefix = "Bearer "

func token(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	if !strings.HasPrefix(string(authHeader), bearerPrefix) {
		return ""
	}

	return string(authHeader[len(bearerPrefix):])
}

// AuthJWT is a gin middleware authenticating the user using signed json token (JWT)
func AuthJWT(keyStore *jwt.KeyStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := jwt.ParseAndValidate(token(c.Request), keyStore.PublicKey)

		if err != nil {
			c.AbortWithError(401, err)
			return
		}

		c.Set("uid", user.UID)
		c.Set("role", user.Role)
		c.Set("level", user.Level)
		c.Set("email", user.Email)
	}
}
