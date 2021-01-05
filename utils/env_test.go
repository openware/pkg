package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetEnv(t *testing.T) {
	assert := assert.New(t)

	t.Run("no env with default value should return default value", func(t *testing.T) {
		myEnv := GetEnv("MY_ENV", "Default Value")
		assert.Equal(myEnv, "Default Value")
	})

	t.Run("have env and default value should return env", func(t *testing.T) {
		os.Setenv("MY_ENV", "Test")
		myEnv := GetEnv("MY_ENV", "Default Value")
		assert.Equal(myEnv, "Test")
	})

	t.Run("empty env and default value should return empty", func(t *testing.T) {
		os.Setenv("MY_ENV", "")
		myEnv := GetEnv("MY_ENV", "")
		assert.Empty(myEnv)
	})
}

func Test_RequireGetEnv(t *testing.T) {
	t.Run("", func(t *testing.T) {
		RequireGetEnv("")
		t.Error("Missing test case")
	})
}

func Test_DefaultStringEmpty(t *testing.T) {
	t.Run("", func(t *testing.T) {
		DefaultStringEmpty("", "")
		t.Error("Missing test case")
	})
}

func Test_SetIfNotEmpty(t *testing.T) {
	t.Run("", func(t *testing.T) {
		str := ""
		SetIfNotEmpty(&str, "")
		t.Error("Missing test case")
	})
}
