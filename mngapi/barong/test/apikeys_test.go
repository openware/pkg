package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/barong"
	"github.com/stretchr/testify/assert"
)

func TestCreateAPIKeys(t *testing.T) {
	t.Run("Successful response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"kid":"133742","algorithm":"HS256","scope":["trade"],"state":"active","secret":"something_in_the_way","created_at":"2021-02-15T10:15:18Z","updated_at":"2021-02-15T10:15:18Z"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateAPIKeyParams{
			UID:       "IDCA2AC08296",
			Algorithm: "HS256",
			Scopes:    "trade",
		}
		apiKey, apiError := client.CreateAPIKey(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(apiKey)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error could not save secret", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "api_key.could_not_save_secret",
			},
		}

		params := CreateAPIKeyParams{
			UID:       "IDCA2AC08296",
			Algorithm: "HS256",
			Scopes:    "trade",
		}
		apiKey, apiError := client.CreateAPIKey(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "api_key.could_not_save_secret")
		assert.Nil(t, apiKey)
	})
}
