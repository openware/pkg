package crypto

import (
	"testing"
)

func Test_SHA256(t *testing.T) {
	t.Run("return SHA256 for string", func(t *testing.T) {
		SHA256("")
		t.Error("Missing test case")
	})
}

func Test_MD5(t *testing.T) {
	t.Run("return SHA256 for string", func(t *testing.T) {
		MD5("")
		t.Error("Missing test case")
	})
}
