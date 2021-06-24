package klog

var (
	levelOfLog  = [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	PREFIX         = "[Kirby: STA]"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger log接口
type Logger interface {
	Debug(args ...interface{})

	Info(args ...interface{})

	Warn(args ...interface{})

	Error(args ...interface{})

	Fatal(args ...interface{})

	// SetLevel 设置输出端日志级别
	SetLevel(level Level)
	// GetLevel 获取输出端日志级别
	GetLevel() string

	setSkip(skip int)
}