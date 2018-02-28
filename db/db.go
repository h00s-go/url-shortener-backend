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

// NewDatabase create new DB Database
func NewDatabase(configPath string) (*Database, error) {
	c, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
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
