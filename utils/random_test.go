package utils

import (
	"testing"
)

func Test_HEX(t *testing.T) {
	t.Run("return HEX ", func(t *testing.T) {
		Hex(1)
		t.Error("Missing test case")
	})
}
