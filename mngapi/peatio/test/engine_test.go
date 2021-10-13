package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestCreateEngine(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"name":"opendax_cloud","driver":"opendax","uid":"UID123123","url":"https://example.com","state":"online"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateEngineParams{
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}

		engine, apiError := client.CreateEngine(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(engine)
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

		params := CreateEngineParams{
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}

		engine, apiError := client.CreateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, engine)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"name": opendax_cloud}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateEngineParams{
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}
		engine, apiError := client.CreateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateEngineParams{
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}
		engine, apiError := client.CreateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})
}

func TestUpdateEngine(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"name":"opendax_cloud","driver":"opendax","uid":"UID123123","url":"https://example.com","state":"online"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateEngineParams{
			ID:     "1",
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}

		engine, apiError := client.UpdateEngine(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(engine)
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

		params := UpdateEngineParams{
			ID:     "1",
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}

		engine, apiError := client.UpdateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, engine)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"name": opendax_cloud}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateEngineParams{
			ID:     "1",
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}
		engine, apiError := client.UpdateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateEngineParams{
			ID:     "1",
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}
		engine, apiError := client.UpdateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})
}

func TestGetEngines(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":1,"name":"opendax_cloud","driver":"opendax","uid":"UID123123","url":"https://example.com","state":"online"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetEngineParams{
			Name: "opendax_cloud",
		}
		engine, apiError := client.GetEngines(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(engine)
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

		params := GetEngineParams{
			Name: "opendax_cloud",
		}
		engine, apiError := client.GetEngines(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, engine)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"name": opendax_cloud}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetEngineParams{
			Name: "opendax_cloud",
		}
		engine, apiError := client.GetEngines(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"-"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetEngineParams{
			Name: "opendax_cloud",
		}
		engine, apiError := client.GetEngines(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})
}
