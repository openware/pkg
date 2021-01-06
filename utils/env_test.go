package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetEnv(t *testing.T) {
	assert := assert.New(t)

	t.Run("no env with default value should return default value", func(t *testing.T) {
		os.Clearenv()
		myEnv := GetEnv("MY_ENV", "Default Value")
		assert.Equal(myEnv, "Default Value")
	})

	t.Run("have env and default value should return env", func(t *testing.T) {
		os.Setenv("MY_ENV", "Test")
		myEnv := GetEnv("MY_ENV", "Default Value")
		assert.Equal(myEnv, "Test")
	})

	t.Run("empty env and default value should return empty", func(t *testing.T) {
		os.Clearenv()
		myEnv := GetEnv("MY_ENV", "")
		assert.Empty(myEnv)
	})
}

func Test_RequireGetEnv(t *testing.T) {
	assert := assert.New(t)

	t.Run("check existing env should return value with exist set to true", func(t *testing.T) {
		os.Setenv("MY_ENV", "Test")
		value, exist := RequireGetEnv("MY_ENV")
		assert.Equal(value, "Test")
		assert.Equal(exist, true)
	})

	t.Run("check non existing env should return empty with exist set to false", func(t *testing.T) {
		os.Clearenv()
		value, exist := RequireGetEnv("MY_ENV")
		assert.Equal(value, "")
		assert.Equal(exist, false)
	})
}

func Test_DefaultStringEmpty(t *testing.T) {
	assert := assert.New(t)

	t.Run("with empty value should return default string", func(t *testing.T) {
		value := DefaultStringEmpty("", "Default Value")
		assert.Equal(value, "Default Value")
	})

	t.Run("with not empty value should return value", func(t *testing.T) {
		value := DefaultStringEmpty("My Value", "Default Value")
		assert.Equal(value, "My Value")
	})
}

func Test_SetIfNotEmpty(t *testing.T) {
	assert := assert.New(t)

	t.Run("set with empty value should do nothing", func(t *testing.T) {
		str := "My Value"
		SetIfNotEmpty(&str, "")
		assert.Equal(str, "My Value")
	})

	t.Run("set with not empty value should replace with new value", func(t *testing.T) {
		str := "My Value"
		SetIfNotEmpty(&str, "New Value")
		assert.Equal(str, "New Value")
	})
}
