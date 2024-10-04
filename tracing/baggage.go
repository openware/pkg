package tracing

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
)

func BaggageItem(ctx context.Context, key string) string {
	b := baggage.FromContext(ctx)
	m := b.Member(key)
	return m.Value()
}
