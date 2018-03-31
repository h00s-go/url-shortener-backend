package link

const sqlInsertActivity = `
INSERT INTO activities (
	link_id, client_address, accessed_at
)
VALUES (
	$1, $2, $3
)
`

const sqlGetLinkActivityStats = `
SELECT COUNT(*)
FROM activities
WHERE link_id = $1
`
