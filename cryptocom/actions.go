package cryptocom

import "fmt"

type formater func(string) string

// Input: ["ETH_BTC", "ETH_CRO"]
func (c *Client) format(markets []string, fn formater) []string {
	channels := []string{}
	for _, v := range markets {
		channels = append(channels, fn(v))
	}

	return channels
}

// SubscribeTrades is subscription trade channel
// Example: SubscribeTrades("ETH_BTC", "ETH_CRO")
func (c *Client) SubscribePublicTrades(markets ...string) {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("trade.%s", s)
	})

	c.subscribePublicChannels(channels)
}

// SubscribeOrderBook is subscription orderbook channel
// Example: SubscribeOrderBook(depth, "ETH_BTC", "ETH_CRO")
// depth: Number of bids and asks to return. Allowed values: 10 or 150
func (c *Client) SubscribePublicOrderBook(depth int, markets ...string) {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("book.%s.%d", s, depth)
	})

	c.subscribePublicChannels(channels)
}

// SubscribeTickers is subscription ticker channel
func (c *Client) SubscribePublicTickers(markets ...string) {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("ticker.%s", s)
	})

	c.subscribePublicChannels(channels)
}

func (c *Client) SubscribePrivateOrders(markets ...string) {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("user.order.%s", s)
	})
	c.subscribePrivateChannels(channels)
}

func (c *Client) SubscribePrivateTrades(markets ...string) {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("user.trade.%s", s)
	})
	c.subscribePrivateChannels(channels)
}

func (c *Client) SubscribePrivateBalanceUpdates() {
	channel := []string{"user.balance"}
	c.subscribePrivateChannels(channel)
}

func (c *Client) respondHeartBeat(scope string, id int) {
	r := c.hearBeatRequest(id)

	switch scope {
	case "private":
		c.sendPrivateRequest(r)
	case "public":
		c.sendPublicRequest(r)
	}
}
