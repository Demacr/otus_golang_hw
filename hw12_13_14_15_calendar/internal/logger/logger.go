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

var StringLevels []string = []string{
	"NONE",
	"DEBUG",
	"INFO ",
	"WARN ",
	"ERROR",
	"FATAL",
}

func (level Level) String() string {
	return StringLevels[level]
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
	switch lowercased {
	case "debug":
		return DEBUG
	case "info":
		return INFORMATIONAL
	case "warn":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return INFORMATIONAL
	}
}
