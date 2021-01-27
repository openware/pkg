package mgntapi

// KeychainData config to store client keychain
type KeychainData struct {
	Algorithm string `yaml:"algorithm"`
	Value     string `yaml:"value"`
}

// Keychain to store keychain for all clients
type Keychain struct {
	Applogic KeychainData `yaml:"applogic"` // Opendax is default client
	Opendax  KeychainData `yaml:"opendax"`  // Opendax is default client
	// TODO: Add more client keychain if needed
}

// Action config to store allowed actions
type Action struct {
	Signatures  []string `yaml:"required_signatures"`
	RequireTOTP bool     `yaml:"requires_barong_totp"`
}

// PeatioActions config to store all allowed actions for Peatio
type PeatioActions struct {
	ReadAccounts  Action `yaml:"read_accounts"`
	WriteAccounts Action `yaml:"write_accounts"`
	ReadOrders    Action `yaml:"read_orders"`
	WriteOrders   Action `yaml:"write_orders"`
	ReadMarkets   Action `yaml:"read_markets"`
	WriteMarkets  Action `yaml:"write_markets"`
	// Add more predefied actions if needed
}

// BarongActions config to store all allowed actions for Barong
type BarongActions struct {
	ReadUsers  Action `yaml:"read_users"`
	WriteUsers Action `yaml:"write_users"`
	// Add more predefined actions if needed
}

// PeatioAPIV2Config config to store root config for Peatio management API services
type PeatioAPIV2Config struct {
	Keychain Keychain      `yaml:"keychain"`
	Actions  PeatioActions `yaml:"actions"`
}

// BarongAPIV2Config config to store root config for Barong management API services
type BarongAPIV2Config struct {
	Keychain Keychain      `yaml:"keychain"`
	Actions  BarongActions `yaml:"actions"`
}

// ManagementAPIV2Config config to store all management config "management_api_v2.yaml"
type ManagementAPIV2Config struct {
	Barong BarongAPIV2Config `yaml:"barong"`
	Peatio PeatioAPIV2Config `yaml:"peatio"`
}
