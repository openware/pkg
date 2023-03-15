package command

import (
	"errors"
	"fmt"

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

func NewSetConfigValueCommand(msgId uint32, config []string) *SetConfigValueCommand {
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

func (command *SetConfigValueCommand) ReadConfig() ([]string, error) {
	if len(command.Params) != 1 {
		return nil, InvalidParamError
	}

	switch command.Params[0].(type) {
	case []string:
		return command.Params[0].([]string), nil
	case []interface{}:
		val := command.Params[0].([]interface{})
		params := make([]string, 2)
		for i, param := range val {
			params[i] = param.(string)
		}

		return params, nil

	default:
		return nil, fmt.Errorf("invalid type")
	}
}
