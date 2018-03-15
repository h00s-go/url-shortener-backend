package link

import (
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
RETURNING id`
const sqlUpdateLinkName = `
"UPDATE links SET name = $1 WHERE id = $2"
`
const sqlGetLinkByID = `
SELECT id, name, url, view_count, client_address, created_at
FROM links
WHERE id = $1"
`
const sqlGetLinkByURL = `
SELECT id, name, url, view_count, client_address, created_at
FROM links
WHERE url = $1"
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

// InsertLink in db. If inserted, return Link struct
func InsertLink(c *Controller, url string, clientAddress string) (*Link, error) {
	l := &Link{}

	id := 0
	err := c.db.Conn.QueryRow(sqlInsertLink, "0", url, 0, "127.0.0.1", "NOW()").Scan(&id)
	if err != nil {
		return l, errors.New("Error while inserting link")
	}

	_, err = c.db.Conn.Exec(sqlUpdateLinkName, getNameFromID(id), id)
	if err != nil {
		return l, errors.New("Error while updating link name")
	}

	err = c.db.Conn.QueryRow(sqlGetLinkByID, id).Scan(&l.ID, &l.Name, &l.ViewCount, &l.ClientAddress, &l.CreatedAt)
	if err != nil {
		return l, errors.New("Error while getting link")
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
