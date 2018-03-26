package db

const sqlCreateSchema = `
CREATE TABLE IF NOT EXISTS schema (
	version integer
)
`

const sqlGetSchema = `
SELECT * FROM schema
`

const sqlInsertSchema = `
INSERT INTO schema (version) VALUES (1)
`

const sqlCreateLinks = `
CREATE TABLE IF NOT EXISTS links (
	id serial PRIMARY KEY,
	name text UNIQUE,
	url text UNIQUE NOT NULL,
	client_address inet NOT NULL,
	created_at timestamp NOT NULL
)
`

// create index for faster link post throttling
const sqlCreateLinksClientAddressIndex = `
CREATE INDEX IF NOT EXISTS links_client_address_idx ON links (client_address)
`

const sqlCreateLinksCreatedAtIndex = `
CREATE INDEX IF NOT EXISTS links_created_at_idx ON links (created_at)
`
