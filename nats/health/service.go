package health

import (
	"time"

	"github.com/openware/pkg/nats"
)

func RunHealthService(publisher nats.EventPublisher, serviceName string, period time.Duration, cb func() string) {
	subjectName := serviceName + ".metric"
	go publishHealthReport(publisher, subjectName, period, cb)
}

func publishHealthReport(publisher nats.EventPublisher, subject string, period time.Duration, cb func() string) {
	for range time.Tick(period) {
		result := cb()
		publisher.Publish(subject, []byte(result))
	}
}
