package link

import (
	"database/sql"
	"errors"
	"strings"
)

// Do not change this after the links are inserted in database
// It will break getting and inserting new links
const validChars = "ABCDEFHJKLMNPRSTUVXYZabcdefgijkmnprstuvxyz23456789"

const sqlInsertLink = `
INSERT INTO links (
	name, url, view_count, client_address, created_at
)
VALUES (
	$1, $2, $3, $4, $5
)
RETURNING id
`
const sqlUpdateLinkName = `
UPDATE links SET name = $1 WHERE id = $2
`
const sqlGetLinkByID = `
SELECT id, name, url, view_count, client_address, created_at
FROM links
WHERE id = $1
`
const sqlGetLinkByName = `
SELECT id, name, url, view_count, client_address, created_at
FROM links
WHERE name = $1
`
const sqlGetLinkByURL = `
SELECT id, name, url, view_count, client_address, created_at
FROM links
WHERE url = $1
`

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

	err := c.db.Conn.QueryRow(sqlGetLinkByURL, url).Scan(&l.ID, &l.Name, &l.URL, &l.ViewCount, &l.ClientAddress, &l.CreatedAt)
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
