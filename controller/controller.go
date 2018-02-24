package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq" //for a postgres
)

// Controller handles DB connections
type Controller struct {
	db *sql.DB
}

type configuration struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

// NewController create new DB Controller
func NewController() (*Controller, error) {
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
	h := &Controller{db: db}
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
