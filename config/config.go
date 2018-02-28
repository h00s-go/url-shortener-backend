package config

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration struct have all fields from configuration JSON file
type Configuration struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
}

// LoadConfiguration loads configuration from path
func LoadConfiguration(path string) (Configuration, error) {
	var c Configuration
	configJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(configJSON, &c)
	return c, err
}
