package logger

import (
	"log"
	"os"
	"strings"
)

var Log Logger

var (
	Debug   func(...interface{})
	Info    func(...interface{})
	Warning func(...interface{})
	Error   func(...interface{})
	Fatal   func(...interface{})
)

type Level int

const (
	NONE Level = iota
	DEBUG
	INFORMATIONAL
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	level Level
	file  *os.File
	l     *log.Logger
}

func (level Level) String() string {
	return []string{
		"NONE",
		"DEBUG",
		"INFO ",
		"WARN ",
		"ERROR",
		"FATAL",
	}[level]
}

func NewLogger(level Level, file *os.File, l *log.Logger) Logger {
	return Logger{
		level: level,
		file:  file,
		l:     l,
	}
}

func GenerateLoggerFunc(target, level Level) func(...interface{}) {
	if target < level {
		return func(...interface{}) {
		}
	}
	if target == FATAL {
		return func(msg ...interface{}) {
			Log.l.Println(target, msg)
			panic(msg)
		}
	}
	return func(msg ...interface{}) {
		Log.l.Println(target, msg)
	}
}

func Close() {
	Log.file.Close()
}

func LevelFromString(level string) Level {
	lowercased := strings.ToLower(level)
	switch {
	case strings.HasPrefix(lowercased, "deb"):
		return DEBUG
	case strings.HasPrefix(lowercased, "info"):
		return INFORMATIONAL
	case strings.HasPrefix(lowercased, "warn"):
		return WARNING
	case strings.HasPrefix(lowercased, "err"):
		return ERROR
	case strings.HasPrefix(lowercased, "fat"):
		return FATAL
	default:
		return INFORMATIONAL
	}
}
