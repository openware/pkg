package barong

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/openware/pkg/mngapi"
)

// Client is barong management api client instance
type Client struct {
	mngapiClient mngapi.DefaultClient
}

// New return barong management api client
func New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey string) (*Client, error) {
	client, err := mngapi.New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	return &Client{
		mngapiClient: client,
	}, nil
}

// CreateServiceAccount call barong management api to create new service account
func (b *Client) CreateServiceAccount(params CreateServiceAccountParams) (*ServiceAccount, *mngapi.APIError) {
	res, apiError := b.mngapiClient.Request(http.MethodPost, "service_accounts/create", params)
	if apiError != nil {
		return nil, apiError
	}

	serviceAccount := &ServiceAccount{}
	err := json.Unmarshal([]byte(res), serviceAccount)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return serviceAccount, nil
}

// CreateAPIKey calls Barong Management Api to create a new API key for a given
func (b *Client) CreateAPIKey(params CreateAPIKeyParams) (*APIKey, *mngapi.APIError) {
	res, apiError := b.mngapiClient.Request(http.MethodPost, "api_keys", params)
	if apiError != nil {
		return nil, apiError
	}

	apiKey := &APIKey{}
	err := json.Unmarshal([]byte(res), apiKey)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return apiKey, nil
}

// DeleteServiceAccountByUID call barong management api to delete service account by uid
func (b *Client) DeleteServiceAccountByUID(uid string) (*ServiceAccount, *mngapi.APIError) {
	// Build parameters
	params := map[string]interface{}{
		"uid": uid,
	}

	res, apiError := b.mngapiClient.Request(http.MethodPost, "service_accounts/delete", params)
	if apiError != nil {
		return nil, apiError
	}

	serviceAccount := &ServiceAccount{}
	err := json.Unmarshal([]byte(res), serviceAccount)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	return serviceAccount, nil
}

// CreateAPIKey calls Barong Management Api to create a new API key for a given
func (b *Client) CreateAttachment(params CreateAttachmentParams) (*Attachment, *mngapi.APIError) {
	res, apiError := b.mngapiClient.Request(http.MethodPost, "attachments", params)
	if apiError != nil {
		return nil, apiError
	}

	attachment := &Attachment{}
	err := json.Unmarshal([]byte(res), attachment)
	if err != nil {
		return nil, &mngapi.APIError{StatusCode: 500, Error: err.Error()}
	}

	fmt.Println(attachment)
	return attachment, nil
}
