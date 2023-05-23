package logger

import (
	"fmt"
	"log"
)

const (
	FATAL = iota
	ERROR
	WARNING
	INFO
	DEBUG
)
const (
	_FATAL   = "FATAL"
	_ERROR   = "ERROR"
	_WARNING = "WARNING"
	_INFO    = "INFO"
	_DEBUG   = "DEBUG"
)

type Logger struct {
	level  int
	on     bool
	name   string
	logger *log.Logger
}

func GetLogger(name string, level int) *Logger {
	lg := log.Default()
	return &Logger{name: name, level: level, on: true, logger: lg}
}

func (log *Logger) Log(args ...any) {
	if log.on {
		log.SetPrefix(_INFO)
		log.logger.Println(args...)
	}
}

func (log *Logger) Debug(args ...any) {
	if log.on && log.level >= DEBUG {
		log.SetPrefix(_DEBUG)
		log.logger.Println(args...)
	}
}

func (log *Logger) Info(args ...any) {
	if log.on && log.level >= INFO {
		log.SetPrefix(_INFO)
		log.logger.Println(args...)
	}
}

func (log *Logger) Error(args ...any) {
	if log.on && log.level >= ERROR {
		log.SetPrefix(_ERROR)
		log.logger.Println(args...)
	}
}

func (log *Logger) WARNING(args ...any) {
	if log.on && log.level >= WARNING {
		log.SetPrefix(_WARNING)
		log.logger.Println(args...)
	}
}

func (log *Logger) Fatal(args ...any) {
	if log.on && log.level >= FATAL {
		log.SetPrefix(_FATAL)
		log.logger.Println(args...)
	}
}

func (log *Logger) SetFlag(flag int) {
	log.logger.SetFlags(flag)
}

func (log *Logger) SetPrefix(prefix string) {
	log.logger.SetPrefix(fmt.Sprintf("[%s-%s]", log.name, prefix))
}
