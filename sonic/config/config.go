package config

// VaultConfig contains Vault-related configuration
type VaultConfig struct {
	Addr  string `yaml:"addr" env:"VAULT_ADDR" env-default:"http://localhost:8200"`
	Token string `yaml:"token" env:"VAULT_TOKEN"`
}

// OpendaxConfig is the configuration for opendax cloud
type OpendaxConfig struct {
	Addr string `yaml:"addr" env:"OPENDAX_ADDR" env-description:"Opendax API endpoint" env-default:"http://opendax:6969"`
}
