package main

import (
	"testing"

	"github.com/h00s/url-shortener-backend/model"
)

func TestGetLinkNameFromID(t *testing.T) {
	name := model.GetLinkNameFromID(1)
	if name != "B" {
		t.Error("Expected B, got ", name)
	}
	name = model.GetLinkNameFromID(50)
	if name != "BA" {
		t.Error("Expected BA, got ", name)
	}
	name = model.GetLinkNameFromID(52)
	if name != "BC" {
		t.Error("Expected AB, got ", name)
	}
	name = model.GetLinkNameFromID(2500)
	if name != "BAA" {
		t.Error("Expected BAA, got ", name)
	}
	id := model.GetLinkIDFromName("BAA")
	if id != 2500 {
		t.Error("Expected 2500, got ", id)
	}
}
