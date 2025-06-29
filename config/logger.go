package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.MultiWriter(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debug:  log.New(writer, "DEBUG: ", logger.Flags()),
		info:   log.New(writer, "INFO: ", logger.Flags()),
		warn:   log.New(writer, "WARN: ", logger.Flags()),
		err:    log.New(writer, "ERROR: ", logger.Flags()),
		writer: writer,
	}
}

// Create non-formated logs
func (logger *Logger) Debug(v ...interface{}) {
	logger.debug.Println(v)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.info.Println(v)
}

func (logger *Logger) Warn(v ...interface{}) {
	logger.warn.Println(v)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.err.Println(v)
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.debug.Printf(format, v)
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.info.Printf(format, v)
}

func (logger *Logger) Warnf(format string, v ...interface{}) {
	logger.warn.Printf(format, v)
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.err.Printf(format, v)
}
