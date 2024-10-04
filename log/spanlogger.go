package log

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type spanLogger struct {
	logger            *zap.SugaredLogger
	span              trace.Span
	spanKeysAndValues []interface{}
}

func (sl spanLogger) Debug(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("debug", msg, keysAndValues...)
	sl.logger.Debugw(msg, append(sl.spanKeysAndValues, keysAndValues...)...)
}

func (sl spanLogger) Info(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("info", msg, keysAndValues...)
	sl.logger.Infow(msg, append(sl.spanKeysAndValues, keysAndValues...)...)
}

func (sl spanLogger) Warn(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("warn", msg, keysAndValues...)
	sl.logger.Warnw(msg, append(sl.spanKeysAndValues, keysAndValues...)...)
}

func (sl spanLogger) Error(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("error", msg, keysAndValues...)
	sl.logger.Errorw(msg, append(sl.spanKeysAndValues, keysAndValues...)...)
}

func (sl spanLogger) Fatal(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("fatal", msg, keysAndValues...)
	sl.span.SetStatus(codes.Error, msg)
	sl.logger.Fatalw(msg, append(sl.spanKeysAndValues, keysAndValues...)...)
}

func (sl spanLogger) Trace(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("trace", msg, keysAndValues...)
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (sl spanLogger) With(keysAndValues ...interface{}) Logger {
	return spanLogger{logger: sl.logger.With(keysAndValues...), span: sl.span, spanKeysAndValues: sl.spanKeysAndValues}
}

func (sl spanLogger) logToSpan(level, msg string, keysAndValues ...interface{}) {
	attributes := []attribute.KeyValue{
		attribute.String("level", level),
	}

	if len(keysAndValues)%2 != 0 {
		keysAndValues = append(keysAndValues, "MISSING")
	}

	for i := 0; i < len(keysAndValues); i += 2 {
		var key string
		s, keyIsStr := keysAndValues[i].(string)
		if keyIsStr {
			key = s
		} else {
			key = "invalidKeysAndValues"
		}

		if !keyIsStr {
			attributes = append(attributes, attribute.String(
				"invalidKeysAndValues",
				fmt.Sprint(keysAndValues[i:]),
			))
			break
		}

		var keyValue attribute.KeyValue
		switch v := keysAndValues[i+1].(type) {
		case bool:
			keyValue = attribute.Bool(key, v)
		case int:
			keyValue = attribute.Int(key, v)
		case int16:
		case int32:
		case int64:
		case uint8:
		case uint16:
		case uint32:
			keyValue = attribute.Int64(key, int64(v))
		case float32:
		case float64:
			keyValue = attribute.Float64(key, float64(v))
		case stringer:
			keyValue = attribute.String(key, v.String())
		default:
			keyValue = attribute.String(key, fmt.Sprint(v))
		}

		attributes = append(attributes, keyValue)
	}

	sl.span.AddEvent(
		msg,
		trace.WithAttributes(attributes...),
	)
}

type stringer interface {
	String() string
}
