package mgntapi

// Action config to store map which consists of pairs: Action name => signature Key ID
// Example:
// 		read_users:
// 				required_signatures: ["applogic"]
//    		requires_barong_totp: false
type Action map[string]interface{}

// Keychain config to stroe map which consists of pairs: Key ID => algorithm and private key value
// Example:
//  	applogic:
// 				algorithm: RS256
// 				value: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBL.....
type Keychain map[string]interface{}

// PeatioAPIV2Config config to store root config for Peatio management API services
type PeatioAPIV2Config struct {
	Keychain []Keychain `yaml:"keychain"`
	Actions  []Action   `yaml:"actions"`
}

// BarongAPIV2Config config to store root config for Barong management API services
type BarongAPIV2Config struct {
	Keychain []Keychain `yaml:"keychain"`
	Actions  []Action   `yaml:"actions"`
}

// ManagementAPIV2Config config to store all management config "management_api_v2.yaml"
type ManagementAPIV2Config struct {
	Barong BarongAPIV2Config `yaml:"barong"`
	Peatio PeatioAPIV2Config `yaml:"peatio"`
}
