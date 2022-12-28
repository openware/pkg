package nats

import (
	"sync"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func initNatsHandlers() (*natsEventPublisher, *NatsEventHandler) {
	nc, _ := InitEmbededNats()
	publisher := NewNatsEventPublisher(nc)
	handler := NewNatsHandler(nc, nil, nil)

	return publisher, handler
}

func TestNatsPubSub(t *testing.T) {
	pub, handler := initNatsHandlers()
	handler.Subscribe("test.*", func(msg *nats.Msg) {
		res := string(msg.Data)
		assert.Equal(t, "test", res)
	})

	handler.Subscribe("test.*", func(msg *nats.Msg) {
		res := string(msg.Data)
		assert.Equal(t, "test", res)
	})

	pub.Publish("test.all", []byte("test"))
}

func TestNatsQueue(t *testing.T) {
	pub, handler := initNatsHandlers()
	wg := sync.WaitGroup{}
	subCount := 0
	handler.SubscribeToQueue("test.*", "grp", func(msg *nats.Msg) {
		subCount++
		wg.Done()
	})

	handler.SubscribeToQueue("test.*", "grp", func(msg *nats.Msg) {
		subCount++
		wg.Done()
	})

	handler.SubscribeToQueue("test.*", "grp2", func(msg *nats.Msg) {
		subCount++
		wg.Done()
	})

	pub.Publish("test.all", []byte("test"))
	wg.Add(2)
	wg.Wait()
	assert.Equal(t, 2, subCount)
}

func TestNatsChannelQueue(t *testing.T) {
	pub, handler := initNatsHandlers()
	msgChan := make(chan *nats.Msg)
	wg := sync.WaitGroup{}
	go func() {
		for msg := range msgChan {
			assert.Equal(t, "test", string(msg.Data))
			wg.Done()
		}
	}()

	handler.SubscribeToQueueUsingChannel("test.*", "grp", msgChan)
	pub.Publish("test.topic", []byte("test"))
	wg.Add(1)
	wg.Wait()
}
