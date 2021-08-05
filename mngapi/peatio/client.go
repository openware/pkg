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

// GetBlockchainCurrencyByID call peatio management api to get blockchain currency information by id
func (p *Client) GetBlockchainCurrencyByID(id string) (*BlockchainCurrency, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, fmt.Sprintf("blockchain_currencies/%v", id), nil)
	if apiError != nil {
		return nil, apiError
	}

	blockchainCurrency := &BlockchainCurrency{}
	err := json.Unmarshal([]byte(res), blockchainCurrency)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return blockchainCurrency, nil
}

// GetCurrenciesList call peatio management api to get currency information by code name
func (p *Client) GetCurrenciesList(params CurrenciesListParams) (*[]Currency, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "currencies/list", params)
	if apiError != nil {
		return nil, apiError
	}

	currencies := []Currency{}
	err := json.Unmarshal([]byte(res), &currencies)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return &currencies, nil
}

func (p *Client) CreateCurrency(params CreateCurrencyParams) (*Currency, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "currencies/create", params)
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

func (p *Client) CreateBlockchainCurrency(params CreateBlockchainCurrencyParams) (*BlockchainCurrency, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "blockchain_currencies/new", params)
	if apiError != nil {
		return nil, apiError
	}

	blockchainCurrency := &BlockchainCurrency{}
	err := json.Unmarshal([]byte(res), blockchainCurrency)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return blockchainCurrency, nil
}

func (p *Client) UpdateCurrency(params UpdateCurrencyParams) (*Currency, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPut, "currencies/update", params)
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

func (p *Client) UpdateBlockchainCurrency(params UpdateBlockchainCurrencyParams) (*BlockchainCurrency, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPut, "blockchain_currencies/update", params)
	if apiError != nil {
		return nil, apiError
	}

	blockchainCurrency := &BlockchainCurrency{}
	err := json.Unmarshal([]byte(res), blockchainCurrency)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return blockchainCurrency, nil
}

// CreateWithdraw call peatio management api to create new withdraw
func (p *Client) CreateWithdraw(params CreateWithdrawParams) (*Withdraw, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "withdraws/new", params)
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

// CreateDeposit call peatio management api to create new deposit
func (p *Client) CreateDeposit(params CreateDepositParams) (*Deposit, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "deposits/new", params)
	if apiError != nil {
		return nil, apiError
	}

	deposit := &Deposit{}
	err := json.Unmarshal([]byte(res), deposit)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return deposit, nil
}

// GetDepositByID call peatio management api to get deposit information by transaction ID
func (p *Client) GetDepositByID(tid string) (*Deposit, *mngapi.APIError) {
	// Build parameters
	params := map[string]interface{}{
		"tid": tid,
	}

	res, apiError := p.mngapiClient.Request(http.MethodPost, "deposits/get", params)
	if apiError != nil {
		return nil, apiError
	}

	deposit := &Deposit{}
	err := json.Unmarshal([]byte(res), deposit)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return deposit, nil
}

// GetDeposits call peatio management api to get deposits as paginated collection
func (p *Client) GetDeposits(params GetDepositsParams) ([]*Deposit, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "deposits", params)
	if apiError != nil {
		return nil, apiError
	}

	deposits := make([]*Deposit, 0)
	err := json.Unmarshal([]byte(res), &deposits)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: fmt.Sprintf("payload: %s; error: %s", res, err.Error())}
	}

	return deposits, nil
}

// CreateEngine call peatio management api to create new engine
func (p *Client) CreateEngine(params CreateEngineParams) (*Engine, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "engines/new", params)
	if apiError != nil {
		return nil, apiError
	}

	engine := &Engine{}
	err := json.Unmarshal([]byte(res), engine)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return engine, nil
}

// UpdateEngine call peatio management api to update engine
func (p *Client) UpdateEngine(params UpdateEngineParams) (*Engine, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "engines/update", params)
	if apiError != nil {
		return nil, apiError
	}

	engine := &Engine{}
	err := json.Unmarshal([]byte(res), engine)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return engine, nil
}

// GetEngines call peatio management api to get engines
func (p *Client) GetEngines(params GetEngineParams) ([]*Engine, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "engines/get", params)
	if apiError != nil {
		return nil, apiError
	}

	engines := make([]*Engine, 0)

	err := json.Unmarshal([]byte(res), &engines)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return engines, nil
}

// GetMarkets call peatio management api to get all markets
func (p *Client) GetMarkets() ([]*Market, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "/markets/list", nil)
	if apiError != nil {
		return nil, apiError
	}

	markets := make([]*Market, 0)

	err := json.Unmarshal([]byte(res), &markets)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return markets, nil
}

// UpdateMarket call peatio management api to update market
func (p *Client) UpdateMarket(params UpdateMarketParams) (*Market, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPut, "/markets/update", params)
	if apiError != nil {
		return nil, apiError
	}

	market := &Market{}
	err := json.Unmarshal([]byte(res), market)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return market, nil
}

func (p *Client) CreateMarket(params CreateMarketParams) (*Market, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "markets/new", params)
	if apiError != nil {
		return nil, apiError
	}

	market := &Market{}
	err := json.Unmarshal([]byte(res), market)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return market, nil
}

func (p *Client) GetMarketByID(id string) (*Market, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, fmt.Sprintf("markets/%v", id), nil)
	if apiError != nil {
		return nil, apiError
	}

	market := &Market{}
	err := json.Unmarshal([]byte(res), market)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return market, nil
}

func (p *Client) CreateMember(params CreateMemberParams) (*Member, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "members", params)
	if apiError != nil {
		return nil, apiError
	}

	member := &Member{}
	err := json.Unmarshal([]byte(res), member)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return member, nil
}

// CreateWallet call peatio management api to create wallet
func (p *Client) CreateWallet(params CreateWalletParams) (*Wallet, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "wallets/new", params)
	if apiError != nil {
		return nil, apiError
	}

	wallet := &Wallet{}
	err := json.Unmarshal([]byte(res), wallet)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return wallet, nil
}

// UpdateWallet call peatio management api to update wallet
func (p *Client) UpdateWallet(params UpdateWalletParams) (*Wallet, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "wallets/update", params)
	if apiError != nil {
		return nil, apiError
	}

	wallet := &Wallet{}
	err := json.Unmarshal([]byte(res), wallet)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return wallet, nil
}

// GetWallets call peatio management api to get wallets
func (p *Client) GetWallets() ([]*Wallet, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, "wallets", nil)
	if apiError != nil {
		return nil, apiError
	}

	wallets := make([]*Wallet, 0)

	err := json.Unmarshal([]byte(res), &wallets)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return wallets, nil
}

func (p *Client) GetWalletByID(id int) (*Wallet, *mngapi.APIError) {
	res, apiError := p.mngapiClient.Request(http.MethodPost, fmt.Sprintf("wallets/%v", id), nil)
	if apiError != nil {
		return nil, apiError
	}

	wallet := &Wallet{}
	err := json.Unmarshal([]byte(res), wallet)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return wallet, nil
}
