package config

import (
	"testing"
)

func TestConfiguration(t *testing.T) {
	_, err := Load("../configuration_test.toml")
	if err != nil {
		t.Error("Unable to load configuration")
	}
}
