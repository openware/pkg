package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountBalance(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"uid":"IDCA2AC08296","balance":"996.23352165725","locked":"0.0"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetAccountBalanceParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		balance, apiError := client.GetAccountBalance(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(balance)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error record not found", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "Couldn't find record.",
			},
		}

		params := GetAccountBalanceParams{
			UID:      "ID1234567890",
			Currency: "bnb",
		}
		balance, apiError := client.GetAccountBalance(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Couldn't find record.")
		assert.Nil(t, balance)
	})

	t.Run("Error invalid currency", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "currency does not have a valid value",
			},
		}

		params := GetAccountBalanceParams{
			UID:      "IDCA2AC08296",
			Currency: "bnbxxx",
		}
		balance, apiError := client.GetAccountBalance(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "currency does not have a valid value")
		assert.Nil(t, balance)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"balance":996.23352165725,locked:0.0}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetAccountBalanceParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		balance, apiError := client.GetAccountBalance(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, balance)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{aaa: 1}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetAccountBalanceParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		balance, apiError := client.GetAccountBalance(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, balance)
	})
}
