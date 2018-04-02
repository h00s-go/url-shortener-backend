package logger

import (
	"errors"
	"log"
	"os"
)

// Logger contains Logger
type Logger struct {
	log *log.Logger
	f   *os.File
}

// New returns new Log
func New(filename string) (*Logger, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, errors.New("Error creating file for logging")
	}

	l := &Logger{log.New(f, "", log.Ldate|log.Ltime), f}
	return l, nil
}

// Close closes file handle
func (l *Logger) Close() error {
	return l.f.Close()
}

// Debug prints debug message to log file
func (l *Logger) Debug(text string) {
	l.log.SetPrefix("[DEBU] ")
	l.log.Println(text)
}

// Error prints error message to log file
func (l *Logger) Error(text string) {
	l.log.SetPrefix("[ERRO] ")
	l.log.Println(text)
}

// Fatal prints fatal message to log file
func (l *Logger) Fatal(text string) {
	l.log.SetPrefix("[ERRO] ")
	l.log.Fatalln(text)
}

// Info prints informational message to log file
func (l *Logger) Info(text string) {
	l.log.SetPrefix("[INFO] ")
	l.log.Println(text)
}

// Warning prints warning message to log file
func (l *Logger) Warning(text string) {
	l.log.SetPrefix("[WARN] ")
	l.log.Println(text)
}
