package log

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type baseLogger struct {
	lg *zap.SugaredLogger
}

func (l baseLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.log(zap.DebugLevel, msg, keysAndValues...)
}

func (l baseLogger) Info(msg string, keysAndValues ...interface{}) {
	l.log(zap.InfoLevel, msg, keysAndValues...)
}

func (l baseLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.log(zap.WarnLevel, msg, keysAndValues...)
}

func (l baseLogger) Error(msg string, keysAndValues ...interface{}) {
	l.log(zap.ErrorLevel, msg, keysAndValues...)
}

func (l baseLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.log(zap.FatalLevel, msg, keysAndValues...)
}

func (l baseLogger) Trace(msg string, keysAndValues ...interface{}) {}

func (l baseLogger) With(keysAndValues ...interface{}) Logger {
	return baseLogger{lg: l.lg.With(keysAndValues...)}
}

func (l baseLogger) For(ctx context.Context) Logger {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return l
	}

	return spanLogger{
		span:   span,
		logger: l.lg,
		spanKeysAndValues: []interface{}{
			"trace_id", span.SpanContext().TraceID().String(),
			"span_id", span.SpanContext().SpanID().String(),
		},
	}
}

func (l baseLogger) log(level zapcore.Level, msg string, keysAndValues ...interface{}) {
	l.lg.Logw(level, msg, withCaller(keysAndValues)...)
}
