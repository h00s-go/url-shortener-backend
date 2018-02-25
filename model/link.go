package model

import (
	"strings"
)

const validChars = "ABCDEFHJKLMNPRSTUVXYZabcdefgijkmnprstuvxyz23456789"

// Link represent one shortened link
type Link struct {
	URL string `json:"url"`
}

// GetLinkNameFromID gets name from numerical ID
func GetLinkNameFromID(id int) string {
	name := ""
	for id > 0 {
		name = string(validChars[id%50]) + name
		id = id / 50
	}
	return name
}

// GetLinkIDFromName gets ID from name
func GetLinkIDFromName(name string) int {
	id := 0
	for i := 0; i < len(name); i++ {
		id = 50*id + (strings.Index(validChars, string(name[i])))
	}
	return id
}
