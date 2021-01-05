package database

import (
	"testing"
)

func Test_ConnectDatabase(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ConnectDatabase("")
		t.Error("Missing test case")
	})
}

func Test_RunMigrations(t *testing.T) {
	t.Run("", func(t *testing.T) {
		RunMigrations("")
		t.Error("Missing test case")
	})
}

func Test_LoadSeeds(t *testing.T) {
	t.Run("", func(t *testing.T) {
		LoadSeeds("")
		t.Error("Missing test case")
	})
}
