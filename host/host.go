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

	err = isBlacklisted(host)
	if err != nil {
		return err
	}

	return nil
}

// IsBlacklisted checks if host is blacklisted in surbl
// Returns nil if host is not blacklisted
// isBlacklisted also checks against whitelisting and url shorteners/redirectors
func isBlacklisted(host string) error {
	domain, _ := publicsuffix.EffectiveTLDPlusOne(host)

	if isWhitelisted(domain) {
		return nil
	}

	if isRedirector(domain) {
		return errors.New("Redirectors are not allowed")
	}

	_, err := net.LookupHost(domain + ".multi.surbl.org")
	if err == nil {
		return errors.New("Host found in surbl.org")
	}
	return nil
}

// isWhitelisted should only be called from isBlacklisted
func isWhitelisted(domain string) bool {
	switch domain {
	case
		"google.com",
		"tinyurl.com":
		return true
	}
	return false
}

// isRedirector should only be called from isBlacklisted
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
