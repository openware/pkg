package mngapi

// Config for management api
type Config struct {
	PeatioURL     string `yaml:"peatio_url" env:"PEATIO_URL" env-description:"Peatio URL endpoint for management API"`
	BarongURL     string `yaml:"barong_url" env:"BARONG_URL" env-description:"Barong URL endpoint for management API"`
	JWTIssuer     string `yaml:"jwt_issuer" env:"JWT_ISSUER" env-description:"JWT issuer name"`
	JWTAlgo       string `yaml:"jwt_algo" env:"JWT_ALGO" env-description:"JWT algorithm (default is RS256)" env-default:"RS256"`
	JWTPrivateKey string `yaml:"jwt_private_key" env:"JWT_PRIVATE_KEY" env-description:"Private key for signing JWT"`
}
