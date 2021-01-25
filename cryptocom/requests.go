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
	mrkts := make(map[string][]string)
	for i, v := range markets {
		markets[i] = "trade." + v
	}
	mrkts["channels"] = markets
	return &Request{
		Id:     12,
		Type:   SubscribeRequest,
		Method: "subscribe",
		Params: mrkts, // Not sure
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
