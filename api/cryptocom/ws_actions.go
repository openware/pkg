package cryptocom

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type formater func(string) string

// Input: ["ETH_BTC", "ETH_CRO"]
func (c *Client) format(markets []string, fn formater) []string {
	channels := []string{}
	for _, v := range markets {
		channels = append(channels, fn(v))
	}

	return channels
}

// SubscribePublicTrades is subscription trade channel
// Example: SubscribeTrades("ETH_BTC", "ETH_CRO")
func (c *Client) SubscribePublicTrades(markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("trade.%s", s)
	})

	err := c.subscribePublicChannels(channels)
	if err != nil {
		return err
	}

	c.recordPublicSubscription(channels)
	return nil
}

// SubscribePublicOrderBook is subscription orderbook channel
// Example: SubscribeOrderBook(depth, "ETH_BTC", "ETH_CRO")
// depth: Number of bids and asks to return. Allowed values: 10 or 150
func (c *Client) SubscribePublicOrderBook(depth int, markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("book.%s.%d", s, depth)
	})

	err := c.subscribePublicChannels(channels)
	if err != nil {
		return err
	}

	c.recordPublicSubscription(channels)
	return nil
}

// SubscribePublicTickers is subscription ticker channel
func (c *Client) SubscribePublicTickers(markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("ticker.%s", s)
	})

	err := c.subscribePublicChannels(channels)
	if err != nil {
		return err
	}
	c.recordPublicSubscription(channels)
	return nil
}

// SubscribePrivateOrders is subscription private order user.order.markets channel
func (c *Client) SubscribePrivateOrders(markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("user.order.%s", s)
	})

	err := c.subscribePrivateChannels(channels)
	if err != nil {
		return err
	}

	c.recordPrivateSubscription(channels)
	return nil
}

// SubscribePrivateTrades is subscription private user.trade channel
func (c *Client) SubscribePrivateTrades(markets ...string) error {
	// channels := c.format(markets, func(s string) string {
	// 	return fmt.Sprintf("user.trade.%s", s)
	// })

	channels := []string{"user.trade"}

	err := c.subscribePrivateChannels(channels)
	if err != nil {
		return err
	}

	c.recordPrivateSubscription(channels)
	return nil
}

func (c *Client) SubscribePrivateBalanceUpdates() error {
	channel := []string{"user.balance"}

	err := c.subscribePrivateChannels(channel)
	if err != nil {
		return err
	}

	c.recordPrivateSubscription(channel)
	return nil
}

func (c *Client) CreateOrder(
	reqID int,
	ask string,
	bid string,
	orderSide string,
	orderType string,
	price decimal.Decimal,
	volume decimal.Decimal,
	uuid uuid.UUID,
) error {
	r := c.createOrderRequest(
		reqID,
		ask,
		bid,
		orderSide,
		orderType,
		price,
		volume,
		uuid,
	)
	return c.sendPrivateRequest(r)
}

func (c *Client) CancelOrder(reqID int, remoteID, market string) error {
	r := c.cancelOrderRequest(
		reqID,
		remoteID,
		market,
	)
	return c.sendPrivateRequest(r)
}

func (c *Client) CancelAllOrders(reqID int, market string) error {
	r := c.cancelAllOrdersRequest(reqID, market)
	return c.sendPrivateRequest(r)
}

func (c *Client) GetOrderDetails(reqID int, remoteID sql.NullString) error {
	r := c.getOrderDetailsRequest(reqID, remoteID)
	return c.sendPrivateRequest(r)
}

func (c *Client) respondHeartBeat(isPrivate bool, id int) {
	r := c.hearBeatRequest(id)

	if isPrivate {
		c.sendPrivateRequest(r)
	} else {
		c.sendPublicRequest(r)
	}
}
