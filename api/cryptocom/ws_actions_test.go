package cryptocom

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type mockRequest struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Nonce  string `json:"nonce"`
	Params map[string]interface{}
}

type testingFunc func(client *Client)

func TestFormat(t *testing.T) {
	client := New("", "", "test", "test")

	markets := []string{"ETH_BTC", "ETH_COV", "XRP_BTC"}
	expected := []string{"trade.ETH_BTC", "trade.ETH_COV", "trade.XRP_BTC"}

	result := client.format(markets, func(s string) string {
		return fmt.Sprintf("trade.%s", s)
	})

	assert.Equal(t, result, expected)
}

func testPublicSubscribe(t *testing.T, expected string, testFunc testingFunc) {
	// prepare expected
	var expectedResponse mockRequest
	err := json.Unmarshal([]byte(expected), &expectedResponse)
	if err != nil {
		t.Fatal("error on parse expected")
	}

	// prepare mock
	client := New("test", "test", "test", "test")
	publiceWritingMessage := bytes.NewBuffer(nil)
	client.connectMock(bytes.NewBuffer(nil), bytes.NewBuffer(nil), bytes.NewBuffer(nil), publiceWritingMessage)

	// start test
	testFunc(client)

	var writingMessage mockRequest
	err = json.Unmarshal(publiceWritingMessage.Bytes(), &writingMessage)
	if err != nil {
		t.Fatal("error on parse writing message")
	}

	assert.Equal(t, expectedResponse.ID, writingMessage.ID)
	assert.Equal(t, expectedResponse.Method, writingMessage.Method)
	assert.Equal(t, expectedResponse.Params, writingMessage.Params)
}

func testPrivateSubscribe(t *testing.T, expected string, testFunc testingFunc) {
	// prepare expected
	var expectedResponse mockRequest
	err := json.Unmarshal([]byte(expected), &expectedResponse)
	if err != nil {
		t.Fatal("error on parse expected")
	}

	// prepare mock
	client := New("test", "test", "test", "test")
	privateWritingMessage := bytes.NewBuffer(nil)
	client.connectMock(bytes.NewBuffer(nil), bytes.NewBuffer(nil), privateWritingMessage, bytes.NewBuffer(nil))

	// start test
	testFunc(client)

	var writingMessage mockRequest
	err = json.Unmarshal(privateWritingMessage.Bytes(), &writingMessage)
	if err != nil {
		t.Fatal("error on parse writing message")
	}

	assert.Equal(t, expectedResponse.ID, writingMessage.ID)
	assert.Equal(t, expectedResponse.Method, writingMessage.Method)
	assert.Equal(t, expectedResponse.Params, writingMessage.Params)
}

func TestPublicOrderBook(t *testing.T) {
	expected := `{"id":1,"method":"subscribe","nonce":"","params":{"channels":["book.ETH_BTC.10"]}}`
	testPublicSubscribe(t, expected, func(client *Client) { client.SubscribePublicOrderBook(10, "ETH_BTC") })
}

func TestPublicTrades(t *testing.T) {
	// prepare expected
	expected := `{"id":1,"method":"subscribe","nonce":"","params":{"channels":["trade.ETH_BTC"]}}`
	testPublicSubscribe(t, expected, func(client *Client) { client.SubscribePublicTrades("ETH_BTC") })
}

func TestPublicTickers(t *testing.T) {
	// prepare expected
	expected := `{"id":1,"method":"subscribe","nonce":"","params":{"channels":["ticker.ETH_BTC"]}}`
	testPublicSubscribe(t, expected, func(client *Client) { client.SubscribePublicTickers("ETH_BTC") })
}

func TestSubscribePrivateOrders(t *testing.T) {
	// prepare expected
	expected := `{"id":1,"method":"subscribe","nonce":"","params":{"channels":["user.order.ETH_BTC"]}}`
	testPrivateSubscribe(t, expected, func(client *Client) { client.SubscribePrivateOrders("ETH_BTC") })
}

func TestSubscribePrivateTrades(t *testing.T) {
	// prepare expected
	expected := `{"id":1,"method":"subscribe","nonce":"","params":{"channels":["user.trade"]}}`
	testPrivateSubscribe(t, expected, func(client *Client) { client.SubscribePrivateTrades("ETH_BTC") })
}

