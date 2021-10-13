package handlers

import (
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/openware/kaigara/pkg/vault"
	"github.com/openware/pkg/jwt"
	"github.com/openware/pkg/sonic/config"
)

type cache struct {
	Mutex sync.RWMutex
	Data  map[string]interface{}
}

type VaultService interface {
	LoadSecrets(appName, scope string) error
	SetSecret(appName, name string, value interface{}, scope string) error
	SaveSecrets(appName, scope string) error
	GetSecrets(appName, scope string) (map[string]interface{}, error)
	ListSecrets(appName, scope string) ([]string, error)
	ListAppNames() ([]string, error)
}

// GetOpendaxConfig helper return kaigara config from gin context
func GetOpendaxConfig(ctx *gin.Context) (*config.OpendaxConfig, error) {
	value, ok := ctx.Get("OpendaxConfig")
	if !ok {
		return nil, fmt.Errorf("Opendax config is not found")
	}
	config := value.(*config.OpendaxConfig)

	return config, nil
}

func GetSonicCtx(ctx *gin.Context) (*SonicContext, error) {
	value, ok := ctx.Get("sctx")
	if !ok {
		return nil, fmt.Errorf("Sonic config is not found")
	}
	sctx := value.(*SonicContext)

	return sctx, nil
}

// GetAuth helper return auth from gin context
func GetAuth(ctx *gin.Context) (*jwt.Auth, error) {
	value, ok := ctx.Get("auth")
	if !ok {
		return nil, fmt.Errorf("Auth is not found")
	}
	auth := value.(*jwt.Auth)

	return auth, nil
}

// GetVaultService helper return global vault service from gin context
func GetVaultService(ctx *gin.Context) (VaultService, error) {
	value, ok := ctx.Get("VaultService")
	if !ok {
		return nil, fmt.Errorf("Global vault service is not found")
	}
	vaultService := value.(VaultService)

	return vaultService, nil
}

// WriteCache read latest vault version and fetch keys values from vault
// 'firstRun' variable will help to run writing to cache on first system start
// as on the start latest and current versions are the same
func WriteCache(vaultService *vault.Service, scope string, firstRun bool) {
	err := vaultService.LoadSecrets("global", scope)
	if err != nil {
		panic(err)
	}

	if memoryCache.Data == nil {
		memoryCache.Data = make(map[string]interface{})
	}

	current, err := vaultService.GetCurrentVersion("global", scope)
	if err != nil {
		panic(err)
	}

	latest, err := vaultService.GetLatestVersion("global", scope)
	if err != nil {
		panic(err)
	}

	if current != latest || firstRun {
		log.Println("Writing to cache")
		keys, err := vaultService.ListSecrets("global", scope)
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			val, err := vaultService.GetSecret("global", key, scope)
			if err != nil {
				panic(err)
			}
			memoryCache.Data[key] = val
		}
	}
}
