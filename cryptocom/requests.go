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

func (c *Client) subscribeTradesRequest(markets []string) *Request {
	params := make([]string, 0)

	for _, v := range markets {
		params = append(params, "trade."+v)
	}

	return &Request{
		Id:     12,
		Type:   SubscribeRequest,
		Method: "subscribe",
		Params: map[string]interface{}{"channels": params},
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
