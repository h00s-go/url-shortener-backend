package link

import (
	"net"
	"net/url"
	"strings"
)

// Do not change this after the links are inserted in database
// It will break getting and inserting new links
const validChars = "ABCDEFHJKLMNPRSTUVXYZabcdefgijkmnprstuvxyz23456789"

// Link represent one shortened link
type Link struct {
	URL string `json:"url"`
}

// getNameFromID gets name from numerical ID
func getNameFromID(id int) string {
	name := ""
	for id > 0 {
		name = string(validChars[id%len(validChars)]) + name
		id = id / len(validChars)
	}
	return name
}

// getIDFromName gets ID from name
func getIDFromName(name string) int {
	id := 0
	for i := 0; i < len(name); i++ {
		id = len(validChars)*id + (strings.Index(validChars, string(name[i])))
	}
	return id
}

// checkURL verifies if it's valid url
// also doing DNS lookup if domain exists
func checkURL(urlToCheck string) bool {
	u, err := url.ParseRequestURI(urlToCheck)
	if err != nil {
		return false
	}
	if u.IsAbs() && (u.Scheme == "http" || u.Scheme == "https" || u.Scheme == "ftp") {
		_, err = net.LookupHost(u.Host)
		if err == nil {
			return true
		}
	}
	return false
}
