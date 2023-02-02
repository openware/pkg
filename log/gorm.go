package log

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlog "gorm.io/gorm/logger"
)

type GORMLogger struct {
	logger *zap.SugaredLogger
	gormlog.Config
}

func NewGORMLogger(conf gormlog.Config) gormlog.Interface {
	return &GORMLogger{
		logger: logger.With(zap.String("pkg", "gorm")),
		Config: conf,
	}
}

func (l *GORMLogger) LogMode(logLevel gormlog.LogLevel) gormlog.Interface {
	var zapLevel zapcore.Level
	switch logLevel {
	case gormlog.Silent:
		zapLevel = zapcore.DPanicLevel
	case gormlog.Error:
		zapLevel = zapcore.ErrorLevel
	case gormlog.Warn:
		zapLevel = zapcore.WarnLevel
	case gormlog.Info:
		zapLevel = zapcore.InfoLevel
	}
	level.SetLevel(zapLevel)

	return l
}

func (l *GORMLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Infof(msg, data...)
}

func (l *GORMLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Warnf(msg, data...)
}

func (l *GORMLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Errorf(msg, data...)
}

func (l *GORMLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if level.Level() >= zapcore.DPanicLevel {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	switch {
	case err != nil && level.Level() <= zapcore.ErrorLevel && (!errors.Is(err, gormlog.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		l.logger.Errorw(sql, "elapsed", elapsed, "rows", rows, "error", err.Error())
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && level.Level() >= zapcore.WarnLevel:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		l.logger.Warnw(sql, "elapsed", elapsed, "rows", rows, "warn", slowLog)
	case level.Level() < zapcore.InfoLevel:
		l.logger.Debugw(sql, "elapsed", elapsed, "rows", rows)
	}
}
