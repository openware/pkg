package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/openware/openfinex/pkg/log"
	"github.com/openware/pkg/api/cryptocom"
	"github.com/shopspring/decimal"
)

// 1138210129647637539
func main() {

	// testWs()
	// testGetOrderInfo(sql.NullString{String: "1160202208225505696", Valid: true})
	// createOrder()
	// cancelOrder()
	// testRest()
	// testMarkets()
	// testMarketTrades()
	// testSubscribeUserOrders()
	testRestOpenOrders()
}

// func testSubscribeUserOrders() {
// 	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
// 	// bidOrder := order.Model{Ask: "BTC", Bid: "ETH", Side: order.Sell, Type: "LIMIT", Price: decimal.NewFromFloat(0.5), Volume: decimal.NewFromFloat(3), UUID: uuid.New()}
// 	// askOrder := order.Model{Ask: "ETH", Bid: "CRO", Side: order.Buy, Type: "LIMIT", Price: decimal.NewFromFloat(0.5), Volume: decimal.NewFromFloat(0.3), UUID: uuid.New()}
// 	client.Connect()
// 	msgs := client.Listen()

// 	time.Sleep(2 * time.Second)
// 	client.SubscribePublicOrderBook(10, "CRO_BTC")
// 	// client.SubscribePrivateOrders("ETH_CRO", "CRO_BTC", "ETH_CRO")
// 	// client.SubscribePrivateTrades("ETH_CRO", "CRO_BTC", "ETH_CRO")
// 	// client.CreateOrder(1, &bidOrder)
// 	// client.CreateOrder(1, &askOrder)

// 	for {
// 		m := <-msgs
// 		fmt.Println("")
// 		fmt.Println("INFO: New message received")
// 		formatted, _ := json.MarshalIndent(m, "", "  ")
// 		fmt.Println(string(formatted))

// 	}
// }
// func testMarketTrades() {
// 	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
// 	resp, err := client.RestGetTrades(1, "CRO_BTC")
// 	if err != nil {
// 		fmt.Println(err)
// 		return

// 	}
// 	formatted, _ := json.MarshalIndent(resp, "", "  ")
// 	fmt.Println(string(formatted))
// }

// func testRest() {
// 	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
// 	order := order.Model{Ask: "ETH", Bid: "CRO", Side: order.Buy, Type: "LIMIT", Price: decimal.NewFromFloat(0.1), Volume: decimal.NewFromFloat(0.00001), UUID: uuid.New(), RemoteID: sql.NullString{String: "1140426162903344128", Valid: true}}
// 	resp, err := client.RestGetOrderDetails(1, &order)
// 	if err != nil {
// 		fmt.Println(err)
// 		return

// 	}
// 	formatted, _ := json.MarshalIndent(resp, "", "  ")

// 	// rawOrdInfo := resp.Result["order_info"].(map[string]interface{})
// 	// remoteID := rawOrdInfo["order_id"].(string)
// 	// amount := decimal.NewFromFloat(rawOrdInfo["quantity"].(float64))
// 	// price := decimal.NewFromFloat(rawOrdInfo["price"].(float64))

// 	// ordInfo := engine.OrderInfo{
// 	// 	RemoteID: remoteID,
// 	// 	Price:    price,
// 	// 	Amount:   amount,
// 	// }

// 	// fmt.Printf("OrderInfo: remoteId - %s, price - %s, amount - %s", ordInfo.RemoteID, ordInfo.Price, ordInfo.Amount)
// 	fmt.Println(string(formatted))
// }

func testRestOpenOrders() {
	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
	resp, err := client.RestOpenOrders(1, "BTC_USDT", 0, 200)
	if err != nil {
		fmt.Println(err)
		return

	}
	formatted, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(formatted))
}

