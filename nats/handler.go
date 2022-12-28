package nats

import (
	"os"
	"sync"

	"github.com/nats-io/nats.go"
)

type EventHandler interface {
	SubscribeToQueueUsingChannel(subject string, group string, msgChannel chan *nats.Msg) error
	SubscribeToQueue(subject string, group string, cb nats.MsgHandler) error
	Subscribe(subject string, cb nats.MsgHandler) error
}

type eventHandlerBase struct {
	subs        []*nats.Subscription
	mutex       sync.Mutex
	termination <-chan os.Signal
}

type handlerConfig struct {
	autoUnsubscribeOnShutdown bool
}

type NatsEventHandler struct {
	eventHandlerBase
	nc     *nats.Conn
	config *handlerConfig
}

type JsEventHandler struct {
	eventHandlerBase
	js     nats.JetStreamContext
	config *handlerConfig
}

// For Checking compatibility
var (
	_ EventHandler = (*NatsEventHandler)(nil)
	_ EventHandler = (*JsEventHandler)(nil)
)

func newHandlerBase(termination <-chan os.Signal) eventHandlerBase {
	return eventHandlerBase{
		termination: termination,
		subs:        make([]*nats.Subscription, 0),
		mutex:       sync.Mutex{},
	}
}

func NewNatsHandler(nc *nats.Conn, termination <-chan os.Signal, config *handlerConfig) *NatsEventHandler {
	if config == nil {
		config = NewHandlerDefaultConfig()
	}

	handler := NatsEventHandler{
		nc:               nc,
		eventHandlerBase: newHandlerBase(termination),
		config:           config,
	}

	go handler.handleShutdown(config.autoUnsubscribeOnShutdown)

	return &handler
}

func NewJsHandler(nc *nats.Conn, termination <-chan os.Signal, config *handlerConfig) (*JsEventHandler, error) {
	if config == nil {
		config = NewHandlerDefaultConfig()
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	handler := &JsEventHandler{
		js:               js,
		eventHandlerBase: newHandlerBase(termination),
		config:           config,
	}

	go handler.handleShutdown(config.autoUnsubscribeOnShutdown)

	return handler, nil
}

func NewHandlerDefaultConfig() *handlerConfig {
	return &handlerConfig{
		autoUnsubscribeOnShutdown: true,
	}
}

func (h *eventHandlerBase) handleShutdown(unsubOnShutdown bool) []error {
	if unsubOnShutdown {
		for _ = range h.termination {
			var errors []error
			for _, sub := range h.subs {
				err := sub.Unsubscribe()
				if err != nil {
					errors = append(errors, err)
				}
			}

			return errors
		}
	}

	return nil
}

func (h *eventHandlerBase) GetSubscriptions() []*nats.Subscription {
	return h.subs
}

func (h *eventHandlerBase) pushSub(sub *nats.Subscription) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.subs = append(h.subs, sub)
}

func (j *JsEventHandler) SubscribeToQueueUsingChannel(subject string, group string, msgChannel chan *nats.Msg) error {
	sub, err := j.js.ChanQueueSubscribe(subject, group, msgChannel, nats.AckExplicit())
	if err != nil {
		return err
	}

	j.pushSub(sub)
	return nil
}

func (j *JsEventHandler) SubscribeToQueue(subject string, group string, cb nats.MsgHandler) error {
	sub, err := j.js.QueueSubscribe(subject, group, cb, nats.AckExplicit())
	if err != nil {
		return err
	}

	j.pushSub(sub)
	return nil
}

func (j *JsEventHandler) Subscribe(subject string, cb nats.MsgHandler) error {
	sub, err := j.js.Subscribe(subject, cb, nats.AckExplicit())
	if err != nil {
		return err
	}

	j.pushSub(sub)
	return nil
}

func (n *NatsEventHandler) SubscribeToQueueUsingChannel(subject string, group string, msgChannel chan *nats.Msg) error {
	sub, err := n.nc.ChanQueueSubscribe(subject, group, msgChannel)
	if err != nil {
		return err
	}

	n.pushSub(sub)
	return nil
}

func (n *NatsEventHandler) SubscribeToQueue(subject string, group string, cb nats.MsgHandler) error {
	sub, err := n.nc.QueueSubscribe(subject, group, cb)
	if err != nil {
		return err
	}

	n.pushSub(sub)
	return nil
}

func (n *NatsEventHandler) Subscribe(subject string, cb nats.MsgHandler) error {
	sub, err := n.nc.Subscribe(subject, cb)

	if err != nil {
		return err
	}

	n.pushSub(sub)
	return nil
}
