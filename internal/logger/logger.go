package logger

import (
	"os"

	charmLogger "github.com/charmbracelet/log"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var log *charmLogger.Logger

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
		log.SetLevel(charmLogger.DebugLevel)
	case InfoLevel:
		log.SetLevel(charmLogger.InfoLevel)
	case WarnLevel:
		log.SetLevel(charmLogger.WarnLevel)
	case ErrorLevel:
		log.SetLevel(charmLogger.ErrorLevel)
	}
}

func init() {
	log = charmLogger.New(os.Stderr)
	log.SetReportTimestamp(false)
	log.SetReportCaller(false)
	log.SetLevel(charmLogger.InfoLevel)
}
