package test

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	. "github.com/openware/pkg/mngapi/peatio"
	"github.com/stretchr/testify/assert"
)

func TestGetMarkets(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":"btcusd","name":"BTC/USD","base_unit":"btc","quote_unit":"usd","min_price":"0.01","max_price":"0.0","min_amount":"0.00000001","amount_precision":8,"price_precision":2,"state":"enabled","position":1,"engine_id":1,"created_at":"2021-03-05T14:52:48+01:00","updated_at":"2021-03-05T14:52:48+01:00"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		markets, apiError := client.GetMarkets()
		assert.Nil(t, apiError)

		result, err := json.Marshal(markets)
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

		markets, apiError := client.GetMarkets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, markets)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"name": BTC/USD}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		markets, apiError := client.GetMarkets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, markets)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"-"}]`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		markets, apiError := client.GetMarkets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, markets)
	})
}

func TestUpdateMarket(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"btcusd","name":"BTC/USD","base_unit":"btc","quote_unit":"usd","min_price":"0.01","max_price":"0.0","min_amount":"0.00000001","amount_precision":8,"price_precision":2,"state":"enabled","position":1,"engine_id":1,"created_at":"2021-03-05T14:52:48+01:00","updated_at":"2021-03-05T14:52:48+01:00"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateMarketParams{
			ID:       "1",
			EngineID: "1",
		}

		market, apiError := client.UpdateMarket(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(market)
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

		params := UpdateMarketParams{
			ID:       "1",
			EngineID: "1",
		}

		market, apiError := client.UpdateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, market)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"name": BTC/USD}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateMarketParams{
			ID:       "1",
			EngineID: "1",
		}
		market, apiError := client.UpdateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateMarketParams{
			ID:       "1",
			EngineID: "1",
		}
		market, apiError := client.UpdateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})
}

func TestCreateMarket(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"btcusd","name":"BTC/USD","base_unit":"btc","quote_unit":"usd","min_price":"0.01","max_price":"0.0","min_amount":"0.00000001","amount_precision":8,"price_precision":2,"state":"enabled","position":1,"engine_id":1,"created_at":"2021-03-05T14:52:48+01:00","updated_at":"2021-03-05T14:52:48+01:00"}`

		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateMarketParams{
			BaseCurrency:    "btc",
			QuoteCurrency:   "usd",
			State:           "disabled",
			EngineName:      "opendax_cloud",
			AmountPrecision: 2,
			PricePrecision:  6,
			MinPrice:        "0.2",
			MaxPrice:        "1.0",
			MinAmount:       "0.1",
			Position:        1,
		}
		market, apiError := client.CreateMarket(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(market)
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
		params := CreateMarketParams{
			BaseCurrency:    "btc",
			QuoteCurrency:   "usd",
			State:           "disabled",
			EngineName:      "opendax_cloud",
			AmountPrecision: 2,
			PricePrecision:  6,
			MinPrice:        "0.2",
			MaxPrice:        "1.0",
			MinAmount:       "0.1",
			Position:        1,
		}
		market, apiError := client.CreateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, market)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"name": opendax_cloud}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateMarketParams{
			BaseCurrency:    "btc",
			QuoteCurrency:   "usd",
			State:           "disabled",
			EngineName:      "opendax_cloud",
			AmountPrecision: 2,
			PricePrecision:  6,
			MinPrice:        "0.2",
			MaxPrice:        "1.0",
			MinAmount:       "0.1",
			Position:        1,
		}
		market, apiError := client.CreateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateMarketParams{
			BaseCurrency:    "btc",
			QuoteCurrency:   "usd",
			State:           "disabled",
			EngineName:      "opendax_cloud",
			AmountPrecision: 2,
			PricePrecision:  6,
			MinPrice:        "0.2",
			MaxPrice:        "1.0",
			MinAmount:       "0.1",
			Position:        1,
		}
		market, apiError := client.CreateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})
}

func TestGetMarketByID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"btcusd","name":"BTC/USD","base_unit":"btc","quote_unit":"usd","min_price":"0.01","max_price":"0.0","min_amount":"0.00000001","amount_precision":8,"price_precision":2,"state":"enabled","position":1,"engine_id":1,"created_at":"2021-03-05T14:52:48+01:00","updated_at":"2021-03-05T14:52:48+01:00"}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		market, apiError := client.GetMarketByID("btcusd")
		assert.Nil(t, apiError)

		result, err := json.Marshal(market)
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
				Error:      "Invalid",
			},
		}

		market, apiError := client.GetMarketByID("btcusd")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "Invalid")
		assert.Nil(t, market)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":btcusd,min_price:123.456}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		market, apiError := client.GetMarketByID("btcusd")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{""}`
		client.MngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		market, apiError := client.GetMarketByID("btcusd")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})
}
