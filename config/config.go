package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config struct have all fields from configuration JSON file
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
}

// LoadConfig loads configuration from path
func LoadConfig(path string) (Config, error) {
	var c Config
	configJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(configJSON, &c)
	return c, err
}
