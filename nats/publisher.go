package nats

import (
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

type eventPublisherBase interface {
	Request(subj string, data []byte) (*nats.Msg, error)
	RequestWithTimeout(subject string, data []byte, timeout time.Duration) (*nats.Msg, error)
	Publish(topic string, payload []byte) error
	publishMultiple(topics []string, payload []byte) []error
}

type EventPublisher interface {
	eventPublisherBase
}

type JsEventPublisher interface {
	eventPublisherBase
	CreateNewEventStream(string, []string) error
	DeleteEventStream(streamName string) error
}

type publisherBase struct {
	nc *nats.Conn
}

var _ eventPublisherBase = (*publisherBase)(nil)

type natsEventPublisher struct {
	publisherBase
}

var _ EventPublisher = (*natsEventPublisher)(nil)

type jsEventPublisher struct {
	publisherBase
	js nats.JetStreamContext
}

var _ JsEventPublisher = (*jsEventPublisher)(nil)

func NewNatsEventPublisher(nc *nats.Conn) *natsEventPublisher {
	dispatcher := natsEventPublisher{
		publisherBase{nc: nc},
	}

	return &dispatcher
}

func NewJsEventPublisher(nc *nats.Conn) (*jsEventPublisher, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	dispatcher := jsEventPublisher{
		publisherBase: publisherBase{nc},
		js:            js,
	}

	return &dispatcher, nil
}

func (p *publisherBase) Request(subject string, data []byte) (*nats.Msg, error) {
	// TODO: maybe modify to something else
	return p.RequestWithTimeout(subject, data, time.Second*3)
}

func (p *publisherBase) RequestWithTimeout(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return p.nc.Request(subject, data, timeout)
}

func (p *publisherBase) Publish(topic string, payload []byte) error {
	return p.nc.Publish(topic, payload)
}

func (p *publisherBase) publishMultiple(topics []string, payload []byte) []error {
	wg := sync.WaitGroup{}
	var errors []error
	for _, topic := range topics {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := p.Publish(topic, payload)
			if err != nil {
				errors = append(errors, err)
			}
		}()
	}
	wg.Wait()

	return errors
}

func (j *jsEventPublisher) CreateNewEventStream(streamName string, subjects []string) error {
	stream, err := j.js.StreamInfo(streamName)
	if err != nil && err != nats.ErrStreamNotFound {
		return err
	}

	if stream == nil {
		_, err := j.js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: subjects,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (j *jsEventPublisher) DeleteEventStream(streamName string) error {
	return j.js.DeleteStream(streamName)
}

func (j *jsEventPublisher) Publish(topic string, payload []byte) error {
	_, err := j.js.Publish(topic, payload)
	return err
}

func (j *jsEventPublisher) publishMultiple(topics []string, payload []byte) []error {
	wg := sync.WaitGroup{}
	var errors []error
	for _, topic := range topics {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := j.js.Publish(topic, payload)
			if err != nil {
				errors = append(errors, err)
			}
		}()
	}
	wg.Wait()

	return errors
}
