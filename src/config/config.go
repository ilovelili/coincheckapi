// Package config configuration related
package config

import (
	"encoding/json"
	"os"
	"path"
)

// GetConfig get config defined in config.json
func GetConfig() (config *Config, err error) {
	pwd, _ := os.Getwd()
	path := path.Join(pwd, "config.json")
	configFile, err := os.Open(path)
	defer configFile.Close()

	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return
}

// API API settings
type API struct {
	Key       string `json:"key"`
	SecretKey string `json:"secret"`
}

// Config config entry
type Config struct {
	API `json:"api"`
}
