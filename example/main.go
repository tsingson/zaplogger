package main

import (
	"go.uber.org/zap"

	"github.com/tsingson/zaplogger"
)

func main() {
	core := zaplogger.NewConsoleDebug()

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("constructed a logger")
}
