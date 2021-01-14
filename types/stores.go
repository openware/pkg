package types

import "time"

// KVStore represent the repositories
type KVStore interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
	Close() error
}
