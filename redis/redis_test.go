package redis

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	client *redis.Client
)

var (
	key = "key"
	val = "val"
)

func TestMain(m *testing.M) {
	// Start redis mock server
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// New redis client point to mock server
	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	// Assert exit code
	code := m.Run()
	os.Exit(code)
}

func TestSet(t *testing.T) {
	exp := time.Duration(0)

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, val, exp).Return(redis.NewStatusResult("", nil))

	r := create(mock)
	err := r.Set(key, val, exp)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(val, nil))

	r := create(mock)
	res, err := r.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, val, res)
}
