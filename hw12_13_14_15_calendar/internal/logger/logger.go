package logger

import (
	"log"
	"os"
	"strings"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
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

func ConfigureLoggerByConfig(cfg *config.Config) {
	fd, err := os.OpenFile(cfg.Log.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	level := levelFromString(cfg.Log.Level)
	Log = Logger{
		file:  fd,
		level: level,
		l:     log.New(fd, "", log.LstdFlags),
	}
	Debug = GenerateLoggerFunc(DEBUG, level)
	Info = GenerateLoggerFunc(INFORMATIONAL, level)
	Warning = GenerateLoggerFunc(WARNING, level)
	Error = GenerateLoggerFunc(ERROR, level)
	Fatal = GenerateLoggerFunc(FATAL, level)
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

func levelFromString(level string) Level {
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
