package klog

var (
	levelOfLog  = [...]string{"Trace", "INFO", "WARN", "ERROR", "FATAL"}
	PREFIX         = "[Kirby: STA]"
)

type Level int

const (
	Trace Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger log接口
type Logger interface {
	Trace(args ...interface{})

	Info(args ...interface{})

	Warn(args ...interface{})

	Errorf(args ...interface{})

	Fatal(args ...interface{})

	SetPrefix(p string)

}