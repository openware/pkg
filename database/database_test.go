package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO
// Test connection to sqlite mem
// Test connection to invalid config, wrong adapter, and more
func Test_Connect(t *testing.T) {
	t.Run("", func(t *testing.T) {
		_, err := Connect(&Config{
			Driver: "memory",
		})
		require.NoError(t, err)
	})
}
