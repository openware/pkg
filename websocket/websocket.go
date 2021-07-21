package websocket

import (
	"errors"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// Client for forex api instance
type Client struct {
	Address   *string
	Query     *string
	Conn      *websocket.Conn
	Listeners *[]func(interface{})
	logger    *log.Logger
}

var reconnectEvents = []int{
	websocket.CloseGoingAway,
	websocket.CloseAbnormalClosure,
}

var reconnectIntervalSeconds int64 = 5

// New : create the weboscket instance
func New(address *string, logger *log.Logger) (*Client, error) {
	if address == nil {
		return nil, errors.New("'address' is missing")
	}
	if logger == nil {
		logger = log.New(os.Stderr, "WebSocket: ", log.LstdFlags)
	}
	return &Client{
		Address:   address,
		logger:    logger,
		Listeners: &[]func(interface{}){},
	}, nil
}

// Connect to websocket endpoint
func (client *Client) Connect() error {
	// connection
	u, err := url.Parse(*client.Address)
	if err != nil {
		return err
	}
	if u.Scheme != "ws" && u.Scheme != "wss" {
		return errors.New(`url scheme must be one of "ws" and "wss"`)
	}
	u.RawQuery = *client.Query

	dialer := websocket.DefaultDialer
	dialer.ReadBufferSize = 1024
	dialer.WriteBufferSize = 1024

	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	client.Conn = c

	// reader
	go func() {
		for {
			var message interface{}
			err := c.ReadJSON(&message)

			if err != nil {
				client.logger.Println("error:", err)

				// Autoreconnect
				if websocket.IsCloseError(err, reconnectEvents...) {
					client.Reconnect()
				}

				break
			}
			go func() {
				for _, listener := range *client.Listeners {
					go listener(message)
				}
			}()
		}
	}()

	return nil
}

// Close the websocket connection
func (client *Client) Close() error {
	if client.IsConnected() {
		if err := client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
			client.logger.Println(err)
		}
		if err := client.Conn.Close(); err != nil {
			return err
		}
	}
	return nil
}

// IsConnected : check the websocket connection exists
func (client *Client) IsConnected() bool {
	return client.Conn != nil
}

// Reconnect the websocket
func (client *Client) Reconnect() {
	client.logger.Println("reconnecting...")
	for {
		err := client.Connect()
		if err == nil {
			break
		} else {
			client.logger.Println("error:", err)
			client.logger.Printf("reconnect failed. waiting for %v seconds to retry...", reconnectIntervalSeconds)
			time.Sleep(time.Second * time.Duration(reconnectIntervalSeconds))
		}
	}
}

// AddListener : add the new listerer function and return total listeners amount
func (client *Client) AddListener(handler func(interface{})) int {
	*client.Listeners = append(*client.Listeners, handler)
	return len(*client.Listeners)
}

// RemoveListener : remove existing listerer function
func (client *Client) RemoveListener(index int) {
	ls := *client.Listeners
	*client.Listeners = append(ls[:index], ls[index+1:]...)
}

// ClearListeners : clean all listeners function
func (client *Client) ClearListeners() {
	client.Listeners = &[]func(interface{}){}
}

// Send the new websocket payload
func (client *Client) Send(payload interface{}) error {
	client.logger.Println("send:", payload)
	err := client.Conn.WriteJSON(payload)
	if err != nil {
		return err
	}
	return nil
}
