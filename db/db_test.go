package db

import (
	"testing"
)

func TestDB(t *testing.T) {
	_, err := NewDatabase("configuration_test.json")
	if err != nil {
		t.Error("Unable to connect to DB")
	}
}
