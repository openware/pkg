package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/openware/pkg/jwt"
	sonic "github.com/openware/pkg/sonic/config"

	"github.com/gin-gonic/gin"
	"github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

type TestVaultService struct {
	storage map[string]interface{}
}

func NewTestVaultService() *TestVaultService {
	return &TestVaultService{
		storage: make(map[string]interface{}),
	}
}

func (tvs *TestVaultService) LoadSecrets(appName, scope string) error {
	if _, ok := tvs.storage[appName]; !ok {
		tvs.storage[appName] = make(map[string]interface{})
	}

	if _, ok := tvs.storage[appName].(map[string]interface{})[scope]; !ok {
		tvs.storage[appName].(map[string]interface{})[scope] = make(map[string]interface{})
	}

	return nil
}

func (tvs *TestVaultService) SetSecret(appName, name string, value interface{}, scope string) error {
	if err := tvs.errorIfNoStorage(appName, scope); err != nil {
		return err
	}

	tvs.storage[appName].(map[string]interface{})[scope].(map[string]interface{})[name] = value

	return nil
}

func (tvs *TestVaultService) SaveSecrets(appName, scope string) error {
	return tvs.errorIfNoStorage(appName, scope)
}

func (tvs *TestVaultService) GetSecrets(appName, scope string) (map[string]interface{}, error) {
	if err := tvs.errorIfNoStorage(appName, scope); err != nil {
		return nil, err
	}

	return tvs.storage[appName].(map[string]interface{})[scope].(map[string]interface{}), nil
}

func (tvs *TestVaultService) ListSecrets(appName, scope string) ([]string, error) {
	if err := tvs.errorIfNoStorage(appName, scope); err != nil {
		return []string{}, nil
	}
	store := tvs.storage[appName].(map[string]interface{})[scope].(map[string]interface{})

	keys := make([]string, len(store))

	i := 0
	for k := range store {
		keys[i] = k
		i++
	}

	return keys, nil
}

func (tvs *TestVaultService) ListAppNames() ([]string, error) {
	keys := make([]string, len(tvs.storage))

	i := 0
	for k := range tvs.storage {
		keys[i] = k
		i++
	}

	return keys, nil
}

func (tvs *TestVaultService) errorIfNoStorage(appName, scope string) error {
	if _, ok := tvs.storage[appName]; !ok {
		return fmt.Errorf("No such app defined: %s", appName)
	}

	if _, ok := tvs.storage[appName].(map[string]interface{})[scope]; !ok {
		return fmt.Errorf("No such scope loaded: %s", scope)
	}

	return nil
}

func testVaultServiceMiddleware(vaultService VaultService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("VaultService", vaultService)
		c.Next()
	}
}

func testAuthMiddleware(auth *jwt.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("auth", auth)
		c.Next()
	}
}

func prepareParams(t *testing.T, params interface{}) io.Reader {
	body, err := json.Marshal(params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	return bytes.NewBuffer(body)
}

func testLicenseCreator(appName string, config *sonic.OpendaxConfig, service VaultService) error {
	return nil
}

func testFetchConfig(client *peatio.Client, opendaxAddr string, platformID string) error {
	return nil
}

func setupTestRouter(addMiddlewares func(*gin.Engine)) *gin.Engine {
	r := gin.Default()
	addMiddlewares(r)
	r.PUT("/api/v2/admin/secret", SetSecret)
	r.GET("/api/v2/admin/secrets", GetSecrets)
	r.POST("/api/v2/admin/platform/new", func(ctx *gin.Context) {
		createPlatform(ctx, testLicenseCreator, testFetchConfig)
	})
	return r
}

func TestSetSecret(t *testing.T) {
	t.Run("Without Vault Service", func(t *testing.T) {
		router := setupTestRouter(func(r *gin.Engine) {})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/v2/admin/secret", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "{\"error\":\"Global vault service is not found\"}", w.Body.String())
	})

	t.Run("Invalid parameters", func(t *testing.T) {
		tvs := NewTestVaultService()
		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(testVaultServiceMiddleware(tvs))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/v2/admin/secret", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"error\":\"invalid request\"}", w.Body.String())
	})

	t.Run("Valid parameters", func(t *testing.T) {
		tvs := NewTestVaultService()
		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(testVaultServiceMiddleware(tvs))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/v2/admin/secret", prepareParams(t, setSecretParams{
			Key:   "key",
			Value: "value",
			Scope: "public",
		}))
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "\"Secret saved successfully\"", w.Body.String())
	})

}

