package link

import (
	"testing"
)

func TestLinkNames(t *testing.T) {
	name := getNameFromID(1)
	if name != "B" {
		t.Error("Expected B, got ", name)
	}
	name = getNameFromID(50)
	if name != "BA" {
		t.Error("Expected BA, got ", name)
	}
	name = getNameFromID(52)
	if name != "BC" {
		t.Error("Expected AB, got ", name)
	}
	name = getNameFromID(2500)
	if name != "BAA" {
		t.Error("Expected BAA, got ", name)
	}
	id := getIDFromName("BAA")
	if id != 2500 {
		t.Error("Expected 2500, got ", id)
	}
}
