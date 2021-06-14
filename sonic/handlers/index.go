package handlers

import (
	"github.com/openware/kaigara/pkg/vault"
	"github.com/openware/pkg/mngapi/peatio"
	"sync"
	"time"
)

var (
	memoryCache  = cache{
		Data:  make(map[string]interface{}),
		Mutex: sync.RWMutex{},
	}
	SonicPublicKey  string
	PeatioPublicKey string
	BarongPublicKey string
)

type SonicContext struct {
	PeatioClient *peatio.Client
}

// StartConfigCaching will fetch latest data from vault every 30 seconds
func StartConfigCaching(vaultService *vault.Service, scope string) {
	for {
		<-time.After(20 * time.Second)

		memoryCache.Mutex.Lock()
		WriteCache(vaultService, scope, false)
		memoryCache.Mutex.Unlock()
	}
}
