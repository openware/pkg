package command

import (
	"time"

	natspkg "github.com/nats-io/nats.go"
	"github.com/openware/pkg/nats"
	"github.com/openware/pkg/nats/protocol"
)

type CommandListener interface {
	ListenToCommands()
	SetServiceRestartCb(cb func())
	SetConfigLoadCb(cb func(string))
	SetConfigValueCb(cb func(string, string))
}

type commandListener struct {
	subjectName    string
	handler        nats.EventHandler
	restartService func()
	loadConfig     func(string)
	setConfigValue func(string, string)
}

func NewCommandListener(serviceName string, handler nats.EventHandler) *commandListener {
	return &commandListener{
		subjectName: serviceName + ".ctrl",
		handler:     handler,
	}
}

func (c *commandListener) ListenToCommands() {
	go c.listenToCommands(3)
}

func (c *commandListener) listenToCommands(tries int32) {
	err := c.handler.Subscribe(c.subjectName, func(msg *natspkg.Msg) {
		request, err := protocol.UnmarshalMessage[protocol.RequestMessage](msg.Data)
		// TODO: handle error
		if err != nil {
			return
		}
		switch request.Method {
		case RESTART_SERVICE_COMMAND:
			c.restartService()
			break
		case LOAD_CONFIG_COMMAND:
			command := (LoadConfigCommand)(*request)
			param, _ := command.ReadParam()
			c.loadConfig(param)
			break
		case SET_CONFIG_VALUE_COMMAND:
			command := (SetConfigValueCommand)(*request)
			param, _ := command.ReadConfig()
			c.setConfigValue(param[0], param[1])
			break
		default:
			// TODO: log
		}
	})

	if err != nil && tries > 0 {
		time.Sleep(100 * time.Millisecond)
	} else if err != nil {
		panic("couldn't listen to commands after 3 tries")
	}
}

func (c *commandListener) SetServiceRestartCb(cb func()) {
	c.restartService = cb
}

func (c *commandListener) SetConfigLoadCb(cb func(string)) {
	c.loadConfig = cb
}

func (c *commandListener) ListenToSetConfigValue(cb func(string, string)) {
	c.setConfigValue = cb
}
