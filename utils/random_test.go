package utils

import (
	"math/rand"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HEX(t *testing.T) {
	assert := assert.New(t)

	t.Run("return valid HEX with parameter > 0", func(t *testing.T) {
		numBytes := rand.Intn(100) + 1
		value := Hex(numBytes)
		match, _ := regexp.MatchString("^[A-Fa-f0-9]+$", value)

		assert.Equal(match, true)
		assert.Equal(len(value), numBytes*2)
	})

	t.Run("return empty string with parameter = 0", func(t *testing.T) {
		numBytes := 0
		value := Hex(numBytes)
		assert.Equal(value, "")
	})
}
