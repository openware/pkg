package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestCreateMember(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"uid":"UID123123","email":"user@barong.io","level":3,"role":"sa_maker","group":"sat-zero","state":"active"}`

		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateMemberParams{
			Email: "user@barong.io",
			UID:   "UID123123",
			Level: 3,
			Role:  "sa_maker",
			State: "active",
			Group: "sat-zero",
		}
		member, apiError := client.CreateMember(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(member)
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
		params := CreateMemberParams{
			Email: "user@barong.io",
			UID:   "UID123123",
			Level: 3,
			Role:  "sa_maker",
			State: "active",
			Group: "sat-zero",
		}
		member, apiError := client.CreateMember(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, member)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"level": opendax_cloud}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateMemberParams{
			Email: "user@barong.io",
			UID:   "UID123123",
			Level: 3,
			Role:  "sa_maker",
			State: "active",
			Group: "sat-zero",
		}
		member, apiError := client.CreateMember(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, member)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateMemberParams{
			Email: "user@barong.io",
			UID:   "UID123123",
			Level: 3,
			Role:  "sa_maker",
			State: "active",
			Group: "sat-zero",
		}
		member, apiError := client.CreateMember(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, member)
	})
}
