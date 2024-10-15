package tracing

import (
	"context"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type MutedExporter struct{}

func (e *MutedExporter) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	return nil
}

func (e *MutedExporter) Shutdown(ctx context.Context) error {
	return nil
}
