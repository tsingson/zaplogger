// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package zapgrpc2 provides a logger that is compatible with grpclog.
package zapgrpc2 // import "go.uber.org/zap/zapgrpc2"
import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

// NewGRPCLoggerV2 converts "*zap.Logger" to "grpclog.LoggerV2".
// It discards all INFO level logging in gRPC, if debug level
// is not enabled in "*zap.Logger".
func NewGRPCLoggerV2(lcfg zap.Config) (grpclog.LoggerV2, error) {
	lg, err := lcfg.Build(zap.AddCallerSkip(1)) // to annotate caller outside of "logutil"
	if err != nil {
		return nil, err
	}
	return &zapGRPCLogger{lg: lg, sugar: lg.Sugar()}, nil
}

// NewLoggerV2 new logger v2
func NewLoggerV2(l *zap.Logger) grpclog.LoggerV2 {
	return &zapGRPCLogger{
		lg:    l,
		sugar: l.Sugar(),
	}
}

// NewGRPCLoggerV2FromZapCore creates "grpclog.LoggerV2" from "zap.Core"
// and "zapcore.WriteSyncer". It discards all INFO level logging in gRPC,
// if debug level is not enabled in "*zap.Logger".
func NewGRPCLoggerV2FromZapCore(cr zapcore.Core, syncer zapcore.WriteSyncer) grpclog.LoggerV2 {
	// "AddCallerSkip" to annotate caller outside of "logutil"
	lg := zap.New(cr, zap.AddCaller(), zap.AddCallerSkip(1), zap.ErrorOutput(syncer))
	return &zapGRPCLogger{lg: lg, sugar: lg.Sugar()}
}

type zapGRPCLogger struct {
	lg    *zap.Logger
	sugar *zap.SugaredLogger
}

func (zl *zapGRPCLogger) Info(args ...interface{}) {
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	zl.sugar.Info(args...)
}

func (zl *zapGRPCLogger) Infoln(args ...interface{}) {
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	zl.sugar.Info(args...)
}

func (zl *zapGRPCLogger) Infof(format string, args ...interface{}) {
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	zl.sugar.Infof(format, args...)
}

func (zl *zapGRPCLogger) Warning(args ...interface{}) {
	zl.sugar.Warn(args...)
}

func (zl *zapGRPCLogger) Warningln(args ...interface{}) {
	zl.sugar.Warn(args...)
}

func (zl *zapGRPCLogger) Warningf(format string, args ...interface{}) {
	zl.sugar.Warnf(format, args...)
}

func (zl *zapGRPCLogger) Error(args ...interface{}) {
	zl.sugar.Error(args...)
}

func (zl *zapGRPCLogger) Errorln(args ...interface{}) {
	zl.sugar.Error(args...)
}

func (zl *zapGRPCLogger) Errorf(format string, args ...interface{}) {
	zl.sugar.Errorf(format, args...)
}

func (zl *zapGRPCLogger) Fatal(args ...interface{}) {
	zl.sugar.Fatal(args...)
}

func (zl *zapGRPCLogger) Fatalln(args ...interface{}) {
	zl.sugar.Fatal(args...)
}

func (zl *zapGRPCLogger) Fatalf(format string, args ...interface{}) {
	zl.sugar.Fatalf(format, args...)
}

func (zl *zapGRPCLogger) V(l int) bool {
	// infoLog == 0
	if l <= 0 { // debug level, then we ignore info level in gRPC
		return !zl.lg.Core().Enabled(zapcore.DebugLevel)
	}
	return true
}
