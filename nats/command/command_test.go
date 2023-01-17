package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadParam(t *testing.T) {
	msg := NewLoadConfigCommand(123, "my-command")

	param, err := msg.ReadParam()
	assert.NoError(t, err)
	assert.Equal(t, LOAD_CONFIG_COMMAND, msg.Method)
	assert.Equal(t, uint32(123), msg.MsgId)
	assert.Equal(t, "my-command", param)
}

func TestReadConfig(t *testing.T) {
	msg := NewSetConfigValueCommand(234, []string{"foo", "baz"})
	assert.Equal(t, SET_CONFIG_VALUE_COMMAND, msg.Method)
	assert.Equal(t, uint32(234), msg.MsgId)

	params, err := msg.ReadConfig()
	assert.NoError(t, err)
	assert.Equal(t, []string{"foo", "baz"}, params)
}
