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
	return getLink(c, sqlGetLinkByName, strings.TrimSpace(name))
}

func getLinkByURL(c *Controller, url string) (*Link, error) {
	return getLink(c, sqlGetLinkByURL, strings.TrimSpace(url))
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
	url = strings.TrimSpace(url)

	err := host.IsValid(url)
	if err != nil {
		return nil, errors.New("Link is invalid: " + err.Error())
	}

	// Check if URL is already in DB
	l, err := getLinkByURL(c, url)
	switch {
	case err != nil:
		return nil, err
	case l != nil:
		return l, nil
	}

	// URL is not in DB, insert it
	l = &Link{}
	id := 0

	tx, err := c.db.Conn.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(sqlInsertLink, nil, url, clientAddress, "NOW()").Scan(&id)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Error while inserting link")
	}

	_, err = tx.Exec(sqlUpdateLinkName, getNameFromID(id), id)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Error while updating link name")
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	l, err = getLinkByID(c, id)
	if err != nil {
		return nil, errors.New("Error while getting link")
	}

	return l, nil
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
