package link

const sqlInsertActivity = `
INSERT INTO activities (
	link_id, client_address, accessed_at
)
VALUES (
	$1, $2, $3
)
`
