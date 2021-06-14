package handlers

import (
	"github.com/openware/pkg/mngapi/peatio"
	"sync"
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
