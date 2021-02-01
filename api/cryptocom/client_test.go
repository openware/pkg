package cryptocom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

type connectionMock struct {
	Buffer       *bytes.Buffer
	ResponseMock *bytes.Buffer
}

func (c *Client) connectMock(privateResponse, publicResponse, privateWriter, publicWriter *bytes.Buffer) {
	c.publicConn = Connection{
		Transport: &connectionMock{Buffer: publicWriter, ResponseMock: publicResponse},
		IsPrivate: false,
	}
	c.privateConn = Connection{
		Transport: &connectionMock{Buffer: privateWriter, ResponseMock: privateResponse},
		IsPrivate: true,
	}
}

func (cm *connectionMock) ReadMessage() (int, []byte, error) {
	return 1, cm.ResponseMock.Bytes(), nil
}

func (cm *connectionMock) WriteMessage(messageType int, data []byte) error {
	cm.Buffer.Write(data)
	return nil
}

func (cm *connectionMock) Close() error {
	cm.Buffer.Reset()
	cm.ResponseMock.Reset()
	return nil
}

// TODO: test all possible requests
func TestConnectionWrite(t *testing.T) {
	client := New("test", "test", "test", "test")
	// publicBuffer := bytes.NewBuffer(nil)
	privateBuffer := bytes.NewBuffer(nil)

	t.Run("auth", func(t *testing.T) {
		client.privateConn = Connection{Transport: &connectionMock{Buffer: privateBuffer}}

		req := client.AuthRequest()
		client.sendPrivateRequest(req)

		expected := fmt.Sprintf("{\"api_key\":\"test\",\"id\":1,\"method\":\"public/auth\",\"nonce\":\"%v\",\"sig\":\"%v\"}", req.Nonce, req.Signature)
		assert.Equal(t, expected, privateBuffer.String())
	})
}

func TestConnectionRead(t *testing.T) {
	// prepare mock
	client := New("", "", "test", "test")
	publicBuffer := bytes.NewBuffer(nil)
	privateBuffer := bytes.NewBuffer(nil)
	client.connectMock(privateBuffer, publicBuffer, bytes.NewBuffer(nil), bytes.NewBuffer(nil))

	// expectations
	expectStr := `{"id":12,"method":"public/auth","code":10002,"message":"UNAUTHORIZED"}`
	var expectedResponse Response
	err := json.Unmarshal([]byte(expectStr), &expectedResponse)
	if err != nil {
		fmt.Println("error on parse expected message")
	}

	// mocked responses
	publicBuffer.WriteString(expectStr)
	privateBuffer.WriteString(expectStr)

	// Running client
	msgs := client.Listen()

	// assertion
	assert.Equal(t, expectedResponse, <-msgs)
}
