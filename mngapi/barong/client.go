package barong

import (
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
