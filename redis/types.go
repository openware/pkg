package redis

import (
    "time"
	"github.com/go-redis/redis"
)

// Config for redis connection
// TODO Set all default values
type Config struct {
	Host string `yaml:"host" env:"REDIS_HOST" env-description:"Redis host"`
	Port string `yaml:"port" env:"REDIS_PORT" env-description:"Redis port"`
	Pass string `env:"REDISPASS" env-description:"Redis user password"`
}

// Store represent the Store model
type Store struct {
	client *redis.Client
}

// KVStore represent the repositories
type KVStore interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
	Close() error
}
