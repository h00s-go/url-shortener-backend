package link

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/h00s/url-shortener-backend/host"
)

// Link represent one shortened link
type Link struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	ViewCount     int    `json:"viewCount"`
	ClientAddress string `json:"clientAddress"`
	CreatedAt     string `json:"createdAt"`
}

// GetLink gets link from db by link name
func GetLink(c *Controller, name string) (*Link, error) {
	l := &Link{}

	err := c.db.Conn.QueryRow(sqlGetLinkByName, name).Scan(&l.ID, &l.Name, &l.URL, &l.ViewCount, &l.ClientAddress, &l.CreatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			return l, errors.New("Error while getting link by URL")
		}
		return l, errors.New("Link not found")
	}

	return l, nil
}

// InsertLink in db. If inserted, return Link struct
func InsertLink(c *Controller, url string, clientAddress string) (*Link, error) {
	l := &Link{}

	if isSpammer(c, clientAddress) {
		return l, errors.New("Too many links posted, please wait couple of minutes")
	}

	err := host.IsValid(url)
	if err != nil {
		return l, errors.New("Link is invalid: " + err.Error())
	}

	err = c.db.Conn.QueryRow(sqlGetLinkByURL, url).Scan(&l.ID, &l.Name, &l.URL, &l.ViewCount, &l.ClientAddress, &l.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		return l, errors.New("Error while getting link by URL")
	}
	if l.ID != 0 {
		return l, nil
	}

	id := 0
	err = c.db.Conn.QueryRow(sqlInsertLink, "0", url, 0, clientAddress, "NOW()").Scan(&id)
	if err != nil {
		return l, errors.New("Error while inserting link")
	}

	_, err = c.db.Conn.Exec(sqlUpdateLinkName, getNameFromID(id), id)
	if err != nil {
		return l, errors.New("Error while updating link name")
	}

	err = c.db.Conn.QueryRow(sqlGetLinkByID, id).Scan(&l.ID, &l.Name, &l.URL, &l.ViewCount, &l.ClientAddress, &l.CreatedAt)
	if err != nil {
		return l, errors.New("Error while getting link by ID")
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
