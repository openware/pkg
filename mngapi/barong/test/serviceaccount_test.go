package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/barong"
	"github.com/stretchr/testify/assert"
)

func TestCreateServiceAccount(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":"test+SI0388B7681C@yellow.com","uid":"SI0388B7681C","role":"service_account","level":3,"state":"active","user":{"email":"test@test.com","uid":"IDCA2AC08296","role":"superadmin","level":3,"otp":true,"state":"active","referral_uid":"","data":"{\"onboarding\":true,\"language\":\"en\"}"},"created_at":"2021-02-15T10:15:18Z","updated_at":"2021-02-15T10:15:18Z"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateServiceAccountParams{
			OwnerUID: "IDCA2AC08296",
			Role:     "service_account",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(serviceAccount)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Success response w/ state", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":"test+SI0388B7681C@yellow.com","uid":"SI0388B7681C","role":"service_account","level":3,"state":"active","user":{"email":"test@test.com","uid":"IDCA2AC08296","role":"superadmin","level":3,"otp":true,"state":"active","referral_uid":"","data":"{\"onboarding\":true,\"language\":\"en\"}"},"created_at":"2021-02-15T10:15:18Z","updated_at":"2021-02-15T10:15:18Z"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateServiceAccountParams{
			OwnerUID: "IDCA2AC08296",
			Role:     "service_account",
			State:    "active",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(serviceAccount)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error user doesn't exist", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "User doesnt exist",
			},
		}

		params := CreateServiceAccountParams{
			OwnerUID: "ID123456789",
			Role:     "service_account",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "User doesnt exist")
		assert.Nil(t, serviceAccount)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":test@test.com}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateServiceAccountParams{
			OwnerUID: "ID123456789",
			Role:     "service_account",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, serviceAccount)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateServiceAccountParams{
			OwnerUID: "ID123456789",
			Role:     "service_account",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, serviceAccount)
	})
}

func TestDeleteServiceAccountByUID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":"test+SI0388B7681C@yellow.com","uid":"SI0388B7681C","role":"service_account","level":3,"state":"active","user":{"email":"test@test.com","uid":"IDCA2AC08296","role":"superadmin","level":3,"otp":true,"state":"disabled","referral_uid":"","data":"{\"onboarding\":true,\"language\":\"en\"}"},"created_at":"2021-02-15T10:15:18Z","updated_at":"2021-02-15T10:15:18Z"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		uid := "SI0388B7681C"
		serviceAccount, apiError := client.DeleteServiceAccountByUID(uid)
		assert.Nil(t, apiError)

		result, err := json.Marshal(serviceAccount)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error record is not found", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.MngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "Record is not found",
			},
		}

		uid := "SI123456789"
		serviceAccount, apiError := client.DeleteServiceAccountByUID(uid)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Record is not found")
		assert.Nil(t, serviceAccount)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":123}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		uid := "SI123456789"
		serviceAccount, apiError := client.DeleteServiceAccountByUID(uid)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, serviceAccount)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{error}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		uid := "SI123456789"
		serviceAccount, apiError := client.DeleteServiceAccountByUID(uid)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, serviceAccount)
	})
}
