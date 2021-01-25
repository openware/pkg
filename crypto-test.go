package main

import (
	"fmt"

	"github.com/openware/pkg/cryptocom"
)

func main() {
	client := cryptocom.New("test", "test")

	err := client.Connect("wss://uat-stream.3ona.co/v2/market")
	if err != nil {
		fmt.Println(err)
	}

	// client.SubscribeTrades("ETH_BTC", "ETH_CRO")
	// client.SubscribeTickers("ETH_BTC", "ETH_CRO")
	client.SubscribeOrderBook(10, "ETH_BTC", "ETH_CRO")
	msgs := client.Listen()

	for m := range msgs {
		fmt.Println(m)
	}

	// r := client.AuthRequest()
	// b, _ := r.Encode()

	// var res interface{}
	// err := json.Unmarshal(b, &res)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// m := res.(map[string]interface{})

	// fmt.Println(m)
}
