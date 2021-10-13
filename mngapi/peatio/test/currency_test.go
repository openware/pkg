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

		expected := `{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","status":"enabled","type":"coin","precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","code":"bnb","networks":[]}`
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

		expected := `{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","status":"enabled","type":"coin","precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","code":"bnb","networks":[]}`
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

		expected := `[{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","status":"enabled","type":"coin","precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","code":"bnb","networks":[]}]`
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

		expected := `[{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","status":"enabled","type":"coin","precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","code":"bnb","networks":[]}]`
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

		expected := `{"id":"bnb","name":"Binance Coin","description":"","homepage":"","price":"23.8","status":"enabled","type":"coin","precision":10,"position":10,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","code":"bnb","networks":[]}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateCurrencyParams{
			ID:       "bnb",
			Position: 10,
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
			ID:       "bnb",
			Position: 10,
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
			ID:       "bnb",
			Position: 10,
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
			ID:       "bnb",
			Position: 10,
		}
		currency, apiError := client.UpdateCurrency(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})
}
