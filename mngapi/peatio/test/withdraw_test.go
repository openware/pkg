package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestCreateWithdraw(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":"TIDE54B7D229E","uid":"ID16421C020A","currency":"btc","note":"","type":"coin","amount":"0.1195","fee":"0.0005","rid":"1CzSHQnuwp52ErrrtM169FW4FuuRhEksMR","state":"skipped","created_at":"2021-01-12T07:27:41+01:00","blockchain_txid":null,"transfer_type":"crypto"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(withdraw)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Success response with empty txid", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":"TIDE54B7D229E","uid":"ID16421C020A","currency":"btc","note":"","type":"coin","amount":"0.1195","fee":"0.0005","rid":"1CzSHQnuwp52ErrrtM169FW4FuuRhEksMR","state":"skipped","created_at":"2021-01-12T07:27:41+01:00","blockchain_txid":"","transfer_type":"crypto"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(withdraw)
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

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, withdraw)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":1234}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, withdraw)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, withdraw)
	})
}

func TestGetWithdrawByID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":"TIDE54B7D229E","uid":"ID16421C020A","currency":"btc","note":"","type":"coin","amount":"0.1195","fee":"0.0005","rid":"1CzSHQnuwp52ErrrtM169FW4FuuRhEksMR","state":"skipped","created_at":"2021-01-12T07:27:41+01:00","blockchain_txid":null,"transfer_type":"crypto"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		withdraw, apiError := client.GetWithdrawByID("TIDE54B7D229E")
		assert.Nil(t, apiError)

		result, err := json.Marshal(withdraw)
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

		withdraw, apiError := client.GetWithdrawByID("TIDXXXX")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Couldn't find record.")
		assert.Nil(t, withdraw)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":1234}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		withdraw, apiError := client.GetWithdrawByID("TIDE54B7D229E")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, withdraw)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{","}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		withdraw, apiError := client.GetWithdrawByID("TIDE54B7D229E")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, withdraw)
	})
}
