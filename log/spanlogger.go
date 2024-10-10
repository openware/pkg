package log

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type spanLogger struct {
	logger            *zap.SugaredLogger
	span              trace.Span
	spanKeysAndValues []interface{}
}

func (sl spanLogger) Debug(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("debug", msg, keysAndValues...)
	sl.log(zapcore.DebugLevel, msg, keysAndValues...)
}

func (sl spanLogger) Info(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("info", msg, keysAndValues...)
	sl.log(zapcore.InfoLevel, msg, keysAndValues...)
}

func (sl spanLogger) Warn(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("warn", msg, keysAndValues...)
	sl.log(zapcore.WarnLevel, msg, keysAndValues...)
}

func (sl spanLogger) Error(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("error", msg, keysAndValues...)
	sl.log(zapcore.ErrorLevel, msg, keysAndValues...)
	sl.span.SetStatus(codes.Error, msg)
}

func (sl spanLogger) Fatal(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("fatal", msg, keysAndValues...)
	sl.span.SetStatus(codes.Error, msg)
	sl.log(zapcore.FatalLevel, msg, keysAndValues...)
}

func (sl spanLogger) Trace(msg string, keysAndValues ...interface{}) {
	sl.logToSpan("trace", msg, keysAndValues...)
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (sl spanLogger) With(keysAndValues ...interface{}) Logger {
	return spanLogger{logger: sl.logger.With(keysAndValues...), span: sl.span, spanKeysAndValues: sl.spanKeysAndValues}
}

func (sl spanLogger) log(level zapcore.Level, msg string, keysAndValues ...interface{}) {
	sl.logger.Logw(level, msg, append(sl.spanKeysAndValues, withCaller(keysAndValues)...)...)
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
		case int16, int32, int64, uint8, uint16, uint32:
			keyValue = attribute.Int64(key, toInt64(v))
		case float32, float64:
			keyValue = attribute.Float64(key, toFloat64(v))
		case fmt.Stringer:
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

// Helper function to convert integer types to int64
func toInt64(value interface{}) int64 {
	switch v := value.(type) {
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	default:
		return 0
	}
}

// Helper function to convert float types to float64
func toFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}