func TestSubscribePrivateBalanceUpdates(t *testing.T) {
	// prepare expected
	expected := `{"id":1,"method":"subscribe","nonce":"","params":{"channels":["user.balance"]}}`
	testPrivateSubscribe(t, expected, func(client *Client) { client.SubscribePrivateBalanceUpdates() })
}

func TestCreateOrder(t *testing.T) {
	// prepare expected
	uuid := uuid.New()
	price := decimal.NewFromFloat(0.01)
	volume := decimal.NewFromFloat(0.0001)

	expected := fmt.Sprintf(
		`{"id":1,"method":"private/create-order","nonce":"","params":{"client_oid":"%s","instrument_name":"ETH_CRO","price":"%s","quantity":"%s","side":"%s","type":"LIMIT"}}`,
		uuid, price.String(), volume.String(), "SELL",
	)
	testPrivateSubscribe(t, expected, func(client *Client) {
		client.CreateOrder(
			1,
			"ETH",
			"CRO",
			"sell",
			"LIMIT",
			decimal.NewFromFloat(0.01),
			decimal.NewFromFloat(0.0001),
			uuid,
		)
	})
}

func TestCancelOrder(t *testing.T) {
	remoteID := sql.NullString{String: "1138210129647637539", Valid: true}

	// prepare expected
	expected := fmt.Sprintf(
		`{"id":1,"method":"private/cancel-order","nonce":"","params":{"instrument_name":"ETH_CRO","order_id":"%s"}}`,
		remoteID.String,
	)
	testPrivateSubscribe(t, expected, func(client *Client) {
		client.CancelOrder(
			1,
			remoteID.String,
			"ETH_CRO",
		)
	})
}

func TestCancelAllOrders(t *testing.T) {
	// prepare expected
	expected := `{"id":1,"method":"private/cancel-all-orders","nonce":"","params":{"instrument_name":"ETH_CRO"}}`
	testPrivateSubscribe(t, expected, func(client *Client) { client.CancelAllOrders(1, "ETH_CRO") })
}

func TestGetOrderDetails(t *testing.T) {
	// prepare expected
	remoteID := sql.NullString{String: "1138210129647637539", Valid: true}
	expected := `{"id":1,"method":"private/get-order-detail","nonce":"","params":{"order_id":"1138210129647637539"}}`
	testPrivateSubscribe(t, expected, func(client *Client) { client.GetOrderDetails(1, remoteID) })
}

func TestRespondHeartBeat(t *testing.T) {
	// prepare mock
	client := New("test", "test", "test", "test")
	privateWritingMessage := bytes.NewBuffer(nil)
	publicWritingMessage := bytes.NewBuffer(nil)
	client.connectMock(bytes.NewBuffer(nil), bytes.NewBuffer(nil), privateWritingMessage, publicWritingMessage)

	t.Run("private", func(t *testing.T) {
		var writingMessage mockRequest
		var expectedResponse mockRequest
		expected := `{"id":1,"method":"public/respond-heartbeat"}`

		// start test
		client.respondHeartBeat(true, 1)
		json.Unmarshal(privateWritingMessage.Bytes(), &writingMessage)
		// prepare expected
		json.Unmarshal([]byte(expected), &expectedResponse)

		assert.Equal(t, expectedResponse.ID, writingMessage.ID)
		assert.Equal(t, expectedResponse.Method, writingMessage.Method)
		assert.Equal(t, expectedResponse.Params, writingMessage.Params)
	})

	t.Run("public", func(t *testing.T) {
		var writingMessage mockRequest
		var expectedResponse mockRequest
		expected := `{"id":1,"method":"public/respond-heartbeat"}`

		// start test
		client.respondHeartBeat(false, 1)
		json.Unmarshal(publicWritingMessage.Bytes(), &writingMessage)
		// prepare expected
		json.Unmarshal([]byte(expected), &expectedResponse)

		assert.Equal(t, expectedResponse.ID, writingMessage.ID)
		assert.Equal(t, expectedResponse.Method, writingMessage.Method)
		assert.Equal(t, expectedResponse.Params, writingMessage.Params)
	})
}
