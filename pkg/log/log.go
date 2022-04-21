package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

type Level int8

func (level Level) toZapLogLevel() zapcore.Level {
	switch level {
	case DebugLevel:
		return zap.DebugLevel
	case InfoLevel:
		return zap.InfoLevel
	case WarnLevel:
		return zap.WarnLevel
	case ErrorLevel:
		return zap.ErrorLevel
	case FatalLevel:
		return zap.FatalLevel
	case PanicLevel:
		return zap.PanicLevel
	}
	return zap.DebugLevel
}

func New(logLevel Level) (Logger, error) {
	c := zap.NewProductionConfig()
	c.Level.SetLevel(logLevel.toZapLogLevel())
	zapLogger, err := c.Build()
	if err != nil {
		return nil, err
	}
	defer zapLogger.Sync()
	sugaredLogger := zapLogger.Sugar()
	return &logger{sugaredLogger}, nil
}

func Must(logger Logger, err error) Logger {
	if err != nil {
		panic(err)
	}
	return logger
}
