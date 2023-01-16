package protocol

type MessageType int32

const (
	Request MessageType = iota
	Response
	Event
)

//go:generate msgp

type MessageBase struct {
	Type  MessageType `msg:"message_type"`
	MsgId uint32      `msg:"msg_id"`
}

type RequestMessage struct {
	MessageBase
	Method string        `msg:"method"`
	Params []interface{} `msg:"params"`
}

type ResponseMessage struct {
	MessageBase
	Error  interface{} `msg:"method"`
	Result interface{} `msg:"result"`
}

type EventMessage struct {
	MessageBase
	Method string        `msg:"method"`
	Params []interface{} `msg:"params"`
}

func NewRequestMessage(msgId uint32, method string, params []interface{}) *RequestMessage {
	return &RequestMessage{
		MessageBase: MessageBase{
			Type:  Request,
			MsgId: msgId,
		},
		Method: method,
		Params: params,
	}
}

func NewResponseMessage(msgId uint32, error any, result interface{}) *ResponseMessage {
	return &ResponseMessage{
		MessageBase: MessageBase{
			Type:  Response,
			MsgId: msgId,
		},
		Error:  nil,
		Result: nil,
	}
}

func NewEventMessage(msgId uint32, method string, params []interface{}) *EventMessage {
	return &EventMessage{
		MessageBase: MessageBase{
			Type:  Event,
			MsgId: msgId,
		},
		Method: method,
		Params: params,
	}
}
