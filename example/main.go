package main

import (
	"github.com/tsingson/zaplogger"
)

func main() {
	logger := zaplogger.ConsoleDebug()

	defer logger.Sync()
	logger.Info("constructed a logger")
}
