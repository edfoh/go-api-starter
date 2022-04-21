package log

type Logger interface {
	Debug(string, ...field)
	Debugf(string, ...interface{})
	Info(string, ...field)
	Infof(string, ...interface{})
	Warn(string, ...field)
	Warnf(string, ...interface{})
	Error(string, ...field)
	Errorf(string, ...interface{})
	Fatal(string, ...field)
	Fatalf(string, ...interface{})
	Panic(string, ...field)
	Panicf(string, ...interface{})
}
