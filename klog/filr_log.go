package klog

import (
	"fmt"
	tm "github.com/sta-golang/go-lib-utils/time"
	"os"
	"runtime"
)

const (
	defDir = "./log"
	defPrefix = "[Kirby:STA]"
	defFileName = "./log/LogAll.log"
)

var globalFile *os.File

func defInit() {
	_ = os.Mkdir(defDir, os.ModePerm)
	file,_ := os.OpenFile(defFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	globalFile = file
}

func InitWithDir(Dir string) {
	_ = os.Mkdir(Dir, os.ModePerm)
	file,_ := os.OpenFile(defFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	globalFile = file
}

type FileLog struct {
	Dir string
	prefix string
}

func NewFileLog() *FileLog {
	defInit()
	return &FileLog{
		Dir: defDir,
		prefix: defPrefix,
	}
}

func NewFileLogWithDir(Dir string,prefix string) *FileLog {
	InitWithDir(Dir)
	return &FileLog{
		Dir: Dir,
		prefix: prefix,
	}
}

func (f *FileLog) println(level Level, args ...interface{}) {
	var logFormat = "%s %s [%s] %s => "
	_, transFile, transLine, _ := runtime.Caller(defStep)
	_, _ = fmt.Fprintf(globalFile, fmt.Sprintf("%s%s\n", fmt.Sprintf(logFormat, f.prefix, tm.GetNowDateTimeStr(),
		levelOfLog[level], fmt.Sprintf("%s:%d", transFile, transLine)), fmt.Sprint(args...)))
}

func (f *FileLog) DeBugf(args...interface{})  {
	f.println(DEBUG,args...)
	Debug(args...)
}

func (f *FileLog) Warnf(args...interface{})  {
	f.println(WARNING,args...)
	Warn(args...)
}

func (f *FileLog) Infof(args...interface{})  {
	f.println(INFO,args...)
	Info(args...)
}

func (f *FileLog) Errorf(args...interface{})  {
	f.println(ERROR,args...)
	Error(args...)
}

func (f *FileLog) Fatalf(args...interface{})  {
	f.println(FATAL,args...)
	Fatal(args...)
}

