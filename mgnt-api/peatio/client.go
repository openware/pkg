package peatio

import (
	"encoding/json"
	"net/http"

	mgntapi "github.com/openware/pkg/mgnt-api"
)

type PeatioMngAPIV2 struct {
	cli *mgntapi.Client
}

func New(cli *mgntapi.Client) *PeatioMngAPIV2 {
	return &PeatioMngAPIV2{
		cli: cli,
	}
}

func (p *PeatioMngAPIV2) GetWithdrawById(tid string) (*Withdraw, error) {
	// TODO: Post with formData
	res, _ := p.cli.Request(http.MethodGet, "withdraws/get", []byte{})
	withdraw := &Withdraw{}
	err := json.Unmarshal([]byte(res), withdraw)

	if err != nil {
		return nil, err
	}

	return withdraw, nil
}
