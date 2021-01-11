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

// repository represent the repository model
type repository struct {
	Client redis.Cmdable
}

// Repository represent the repositories
type Repository interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
}

// Set attaches the redis repository and set the data
func (r *repository) Set(key string, value interface{}, exp time.Duration) error {
	return r.Client.Set(key, value, exp).Err()
}

// Get attaches the redis repository and get the data
func (r *repository) Get(key string) (string, error) {
	get := r.Client.Get(key)
	return get.Result()
}
