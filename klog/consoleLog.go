package klog

import "io"

type ConsoleLog struct {
	logConfig *globalLog
	w io.Writer
	prefix string
}

func NewConsoleLog(w io.Writer, prefix string) *ConsoleLog {
	return &ConsoleLog{
		logConfig: newLog(w,prefix),
		w: w,
		prefix: prefix,
	}
}

func (l *ConsoleLog) Trace(args ...interface{}) {
	l.logConfig.trace(args)
}

func (l *ConsoleLog) Info(args ...interface{}) {
	l.logConfig.info(args)
}


func (l *ConsoleLog) Warn(args ...interface{}) {
	l.logConfig.warn(args)
}


func (l *ConsoleLog) Errorf(args ...interface{}) {
	l.logConfig.error(args)
}


func (l *ConsoleLog) Fatal(args ...interface{}) {
	l.logConfig.fatal(args)
}

func (l *ConsoleLog) SetPrefix(p string) {
	l.logConfig.setPrefix(p)
	l.prefix = p
}
