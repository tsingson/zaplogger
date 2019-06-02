package zaplogger

import (
	"os"
	"time"

	"github.com/spf13/afero"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/rs/zerolog/diode"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel log level
var LogLevel = zap.NewAtomicLevelAt(zap.DebugLevel)

// NewLogger new zap logger
func NewLogger() *zap.Logger {
	p, _ := getCurrentExecDir()
	p = p + "/log"
	path, _ := buildLogPath(p)
	return NewZapLog(path, "default", false)
}

// NewZapLog  init a log
func NewZapLog(path, prefix string, stdoutFlag bool) (log *zap.Logger) {

	if stdoutFlag {
		// opts = append(opts, zap.AddCaller())
		// opts = append(opts, zap.AddStacktrace(zap.WarnLevel))
		opts := []zap.Option{}
		std := newStdoutCore(zapcore.DebugLevel)
		debug := newZapCore(path, prefix)

		log = zap.New(zapcore.NewTee(std, debug), opts...)
	} else {

		log = zap.New(newZapCore(path, prefix))
	}
	return

}

// NewZapLog  initial a zap logger
func newZapCore(path, prefix string) zapcore.Core {

	dataTimeFmtInFileName := time.Now().Format("2006-01-02-15")
	var err error
	var logPath string

	logPath, err = buildLogPath(path)
	if err != nil {
		// TODO: handle error
	}
	var w zapcore.WriteSyncer
	var logFilename string
	if len(prefix) == 0 {
		// 	logFilename = logpath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"
		// logFilename = logPath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"

		wdiode := diode.NewWriter(os.Stdout, 1024*1024*4, 50*time.Millisecond, func(missed int) {
			// 	fmt.Printf("Logger Dropped %d messages", missed)
		})

		// lumberjack.Logger is already safe for concurrent use, so we don't need to
		// lock it.

		w = zapcore.AddSync(wdiode)
	} else {
		// 	logFilename = logpath + "/" + prefix + "-pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"
		logFilename = logPath + "/" + prefix + "-" + dataTimeFmtInFileName + ".zlog"

		var LumberLogger = &lumberjack.Logger{
			Filename:   logFilename,
			MaxSize:    100, // megabytes
			MaxBackups: 31,
			MaxAge:     31,    // days
			Compress:   false, // 开发时不压缩
		}

		wdiode := diode.NewWriter(LumberLogger, 1024*1024*4, 50*time.Millisecond, func(missed int) {
			// 	fmt.Printf("Logger Dropped %d messages", missed)
		})

		// lumberjack.Logger is already safe for concurrent use, so we don't need to
		// lock it.

		w = zapcore.AddSync(wdiode)
	}

	return newCore(true, w)

}

func newStdoutCore(level zapcore.Level) zapcore.Core {

	wdiode := diode.NewWriter(os.Stdout, 1024*1024*4, 50*time.Millisecond, func(missed int) {
		// 	fmt.Printf("Logger Dropped %d messages", missed)
	})

	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.

	var w = zapcore.AddSync(wdiode)

	return newCore(true, w)
}

// newZapLogger
func newCore(jsonFlag bool, output zapcore.WriteSyncer) zapcore.Core {

	cfg := zapcore.EncoderConfig{
		TimeKey:        "logtime",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder //
	if jsonFlag {
		encoder = zapcore.NewJSONEncoder(cfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(cfg)
	}

	return zapcore.NewCore(encoder, output, LogLevel)
}

// buildLogPath
func buildLogPath(path ...string) (logPath string, err error) {
	var p string
	if len(path[0]) == 0 {
		p, _ = getCurrentExecDir()
	} else {
		p = path[0]
	}
	logPath = p // + "/log"

	afs := afero.NewOsFs()
	check, _ := afero.DirExists(afs, logPath)
	if !check {
		err := afs.MkdirAll(logPath, 0755)
		if err != nil {
			return "", err
		}
	}

	// tf := logPath + "/test.log"
	// err = afero.WriteFile(afs, tf, []byte("file b"), 0644)
	// if err != nil {
	// 	return "", err
	// } else {
	// 	_ = afs.Remove(tf)
	// }

	return logPath, nil
}
