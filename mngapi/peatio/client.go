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
func New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey string) (*Client, error) {
	client, err := mngapi.New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	return &Client{
		mngapiClient: client,
	}, nil
}

// GetCurrencyByCode call peatio management api to get currency information by code name
func (p *Client) GetCurrencyByCode(code string) (*Currency, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, fmt.Sprintf("currencies/%v", code), nil)
	if apiError != nil {
		return nil, apiError
	}

	currency := &Currency{}
	err := json.Unmarshal([]byte(res), currency)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return currency, nil
}

// CreateWithdraw call peatio management api to create new withdraw
func (p *Client) CreateWithdraw(params CreateWithdrawParams) (*Withdraw, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "withdraw/new", params)
	if apiError != nil {
		return nil, apiError
	}

	withdraw := &Withdraw{}
	err := json.Unmarshal([]byte(res), withdraw)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return withdraw, nil
}

// GetWithdrawByID call peatio management api to get withdraw information by transaction ID
func (p *Client) GetWithdrawByID(tid string) (*Withdraw, *mngapi.APIError) {
	// Build parameters
	params := map[string]interface{}{
		"tid": tid,
	}

	res, apiError := p.mngapiClient.Request(http.MethodPost, "withdraws/get", params)
	if apiError != nil {
		return nil, apiError
	}

	withdraw := &Withdraw{}
	err := json.Unmarshal([]byte(res), withdraw)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return withdraw, nil
}

// GetAccountBalance call peatio management api to get account balance
func (p *Client) GetAccountBalance(params GetAccountBalanceParams) (*Balance, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "accounts/balance", params)
	if apiError != nil {
		return nil, apiError
	}

	balance := &Balance{}
	err := json.Unmarshal([]byte(res), balance)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return balance, nil
}

// GenerateDepositAddress call peatio management api to generate new deposit address
func (p *Client) GenerateDepositAddress(params GenerateDepositAddressParams) (*PaymentAddress, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "deposit_address/new", params)
	if apiError != nil {
		return nil, apiError
	}

	paymentAddress := &PaymentAddress{}
	err := json.Unmarshal([]byte(res), paymentAddress)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return paymentAddress, nil
}
