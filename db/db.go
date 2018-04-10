package db

import (
	"database/sql"
	"errors"
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
func Connect(db config.Database) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", db.Host, db.Port, db.User, db.Password, db.Name)
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
		return errors.New("Unable to start transaction: " + err.Error())
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			err = errors.New("There was an error in migrate transaction: " + err.Error())
		}
		err = tx.Commit()
		if err != nil {
			err = errors.New("There was an error in commiting migrate transaction: " + err.Error())
		}
	}()

	_, err = tx.Exec(sqlCreateSchema)
	if err != nil {
		return
	}

	version := 0
	err = tx.QueryRow(sqlGetSchema).Scan(&version)
	switch {
	case err == sql.ErrNoRows:
		_, err = tx.Exec(sqlInsertSchema)
		if err != nil {
			return
		}
	case err != nil:
		return
	}

	_, err = tx.Exec(sqlCreateLinks)
	if err != nil {
		return
	}

	_, err = tx.Exec(sqlCreateLinksClientAddressIndex)
	if err != nil {
		return
	}

	_, err = tx.Exec(sqlCreateLinksCreatedAtIndex)
	if err != nil {
		return
	}

	_, err = tx.Exec(sqlCreateActivities)
	if err != nil {
		return
	}

	_, err = tx.Exec(sqlCreateActivitiesLinkIDIndex)
	if err != nil {
		return
	}

	return
}
