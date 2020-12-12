package logger

type LogWriter interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Debug(msg string)
	Info(msg string)
	Print(msg string)
	Warn(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)
	Close()
}
