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
func (db *Database) Migrate() (err error) {
	tx, err := db.Conn.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec(sqlCreateSchema)
	if err != nil {
		return err
	}

	err = tx.QueryRow(sqlGetSchema).Scan()
	switch {
	case err == sql.ErrNoRows:
		_, err = tx.Exec(sqlInsertSchema)
		if err != nil {
			return
		}
	case err != nil:
		return
	}

	_, err = db.Conn.Exec(sqlCreateLinks)
	if err != nil {
		return err
	}

	_, err = db.Conn.Exec(sqlCreateLinksClientAddressIndex)
	if err != nil {
		return err
	}

	_, err = db.Conn.Exec(sqlCreateLinksCreatedAtIndex)
	if err != nil {
		return err
	}

	return
}
