package barong

import mgntapi "github.com/openware/pkg/mgnt-api"

type BarongAPIV2 struct {
	apiClient *mgntapi.ManagementAPIV2
}

func New(config *mgntapi.BarongAPIV2Config) *BarongAPIV2 {
	return &BarongAPIV2{
		apiClient: mgntapi.New("", "", &config.Keychain.Opendax),
	}
}
