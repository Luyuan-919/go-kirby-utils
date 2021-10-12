package klog


import (
"io"
"log"
"os"
)

var (
	defWriter = os.Stdout
	defPrefix = "[kirby:STA] "
	GlobalLog *globalLog
)

type globalLog struct {
	Trace *log.Logger
	Info *log.Logger
	Warn *log.Logger
	Error *log.Logger
	Fatal *log.Logger
}

func init() {
	GlobalLog = newLog(defWriter,defPrefix)
}


func newLog(w io.Writer,p string) *globalLog{
	if w == nil{w = defWriter}
	if p == ""{p = defPrefix}
	l := &globalLog{}
	l.Trace = log.New(w,p,log.Ldate | log.Ltime )
	l.Info = log.New(w,p,log.Ldate | log.Ltime )
	l.Warn = log.New(w,p,log.Ldate | log.Ltime )
	l.Error = log.New(w,p,log.Ldate | log.Ltime)
	l.Fatal = log.New(w,p,log.Ldate | log.Ltime )
	return l
}

func (l *globalLog) trace(args ...interface{})  {
	l.Trace.Println(args)
}

func (l *globalLog) info(args ...interface{})  {
	l.Info.Println(args)
}

func (l *globalLog) warn(args ...interface{})  {
	l.Warn.Println(args)
}

func (l *globalLog) error(args ...interface{})  {
	l.Error.Println(args)
}

func (l *globalLog) fatal(args ...interface{})  {
	l.Fatal.Println(args)
	panic(args)
}

func (l *globalLog) setPrefix(p string) {
	l.Trace.SetPrefix(p)
	l.Info.SetPrefix(p)
	l.Warn.SetPrefix(p)
	l.Error.SetPrefix(p)
	l.Fatal.SetPrefix(p)
}

func Tracef(args ...interface{}) {
	GlobalLog.Trace.Println(args)
}

func Info(args ...interface{}) {
	GlobalLog.Info.Println(args)
}

func Warn(args ...interface{}) {
	GlobalLog.Warn.Println(args)
}

func Error(args ...interface{}) {
	GlobalLog.Error.Println(args)
}

func Fatal(args ...interface{}) {
	GlobalLog.Fatal.Println(args)
}


