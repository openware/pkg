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

type PeatioMngAPIV2Client interface {
	GetCurrencyByCode(code string) (*Currency, *mngapi.APIError)
	CreateWithdraw(params CreateWithdrawParams) (*Withdraw, *mngapi.APIError)
	GetWithdrawByID(tid string) (*Withdraw, *mngapi.APIError)
	GetAccountBalance(params GetAccountBalanceParams) (*Balance, *mngapi.APIError)
	GenerateDepositAddress(params GenerateDepositAddressParams) (*PaymentAddress, *mngapi.APIError)
}

func New(rootAPIUrl, endpointPrefix, jwtIssuer, jwtPrivateKey string) PeatioMngAPIV2Client {
	cli, _ := mngapi.New(rootAPIUrl, endpointPrefix, jwtIssuer, "RS256", jwtPrivateKey)

	return &PeatioMngAPIV2{
		cli: cli,
	}
}

func (p *PeatioMngAPIV2) GetCurrencyByCode(code string) (*Currency, *mngapi.APIError) {
	res, apiError := p.cli.Request(http.MethodPost, fmt.Sprintf("currencies/%v", code), nil)
	if apiError != nil {
		return nil, apiError
	}

	currency := &Currency{}
	_ = json.Unmarshal([]byte(res), currency)

	return currency, nil
}

func (p *PeatioMngAPIV2) CreateWithdraw(params CreateWithdrawParams) (*Withdraw, *mngapi.APIError) {
	res, apiError := p.cli.Request(http.MethodPost, "withdraw/new", params)
	if apiError != nil {
		return nil, apiError
	}

	withdraw := &Withdraw{}
	_ = json.Unmarshal([]byte(res), withdraw)

	return withdraw, nil
}

func (p *PeatioMngAPIV2) GetWithdrawByID(tid string) (*Withdraw, *mngapi.APIError) {
	// Build parameters
	params := map[string]interface{}{
		"tid": tid,
	}

	res, apiError := p.cli.Request(http.MethodPost, "withdraws/get", params)
	if apiError != nil {
		return nil, apiError
	}

	withdraw := &Withdraw{}
	_ = json.Unmarshal([]byte(res), withdraw)

	return withdraw, nil
}

func (p *PeatioMngAPIV2) GetAccountBalance(params GetAccountBalanceParams) (*Balance, *mngapi.APIError) {
	res, apiError := p.cli.Request(http.MethodPost, "accounts/balance", params)
	if apiError != nil {
		return nil, apiError
	}

	balance := &Balance{}
	_ = json.Unmarshal([]byte(res), balance)

	return balance, nil
}

func (p *PeatioMngAPIV2) GenerateDepositAddress(params GenerateDepositAddressParams) (*PaymentAddress, *mngapi.APIError) {
	res, apiError := p.cli.Request(http.MethodPost, "deposit_address/new", params)
	if apiError != nil {
		return nil, apiError
	}

	paymentAddress := &PaymentAddress{}
	_ = json.Unmarshal([]byte(res), paymentAddress)

	return paymentAddress, nil
}
