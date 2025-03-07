package logger

import (
	"os"

	l "github.com/charmbracelet/log"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var log *l.Logger

func Debug(msg any, keys ...any) {
	log.Debug(msg, keys...)
}

func Info(msg any, keys ...any) {
	log.Info(msg, keys...)
}

func Warn(msg any, keys ...any) {
	log.Warn(msg, keys...)
}

func Error(msg any, keys ...any) {
	log.Error(msg, keys...)
	os.Exit(1)
}

func SetLevel(level Level) {
	switch level {
	case DebugLevel:
		log.SetLevel(l.DebugLevel)
	case InfoLevel:
		log.SetLevel(l.InfoLevel)
	case WarnLevel:
		log.SetLevel(l.WarnLevel)
	case ErrorLevel:
		log.SetLevel(l.ErrorLevel)
	}
}

func init() {
	log = l.New(os.Stderr)
	log.SetReportTimestamp(false)
	log.SetReportCaller(false)

	// default is debug because config package uses some debug logs and verbose flag is loaded afterwards
	log.SetLevel(l.DebugLevel)
}
