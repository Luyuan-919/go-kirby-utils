package klog

import "os"

var globalLogger Logger

const (
	defStep    = 2
	globalSkip = 3
)

func init() {
	globalLogger = NewConsoleLog(DEBUG, PREFIX)
	globalLogger.setSkip(globalSkip)
}

func SetGlobalLogger(logger Logger) {
	globalLogger = logger
	globalLogger.setSkip(globalSkip)
}

func SetPrefix(str string) {
	PREFIX = str
}

func SetLevel(level Level) {
	globalLogger.SetLevel(level)
}

// GetLevel 获取输出端日志级别
func GetLevel() string {
	return globalLogger.GetLevel()
}


func Debug(args ...interface{}) {
	globalLogger.Debug(args...)
}


func Warn(args ...interface{}) {
	globalLogger.Warn(args...)
}


func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	globalLogger.Fatal(args...)
	os.Exit(0)
}