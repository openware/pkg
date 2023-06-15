package forex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func prepareHandler(t *testing.T, handler func(c *websocket.Conn)) *httptest.Server {
	var upgrader = websocket.Upgrader{}
	h := func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		assert.NoError(t, err)
		defer c.Close()
		handler(c)
	}
	s := httptest.NewServer(http.HandlerFunc(h))
	return s
}

func prepareTest(t *testing.T, handler func(c *websocket.Conn)) (*httptest.Server, *Client) {
	s := prepareHandler(t, handler)
	u := "ws://" + strings.TrimPrefix(s.URL, "http://")
	f, err := New("platform_id", u, nil)
	assert.NoError(t, err)

	return s, f
}

func TestGetCorrectData(t *testing.T) {
	s, f := prepareTest(t, func(c *websocket.Conn) {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			var req []interface{}
			if err := json.Unmarshal(message, &req); err != nil {
				fmt.Println(err)
				break
			}
			market := req[3].([]interface{})[0].(string)
			d := []byte(`[3,"forex",["` + market + `","100.0",1,2]]`)
			err = c.WriteMessage(mt, d)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	})
	defer s.Close()
	defer f.Close()

	ch := make(chan PriceResponse)
	err := f.Connect(func(p *PriceResponse) {
		ch <- *p
	})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	err = f.Subscribe("testusdt")
	assert.NoError(t, err)

	d := <-ch
	assert.Equal(t, d.Market, "testusdt")
	assert.Equal(t, d.Price, "100.0")
	assert.Equal(t, d.CreatedAt, 1.0)
	assert.Equal(t, d.UpdatedAt, 2.0)
}

func TestOnlySubscribeOnce(t *testing.T) {
	ch := make(chan int)
	count := 0
	s, f := prepareTest(t, func(c *websocket.Conn) {
		for {
			_, _, err := c.ReadMessage()
			if err == nil {
				count++
			} else {
				ch <- count
				break
			}
		}
	})
	defer s.Close()
	defer f.Close()

	err := f.Connect(func(p *PriceResponse) {})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, len(f.Streams), 0)

	err = f.Subscribe("testusdt")
	assert.NoError(t, err)
	err = f.Subscribe("testusdt")
	assert.NoError(t, err)

	f.Close()
	d := <-ch
	assert.Equal(t, 1, len(f.Streams))
	assert.Equal(t, 1, d)
}
