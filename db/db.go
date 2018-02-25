package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq" //for a postgres
)

// Database handles DB connections
type Database struct {
	conn *sql.DB
}

type configuration struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

// NewDatabase create new DB Database
func NewDatabase() (*Database, error) {
	c, err := loadConfiguration("configuration.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", c.DBHost, c.DBUser, c.DBPassword, c.DBName)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Database{conn: conn}, nil
}

func loadConfiguration(path string) (configuration, error) {
	var config configuration
	configJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(configJSON, &config)
	return config, err
}
