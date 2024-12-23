package tracing

import (
	"context"
	"fmt"
	"sync"
	"time"

	gcptrace "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/pkg/otelsemconv"
	"github.com/openware/pkg/log"
)

var once sync.Once

func InitOTEL(serviceName string, exporterType string, logger log.CtxLogger) trace.TracerProvider {
	once.Do(func() {
		otel.SetTextMapPropagator(
			propagation.NewCompositeTextMapPropagator(
				propagation.TraceContext{},
				propagation.Baggage{},
			))
	})

	exp, err := createOtelExporter(exporterType)
	if err != nil {
		logger.Fatal("cannot create exporter", "error", err, "exporterType", exporterType)
	}
	logger.Debug("using trace exporter", "type", exporterType)

	res, err := resource.New(
		context.Background(),
		resource.WithSchemaURL(otelsemconv.SchemaURL),
		resource.WithAttributes(otelsemconv.ServiceNameKey.String(serviceName)),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithOSType(),
	)
	if err != nil {
		logger.Fatal("resource creation failed", zap.Error(err))
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp, sdktrace.WithBatchTimeout(1000*time.Millisecond)),
		sdktrace.WithResource(res),
	)
	logger.Debug("created OTEL tracer", "serviceName", serviceName)
	return tp
}

func createOtelExporter(exporterType string) (sdktrace.SpanExporter, error) {
	var exporter sdktrace.SpanExporter
	var err error
	switch exporterType {
	case "otlp":
		exporter, err = otlptrace.New(
			context.Background(),
			otlptracehttp.NewClient(),
		)
	case "gcp":
		exporter, err = gcptrace.New()
	case "stdout":
		exporter, err = stdouttrace.New()
	case "muted":
		exporter = &MutedExporter{}
	default:
		return nil, fmt.Errorf("unrecognized exporter type %s", exporterType)
	}
	return exporter, err
}
