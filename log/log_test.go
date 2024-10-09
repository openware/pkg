package log

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"
)

func TestLogger(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	rootPath = filepath.Join(filepath.Dir(filename), "../..")
	SetRootPath(rootPath)
	logger := NewLogger("test")

	logger.Debug("test debug")
	logger.Info("test info")
	logger.Warn("test warn")
	logger.Error("test error")
	logger.Trace("test fatal")

	ctx := context.Background()
	logger.For(ctx).Debug("test debug")
	logger.For(ctx).Info("test info")
	logger.For(ctx).Warn("test warn")
	logger.For(ctx).Error("test error")
	logger.For(ctx).Trace("test fatal")
}
