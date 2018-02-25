package main

import (
	"testing"

	"github.com/h00s/url-shortener-backend/link"
)

func TestGetLinkNameFromID(t *testing.T) {
	name := link.GetNameFromID(1)
	if name != "B" {
		t.Error("Expected B, got ", name)
	}
	name = link.GetNameFromID(50)
	if name != "BA" {
		t.Error("Expected BA, got ", name)
	}
	name = link.GetNameFromID(52)
	if name != "BC" {
		t.Error("Expected AB, got ", name)
	}
	name = link.GetNameFromID(2500)
	if name != "BAA" {
		t.Error("Expected BAA, got ", name)
	}
	id := link.GetIDFromName("BAA")
	if id != 2500 {
		t.Error("Expected 2500, got ", id)
	}
}
