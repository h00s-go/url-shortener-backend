package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq" //for a postgres
)

// Handler handles DB connections
type Handler struct {
	db *sql.DB
}

type configuration struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

// NewHandler create new DB handler
func NewHandler() (*Handler, error) {
	c, err := loadConfiguration("configuration.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", c.DBHost, c.DBUser, c.DBPassword, c.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	h := &Handler{db: db}
	return h, nil
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
