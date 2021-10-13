package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestGetWallets(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":1,"name":"BTC Wallet","kind":"deposit","currencies":["btc","xtz"],"address":"address","gateway":"opendax-cloud","max_balance":"0.00000001","balance":{"btc":"8","xtz":"N/A"},"blockchain_key":"opendax_cloud","status":"active"}]`

		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		wallets, apiError := client.GetWallets()
		assert.Nil(t, apiError)

		result, err := json.Marshal(wallets)
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

		wallets, apiError := client.GetWallets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, wallets)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id": "test"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		wallets, apiError := client.GetWallets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, wallets)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"-"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		wallets, apiError := client.GetWallets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, wallets)
	})
}

func TestUpdateWallet(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"name":"BTC Wallet","kind":"deposit","currencies":["btc","xtz"],"address":"address","gateway":"opendax-cloud","max_balance":"0.00000001","balance":{"btc":"8","xtz":"N/A"},"blockchain_key":"opendax_cloud","status":"active"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateWalletParams{
			BlockchainKey: "opendax-cloud",
			Name:          "BTC Deposit Wallet",
			Kind:          "active",
			Gateway:       "opendax_cloud",
			Address:       "address",
			Currencies:    []string{},
			Settings:      Settings{},
			MaxBalance:    "0.0",
			Status:        "active",
		}

		wallet, apiError := client.UpdateWallet(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(wallet)
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

		params := UpdateWalletParams{
			BlockchainKey: "opendax-cloud",
			Name:          "BTC Deposit Wallet",
			Kind:          "active",
			Gateway:       "opendax_cloud",
			Address:       "address",
			Currencies:    []string{},
			Settings:      Settings{},
			MaxBalance:    "0.0",
			Status:        "active",
		}

		wallet, apiError := client.UpdateWallet(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, wallet)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"currencies": fds}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateWalletParams{
			BlockchainKey: "opendax-cloud",
			Name:          "BTC Deposit Wallet",
			Kind:          "active",
			Gateway:       "opendax_cloud",
			Address:       "address",
			Currencies:    []string{},
			Settings:      Settings{},
			MaxBalance:    "0.0",
			Status:        "active",
		}
		wallet, apiError := client.UpdateWallet(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, wallet)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateWalletParams{
			BlockchainKey: "opendax-cloud",
			Name:          "BTC Deposit Wallet",
			Kind:          "active",
			Gateway:       "opendax_cloud",
			Address:       "address",
			Currencies:    []string{},
			Settings:      Settings{},
			MaxBalance:    "0.0",
			Status:        "active",
		}
		wallet, apiError := client.UpdateWallet(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, wallet)
	})
}

func TestCreateWallet(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"name":"BTC Wallet","kind":"deposit","currencies":["btc","xtz"],"address":"address","gateway":"opendax-cloud","max_balance":"0.00000001","balance":{"btc":"8","xtz":"N/A"},"blockchain_key":"opendax_cloud","status":"active"}`

		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWalletParams{
			BlockchainKey: "opendax-cloud",
			Name:          "BTC Deposit Wallet",
			Kind:          "active",
			Gateway:       "opendax_cloud",
			Address:       "address",
			Currencies:    []string{},
			Settings:      Settings{},
			MaxBalance:    "0.0",
			Status:        "active",
		}
		wallet, apiError := client.CreateWallet(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(wallet)
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
		params := CreateWalletParams{
			BlockchainKey: "opendax-cloud",
			Name:          "BTC Deposit Wallet",
			Kind:          "active",
			Gateway:       "opendax_cloud",
			Address:       "address",
			Currencies:    []string{},
			Settings:      Settings{},
			MaxBalance:    "0.0",
			Status:        "active",
		}
		wallet, apiError := client.CreateWallet(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, wallet)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id": opendax_cloud}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWalletParams{
			BlockchainKey: "opendax-cloud",
			Name:          "BTC Deposit Wallet",
			Kind:          "active",
			Gateway:       "opendax_cloud",
			Address:       "address",
			Currencies:    []string{},
			Settings:      Settings{},
			MaxBalance:    "0.0",
			Status:        "active",
		}
		wallet, apiError := client.CreateWallet(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, wallet)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWalletParams{
			BlockchainKey: "opendax-cloud",
			Name:          "BTC Deposit Wallet",
			Kind:          "active",
			Gateway:       "opendax_cloud",
			Address:       "address",
			Currencies:    []string{},
			Settings:      Settings{},
			MaxBalance:    "0.0",
			Status:        "active",
		}
		wallet, apiError := client.CreateWallet(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, wallet)
	})
}

func TestGetWalletByID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"name":"BTC Wallet","kind":"deposit","currencies":["btc","xtz"],"address":"address","gateway":"opendax-cloud","max_balance":"0.00000001","balance":{"btc":"8","xtz":"N/A"},"blockchain_key":"opendax_cloud","status":"active"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		wallet, apiError := client.GetWalletByID(1)
		assert.Nil(t, apiError)

		result, err := json.Marshal(wallet)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Success response with no linked currencies", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"name":"BTC Wallet","kind":"deposit","currencies":[],"address":"address","gateway":"opendax-cloud","max_balance":"0.00000001","balance":null,"blockchain_key":"opendax_cloud","status":"active"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		wallet, apiError := client.GetWalletByID(1)
		assert.Nil(t, apiError)

		result, err := json.Marshal(wallet)
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

		wallet, apiError := client.GetWalletByID(1)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "Invalid")
		assert.Nil(t, wallet)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,name:123.456}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		market, apiError := client.GetWalletByID(1)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{""}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		wallet, apiError := client.GetWalletByID(1)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, wallet)
	})
}
