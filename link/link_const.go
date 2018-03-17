package link

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
