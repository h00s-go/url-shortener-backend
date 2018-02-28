package main

import (
	"testing"

	"github.com/h00s/url-shortener-backend/config"
)

func TestConfiguration(t *testing.T) {
	_, err := config.LoadConfiguration("configuration_test.json")
	if err != nil {
		t.Error("Unable to load configuration")
	}
}
