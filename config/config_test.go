package config

import (
	"testing"
)

func TestConfiguration(t *testing.T) {
	_, err := LoadConfiguration("../configuration_test.json")
	if err != nil {
		t.Error("Unable to load configuration")
	}
}
