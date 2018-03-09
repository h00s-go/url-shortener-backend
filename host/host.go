package host

import (
	"errors"
	"fmt"
	"net"
	"net/url"
)

// IsBlacklisted checks if host is blacklisted in surbl
func IsBlacklisted(host string) error {
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
