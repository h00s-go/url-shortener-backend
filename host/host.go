package host

import (
	"errors"
	"net"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

// IsBlacklisted checks if host is blacklisted in surbl
func IsBlacklisted(host string) error {
	domain, _ := publicsuffix.EffectiveTLDPlusOne(host)
	_, err := net.LookupHost(domain + ".multi.surbl.org")
	if err == nil {
		return errors.New("Host found in surbl.org")
	}
	return nil
}

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

	_, err = net.LookupHost(u.Host)
	if err != nil {
		return errors.New("Domain doesn't exist")
	}

	err = IsBlacklisted(u.Host)
	if err != nil {
		return err
	}

	return nil
}
