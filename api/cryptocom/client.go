package cryptocom

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const (
	userEndpoint   = "/v2/user"
	marketEndpoint = "/v2/market"
)

type Transport interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
	Close() error
}

type HTTPClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
}

type Connection struct {
	Endpoint  string
	IsPrivate bool
	Transport
}

type Client struct {
	publicConn  Connection
	privateConn Connection
	wsRootURL   string
	restRootURL string
	key         string
	secret      string
	privateSubs []string
	publicSubs  []string
	httpClient  HTTPClient
	outbox      chan Response
}

// New returns a pointer of Client struct
func New(wsRootURL, restRootURL, key, secret string) *Client {
	return &Client{
		key:         key,
		secret:      secret,
		wsRootURL:   wsRootURL,
		restRootURL: restRootURL,
		outbox:      make(chan Response),
		privateSubs: make([]string, 0),
		publicSubs:  make([]string, 0),
		httpClient:  &http.Client{},
	}
}

// Connect instanciate WS Connections
func (c *Client) Connect() error {
	publicWsEndpoint := c.wsRootURL + marketEndpoint
	privateWsEndpoint := c.wsRootURL + userEndpoint

	err := c.createConnection(publicWsEndpoint, false)
	if err != nil {
		return err
	}

	err = c.createConnection(privateWsEndpoint, true)
	if err != nil {
		return err
	}

	time.Sleep(3 * time.Second) // Cryptocom requires this sleep.
	c.authenticate()

	return nil
}

func (c *Client) Listen() <-chan Response {
	go c.readConnection(c.publicConn)
	go c.readConnection(c.privateConn)
	return c.outbox
}

func (c *Client) Shutdown() {
	c.privateConn.Close()
	c.publicConn.Close()
	//FIXME: Add wait group to wait for go routine finish
	close(c.outbox)
}

func (c *Client) createConnection(endpoint string, isPrivate bool) error {
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, http.Header{})
	if err != nil {
		return err
	}

	cnx := Connection{Endpoint: endpoint, IsPrivate: isPrivate, Transport: conn}

	if isPrivate {
		c.privateConn = cnx
	} else {
		c.publicConn = cnx
	}

	return nil
}

func (c *Client) readConnection(cnx Connection) {
	fmt.Println("Start listening connection ...", cnx.Endpoint)
	for {
		_, m, err := cnx.ReadMessage()
		if err != nil {
			fmt.Printf("error on read message. Private cnx - %t\n", cnx.IsPrivate)
			for {
				conn, _, err := websocket.DefaultDialer.Dial(cnx.Endpoint, http.Header{})
				if err != nil {
					continue
				}

				time.Sleep(3 * time.Second) // Cryptocom requires this sleep
				newCnx := Connection{Endpoint: cnx.Endpoint, IsPrivate: cnx.IsPrivate, Transport: conn}

				if newCnx.IsPrivate {
					c.privateConn = newCnx
					c.authenticate()
					if len(c.privateSubs) > 0 {
						c.subscribePrivateChannels(c.privateSubs)
					}
				} else {
					c.publicConn = newCnx
					if len(c.publicSubs) > 0 {
						c.subscribePublicChannels(c.publicSubs)
					}
				}

				cnx.Close()
				cnx = newCnx
				break
			}

			continue
		}

		// fmt.Printf("Received: Private Cnx - %t, msg: %s\n", cnx.IsPrivate, string(m))

		var parsed Response
		err = json.Unmarshal(m, &parsed)
		if err != nil {
			fmt.Println("error on parse message")
			continue
		}

		if parsed.Method == "public/heartbeat" {
			c.respondHeartBeat(cnx.IsPrivate, parsed.Id)
			continue
		}

		c.outbox <- parsed
	}
}

func (c *Client) generateSignature(r *Request) {
	secret := c.secret
	params := ""

	// FIXME: Order by alphabet
	for k, v := range r.Params {
		params += k + v.(string)
	}

	data := r.Method + strconv.Itoa(r.Id) + r.ApiKey + string(params) + r.Nonce
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	r.Signature = sha
}

func (c *Client) authenticate() {
	r := c.AuthRequest()
	c.sendPrivateRequest(r)
}

func (c *Client) subscribePrivateChannels(channels []string) error {
	r := c.subscribeRequest(channels)
	return c.sendPrivateRequest(r)
}

func (c *Client) subscribePublicChannels(channels []string) error {
	r := c.subscribeRequest(channels)
	return c.sendPublicRequest(r)
}

func (c *Client) sendPrivateRequest(r *Request) error {
	b, err := r.Encode()

	if err != nil {
		return err
	}
	// fmt.Printf("Sending private: %s\n", string(b))
	return c.privateConn.WriteMessage(websocket.TextMessage, b)
}

func (c *Client) sendPublicRequest(r *Request) error {
	b, err := r.Encode()

	if err != nil {
		return err
	}

	// fmt.Printf("Sending public: %s\n", string(b))
	return c.publicConn.WriteMessage(websocket.TextMessage, b)
}

func (c *Client) recordPublicSubscription(channels []string) {
	for _, ch := range channels {
		c.publicSubs = append(c.publicSubs, ch)
	}
}

func (c *Client) recordPrivateSubscription(channels []string) {
	for _, ch := range channels {
		c.privateSubs = append(c.privateSubs, ch)
	}
}