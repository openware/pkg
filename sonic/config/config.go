package config

import (
	"github.com/openware/pkg/mngapi"
	"github.com/openware/pkg/sonic/database"
)

// Config is the application configuration structure
type Config struct {
	Database database.Config `yaml:"database"`
	// TODO Create a redis and vault package
	Redis struct {
		Host string `yaml:"host" env:"REDIS_HOST" env-description:"Redis Server host" env-default:"localhost"`
		Port string `yaml:"port" env:"REDIS_PORT" env-description:"Redis Server port" env-default:"6379"`
	} `yaml:"redis"`
	Port                string        `env:"APP_PORT" env-description:"Port for HTTP service" env-default:"6009"`
	MngAPI              mngapi.Config `yaml:"mngapi"`
	Vault               VaultConfig   `yaml:"vault"`
	DeploymentID        string        `yaml:"deploymentID" env:"DEPLOYMENT_ID"`
	Opendax             OpendaxConfig `yaml:"opendax"`
	MarketsBlacklist    string        `yaml:"markets_blacklist" env:"MARKETS_BLACKLIST"`
	CurrenciesBlacklist string        `yaml:"currencies_blacklist" env:"CURRENCIES_BLACKLIST"`
}

// VaultConfig contains Vault-related configuration
type VaultConfig struct {
	Addr  string `yaml:"addr" env:"VAULT_ADDR" env-default:"http://localhost:8200"`
	Token string `yaml:"token" env:"VAULT_TOKEN"`
}

// OpendaxConfig is the configuration for opendax cloud
type OpendaxConfig struct {
	Addr string `yaml:"addr" env:"OPENDAX_ADDR" env-description:"Opendax API endpoint" env-default:"http://opendax:6969"`
}
