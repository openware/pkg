package cryptocom

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fasthttp/websocket"
)

type Client struct {
	conn *websocket.Conn

	key    string
	secret string

	done chan struct{}
	msgs chan interface{}
}

func New(key, secret string) *Client {
	return &Client{
		key:    key,
		secret: secret,
		done:   make(chan struct{}),
		msgs:   make(chan interface{}),
	}
}

func (c *Client) generateSignature(r *Request) {
	secret := c.secret
	data := r.Method + strconv.Itoa(r.Id) + r.ApiKey + r.Nonce

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	r.Signature = sha
}

func (c *Client) Connect(url string) error {
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		return err
	}

	c.conn = conn

	// c.authenticate()

	return nil
}

func (c *Client) Listen() <-chan interface{} {
	go func() {
		defer func() {
			close(c.done)
			close(c.msgs)
		}()

		for {
			_, m, err := c.conn.ReadMessage()
			if err != nil {
				fmt.Println("error on read message")
				return
			}

			var parsed Response
			err = json.Unmarshal(m, &parsed)
			if err != nil {
				fmt.Println("error on parse message")
				continue
			}

			if parsed.Method == "public/heartbeat" {
				c.respondHeartBeat(parsed.Id)
				continue
			}

			c.msgs <- parsed
		}
	}()

	return c.msgs
}

func (c *Client) authenticate() {
	r := c.AuthRequest()
	c.sendRequest(r)
}

// SubscribeTrades is subscription trade channel
// Example: SubscribeTrades("ETH_BTC", "ETH_CRO")
func (c *Client) SubscribeTrades(markets ...string) {
	channels := c.format(markets, func (s string) string {
		return fmt.Sprintf("trade.%s", s)
	})

	c.SubscribeChannel(channels)
}

// SubscribeOrderBook is subscription orderbook channel
// Example: SubscribeOrderBook(depth, "ETH_BTC", "ETH_CRO")
func (c *Client) SubscribeOrderBook(depth int, markets ...string) {
	channels := c.format(markets, func (s string) string {
		return fmt.Sprintf("book.%s.%d", s, depth)
	})

	c.SubscribeChannel(channels)
}

// SubscribeTickers is subscription ticker channel
// Example: SubscribeTickers(depth, "ETH_BTC", "ETH_CRO")
// depth: Number of bids and asks to return. Allowed values: 10 or 150
func (c *Client) SubscribeTickers(markets ...string) {
	channels := c.format(markets, func (s string) string {
		return fmt.Sprintf("ticker.%s", s)
	})

	c.SubscribeChannel(channels)
}

// SubscribeChannel is subscription channels by parameter
func (c *Client) SubscribeChannel(channels []string) {
	r := c.subscribeRequest(channels)
	c.sendRequest(r)
}

func (c *Client) respondHeartBeat(id int) {
	r := c.hearBeatRequest(id)
	c.sendRequest(r)
}

func (c *Client) sendRequest(r *Request) error {
	b, err := r.Encode()

	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.TextMessage, b)
}

type formater func(string) string
// Input: ["ETH_BTC", "ETH_CRO"]
func (c *Client) format(markets []string, fn formater) []string {
	channels := []string{}
	for _, v := range markets {
		channels = append(channels, fn(v))
	}

	return channels
}
