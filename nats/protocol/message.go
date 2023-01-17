package protocol

import "encoding/json"

type MessageType uint8

const (
	Request MessageType = iota
	Response
	Event
)

type messageBase struct {
	Type  MessageType `json:"message_type"`
	MsgId uint32      `json:"msg_id"`
}

type RequestMessage struct {
	messageBase
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type ResponseMessage struct {
	messageBase
	Error  interface{} `json:"method"`
	Result interface{} `json:"result"`
}

type EventMessage struct {
	messageBase
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

func NewRequestMessage(msgId uint32, method string, params []interface{}) *RequestMessage {
	return &RequestMessage{
		messageBase: messageBase{
			Type:  Request,
			MsgId: msgId,
		},
		Method: method,
		Params: params,
	}
}

func NewResponseMessage(msgId uint32, error any, result interface{}) *ResponseMessage {
	return &ResponseMessage{
		messageBase: messageBase{
			Type:  Response,
			MsgId: msgId,
		},
		Error:  error,
		Result: result,
	}
}

func NewEventMessage(msgId uint32, method string, params []interface{}) *EventMessage {
	return &EventMessage{
		messageBase: messageBase{
			Type:  Event,
			MsgId: msgId,
		},
		Method: method,
		Params: params,
	}
}

func MarshalMessage[T RequestMessage | ResponseMessage | EventMessage](message *T) ([]byte, error) {
	return json.Marshal(message)
}

func UnmarshalMessage[T RequestMessage | ResponseMessage | EventMessage](data []byte) (*T, error) {
	message := new(T)
	err := json.Unmarshal(data, message)
	return message, err
}
