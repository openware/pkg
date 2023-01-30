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
	), zap.AddCaller()).Sugar()
}

func NewLogger(pkg string) *zap.SugaredLogger {
	return logger.With("pkg", pkg)
}

func ExtendLogger(logger *zap.SugaredLogger, label string, value interface{}) *zap.SugaredLogger {
	return logger.With(label, value)
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
