package kli

import (
	"errors"
	"testing"
)

func TestCli(t *testing.T) {
	c := NewCli("test", "description", "0")

	t.Run("Run SetBannerFunction()", func(t *testing.T) {
		c.SetBannerFunction(func(*Cli) string { return "" })
	})

	// t.Run("Run Abort()", func(t *testing.T) {
	// 	cl := NewCli("test", "description", "0")
	// 	cl.Abort(errors.New("test error"))
	// })

	t.Run("Run AddCommand()", func(t *testing.T) {
		c.AddCommand(&Command{name: "test"})
	})

	t.Run("Run PrintBanner()", func(t *testing.T) {
		c.PrintBanner()
	})

	t.Run("Run Run()", func(t *testing.T) {
		c.Run("test")
		c.Run()

		c.preRunCommand = func(*Cli) error { return errors.New("testing coverage") }
		c.Run("test")
	})

	t.Run("Run DefaultCommand()", func(t *testing.T) {
		c.DefaultCommand(&Command{})
	})

	t.Run("Run NewSubCommand()", func(t *testing.T) {
		c.NewSubCommand("name", "description")
	})

	t.Run("Run PreRun()", func(t *testing.T) {
		c.PreRun(func(*Cli) error { return nil })
	})

	t.Run("Run BoolFlag()", func(t *testing.T) {
		var variable bool
		c.BoolFlag("bool", "description", &variable)
	})

	t.Run("Run StringFlag()", func(t *testing.T) {
		var variable string
		c.StringFlag("string", "description", &variable)
	})

	t.Run("Run IntFlag()", func(t *testing.T) {
		var variable int
		c.IntFlag("int", "description", &variable)
	})

	t.Run("Run Action()", func(t *testing.T) {
		c.Action(func() error { return nil })
	})

	t.Run("Run LongDescription()", func(t *testing.T) {
		c.LongDescription("long description")
	})
}

func TestCliEmpty(t *testing.T) {
	var mockCli *Cli
	t.Run("Run NewCli()", func(t *testing.T) {
		mockCli = NewCli("name", "description", "version")
		t.Log(mockCli)
	})

	t.Run("Run defaultBannerFunction()", func(t *testing.T) {
		err := defaultBannerFunction(mockCli)
		t.Log(err)
	})
}

func TestCliGlobalFlag(t *testing.T) {
	var rootCli *Cli

	t.Run("rcli -config app.yml create", func(t *testing.T) {
		var cnf string

		rootCli = NewCli("rcli", "description", "version")
		rootCli.StringFlag("config", "Application yaml configuration file", &cnf)
		rootCli.Action(func() error {
			t.Fail()
			return nil
		})
		create := rootCli.NewSubCommand("create", "Create a sonic application")
		create.Action(func() error {
			t.Log("Running Create")
			return nil
		})
		rootCli.Run("-config", "app.yml", "create")
		t.Logf("After Execution cnf: %s", cnf)
	})

	t.Run("rcli -config app.yml create -v", func(t *testing.T) {
		var verb bool
		var cnf string

		rootCli = NewCli("rcli", "description", "version")
		rootCli.StringFlag("config", "Application yaml configuration file", &cnf)
		rootCli.Action(func() error {
			t.Fail()
			return nil
		})
		create := rootCli.NewSubCommand("create", "Create a sonic application")
		create.BoolFlag("v", "Activate verbose", &verb)
		create.Action(func() error {
			if verb == false {
				t.Fail()
			} else {
				t.Log("Running Create with verbose")
			}
			return nil
		})
		rootCli.Run("-config", "app.yml", "create", "-v")
		t.Logf("After Execution")
	})

	t.Run("rcli create -f file.txt", func(t *testing.T) {
		var file string

		rootCli = NewCli("rcli", "description", "version")
		create := rootCli.NewSubCommand("create", "Create a sonic application")
		create.StringFlag("f", "Pass file", &file)
		create.Action(func() error {
			if file == "file.txt" {
				t.Log("Running Create with -f params")
			} else {
				t.Fail()
			}
			return nil
		})
		rootCli.Run("create", "-f", "file.txt")
		t.Logf("After Execution")
	})

	t.Run("rcli -c app.yml create -f file.txt", func(t *testing.T) {
		var cnf string
		var file string

		rootCli = NewCli("rcli", "description", "version")
		rootCli.StringFlag("c", "Pass file", &cnf)
		create := rootCli.NewSubCommand("create", "Create a sonic application")
		create.StringFlag("f", "Pass file", &file)
		create.Action(func() error {
			if file == "file.txt" && cnf == "app.yml" {
				t.Log("Running Create with -f -c flags")
			} else {
				t.Fail()
			}
			return nil
		})
		rootCli.Run("-c", "app.yml", "create", "-f", "file.txt")
		t.Logf("After Execution")
	})
}
