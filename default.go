package zaplogger

import (
	"fmt"

	"go.uber.org/zap"
)

var Default *Logger

func init() {
	log := zap.New(newStdoutCore(zap.DebugLevel))
	Default = &Logger{
		Log: log,
	}
}

// Debug logs a message at level Debug on the ZapLogger.
func Debug(args ...interface{}) {
	Default.Log.Debug(fmt.Sprint(args...))
}

// Debugf logs a message at level Debug on the ZapLogger.
func Debugf(template string, args ...interface{}) {
	Default.Log.Debug(fmt.Sprintf(template, args...))
}

// Info logs a message at level Info on the ZapLogger.
func Info(args ...interface{}) {
	Default.Log.Info(fmt.Sprint(args...))
}

// Infof logs a message at level Info on the ZapLogger.
func Infof(template string, args ...interface{}) {
	Default.Log.Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at level Warn on the ZapLogger.
func Warn(args ...interface{}) {
	Default.Log.Warn(fmt.Sprint(args...))
}

// Warnf logs a message at level Warn on the ZapLogger.
func Warnf(template string, args ...interface{}) {

	Default.Log.Warn(fmt.Sprintf(template, args...))
}

// Warningf logs a message at level Warn on the ZapLogger.
func Warningf(template string, args ...interface{}) {

	Default.Log.Warn(fmt.Sprintf(template, args...))
}

// Warning logs a message at level Warn on the ZapLogger.
func Warning(args ...interface{}) {

	Default.Log.Warn(fmt.Sprint(args...))
}

// Error logs a message at level Error on the ZapLogger.
func Error(args ...interface{}) {
	Default.Log.Error(fmt.Sprint(args...))
}

// Errorf logs a message at level Warn on the ZapLogger.
func Errorf(template string, args ...interface{}) {

	Default.Log.Error(fmt.Sprintf(template, args...))
}

// Fatal logs a message at level Fatal on the ZapLogger.
func Fatal(args ...interface{}) {
	Default.Log.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a message at level Warn on the ZapLogger.
func Fatalf(template string, args ...interface{}) {

	Default.Log.Fatal(fmt.Sprintf(template, args...))
}

// Panic logs a message at level Painc on the ZapLogger.
func Panic(args ...interface{}) {
	Default.Log.Panic(fmt.Sprint(args...))
}

// Panicf  logs a message at level Warn on the ZapLogger.
func Panicf(template string, args ...interface{}) {

	Default.Log.Panic(fmt.Sprintf(template, args...))
}

// Printf logs a message at level Info on the ZapLogger.
func Printf(format string, args ...interface{}) {
	Default.Log.Info(fmt.Sprintf(format, args...))
}

// Print logs a message at level Info on the ZapLogger.
func Print(args ...interface{}) {
	Default.Log.Info(fmt.Sprint(args...))
}

// Println logs a message at level Info on the ZapLogger.
func Println(args ...interface{}) {
	Default.Log.Info(fmt.Sprint(args...))
}
