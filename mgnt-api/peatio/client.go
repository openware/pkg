package peatio

import (
	"encoding/json"
	"net/http"

	mgntapi "github.com/openware/pkg/mgnt-api"
)

type PeatioAPIV2 struct {
	apiClient *mgntapi.ManagementAPIV2
}

func New(config *mgntapi.PeatioAPIV2Config) *PeatioAPIV2 {
	return &PeatioAPIV2{
		apiClient: mgntapi.New("https://dev.yellow.openware.work", "/api/v2/peatio/management/", &config.Keychain.Opendax),
	}
}

func (p *PeatioAPIV2) GetWithdrawById(tid string) (*Withdraw, error) {
	// TODO: Post with formData
	res, _ := p.apiClient.Request(http.MethodGet, "withdraws/get", []byte{})
	withdraw := &Withdraw{}
	err := json.Unmarshal([]byte(res), withdraw)

	if err != nil {
		return nil, err
	}

	return withdraw, nil
}
