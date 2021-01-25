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
	publicConn  *websocket.Conn
	privateConn *websocket.Conn

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

// Connect("wss://uat-stream.3ona.co")
func (c *Client) Connect(url string) error {
	publicEndpoint := url + "/v2/market"
	privateEndpoint := url + "/v2/user"
	conn, _, err := websocket.DefaultDialer.Dial(publicEndpoint, http.Header{})
	if err != nil {
		return err
	}

	c.publicConn = conn

	conn, _, err = websocket.DefaultDialer.Dial(privateEndpoint, http.Header{})
	if err != nil {
		return err
	}

	c.privateConn = conn

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
			_, m, err := c.publicConn.ReadMessage()
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
				c.respondHeartBeat("public", parsed.Id)
				continue
			}

			c.msgs <- parsed
		}
	}()

	go func() {
		defer func() {
			close(c.done)
			close(c.msgs)
		}()

		for {
			_, m, err := c.privateConn.ReadMessage()
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
				c.respondHeartBeat("private", parsed.Id)
				continue
			}

			c.msgs <- parsed
		}
	}()

	return c.msgs
}

func (c *Client) authenticate() {
	r := c.AuthRequest()
	c.sendPrivateRequest(r)
}

func (c *Client) sendPrivateRequest(r *Request) error {
	b, err := r.Encode()

	if err != nil {
		return err
	}
	return c.privateConn.WriteMessage(websocket.TextMessage, b)
}

func (c *Client) sendPublicRequest(r *Request) error {
	b, err := r.Encode()

	if err != nil {
		return err
	}
	return c.publicConn.WriteMessage(websocket.TextMessage, b)
}

func (c *Client) subscribePrivateChannels(channels []string) {
	r := c.subscribeRequest(channels)
	c.sendPrivateRequest(r)
}

func (c *Client) subscribePublicChannels(channels []string) {
	r := c.subscribeRequest(channels)
	c.sendPublicRequest(r)
}
