package database_test

import (
	"testing"

	"github.com/openware/pkg/database"
)

// TODO
// Test connection to invalid config, wrong adapter, and more
func Test_Connect(t *testing.T) {
	t.Run("Connection to sqlite", func(t *testing.T) {
		_, err := database.Connect(&database.Config{
			Driver: "memory",
			Pool:   5,
		})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Wrong driver", func(t *testing.T) {
		_, err := database.Connect(&database.Config{
			Driver: "foo1",
		})

		if err == nil {
			t.Error("Has to fail with unsupported driver")
		}
	})

	t.Run("Empty configuration", func(t *testing.T) {
		_, err := database.Connect(&database.Config{})
		if err == nil {
			t.Error(err)
		}
	})
}
