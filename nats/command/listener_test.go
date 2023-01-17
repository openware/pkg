package command

import (
	"sync"
	"testing"

	"github.com/openware/pkg/nats"
	"github.com/openware/pkg/nats/protocol"
	"github.com/stretchr/testify/assert"
)

func setupCommandListener() *commandListener {
	nc, _ := nats.InitEmbededNats()
	handler := nats.NewNatsHandler(nc, nil, nil)

	return NewCommandListener("foo", handler)
}

func TestNewCommandListener(t *testing.T) {
	nc, _ := nats.InitEmbededNats()
	handler := nats.NewNatsHandler(nc, nil, nil)
	listener := NewCommandListener("foo", handler)

	assert.Equal(t, "foo.ctrl", listener.subjectName)
}

func TestListenToCommands(t *testing.T) {
	comListener := setupCommandListener()
	wg := sync.WaitGroup{}
	shouldRestart := false
	configLoadout := ""
	var configValues []string

	comListener.SetServiceRestartCb(func() {
		shouldRestart = true
		wg.Done()
	})
	comListener.SetConfigLoadCb(func(s string) {
		configLoadout = s
		wg.Done()
	})
	comListener.ListenToSetConfigValue(func(key string, val string) {
		configValues = []string{key, val}
		wg.Done()
	})
	comListener.ListenToCommands()

	nc, _ := nats.InitEmbededNats()

	command1 := NewServiceRestartCommand(123)
	data, _ := protocol.MarshalMessage[protocol.RequestMessage]((*protocol.RequestMessage)(command1))
	nc.Publish("foo.ctrl", data)

	command2 := NewLoadConfigCommand(234, "test-param")
	data, _ = protocol.MarshalMessage[protocol.RequestMessage]((*protocol.RequestMessage)(command2))
	nc.Publish("foo.ctrl", data)

	command3 := NewSetConfigValueCommand(234, []string{"key", "value"})
	data, _ = protocol.MarshalMessage[protocol.RequestMessage]((*protocol.RequestMessage)(command3))
	nc.Publish("foo.ctrl", data)

	wg.Add(3)
	wg.Wait()

	assert.Equal(t, true, shouldRestart)
	assert.Equal(t, "test-param", configLoadout)
	assert.Equal(t, []string{"key", "value"}, configValues)
}
