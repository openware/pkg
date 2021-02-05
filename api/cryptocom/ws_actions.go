package cryptocom

import (
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

func subscribePublic(c *Client, channels []string) error {
	err := c.subscribePublicChannels(channels)
	if err != nil {
		return err
	}

	c.publicSubs = append(c.publicSubs, channels...)
	return nil
}

func subscribePrivate(c *Client, channels []string) error {
	err := c.subscribePrivateChannels(channels)
	if err != nil {
		return err
	}

	c.privateSubs = append(c.privateSubs, channels...)
	return nil
}

// SubscribePublicTrades is subscription trade channel
// Example: SubscribeTrades("ETH_BTC", "ETH_CRO")
func (c *Client) SubscribePublicTrades(markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("trade.%s", s)
	})

	return subscribePublic(c, channels)
}

// SubscribePublicOrderBook is subscription orderbook channel
// Example: SubscribeOrderBook(depth, "ETH_BTC", "ETH_CRO")
// depth: Number of bids and asks to return. Allowed values: 10 or 150
func (c *Client) SubscribePublicOrderBook(depth int, markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("book.%s.%d", s, depth)
	})

	return subscribePublic(c, channels)
}

// SubscribePublicTickers is subscription ticker channel
func (c *Client) SubscribePublicTickers(markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("ticker.%s", s)
	})

	return subscribePublic(c, channels)
}

// SubscribePrivateOrders is subscription private order user.order.markets channel
func (c *Client) SubscribePrivateOrders(markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("user.order.%s", s)
	})

	return subscribePrivate(c, channels)
}

// SubscribePrivateTrades is subscription private user.trade channel
func (c *Client) SubscribePrivateTrades(markets ...string) error {
	channels := c.format(markets, func(s string) string {
		return fmt.Sprintf("user.trade.%s", s)
	})

	return subscribePrivate(c, channels)
}

func (c *Client) SubscribePrivateBalanceUpdates() error {
	channels := []string{"user.balance"}
	return subscribePrivate(c, channels)
}

// For MARKET BUY orders, amount is notional (https://exchange-docs.crypto.com/#private-create-order).
func (c *Client) CreateLimitOrder(
	reqID int,
	ask string,
	bid string,
	orderSide string,
	price decimal.Decimal,
	amount decimal.Decimal,
	uuid uuid.UUID,
) error {
	r := c.createOrderLimitRequest(
		reqID,
		ask,
		bid,
		orderSide,
		price,
		amount,
		uuid,
	)
	return c.sendPrivateRequest(r)
}

func (c *Client) CreateMarketOrder(
	reqID int,
	ask string,
	bid string,
	orderSide string,
	amount decimal.Decimal,
	uuid uuid.UUID,
) error {
	r := c.createOrderMarketRequest(
		reqID,
		ask,
		bid,
		orderSide,
		amount,
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

func (c *Client) GetOrderDetails(reqID int, remoteID string) error {
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
