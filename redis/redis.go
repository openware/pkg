package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func create(client redis.Cmdable) Repository {
	return &repository{client}
}

// Connect : Connect the redis with default database
func Connect(cnf *Config) Repository {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cnf.Host, cnf.Port),
		Password: cnf.Pass,
		DB:       0, // use default DB
	})
	return create(client)
}
