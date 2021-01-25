package cryptocom

func (c *Client) AuthRequest() *Request {
	r := &Request{
		Id:     12,
		Type:   AuthRequest,
		Method: "public/auth",
		ApiKey: c.key,
		Nonce:  generateNonce(),
	}

	c.generateSignature(r)
	return r
}

func (c *Client) subscribeRequest(channels []string) *Request {
	return &Request{
		Id:     12,
		Type:   SubscribeRequest,
		Method: "subscribe",
		Params: map[string]interface{}{"channels": channels},
		Nonce:  generateNonce(),
	}
}

func (c *Client) hearBeatRequest(id int) *Request {
	return &Request{
		Id:     id,
		Type:   HeartBeat,
		Method: "public/respond-heartbeat",
	}
}

func (c *Client) createOrderRequest(market string, side string, oType string, price )
