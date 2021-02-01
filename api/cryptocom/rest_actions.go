package cryptocom

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

const apiVersionSuffix = "/v2/"

func (c *Client) RestGetOrderDetails(reqID int, remoteID sql.NullString) (Response, error) {
	r := c.restGetOrderDetailsRequest(reqID, remoteID)
	body, err := r.Encode()
	if err != nil {
		return Response{}, err
	}

	endpoint := c.restRootURL + apiVersionSuffix + r.Method
	resp, err := c.httpClient.Post(endpoint, "application/json", bytes.NewBuffer(body)) // Need to handle 503 response

	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var parsed Response
	err = json.NewDecoder(resp.Body).Decode(&parsed)

	return parsed, err
}

func (c *Client) RestGetBalance(reqID int) (Response, error) {
	r := c.restGetBalanceRequest(reqID)
	body, err := r.Encode()
	if err != nil {
		return Response{}, err
	}

	resp, err := c.httpClient.Post(c.restRootURL+apiVersionSuffix+r.Method, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var parsed Response
	err = json.NewDecoder(resp.Body).Decode(&parsed)

	return parsed, err
}

func (c *Client) RestGetTrades(reqID int, market string) (Response, error) {
	r := c.restGetTradesRequest(reqID, market)
	body, err := r.Encode()
	if err != nil {
		return Response{}, err
	}

	resp, err := c.httpClient.Post(c.restRootURL+apiVersionSuffix+r.Method, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var parsed Response
	err = json.NewDecoder(resp.Body).Decode(&parsed)

	return parsed, err
}

func (c *Client) RestOpenOrders(reqID int, market string, pageNumber int, pageSize int) (Response, error) {
	r := c.restOpenOrdersRequest(reqID, market, pageNumber, pageSize)
	body, err := r.Encode()
	if err != nil {
		return Response{}, err
	}

	resp, err := c.httpClient.Post(c.restRootURL+apiVersionSuffix+r.Method, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var parsed Response
	err = json.NewDecoder(resp.Body).Decode(&parsed)

	return parsed, err
}
