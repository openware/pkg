package peatio

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/openware/pkg/mngapi"
)

type PeatioMngAPIV2 struct {
	cli *mngapi.Client
}

func New(rootAPIUrl, endpointPrefix, jwtIssuer, jwtPrivateKey string) *PeatioMngAPIV2 {
	cli, _ := mngapi.New(rootAPIUrl, endpointPrefix, jwtIssuer, "RS256", jwtPrivateKey)

	return &PeatioMngAPIV2{
		cli: cli,
	}
}

func (p *PeatioMngAPIV2) GetCurrencyByCode(code string) (*Currency, *mngapi.APIError) {
	path := fmt.Sprintf("/currencies/%v", code)
	params := []byte{}
	res, apiError := p.cli.Request(http.MethodPost, path, params)
	if apiError != nil {
		return nil, apiError
	}

	currency := &Currency{}
	_ = json.Unmarshal([]byte(res), currency)

	return currency, nil
}

func (p *PeatioMngAPIV2) GetWithdrawById(tid string) (*Withdraw, *mngapi.APIError) {
	params := fmt.Sprintf(`"tid":"%v"`, tid)
	res, apiError := p.cli.Request(http.MethodPost, "withdraws/get", []byte(params))

	if apiError != nil {
		return nil, apiError
	}

	withdraw := &Withdraw{}
	_ = json.Unmarshal([]byte(res), withdraw)

	return withdraw, nil
}
