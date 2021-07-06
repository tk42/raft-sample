package main

import (
	"fmt"

	"github.com/lni/dragonboat/v3/logger"
	stackdriver "github.com/tommy351/zap-stackdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger.ILogger
	log *zap.Logger
	cfg zap.Config
}

func GetLogger(pkgName string) *Logger {
	logger := Logger{}
	level := zap.NewAtomicLevel()
	cfg := zap.Config{
		Level:            level,
		Encoding:         "json",
		EncoderConfig:    stackdriver.EncoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"Name": pkgName,
		},
	}

	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	logger.log = log
	logger.cfg = cfg
	return &logger
}

func (l *Logger) SetLevel(level logger.LogLevel) {
	switch level {
	case logger.CRITICAL:
		l.cfg.Level.SetLevel(zapcore.FatalLevel)
	case logger.ERROR:
		l.cfg.Level.SetLevel(zapcore.ErrorLevel)
	case logger.WARNING:
		l.cfg.Level.SetLevel(zapcore.WarnLevel)
	case logger.INFO:
		l.cfg.Level.SetLevel(zapcore.InfoLevel)
	case logger.DEBUG:
		l.cfg.Level.SetLevel(zapcore.DebugLevel)
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.log.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.log.Fatal(fmt.Sprintf(format, args...))
}
