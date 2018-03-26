package link

// Do not change this after the links are inserted in database
// It will break getting and inserting new links
const validChars = "ABCDEFHJKLMNPRSTUVXYZabcdefgijkmnprstuvxyz23456789"

const sqlInsertLink = `
INSERT INTO links (
	name, url, client_address, created_at
)
VALUES (
	$1, $2, $3, $4
)
RETURNING id
`

const sqlUpdateLinkName = `
UPDATE links SET name = $1 WHERE id = $2
`

const sqlGetLinkByID = `
SELECT id, name, url, client_address, created_at
FROM links
WHERE id = $1
`

const sqlGetLinkByName = `
SELECT id, name, url, client_address, created_at
FROM links
WHERE name = $1
`

const sqlGetLinkByURL = `
SELECT id, name, url, client_address, created_at
FROM links
WHERE url = $1
`

// get link counts by ip address and interval - used for throttling
const sqlGetPostCountInLastMinutes = `
SELECT COUNT(*)
FROM links
WHERE client_address = $1 AND created_at > current_timestamp - ($2 || ' minutes')::interval
`
