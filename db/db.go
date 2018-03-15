package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/h00s/url-shortener-backend/config"
	_ "github.com/lib/pq" //for a postgres
)

// Database handles DB connections
type Database struct {
	Conn *sql.DB
}

// Connect create new Database struct and connects to DB
func Connect(c config.Configuration) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Name)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Database{Conn: conn}, nil
}

// Migrate migrates database to valid state
func (db *Database) Migrate() error {
	sqlCreateSchema := `
	CREATE TABLE IF NOT EXISTS schema (
		version integer
	);
	`
	_, err := db.Conn.Exec(sqlCreateSchema)
	if err != nil {
		return err
	}

	err = db.Conn.QueryRow("SELECT * FROM schema;").Scan()
	if err == sql.ErrNoRows {
		db.Conn.Exec("INSERT INTO schema (version) VALUES (1);")
	}

	sqlCreateLinks := `
	CREATE TABLE IF NOT EXISTS links (
		id serial PRIMARY KEY,
		name text UNIQUE NOT NULL,
		url text UNIQUE NOT NULL,
		view_count integer DEFAULT 0,
		client_address inet NOT NULL,
		created_at timestamp NOT NULL
	);
	`
	_, err = db.Conn.Exec(sqlCreateLinks)
	return err
}
