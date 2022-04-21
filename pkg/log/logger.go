package log

import "go.uber.org/zap"

type logger struct {
	zapLogger *zap.SugaredLogger
}

func (l *logger) Debug(msg string, fields ...field) {
	l.zapLogger.Debugw(msg, l.mergeKeyValues(fields...))
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.zapLogger.Debugf(template, args)
}

func (l *logger) Info(msg string, fields ...field) {
	l.zapLogger.Infow(msg, l.mergeKeyValues(fields...))
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.zapLogger.Infof(template, args)
}

func (l *logger) Warn(msg string, fields ...field) {
	l.zapLogger.Warnw(msg, l.mergeKeyValues(fields...))
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.zapLogger.Warnf(template, args)
}

func (l *logger) Error(msg string, fields ...field) {
	l.zapLogger.Errorw(msg, l.mergeKeyValues(fields...))
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.zapLogger.Errorf(template, args)
}

func (l *logger) Fatal(msg string, fields ...field) {
	l.zapLogger.Fatalw(msg, l.mergeKeyValues(fields...))
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.zapLogger.Fatalf(template, args)
}

func (l *logger) Panic(msg string, fields ...field) {
	l.zapLogger.Panicw(msg, l.mergeKeyValues(fields...))
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.zapLogger.Panicf(template, args)
}

func (l *logger) mergeKeyValues(fs ...field) []interface{} {
	var res fields
	for _, f := range fs {
		res = append(res, f)
	}
	return res.ToSlice()
}
