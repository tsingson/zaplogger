package zaplogger

import (
	"fmt"

	"github.com/spf13/afero"
)

var defaultzap *ZapLogger

func init() {
	var err error
	var afs = afero.NewOsFs()
	path, _ := getCurrentExecDir()
	LogFullPath := path + "/log"
	check, _ := afero.DirExists(afs, LogFullPath)
	if !check {
		err = afs.MkdirAll(LogFullPath, 0755)
		if err != nil {
			panic("mkdir log path fail")
		}
	}

	log := NewZapLog(LogFullPath, "debug-", true)
	defaultzap = &ZapLogger{
		Log: log,
	}

}

// Debug logs a message at level Debug on the ZapLogger.
func Debug(args ...interface{}) {
	defaultzap.Log.Debug(fmt.Sprint(args...))
}

// Debugf logs a message at level Debug on the ZapLogger.
func Debugf(template string, args ...interface{}) {
	defaultzap.Log.Debug(fmt.Sprintf(template, args...))
}

// Info logs a message at level Info on the ZapLogger.
func Info(args ...interface{}) {
	defaultzap.Log.Info(fmt.Sprint(args...))
}

// Infof logs a message at level Info on the ZapLogger.
func Infof(template string, args ...interface{}) {

	defaultzap.Log.Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at level Warn on the ZapLogger.
func Warn(args ...interface{}) {
	defaultzap.Log.Warn(fmt.Sprint(args...))
}

// Warnf logs a message at level Warn on the ZapLogger.
func Warnf(template string, args ...interface{}) {

	defaultzap.Log.Warn(fmt.Sprintf(template, args...))
}

// Warningf logs a message at level Warn on the ZapLogger.
func Warningf(template string, args ...interface{}) {

	defaultzap.Log.Warn(fmt.Sprintf(template, args...))
}

// Warning logs a message at level Warn on the ZapLogger.
func Warning(args ...interface{}) {

	defaultzap.Log.Warn(fmt.Sprint(args...))
}

// Error logs a message at level Error on the ZapLogger.
func Error(args ...interface{}) {
	defaultzap.Log.Error(fmt.Sprint(args...))
}

// Errorf logs a message at level Warn on the ZapLogger.
func Errorf(template string, args ...interface{}) {

	defaultzap.Log.Error(fmt.Sprintf(template, args...))
}

// Fatal logs a message at level Fatal on the ZapLogger.
func Fatal(args ...interface{}) {
	defaultzap.Log.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a message at level Warn on the ZapLogger.
func Fatalf(template string, args ...interface{}) {

	defaultzap.Log.Fatal(fmt.Sprintf(template, args...))
}

// Panic logs a message at level Painc on the ZapLogger.
func Panic(args ...interface{}) {
	defaultzap.Log.Panic(fmt.Sprint(args...))
}

// Panicf  logs a message at level Warn on the ZapLogger.
func Panicf(template string, args ...interface{}) {

	defaultzap.Log.Panic(fmt.Sprintf(template, args...))
}

// Printf logs a message at level Info on the ZapLogger.
func Printf(format string, args ...interface{}) {
	defaultzap.Log.Info(fmt.Sprintf(format, args...))
}

// Print logs a message at level Info on the ZapLogger.
func Print(args ...interface{}) {
	defaultzap.Log.Info(fmt.Sprint(args...))
}
