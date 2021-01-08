package crypto

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SHA256(t *testing.T) {
	assert := assert.New(t)

	t.Run("return 64 hexadecimal SHA256 for input string", func(t *testing.T) {
		value := SHA256("Test String")
		match, _ := regexp.MatchString("^[A-Fa-f0-9]+$", value)
		assert.Equal(match, true)
		assert.Equal(len(value), 64)
	})
}

func Test_MD5(t *testing.T) {
	assert := assert.New(t)

	t.Run("return 32 hexadecimal MD5 for input string", func(t *testing.T) {
		value := MD5("Test String")
		match, _ := regexp.MatchString("^[A-Fa-f0-9]+$", value)
		assert.Equal(match, true)
		assert.Equal(len(value), 32)
	})
}
