package log

import (
	"os"

	"github.com/charmbracelet/log"
)

type Logger struct {
	log *log.Logger
}

func (l Logger) Debug(msg any, keyvals ...any) {
	l.log.Debug(msg, keyvals...)
}

func (l Logger) Info(msg any, keyvals ...any) {
	l.log.Info(msg, keyvals...)
}

func (l Logger) Warn(msg any, keyvals ...any) {
	l.log.Warn(msg, keyvals...)
}

func (l Logger) Error(msg any, keyvals ...any) {
	l.log.Error(msg, keyvals...)
	os.Exit(1)
}

func New() Logger {
	logger := log.New(os.Stderr)
	logger.SetReportTimestamp(false)
	logger.SetReportCaller(false)
	logger.SetLevel(log.DebugLevel)

	return Logger{
		log: logger,
	}
}
