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

func TestURLCheck(t *testing.T) {
	validURLs := []string{"http://www.foo.com", "ftp://www.foo.com", "http://puresafesupply.ru"}
	invalidURLs := []string{"http://www.foo.kom", "htp://www.foo.com", "foo/bar", "foo", "www.foo.com", "http://vcruut.info"}

	for _, validURL := range validURLs {
		err := checkURL(validURL)
		if err != nil {
			t.Error(validURL, "is not valid", err)
		}
	}

	for _, invalidURL := range invalidURLs {
		err := checkURL(invalidURL)
		if err == nil {
			t.Error(invalidURL, "is valid", err)
		}
	}
}
