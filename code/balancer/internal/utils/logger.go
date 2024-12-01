package utils

import (
	"log"
	"os"
)

// Logger предоставляет унифицированное логирование.
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLogger создает новый экземпляр логгера.
func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info выводит информационное сообщение.
func (l *Logger) Info(msg string) {
	l.infoLogger.Println(msg)
}

// Error выводит сообщение об ошибке.
func (l *Logger) Error(msg string) {
	l.errorLogger.Println(msg)
}

// Fatal выводит сообщение об ошибке и завершает программу.
func (l *Logger) Fatal(msg string) {
	l.errorLogger.Fatal(msg)
}
