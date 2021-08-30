package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LoggerLevel uint16

const (
	DEBUG LoggerLevel = iota
	INFO
	WARING
	ERROR
	FATAL
)

//Logger service
type Logger struct {
	level LoggerLevel
}

func parseLevel(levelSting string) LoggerLevel {
	switch levelSting {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "waring":
		return WARING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return DEBUG
	}
}

//NewConsoleLog service
func NewConsoleLog(levelSting string) *Logger {
	return &Logger{
		level: parseLevel(levelSting),
	}
}

func getLogLevel(lv LoggerLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARING:
		return "WARING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(skip int) (funName string, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("运行出错~")
	}
	fileName = path.Base(file)
	funName = runtime.FuncForPC(pc).Name()
	funName = strings.Split(fileName, ".")[1]
	return

}

func (l Logger) enable(level LoggerLevel) bool {
	return l.level <= level
}

func (l Logger) printLog(lv LoggerLevel, msg string) {
	if l.enable(lv) {
		time := time.Now().Format("2006-01-02 15:04:05")
		funName, fileName, line := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d]  这是一条%s日志\n", time, getLogLevel(lv), funName, fileName, line, msg)
	}
}

func (l Logger) Debug(msg string) {
	l.printLog(DEBUG, msg)
}
func (l Logger) Info(msg string) {
	l.printLog(INFO, msg)
}
func (l Logger) Waring(msg string) {
	l.printLog(WARING, msg)
}
func (l Logger) Error(msg string) {
	l.printLog(ERROR, msg)
}
func (l Logger) Fatal(msg string) {
	l.printLog(FATAL, msg)
}
