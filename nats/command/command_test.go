package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadParam(t *testing.T) {
	msg := &LoadConfigCommand{
		Method: "test",
		Params: []interface{}{"my-command"},
	}

	param, err := msg.ReadParam()
	assert.NoError(t, err)
	assert.Equal(t, "my-command", param)
}

func TestReadConfig(t *testing.T) {
	msg := &SetConfigValueCommand{
		Method: "test",
		Params: []interface{}{
			[2]string{"foo", "baz"},
		},
	}

	params, err := msg.ReadConfig()
	assert.NoError(t, err)
	assert.Equal(t, [2]string{"foo", "baz"}, params)
}
