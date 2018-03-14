package host

import (
	"errors"
	"net"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// IsValid verifies if it's valid url
// also doing DNS lookup if domain exists
// if domain exists, checking for whitelisting and blacklisting
func IsValid(uri string) error {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return errors.New("Invalid URL")
	}

	if !u.IsAbs() {
		return errors.New("URL is not absolute")
	}

	if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "ftp" {
		return errors.New("URL does not have http(s) prefix")
	}

	host := strings.ToLower(u.Host)

	_, err = net.LookupHost(host)
	if err != nil {
		return errors.New("Host doesn't exist")
	}

	domain, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return errors.New("Error while getting domain")
	}

	if isWhitelisted(domain) {
		return nil
	}
	if isRedirector(domain) {
		return errors.New("Domain found in redirectors list")
	}
	if isBlacklisted(domain) {
		return errors.New("Domain is blacklisted")
	}

	return nil
}

func isWhitelisted(domain string) bool {
	switch domain {
	case
		"google.com",
		"yahoo.com":
		return true
	}
	return false
}

func isRedirector(domain string) bool {
	switch domain {
	case
		"adf.ly",
		"bc.vc",
		"bit.do",
		"bit.ly",
		"budurl.com",
		"buff.ly",
		"clicky.me",
		"goo.gl",
		"is.gd",
		"mcaf.ee",
		"ow.ly",
		"s2r.co",
		"soo.gd",
		"short.to",
		"tiny.cc",
		"tinyurl.com":
		return true
	}
	return false
}

// IsBlacklisted checks if host is blacklisted in SURBL
func isBlacklisted(domain string) bool {
	_, err := net.LookupHost(domain + ".multi.surbl.org")
	if err == nil {
		return true
	}
	return false
}
