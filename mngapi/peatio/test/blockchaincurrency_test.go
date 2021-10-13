package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestCreateBlockchainCurrency(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"1","currency_id":"btc","blockchain_key":"btc-testnet","parent_id":"","status":"enabled","deposit_enabled":false,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.0","withdraw_fee":"0.0000000002557544","min_withdraw_amount":"0.0000000025575447","base_factor":1000000000000000000,"min_collection_amount":"123","options":null}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateBlockchainCurrencyParams{
			CurrencyID:    "bnb",
			BlockchainKey: "eth-rinkeby",
			ParentID:      "eth",
		}
		network, apiError := client.CreateBlockchainCurrency(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(network)
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

		params := CreateBlockchainCurrencyParams{
			CurrencyID:    "bnb",
			BlockchainKey: "eth-rinkeby",
			ParentID:      "eth",
		}
		network, apiError := client.CreateBlockchainCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, network)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"blockchain_key":1234}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateBlockchainCurrencyParams{
			CurrencyID:    "bnb",
			BlockchainKey: "eth-rinkeby",
			ParentID:      "eth",
		}
		network, apiError := client.CreateBlockchainCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, network)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateBlockchainCurrencyParams{
			CurrencyID:    "bnb",
			BlockchainKey: "eth-rinkeby",
			ParentID:      "eth",
		}

		network, apiError := client.CreateBlockchainCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, network)
	})
}

func TestUpdateBlockchainCurrency(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"1","currency_id":"btc","blockchain_key":"btc-testnet","parent_id":"","status":"enabled","deposit_enabled":false,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.0","withdraw_fee":"0.0000000002557544","min_withdraw_amount":"0.0000000025575447","base_factor":1000000000000000000,"min_collection_amount":"123","options":null}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateBlockchainCurrencyParams{
			ID:     "1",
			Status: "hidden",
		}
		network, apiError := client.UpdateBlockchainCurrency(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(network)
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

		params := UpdateBlockchainCurrencyParams{
			ID:     "1",
			Status: "hidden",
		}
		network, apiError := client.UpdateBlockchainCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, network)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"status":1234}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateBlockchainCurrencyParams{
			ID:     "1",
			Status: "hidden",
		}
		network, apiError := client.UpdateBlockchainCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, network)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateBlockchainCurrencyParams{
			ID:     "1",
			Status: "hidden",
		}
		network, apiError := client.UpdateBlockchainCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, network)
	})
}

func TestGetBlockchainCurrencyByID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"1","currency_id":"btc","blockchain_key":"btc-testnet","parent_id":"","status":"enabled","deposit_enabled":false,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.0","withdraw_fee":"0.0000000002557544","min_withdraw_amount":"0.0000000025575447","base_factor":1000000000000000000,"min_collection_amount":"123","options":null}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		network, apiError := client.GetBlockchainCurrencyByID("1")
		assert.Nil(t, apiError)

		result, err := json.Marshal(network)
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

		network, apiError := client.GetBlockchainCurrencyByID("bnb")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "Invalid")
		assert.Nil(t, network)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		network, apiError := client.GetBlockchainCurrencyByID("1")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, network)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{""}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		network, apiError := client.GetBlockchainCurrencyByID("bnb")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, network)
	})
}
