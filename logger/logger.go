package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

func New() *log.Logger {
	logger := log.New(os.Stderr)
	logger.SetReportTimestamp(false)
	logger.SetReportCaller(false)
	logger.SetLevel(log.DebugLevel)

	return logger
}
