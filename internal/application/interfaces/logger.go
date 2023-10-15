package interfaces

type Logger interface{
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Debug(args ...any)
	Debugf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Panic(args ...any)
	Panicf(format string, args ...any)
}
