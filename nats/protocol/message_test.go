package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUnmarshal(t *testing.T) {
	msg1 := NewRequestMessage(uint32(123), "test", []interface{}{
		[]int32{1, 2, 3}, struct {
			Param string
		}{
			Param: "test-param",
		},
		"string",
	})

	bytes, err := MarshalMessage[RequestMessage](msg1)
	assert.NoError(t, err)

	msg1_1, err := UnmarshalMessage[RequestMessage](bytes)
	assert.Equal(t, &RequestMessage{
		messageBase: messageBase{
			Type:  Request,
			MsgId: 123,
		},
		Method: "test",
		Params: []interface{}{
			[]interface{}{1, 2, 3},
			map[string]interface{}{
				"Param": "test-param",
			},
			"string",
		},
	}, msg1_1)
}
