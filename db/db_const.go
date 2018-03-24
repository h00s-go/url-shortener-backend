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
	name text UNIQUE NOT NULL,
	url text UNIQUE NOT NULL,
	view_count integer DEFAULT 0,
	client_address inet NOT NULL,
	created_at timestamp NOT NULL
)
`

// create index for faster link post throttling
const sqlCreateLinksIndex = `
CREATE INDEX IF NOT EXISTS links_client_address_idx ON links (client_address)
`