func TestGetSecrets(t *testing.T) {
	t.Run("Without Vault Service", func(t *testing.T) {
		router := setupTestRouter(func(r *gin.Engine) {})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v2/admin/secrets", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "{\"error\":\"Global vault service is not found\"}", w.Body.String())
	})

	t.Run("Empty storage", func(t *testing.T) {
		tvs := NewTestVaultService()
		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(testVaultServiceMiddleware(tvs))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v2/admin/secrets", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{}", w.Body.String())
	})

	t.Run("Public keys", func(t *testing.T) {
		tvs := NewTestVaultService()
		tvs.LoadSecrets("", "public")
		tvs.SetSecret("", "name", "openware", "public")

		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(testVaultServiceMiddleware(tvs))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v2/admin/secrets", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{\"\":{\"private\":{},\"public\":{\"name\":\"openware\"},\"secret\":{}}}", w.Body.String())
	})

	t.Run("Secret keys", func(t *testing.T) {
		tvs := NewTestVaultService()
		tvs.LoadSecrets("", "public")
		tvs.SetSecret("", "name", "openware", "public")
		tvs.LoadSecrets("", "secret")
		tvs.SetSecret("", "secret_name", "yellow", "secret")

		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(testVaultServiceMiddleware(tvs))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v2/admin/secrets", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{\"\":{\"private\":{},\"public\":{\"name\":\"openware\"},\"secret\":{\"secret_name\":\"******\"}}}", w.Body.String())
	})
}

func TestCreatePlatform(t *testing.T) {
	t.Run("Without Opendax configuration", func(t *testing.T) {
		router := setupTestRouter(func(r *gin.Engine) {})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v2/admin/platform/new", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
		assert.Equal(t, "{\"error\":\"Opendax config is not found\"}", w.Body.String())
	})

	t.Run("Without Vault Service", func(t *testing.T) {
		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(OpendaxConfigMiddleware(&sonic.OpendaxConfig{Addr: "http://opendax:6969"}))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v2/admin/platform/new", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
		assert.Equal(t, "{\"error\":\"Global vault service is not found\"}", w.Body.String())
	})

	t.Run("Without Auth Service", func(t *testing.T) {
		tvs := NewTestVaultService()
		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(OpendaxConfigMiddleware(&sonic.OpendaxConfig{Addr: "http://opendax:6969"}))
			r.Use(testVaultServiceMiddleware(tvs))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v2/admin/platform/new", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, "{\"error\":\"Unauthorized\"}", w.Body.String())
	})

	t.Run("Without superadmin role", func(t *testing.T) {
		tvs := NewTestVaultService()
		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(OpendaxConfigMiddleware(&sonic.OpendaxConfig{Addr: "http://opendax:6969"}))
			r.Use(testVaultServiceMiddleware(tvs))
			r.Use(testAuthMiddleware(&jwt.Auth{Role: "admin"}))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v2/admin/platform/new", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, "{\"error\":\"Unauthorized\"}", w.Body.String())
	})

	t.Run("Without params", func(t *testing.T) {
		tvs := NewTestVaultService()
		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(OpendaxConfigMiddleware(&sonic.OpendaxConfig{Addr: "http://opendax:6969"}))
			r.Use(testVaultServiceMiddleware(tvs))
			r.Use(testAuthMiddleware(&jwt.Auth{Role: "superadmin"}))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v2/admin/platform/new", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 422, w.Code)
		assert.Equal(t, "{\"error\":\"invalid request\"}", w.Body.String())
	})

	t.Run("Invalid params", func(t *testing.T) {
		tvs := NewTestVaultService()
		router := setupTestRouter(func(r *gin.Engine) {
			r.Use(OpendaxConfigMiddleware(&sonic.OpendaxConfig{Addr: "http://opendax:6969"}))
			r.Use(testVaultServiceMiddleware(tvs))
			r.Use(testAuthMiddleware(&jwt.Auth{Role: "superadmin"}))
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v2/admin/platform/new", prepareParams(t, CreatePlatformParams{
			PlatformName: "finex",
			PlatformURL:  "example.com:8332",
		}))
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
		assert.Equal(t, "{\"error\":\"Post \\\"http://opendax:6969/api/v2/opx/platforms/new\\\": dial tcp: lookup opendax: No address associated with hostname\"}", w.Body.String())
	})
}
