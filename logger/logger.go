package logger

import (
	"errors"
	"log"
	"os"
)

// Log contains Logger
type Log struct {
	log *log.Logger
	f   *os.File
}

// New returns new Log
func New(filename string) (*Log, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, errors.New("Error creating file for logging")
	}

	l := &Log{log.New(f, "", log.Lshortfile), f}
	return l, nil
}

// Close closes file handle
func (l *Log) Close() error {
	return l.f.Close()
}

// Error prints error message to log file
func (l *Log) Error(text string) {
	l.log.SetPrefix("ERROR: ")
	l.log.Println(text)
}

// Fatal prints fatal message to log file
func (l *Log) Fatal(text string) {
	l.log.SetPrefix("FATAL: ")
	l.log.Fatalln(text)
}

// Info prints informational message to log file
func (l *Log) Info(text string) {
	l.log.SetPrefix("INFO: ")
	l.log.Println(text)
}

// Warning prints warning message to log file
func (l *Log) Warning(text string) {
	l.log.SetPrefix("WARNING: ")
	l.log.Println(text)
}
