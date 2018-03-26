package link

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/h00s/url-shortener-backend/host"
)

// Link represent one shortened link
type Link struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	ClientAddress string `json:"clientAddress"`
	CreatedAt     string `json:"createdAt"`
}

func getLinkByID(c *Controller, id int) (*Link, error) {
	return getLink(c, sqlGetLinkByID, fmt.Sprint(id))
}

func getLinkByName(c *Controller, name string) (*Link, error) {
	return getLink(c, sqlGetLinkByName, name)
}

func getLinkByURL(c *Controller, url string) (*Link, error) {
	return getLink(c, sqlGetLinkByURL, url)
}

func getLink(c *Controller, query string, param string) (*Link, error) {
	l := &Link{}

	err := c.db.Conn.QueryRow(query, param).Scan(&l.ID, &l.Name, &l.URL, &l.ClientAddress, &l.CreatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.New("Error while getting link (" + err.Error() + ")")
	}
	return l, nil
}

// InsertLink in db. If inserted, return Link struct
func insertLink(c *Controller, url string, clientAddress string) (*Link, error) {
	if isSpammer(c, clientAddress) {
		return nil, errors.New("Too many links posted, please wait couple of minutes")
	}

	err := host.IsValid(url)
	if err != nil {
		return nil, errors.New("Link is invalid: " + err.Error())
	}

	l, err := getLinkByURL(c, url)
	if err != nil {
		return nil, err
	}

	if l == nil {
		l = &Link{}
		id := 0
		err = c.db.Conn.QueryRow(sqlInsertLink, "0", url, clientAddress, "NOW()").Scan(&id)
		if err != nil {
			return nil, errors.New("Error while inserting link")
		}

		_, err = c.db.Conn.Exec(sqlUpdateLinkName, getNameFromID(id), id)
		if err != nil {
			return nil, errors.New("Error while updating link name")
		}

		l, err = getLinkByID(c, id)
		if err != nil {
			return nil, errors.New("Error while getting link by ID")
		}
	}

	return l, nil
}

func isSpammer(c *Controller, clientAddress string) bool {
	linkCount := 0
	c.db.Conn.QueryRow(sqlGetPostCountInLastMinutes, clientAddress, 10).Scan(&linkCount)

	if linkCount >= 10 {
		return true
	}
	return false
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
