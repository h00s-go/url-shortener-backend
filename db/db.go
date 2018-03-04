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
	conn *sql.DB
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

	return &Database{conn: conn}, nil
}

// Migrate migrates database to valid state
func (db *Database) Migrate() error {
	sqlCreateSchema := `
	CREATE TABLE IF NOT EXISTS schema (
		version integer
	);
	`
	_, err := db.conn.Exec(sqlCreateSchema)
	if err != nil {
		return err
	}

	err = db.conn.QueryRow("SELECT * FROM schema;").Scan()
	if err == sql.ErrNoRows {
		db.conn.Exec("INSERT INTO schema (version) VALUES (1);")
	}

	sqlCreateLinks := `
	CREATE TABLE IF NOT EXISTS links (
		id SERIAL PRIMARY KEY,
		name TEXT,
		url TEXT
	);
	`
	_, err = db.conn.Exec(sqlCreateLinks)
	return err
}
