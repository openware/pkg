package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Connect the redis with default database
func Connect(cnf *Config) KVStore {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cnf.Host, cnf.Port),
		Password: cnf.Pass,
		DB:       0, // use default DB
	})
	return &Store{client}
}

// Set attaches the redis repository and set the data
func (s *Store) Set(key string, value interface{}, exp time.Duration) error {
	return s.client.Set(key, value, exp).Err()
}

// Get attaches the redis repository and get the data
func (s *Store) Get(key string) (string, error) {
	return s.client.Get(key).Result()
}

// Close attaches the redis repository
func (s *Store) Close() error {
	return s.client.Close()
}
