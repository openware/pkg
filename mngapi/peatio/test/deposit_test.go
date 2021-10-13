package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDepositAddress(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"uid":"IDCA2AC08296","address":"0x5b89a2a38b7398c71cfc420a6ed3b5f2a1a01a3e","blockchain_key":"btc-testnet","currencies":["usdt","bnb","uni"],"state":"active","remote":false}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GenerateDepositAddressParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		paymentAddress, apiError := client.GenerateDepositAddress(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(paymentAddress)
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
				Error:      "Couldn't find record.",
			},
		}

		params := GenerateDepositAddressParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		paymentAddress, apiError := client.GenerateDepositAddress(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Couldn't find record.")
		assert.Nil(t, paymentAddress)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"uid":"IDCA2AC08296","address":0x5b89a2a38b7398c71cfc420a6ed3b5f2a1a01a3e,"currencies":["usdt","bnb","uni"],"state":"active","remote":false}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GenerateDepositAddressParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		paymentAddress, apiError := client.GenerateDepositAddress(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, paymentAddress)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{[]}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GenerateDepositAddressParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		paymentAddress, apiError := client.GenerateDepositAddress(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, paymentAddress)
	})
}

func TestCreateDeposit(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"tid":"TIDBD6B265303","blockchain_key":"","currency":"usd","address":"","uid":"ID732785AC58","type":"fiat","amount":"750.77","state":"submitted","created_at":"2021-03-02T07:33:02+01:00","completed_at":null,"transfer_type":"fiat"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateDepositParams{
			UID:      "ID732785AC58",
			Currency: "usd",
			Amount:   10.0,
		}
		deposit, apiError := client.CreateDeposit(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(deposit)
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

		params := CreateDepositParams{
			UID:      "ID732785AC58",
			Currency: "bnb",
			Amount:   10.0,
		}
		deposit, apiError := client.CreateDeposit(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, deposit)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":1234}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateDepositParams{
			UID:      "ID732785AC58",
			Currency: "bnb",
			Amount:   10.0,
		}
		deposit, apiError := client.CreateDeposit(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposit)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateDepositParams{
			UID:      "ID732785AC58",
			Currency: "bnb",
			Amount:   10.0,
		}
		deposit, apiError := client.CreateDeposit(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposit)
	})
}

func TestGetDepositByID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"tid":"TIDF6289303E1","blockchain_key":"","currency":"btc","address":"","uid":"ID6CBD4E84C7","type":"coin","amount":"6346.0","state":"submitted","created_at":"2021-03-02T05:54:52+01:00","completed_at":null,"blockchain_txid":"56bzwdd359kxd0r3qt3mz1cbcrc8o3r5hshlgbag42z7ka2o9hd4b5me80hh0khb","blockchain_confirmations":"711753","transfer_type":"crypto"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		deposit, apiError := client.GetDepositByID("TIDF6289303E1")
		assert.Nil(t, apiError)

		result, err := json.Marshal(deposit)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error record not found", func(t *testing.T) {
		client, _ := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "Couldn't find record.",
			},
		}

		deposit, apiError := client.GetDepositByID("TIDXXXX")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Couldn't find record.")
		assert.Nil(t, deposit)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":1234}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		deposit, apiError := client.GetDepositByID("TIDF6289303E1")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposit)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{","}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		deposit, apiError := client.GetDepositByID("TIDF6289303E1")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposit)
	})
}

func TestGetDeposits(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":1,"tid":"TID9119EEAE36","blockchain_key":"","currency":"usd","address":"","uid":"ID9C5C7208EB","type":"fiat","amount":"8423.0","state":"collected","created_at":"2021-03-02T04:40:06+01:00","completed_at":"2021-03-02T04:40:06+01:00","transfer_type":"fiat"},{"id":2,"tid":"TID17505F194C","blockchain_key":"btc-testnet","currency":"btc","address":"","uid":"ID0B0C77487A","type":"coin","amount":"191.0","state":"fee_processing","created_at":"2021-03-02T04:40:06+01:00","completed_at":"2021-03-02T04:40:06+01:00","blockchain_txid":"wfmvae8elj0egr309u9oodl58ypzifdfjz9vd1i82t3ng4uepmokagack0shfsif","blockchain_confirmations":"367597","transfer_type":"crypto"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetDepositsParams{
			UID: "IDCA2AC08296",
		}
		deposits, apiError := client.GetDeposits(params)
		assert.Nil(t, apiError)
		assert.NotNil(t, deposits)

		result, err := json.Marshal(deposits)
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
				Error:      "Error",
			},
		}

		params := GetDepositsParams{
			UID: "IDCA2AC08296",
		}
		deposits, apiError := client.GetDeposits(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "Error")
		assert.Nil(t, deposits)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":1,"tid":"TID9119EEAE36","currency":"usd","address":"","uid":"ID9C5C7208EB","type":"fiat","amount":8423.0,"state":"collected","created_at":"2021-03-02T04:40:06+01:00","completed_at":"2021-03-02T04:40:06+01:00","transfer_type":"fiat"},{"id":2,"tid":"TID17505F194C","currency":"btc","address":"","uid":"ID0B0C77487A","type":"coin","amount":"191.0","state":"fee_processing","created_at":"2021-03-02T04:40:06+01:00","completed_at":"2021-03-02T04:40:06+01:00","blockchain_txid":"wfmvae8elj0egr309u9oodl58ypzifdfjz9vd1i82t3ng4uepmokagack0shfsif","blockchain_confirmations":"367597","transfer_type":"crypto"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetDepositsParams{
			UID: "IDCA2AC08296",
		}
		deposits, apiError := client.GetDeposits(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposits)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetDepositsParams{
			UID: "IDCA2AC08296",
		}
		deposits, apiError := client.GetDeposits(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposits)
	})
}
