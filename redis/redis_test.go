package redis

import (
	"strings"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	key = "key"
	val = "val"
)

func TestConnect(t *testing.T) {
	mr, _ := miniredis.Run()

	// Split mock server address to host, port
	addrs := strings.Split(mr.Addr(), ":")

	// New redis client point to mock server
	store := Connect(&Config{
		Host: addrs[0],
		Port: addrs[1],
	})
	err := store.Close()
	require.NoError(t, err)
}

func TestSet(t *testing.T) {
	mr, _ := miniredis.Run()

	// Split mock server address to host, port
	addrs := strings.Split(mr.Addr(), ":")

	// New redis client point to mock server
	store := Connect(&Config{
		Host: addrs[0],
		Port: addrs[1],
	})

	err := store.Set(key, val, time.Duration(0))
	require.NoError(t, err)
	actual, _ := mr.Get(key)
	assert.Equal(t, val, actual)
}

func TestGet(t *testing.T) {
	mr, _ := miniredis.Run()
	mr.Set(key, val)

	// Split mock server address to host, port
	addrs := strings.Split(mr.Addr(), ":")

	// New redis client point to mock server
	store := Connect(&Config{
		Host: addrs[0],
		Port: addrs[1],
	})

	result, err := store.Get(key)
	require.NoError(t, err)
	assert.Equal(t, val, result)
}
