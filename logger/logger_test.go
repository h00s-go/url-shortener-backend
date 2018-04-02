package logger

import (
	"log"
	"testing"

	"github.com/h00s/url-shortener-backend/logger"
)

func TestLogger(t *testing.T) {
	l, err := logger.New("logger_test.log")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	l.Debug("Debug test")
	l.Error("Error test")
	l.Info("Info test")
	l.Warning("Warning test")
}
