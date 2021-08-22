package klog

import (
	"fmt"
	"io"
	"os"
	"runtime"

	tm "github.com/sta-golang/go-lib-utils/time"
)

var (
	console io.Writer = os.Stdout
)

type consoleLog struct {
	prefix   string
	level    Level
	skip     int
}

func NewConsoleLog(level Level, prefix string) Logger {
	if prefix == "" {
		prefix = PREFIX
	}
	return &consoleLog{
		level:    level,
		prefix:   prefix,
		skip:     defStep,
	}
}

func (cl *consoleLog) setSkip(skip int) {
	cl.skip = skip
}


func (cl *consoleLog) print(level Level, format string, args ...interface{}) {
	if level < cl.level {
		return
	}
	var logFormat = "%s %s [%s] %s => %s\n"
	_, transFile, transLine, _ := runtime.Caller(cl.skip)
	_, _ = fmt.Fprintf(console, logFormat, cl.prefix, tm.GetNowDateTimeStr(),
		levelOfLog[level], fmt.Sprintf("%s:%d", transFile, transLine), fmt.Sprintf(format, args...))
}


func (cl *consoleLog) println(level Level, args ...interface{}) {
	if level < cl.level {
		return
	}
	var logFormat = "%s %s [%s] %s => "
	_, transFile, transLine, _ := runtime.Caller(cl.skip)
	_, _ = fmt.Fprintf(console, fmt.Sprintf("%s%s\n", fmt.Sprintf(logFormat, cl.prefix, tm.GetNowDateTimeStr(),
		levelOfLog[level], fmt.Sprintf("%s:%d", transFile, transLine)), fmt.Sprint(args...)))
}


func (cl *consoleLog) SetLevel(level Level) {
	if level < DEBUG || level > FATAL {
		return
	}
	cl.level = level
}

// GetLevel 获取输出端日志级别
func (cl *consoleLog) GetLevel() string {
	return levelOfLog[cl.level]
}

func (cl *consoleLog) Debug(args ...interface{}) {
	cl.println(DEBUG, args...)
}


func (cl *consoleLog) Warn(args ...interface{}) {
	cl.println(WARNING, args...)
}


func (cl *consoleLog) Info(args ...interface{}) {
	cl.println(INFO, args...)
}


func (cl *consoleLog) Error(args ...interface{}) {
	cl.println(ERROR, args...)
}


func (cl *consoleLog) Fatal(args ...interface{}) {
	cl.println(FATAL, args...)
}