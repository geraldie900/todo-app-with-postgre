package logger

import (
	"github.com/geraldie900/todo-app/config"
	"log"
	"strconv"
)

/*
		ALL < DEBUG < INFO < WARN < ERROR < OFF
	     0      1      2      3       4      5
*/
const (
	DEBUG = 1
	INFO  = 2
	WARN  = 3
	ERROR = 4
)

var logLevel = 2

type Logger struct {
	FunctionName string
}

func InitLogging() {
	//-> default to INFO
	if len(config.AppConfig.LogLevel) != 0 {
		logLevel, _ = strconv.Atoi(config.AppConfig.LogLevel)
	}
	log.Println("LOG LEVEL =", logLevel)
}

func (l Logger) LogDebug(msg string, obj ...interface{}) {
	if DEBUG >= logLevel {
		log.Printf("| DEBUG | %s | %s :: %v", l.FunctionName, msg, obj)
	}
}

func (l Logger) LogInfo(msg string, obj ...interface{}) {
	if INFO >= logLevel {
		log.Printf("| INFO | %s | %s :: %v", l.FunctionName, msg, obj)
	}
}

func (l Logger) LogWarn(msg string, obj ...interface{}) {
	if WARN >= logLevel {
		log.Printf("| WARN | %s | %s :: %v", l.FunctionName, msg, obj)
	}
}

func (l Logger) LogError(msg string, obj ...interface{}) {
	if ERROR >= logLevel {
		log.Printf("| ERROR | %s | %s :: %v", l.FunctionName, msg, obj)
	}
}
