package cryptocom

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	// Types
	AuthRequest      = 1
	SubscribeRequest = 2
	HeartBeat        = 3
)

type Request struct {
	Id        int
	Type      uint8
	Method    string
	ApiKey    string
	Signature string
	Nonce     string
	Params    map[string]interface{}
}

type Response struct {
	Id      int
	Method  string
	Code    int
	Message string
	Result  map[string]interface{}
}

func generateNonce() string {
	return fmt.Sprintf("%d", time.Now().Unix()*1000)
}

func (r *Request) Encode() ([]byte, error) {
	switch r.Type {
	case AuthRequest:
		return json.Marshal(map[string]interface{}{
			"id":      r.Id,
			"method":  r.Method,
			"api_key": r.ApiKey,
			"sig":     r.Signature,
			"nonce":   r.Nonce,
		})
	case SubscribeRequest:
		return json.Marshal(map[string]interface{}{
			"id":     r.Id,
			"method": r.Method,
			"params": r.Params,
			"nonce":  r.Nonce,
		})

	case HeartBeat:
		return json.Marshal(map[string]interface{}{
			"id":     r.Id,
			"method": r.Method,
		})
	default:
		return nil, errors.New("invalid type")
	}
}
