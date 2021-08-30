package logger

import (
	"fmt"
	"os"
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
	level        LoggerLevel
	filePath     string
	fileName     string
	fileObj      *os.File
	errorFileObj *os.File
	maxSize      int64
	errorFile    bool
}

func parseLevel(levelSting string) LoggerLevel {
	switch levelSting {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARING":
		return WARING
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return DEBUG
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

//NewFileLog service
func NewFileLog(filePath, fileName, levelSting string, maxSize int64, errorFile bool) *Logger {
	f := &Logger{
		filePath:  filePath,
		fileName:  fileName,
		maxSize:   maxSize,
		errorFile: errorFile,
		level:     parseLevel(levelSting),
	}
	err := f.initFile()
	if err != nil {
		panic(err)
	}
	return f
}

func (l *Logger) initFile() error {
	fullFileName := path.Join(l.filePath, l.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	ErrorfullFileName := path.Join(l.filePath, "error."+l.fileName)
	errorFileObj, err := os.OpenFile(ErrorfullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	l.fileObj = fileObj
	l.errorFileObj = errorFileObj
	return nil
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

func (l Logger) enable(lv LoggerLevel) bool {
	return l.level <= lv
}

func (l Logger) printLog(lv LoggerLevel, msg string) {
	if l.enable(lv) {
		time := time.Now().Format("2006-01-02 15:04:05")
		funName, fileName, line := getInfo(3)
		if l.errorFile && lv == ERROR {
			fmt.Fprintf(l.errorFileObj, "[%s] [%s] [%s:%s:%d]  这是一条%s日志\n", time, getLogLevel(lv), funName, fileName, line, msg)
		} else {
			fmt.Fprintf(l.fileObj, "[%s] [%s] [%s:%s:%d]  这是一条%s日志\n", time, getLogLevel(lv), funName, fileName, line, msg)
		}

	}
}

func (l *Logger) Debug(msg string) {
	l.printLog(DEBUG, msg)
}
func (l *Logger) Info(msg string) {
	l.printLog(INFO, msg)
}
func (l *Logger) Waring(msg string) {
	l.printLog(WARING, msg)
}
func (l *Logger) Error(msg string) {
	l.printLog(ERROR, msg)
}
func (l *Logger) Fatal(msg string) {
	l.printLog(FATAL, msg)
}

func (l *Logger) CloseFile() {
	l.fileObj.Close()
	l.errorFileObj.Close()
}
