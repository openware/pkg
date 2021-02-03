package mngapi

// Config for management api
type Config struct {
	RootURL       string `yaml:"root_url" env:"ROOT_URL" env-description:"Base URL endpoint for management API"`
	PeatioPrefix  string `yaml:"peatio_prefix" env:"PEATIO_PREFIX" env-description:"Peatio manangement API prefix path"`
	BarongPrefix  string `yaml:"barong_prefix" env:"BARONG_PREFIX" env-description:"Barong manangement API prefix path"`
	JWTIssuer     string `yaml:"jwt_issuer" env:"JWT_ISSUER" env-description:"JWT issuer name"`
	JWTAlgo       string `yaml:"jwt_algo" env:"JWT_ALGO" env-description:"JWT algorithm (default is RS256)" env-default:"RS256"`
	JWTPrivateKey string `yaml:"jwt_private_key" env:"JWT_PRIVATE_KEY" env-description:"Private key for signing JWT"`
}
