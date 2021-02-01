package peatio

import (
	"testing"

	"github.com/openware/pkg/mngapi"
	"github.com/stretchr/testify/assert"
)

type MockPeatioClient struct{}

func (m *MockPeatioClient) GetCurrencyByCode(code string) (*Currency, *mngapi.APIError) {
	return &Currency{
		ID:   "bnb",
		Name: "Binance Coin",
	}, nil
}

func TestGetCurrencyByCode(t *testing.T) {
	client := &MockPeatioClient{}
	currency, apierror := client.GetCurrencyByCode("bnb")

	assert.Equal(t, apierror, (*mngapi.APIError)(nil))
	assert.Equal(t, currency.ID, "bnb")
	assert.Equal(t, currency.Name, "Binance Coin")
}
