package link

import (
	"strings"
)

// Do not change this after the links are inserted in database
// It will break getting and inserting new links
const validChars = "ABCDEFHJKLMNPRSTUVXYZabcdefgijkmnprstuvxyz23456789"

// Link represent one shortened link
type Link struct {
	URL string `json:"url"`
}

// GetNameFromID gets name from numerical ID
func GetNameFromID(id int) string {
	name := ""
	for id > 0 {
		name = string(validChars[id%len(validChars)]) + name
		id = id / len(validChars)
	}
	return name
}

// GetIDFromName gets ID from name
func GetIDFromName(name string) int {
	id := 0
	for i := 0; i < len(name); i++ {
		id = len(validChars)*id + (strings.Index(validChars, string(name[i])))
	}
	return id
}