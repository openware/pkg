package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencyByCode(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","parent_id":"","explorer_transaction":"https://kovan.etherscan.io/tx/#{txid}","explorer_address":"https://kovan.etherscan.io/address/#{address}","type":"coin","deposit_enabled":true,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.3455425","withdraw_fee":"0.0","min_withdraw_amount":"0.3455425","withdraw_limit_24h":"100000.0","withdraw_limit_72h":"200000.0","base_factor":1000000000000000000,"precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","min_confirmations":10,"code":"bnb","min_collection_amount":"0.3455425","visible":true,"subunits":18,"options":{"erc20_contract_address":"0xb8c77482e45f1f44de1745f52c74426c631bdd52"},"created_at":"2020-02-24T15:34:03+01:00","updated_at":"2020-12-02T10:42:33+01:00"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		currency, apiError := client.GetCurrencyByCode("bnb")
		assert.Nil(t, apiError)

		result, err := json.Marshal(currency)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "Invalid",
			},
		}

		currency, apiError := client.GetCurrencyByCode("bnb")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "Invalid")
		assert.Nil(t, currency)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":bnb,price:123.456}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		currency, apiError := client.GetCurrencyByCode("bnb")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{""}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		currency, apiError := client.GetCurrencyByCode("bnb")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})
}

func TestCreateCurrency(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","parent_id":"","explorer_transaction":"https://kovan.etherscan.io/tx/#{txid}","explorer_address":"https://kovan.etherscan.io/address/#{address}","type":"coin","deposit_enabled":true,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.3455425","withdraw_fee":"0.0","min_withdraw_amount":"0.3455425","withdraw_limit_24h":"100000.0","withdraw_limit_72h":"200000.0","base_factor":1000000000000000000,"precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","min_confirmations":10,"code":"bnb","min_collection_amount":"0.3455425","visible":true,"subunits":18,"options":{"erc20_contract_address":"0xb8c77482e45f1f44de1745f52c74426c631bdd52"},"created_at":"2020-02-24T15:34:03+01:00","updated_at":"2020-12-02T10:42:33+01:00"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateCurrencyParams{
			Code:    "bnb",
			Type:    "coin",
			Price:   "10.0",
			Visible: true,
		}
		currency, apiError := client.CreateCurrency(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(currency)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		params := CreateCurrencyParams{
			Code:  "bnb",
			Type:  "coin",
			Price: "10.0",
		}
		currency, apiError := client.CreateCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, currency)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"code":1234}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateCurrencyParams{
			Code:  "bnb",
			Type:  "coin",
			Price: "10.0",
		}
		currency, apiError := client.CreateCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateCurrencyParams{
			Code:  "bnb",
			Type:  "coin",
			Price: "10.0",
		}
		currency, apiError := client.CreateCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})
}

func TestCurrenciesList(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","parent_id":"","explorer_transaction":"https://kovan.etherscan.io/tx/#{txid}","explorer_address":"https://kovan.etherscan.io/address/#{address}","type":"coin","deposit_enabled":true,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.3455425","withdraw_fee":"0.0","min_withdraw_amount":"0.3455425","withdraw_limit_24h":"100000.0","withdraw_limit_72h":"200000.0","base_factor":1000000000000000000,"precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","min_confirmations":10,"code":"bnb","min_collection_amount":"0.3455425","visible":true,"subunits":18,"options":{"erc20_contract_address":"0xb8c77482e45f1f44de1745f52c74426c631bdd52"},"created_at":"2020-02-24T15:34:03+01:00","updated_at":"2020-12-02T10:42:33+01:00"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CurrenciesListParams{}

		currency, apiError := client.GetCurrenciesList(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(currency)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Success response with type", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","parent_id":"","explorer_transaction":"https://kovan.etherscan.io/tx/#{txid}","explorer_address":"https://kovan.etherscan.io/address/#{address}","type":"coin","deposit_enabled":true,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.3455425","withdraw_fee":"0.0","min_withdraw_amount":"0.3455425","withdraw_limit_24h":"100000.0","withdraw_limit_72h":"200000.0","base_factor":1000000000000000000,"precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","min_confirmations":10,"code":"bnb","min_collection_amount":"0.3455425","visible":true,"subunits":18,"options":{"erc20_contract_address":"0xb8c77482e45f1f44de1745f52c74426c631bdd52"},"created_at":"2020-02-24T15:34:03+01:00","updated_at":"2020-12-02T10:42:33+01:00"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CurrenciesListParams{
			Type: "coin",
		}

		currency, apiError := client.GetCurrenciesList(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(currency)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Success response with type (empty)", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CurrenciesListParams{
			Type: "fiat",
		}

		currency, apiError := client.GetCurrenciesList(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(currency)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "management.currency.invalid_type",
			},
		}

		params := CurrenciesListParams{
			Type: "smth",
		}

		currency, apiError := client.GetCurrenciesList(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "management.currency.invalid_type")
		assert.Nil(t, currency)
	})
}

func TestUpdateCurrency(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","parent_id":"","explorer_transaction":"https://kovan.etherscan.io/tx/#{txid}","explorer_address":"https://kovan.etherscan.io/address/#{address}","type":"coin","deposit_enabled":true,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.3455425","withdraw_fee":"0.0","min_withdraw_amount":"0.3455425","withdraw_limit_24h":"100000.0","withdraw_limit_72h":"200000.0","base_factor":1000000000000000000,"precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","min_confirmations":10,"code":"bnb","min_collection_amount":"0.3455425","visible":true,"subunits":18,"options":{"erc20_contract_address":"0xb8c77482e45f1f44de1745f52c74426c631bdd52"},"created_at":"2020-02-24T15:34:03+01:00","updated_at":"2020-12-02T10:42:33+01:00"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateCurrencyParams{
			ID:                "bnb",
			MinWithdrawAmount: "2",
			Price:             "10.0",
			Visible:           true,
		}
		currency, apiError := client.UpdateCurrency(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(currency)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		params := UpdateCurrencyParams{
			ID:                "bnb",
			MinWithdrawAmount: "2",
			Price:             "10.0",
		}
		currency, apiError := client.UpdateCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, currency)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"code":1234}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateCurrencyParams{
			ID:                "bnb",
			MinWithdrawAmount: "2",
			Price:             "10.0",
		}
		currency, apiError := client.UpdateCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateCurrencyParams{
			ID:                "bnb",
			MinWithdrawAmount: "2",
			Price:             "10.0",
		}
		currency, apiError := client.UpdateCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})

}
