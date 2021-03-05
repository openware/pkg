package forex

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/openware/pkg/websocket"
)

const (
	responseCode = 3
)

const (
	priceResponseType = "forex"
)

// Client for forex api instance
type Client struct {
	ID      string
	Streams map[string]struct{}
	WS      *websocket.Client
	logger  *log.Logger
}

// PriceResponse from websocket
type PriceResponse struct {
	Market    string
	Price     string
	CreatedAt float64
	UpdatedAt float64
}

// New Forex client
func New(id string, url string, streams []string) (*Client, error) {
	if url == "" {
		return nil, errors.New("'url' is missing")
	}
	if streams == nil {
		streams = []string{}
	}
	logger := log.New(os.Stderr, fmt.Sprintf("[%s] Forex: ", id), log.LstdFlags)

	ws, err := websocket.New(&url, logger)
	if err != nil {
		return nil, err
	}

	streamsMap := map[string]struct{}{}
	for _, stream := range streams {
		streamsMap[stream] = struct{}{}
	}

	return &Client{
		ID:      id,
		WS:      ws,
		Streams: streamsMap,
		logger:  logger,
	}, nil
}

// Connect to Forex websocket endpoint
func (c *Client) Connect(listener func(*PriceResponse)) error {
	if err := c.Close(); err != nil {
		return err
	}

	c.WS.AddListener(func(message interface{}) {
		data, err := c.parsePriceData(message)
		if err != nil {
			c.logger.Println(err)
		} else if data != nil {
			listener(data)
		}
	})

	c.refreshStreamsQuery()

	if err := c.WS.Connect(); err != nil {
		return err
	}

	return nil
}

// Close websocket connection
func (c *Client) Close() error {
	err := c.WS.Close()
	return err
}

// Subscribe to Forex markets
func (c *Client) Subscribe(market string) error {
	if !c.WS.IsConnected() {
		return errors.New("WebSocket is not connected")
	}

	if _, exists := c.Streams[market]; exists {
		return nil
	}

	payload := c.getPayload("subscribe", []string{market})
	if err := c.WS.Send(*payload); err != nil {
		return err
	}

	c.Streams[market] = struct{}{}
	c.refreshStreamsQuery()
	return nil
}

// Unsubscribe to Forex markets
func (c *Client) Unsubscribe(market string) error {
	_, exists := c.Streams[market]
	if !exists {
		return nil
	}

	if !c.WS.IsConnected() {
		return errors.New("WebSocket is not connected")
	}

	payload := c.getPayload("unsubscribe", []string{market})
	if err := c.WS.Send(*payload); err != nil {
		return err
	}

	delete(c.Streams, market)
	c.refreshStreamsQuery()
	return nil
}

func (c *Client) refreshStreamsQuery() {
	markets := &[]string{}
	for stream := range c.Streams {
		*markets = append(*markets, stream)
	}
	query := "stream=" + strings.Join(*markets, ",")

	c.WS.Query = &query
}

func (c *Client) getPayload(action string, data interface{}) *[]interface{} {
	return &[]interface{}{
		1,
		c.ID,
		action,
		data,
	}
}

func (c *Client) parsePriceData(message interface{}) (*PriceResponse, error) {
	if res, ok := message.([]interface{}); ok {
		code := res[0].(float64)
		rType := res[1].(string)
		if code == responseCode && rType == priceResponseType { // TODO: need to handle other response types later.
			data := res[2].([]interface{})
			pRes := &PriceResponse{
				Market:    data[0].(string),
				Price:     data[1].(string),
				CreatedAt: data[2].(float64),
				UpdatedAt: data[3].(float64),
			}
			return pRes, nil
		}
		return nil, nil
	}
	return nil, errors.New(`Can not parse message`)
}
