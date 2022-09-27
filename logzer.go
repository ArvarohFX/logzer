package logzer

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sugaredLogger struct {
	*zap.SugaredLogger
}

func New(debug bool) *sugaredLogger {
	loggingLevel := zap.InfoLevel
	if debug {
		loggingLevel = zap.DebugLevel
	}

	cfgEncoder := zap.NewProductionEncoderConfig()
	cfgProd := zap.NewProductionConfig()

	cfgEncoder.TimeKey = "ts"
	cfgEncoder.MessageKey = "message"
	cfgEncoder.LevelKey = "level"
	cfgEncoder.EncodeTime = zapcore.ISO8601TimeEncoder
	cfgProd.Level = zap.NewAtomicLevelAt(loggingLevel)
	cfgProd.EncoderConfig = cfgEncoder
	cfgProd.OutputPaths = []string{"stdout"}

	logger, _ := cfgProd.Build()
	if logger != nil {
		defer func() {
			_ = logger.Sync()
		}()
	}
	return &sugaredLogger{logger.Sugar()}
}

func (l *sugaredLogger) Error(message string, kvs ...interface{}) {
	l.Errorw(message, kvs...)
}

func (l *sugaredLogger) Warn(message string, kvs ...interface{}) {
	l.Warnw(message, kvs...)
}

func (l *sugaredLogger) Info(message string, kvs ...interface{}) {
	l.Infow(message, kvs...)
}

func (l *sugaredLogger) Debug(message string, kvs ...interface{}) {
	l.Debugw(message, kvs...)
}

func (l *sugaredLogger) Fatal(message string, kvs ...interface{}) {
	l.Fatalw(message, kvs...)
}