func createOrder() {
	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
	// order1 := order.Model{Ask: "ETH", Bid: "CRO", Side: order.Buy, Type: "LIMIT", Price: decimal.NewFromFloat(30000), Volume: decimal.NewFromFloat(1), UUID: uuid.New()}
	// order2 := order.Model{Ask: "CRO", Bid: "BTC", Side: order.Buy, Type: "LIMIT", Price: decimal.NewFromFloat(0.0003), Volume: decimal.NewFromFloat(3), UUID: uuid.New()}
	// order3 := order.Model{Ask: "CRO", Bid: "BTC", Side: order.Buy, Type: "LIMIT", Price: decimal.NewFromFloat(0.0019), Volume: decimal.NewFromFloat(1), UUID: uuid.New()}
	// cnx := engine.NewCryptcomCnx("fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
	// cnx.Start()
	// cnx.SubmitOrder(&order)

	client.Connect()
	time.Sleep(3 * time.Second)
	client.CreateLimitOrder(1, "BTC", "USDT", "sell", decimal.NewFromFloat(50000), decimal.NewFromFloat(0.001), uuid.New())
	// client.CreateMarketOrder(1, "BTC", "USDT", "buy", decimal.NewFromFloat(100), uuid.New())
	// client.CreateOrder(1, &order2)
	msgs := client.Listen()

	// client.CreateOrder(1, &order3)

	// client.SubscribePublicOrderBook(10, "CRO_BTC")
	client.SubscribePrivateTrades("BTC_USDT")

	for {
		m := <-msgs
		fmt.Println("")
		fmt.Println("INFO: New message received")
		formatted, _ := json.MarshalIndent(m, "", "  ")
		fmt.Println(string(formatted))
	}
}

func cancelOrder() {
	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
	// order := order.Model{Ask: "ETH", Bid: "CRO", Side: order.Buy, Type: "LIMIT", Price: decimal.NewFromFloat(0.01), Volume: decimal.NewFromFloat(0.0001), UUID: uuid.New(), RemoteID: sql.NullString{String: "1155115558389273696", Valid: true}}
	// cnx := engine.NewCryptcomCnx("fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
	// cnx.Start()
	// cnx.SubmitOrder(&order)

	client.Connect()
	time.Sleep(3 * time.Second)
	msgs := client.Listen()
	// client.SubscribePrivateOrders("ETH_CRO")
	client.CancelOrder(1, "1160206888191511040", "BTC_USDT")

	for {
		m := <-msgs
		fmt.Println("")
		fmt.Println("INFO: New message received")
		formatted, _ := json.MarshalIndent(m, "", "  ")
		fmt.Println(string(formatted))
	}
}

func testWs() {
	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
	// order := order.Model{Ask: "ETH", Bid: "CRO", Side: order.Sell, Type: "LIMIT", Price: decimal.NewFromFloat(0.01), Volume: decimal.NewFromFloat(0.0001), UUID: uuid.New(), RemoteID: sql.NullString{String: "1138210129647637539", Valid: true}}
	err := client.Connect()
	if err != nil {
		log.Debug("Startup failed")
		return

	}
	msgs := client.Listen()

	client.SubscribePublicOrderBook(150, "BTC_USDT")
	// client.SubscribePublicTrades("ETH_CRO")
	// client.SubscribePrivateOrders("ETH_CRO")
	// client.CancelOrder(1, &order)

	for {
		m := <-msgs
		fmt.Println("")
		fmt.Println("INFO: New message received")
		formatted, _ := json.MarshalIndent(m, "", "  ")
		fmt.Println(string(formatted))
	}
}

func testMarkets() {
	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
	resp, err := client.RestGetBalance(1)
	if err != nil {
		fmt.Println(err)
		return

	}
	formatted, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(formatted))
}

func testGetOrderInfo(orderId sql.NullString) {
	client := cryptocom.New("wss://uat-stream.3ona.co", "https://uat-api.3ona.co", "fFv2PB8TF7QfRBkmnkvSPr", "oiN9DGDeqEMivgGry28Sm4")
	// order := order.Model{Ask: "ETH", Bid: "CRO", Side: order.Sell, Type: "LIMIT", Price: decimal.NewFromFloat(0.01), Volume: decimal.NewFromFloat(0.0001), UUID: uuid.New(), RemoteID: sql.NullString{String: "1155100964112886592", Valid: true}}
	resp, err := client.RestGetOrderDetails(1, orderId)
	if err != nil {
		fmt.Println(err)
		return

	}
	formatted, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(formatted))
}
