package log

import (
	"os"
	"strings"
	"time"

	zaplogfmt "github.com/jsternberg/zap-logfmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level = zap.NewAtomicLevel()
var logger *zap.SugaredLogger

func init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.UTC().Format(time.RFC3339))
	}

	logger = zap.New(zapcore.NewCore(
		zaplogfmt.NewEncoder(config),
		os.Stdout,
		level,
	)).Sugar()
}

func NewLogger(pkg string) CtxLogger {
	return baseLogger{
		logger.With("pkg", pkg),
		0,
	}
}

func NewLoggerWithCallerLevel(pkg string, callerLevel int) CtxLogger {
	return baseLogger{
		logger.With("pkg", pkg),
		callerLevel,
	}
}

func SetLogLevel(logLevel string) {
	logLevel = strings.ToLower(logLevel)

	var zapLevel zapcore.Level
	switch logLevel {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "dpanic":
		zapLevel = zapcore.DPanicLevel
	case "panic":
		zapLevel = zapcore.PanicLevel
	}
	level.SetLevel(zapLevel)
}
