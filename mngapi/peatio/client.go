package peatio

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/openware/pkg/mngapi"
)

// Client is peatio management api client instance
type Client struct {
	mngapiClient mngapi.DefaultClient
}

// New return peatio management api client
func New(rootAPIUrl, endpointPrefix, jwtIssuer, jwtAlgo, jwtPrivateKey string) (*Client, error) {
	client, err := mngapi.New(rootAPIUrl, endpointPrefix, jwtIssuer, jwtAlgo, jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	return &Client{
		mngapiClient: client,
	}, nil
}

// GetCurrencyByCode call peatio management api to get currency information by code name
func (p *Client) GetCurrencyByCode(code string) (*Currency, *mngapi.APIError) {
	res, err := p.mngapiClient.Request(http.MethodPost, fmt.Sprintf("currencies/%v", code), nil)
	if err != nil {
		return nil, err
	}

	currency := &Currency{}
	_ = json.Unmarshal([]byte(res), currency)

	return currency, nil
}

// CreateWithdraw call peatio management api to create new withdraw
func (p *Client) CreateWithdraw(params CreateWithdrawParams) (*Withdraw, *mngapi.APIError) {
	res, err := p.mngapiClient.Request(http.MethodPost, "withdraw/new", params)
	if err != nil {
		return nil, err
	}

	withdraw := &Withdraw{}
	_ = json.Unmarshal([]byte(res), withdraw)

	return withdraw, nil
}

// GetWithdrawByID call peatio management api to get withdraw information by transaction ID
func (p *Client) GetWithdrawByID(tid string) (*Withdraw, *mngapi.APIError) {
	// Build parameters
	params := map[string]interface{}{
		"tid": tid,
	}

	res, err := p.mngapiClient.Request(http.MethodPost, "withdraws/get", params)
	if err != nil {
		return nil, err
	}

	withdraw := &Withdraw{}
	_ = json.Unmarshal([]byte(res), withdraw)

	return withdraw, nil
}

// GetAccountBalance call peatio management api to get account balance
func (p *Client) GetAccountBalance(params GetAccountBalanceParams) (*Balance, *mngapi.APIError) {
	res, err := p.mngapiClient.Request(http.MethodPost, "accounts/balance", params)
	if err != nil {
		return nil, err
	}

	balance := &Balance{}
	_ = json.Unmarshal([]byte(res), balance)

	return balance, nil
}

// GenerateDepositAddress call peatio management api to generate new deposit address
func (p *Client) GenerateDepositAddress(params GenerateDepositAddressParams) (*PaymentAddress, *mngapi.APIError) {
	res, err := p.mngapiClient.Request(http.MethodPost, "deposit_address/new", params)
	if err != nil {
		return nil, err
	}

	paymentAddress := &PaymentAddress{}
	_ = json.Unmarshal([]byte(res), paymentAddress)

	return paymentAddress, nil
}
