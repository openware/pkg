package database_test

import (
	"io/ioutil"
	"os"

	"github.com/openware/pkg/database"
	"github.com/openware/pkg/ika"
)

// In this example we read database configuration from YAML file using ika
// and connect to the database
func Example() {
	const configFile = `
port: 6009
database:
  driver: memory
  host: localhost
  port: 3006
  name: opendax
  user: user
  pass: changeme
`

	type config struct {
		ServerPort string           `yaml:"port"`
		DbConfig   *database.Config `yaml:"database"` // Very important to pass inside pointer to database.Config
	}

	// In your application you can just create the file ( "config/config.yml" for example )
	// And pass the path to the file in ika.ReadConfig
	tmpFile, err := ioutil.TempFile(os.TempDir(), "*.yml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	input := []byte(configFile)
	if _, err = tmpFile.Write(input); err != nil {
		panic(err)
	}

	cfg := config{}
	ika.ReadConfig(tmpFile.Name(), &cfg)

	client, err := database.Connect(cfg.DbConfig)
	if err != nil {
		panic(err)
	}
}
