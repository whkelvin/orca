package logger

import (
	"github.com/charmbracelet/log"
)

func Debug(msg interface{}, keyvals ...interface{}) {
	log.Debug(msg, keyvals...)
}

func Info(msg interface{}, keyvals ...interface{}) {
	log.Info(msg, keyvals...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	log.Info(msg, keyvals...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	log.Error(msg, keyvals...)
}

func Print(msg interface{}, keyvals ...interface{}) {
	log.Print(msg, keyvals...)
}
