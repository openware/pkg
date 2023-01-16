package command

import (
	"errors"

	"github.com/openware/pkg/nats/protocol"
)

const (
	RESTART_SERVICE_COMMAND  = "service_restart"
	LOAD_CONFIG_COMMAND      = "load_config"
	SET_CONFIG_VALUE_COMMAND = "set_config"
)

var (
	InvalidParamError = errors.New("invalid param in payload")
)

type (
	ServiceRestartCommand protocol.RequestMessage
	LoadConfigCommand     protocol.RequestMessage
	SetConfigValueCommand protocol.RequestMessage
)

func NewServiceRestartCommand(msgId uint32) *ServiceRestartCommand {
	return (*ServiceRestartCommand)(protocol.NewRequestMessage(msgId, RESTART_SERVICE_COMMAND, nil))
}

func NewLoadConfigCommand(msgId uint32, param string) *LoadConfigCommand {
	return (*LoadConfigCommand)(protocol.NewRequestMessage(msgId, LOAD_CONFIG_COMMAND, []interface{}{param}))
}

func NewSetConfigValueCommand(msgId uint32, config [2]string) *SetConfigValueCommand {
	return (*SetConfigValueCommand)(protocol.NewRequestMessage(msgId, SET_CONFIG_VALUE_COMMAND, []interface{}{config}))
}

func (command *LoadConfigCommand) ReadParam() (string, error) {
	if len(command.Params) != 1 {
		return "", InvalidParamError
	}

	val, ok := command.Params[0].(string)
	if !ok {
		return "", InvalidParamError
	}

	return val, nil
}

func (command *SetConfigValueCommand) ReadConfig() ([2]string, error) {
	if len(command.Params) != 1 {
		return [2]string{}, InvalidParamError
	}

	val, ok := command.Params[0].([2]string)
	if !ok {
		return [2]string{}, InvalidParamError
	}

	return val, nil
}
