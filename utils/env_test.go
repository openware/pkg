package utils

import (
	"testing"
)

func Test_GetEnv(t *testing.T) {
	t.Run("", func(t *testing.T) {
		GetEnv("", "")
		t.Error("Missing test case")
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
