package logger

import (
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"

	charmLogger "github.com/charmbracelet/log"
)

var log *charmLogger.Logger

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (level Level) string() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	}

	return ""
}

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
	styles := charmLogger.DefaultStyles()
	styles.Levels[charmLogger.DebugLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(DebugLevel.string())).
		Bold(true).
		Padding(0, 1, 0, 1).
		Width(7).
		Foreground(lipgloss.Color("#5f5fff"))
	styles.Levels[charmLogger.InfoLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(InfoLevel.string())).
		Bold(true).
		Padding(0, 1, 0, 1).
		Width(7).
		Foreground(lipgloss.Color("#008000"))
	styles.Levels[charmLogger.WarnLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(WarnLevel.string())).
		Bold(true).
		Padding(0, 1, 0, 1).
		Width(7).
		Foreground(lipgloss.Color("#ff8800"))
	styles.Levels[charmLogger.ErrorLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(ErrorLevel.string())).
		Bold(true).
		Padding(0, 1, 0, 1).
		Width(7).
		Foreground(lipgloss.Color("#ff0000"))

	log = charmLogger.New(os.Stderr)
	log.SetStyles(styles)
	log.SetLevel(charmLogger.InfoLevel)
}
