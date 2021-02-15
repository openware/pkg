package barong

import (
	"encoding/json"
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
	res, err := b.mngapiClient.Request(http.MethodPost, "service_accounts/create", params)
	if err != nil {
		return nil, err
	}

	serviceAccount := &ServiceAccount{}
	_ = json.Unmarshal([]byte(res), serviceAccount)

	return serviceAccount, nil
}

// DeleteServiceAccountByUID call barong management api to delete service account by uid
func (b *Client) DeleteServiceAccountByUID(uid string) (*ServiceAccount, *mngapi.APIError) {
	// Build parameters
	params := map[string]interface{}{
		"uid": uid,
	}

	res, err := b.mngapiClient.Request(http.MethodPost, "service_accounts/delete", params)
	if err != nil {
		return nil, err
	}

	serviceAccount := &ServiceAccount{}
	_ = json.Unmarshal([]byte(res), serviceAccount)

	return serviceAccount, nil
}
