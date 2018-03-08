package link

import (
	"errors"
	"fmt"
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

func checkBadHost(host string) error {
	addresses, err := net.LookupIP(host)
	if err != nil {
		return err
	}
	for _, a := range addresses {
		if a.To4() != nil {
			checkHost := fmt.Sprintf("%d.%d.%d.%d.zen.spamhaus.org", a[15], a[14], a[13], a[12])
			_, err = net.LookupHost(checkHost)
			if err == nil {
				return errors.New("Host found in zen.spamhaus.org")
			}
		}
	}
	return nil
}

// checkURL verifies if it's valid url
// also doing DNS lookup if domain exists
func checkURL(urlToCheck string) error {
	u, err := url.ParseRequestURI(urlToCheck)
	if err != nil {
		return errors.New("Invalid URL")
	}

	if !u.IsAbs() {
		return errors.New("URL is not absolute")
	}

	if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "ftp" {
		return errors.New("URL does not have http(s) prefix")
	}

	_, err = net.LookupHost(u.Host)
	if err != nil {
		return errors.New("Domain doesn't exist")
	}

	err = checkBadHost(u.Host)
	if err != nil {
		return err
	}

	return nil
}
