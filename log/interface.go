package log

import "context"

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
	Trace(msg string, keysAndValues ...interface{})
	With(keysAndValues ...interface{}) Logger
}

type CtxLogger interface {
	Logger
	For(ctx context.Context) Logger
}
