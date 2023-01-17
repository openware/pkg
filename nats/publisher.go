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
	PublishMultiple(topics []string, payload []byte) []error
}

// EventPublisher nats event publisher interface
type EventPublisher interface {
	eventPublisherBase
}

// JsEventPublisher jetstream event publisher and event stream manager
type JsEventPublisher interface {
	eventPublisherBase
	CreateNewEventStream(string, []string) error
	DeleteEventStream(streamName string) error
}

// publisherBase Base publisher stuct. it has base implementation for publishing events
type publisherBase struct {
	nc *nats.Conn
}

type natsEventPublisher struct {
	publisherBase
}

type jsEventPublisher struct {
	publisherBase
	js nats.JetStreamContext
}

// NewNatsEventPublisher initialize new nats event publisher
func NewNatsEventPublisher(nc *nats.Conn) *natsEventPublisher {
	dispatcher := natsEventPublisher{
		publisherBase{nc: nc},
	}

	return &dispatcher
}

// NewJsEventPublisher initialize new jetstream event publisher
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

// Request make a request request to specific subject. (default timeout is set on 3 seconds)
func (p *publisherBase) Request(subject string, data []byte) (*nats.Msg, error) {
	// TODO: maybe modify to something else
	return p.RequestWithTimeout(subject, data, time.Second*3)
}

// RequestWithTimeout make a request to specific subject and specify timeout
func (p *publisherBase) RequestWithTimeout(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return p.nc.Request(subject, data, timeout)
}

// Publish publish an event for specific topic
func (p *publisherBase) Publish(topic string, payload []byte) error {
	return p.nc.Publish(topic, payload)
}

// PublishMultiple publish an event for multiple payload. because each publish may result with error we return error array
func (p *publisherBase) PublishMultiple(topics []string, payload []byte) []error {
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

// CreateNewEventStream create stream for specific subject for jetstream. if stream already exists nothing happens.
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

// DeleteEventStream delete a stream in jetstream
func (j *jsEventPublisher) DeleteEventStream(streamName string) error {
	return j.js.DeleteStream(streamName)
}

// Publish publish a new event
func (j *jsEventPublisher) Publish(topic string, payload []byte) error {
	_, err := j.js.Publish(topic, payload)
	return err
}

// PublishMultiple publish an event for multiple payload. because each publish may result with error we return error array
func (j *jsEventPublisher) PublishMultiple(topics []string, payload []byte) []error {
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
