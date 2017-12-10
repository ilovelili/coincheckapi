// Package config configuration related
package config

import (
	"encoding/json"
	"os"
	"path"
)

// GetConfig get config defined in config.json
func GetConfig() (config *Config) {
	pwd, _ := os.Getwd()
	path := path.Join(pwd, "config.json")
	configFile, err := os.Open(path)
	defer configFile.Close()

	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		panic(err)
	}

	return
}

// API API settings
type API struct {
	Key       string `json:"key"`
	SecretKey string `json:"secret"`
}

// Order bid related settings
type Order struct {
	DailyBudget float32 `json:"dailybudget"`
	Threshold   float32 `json:"threshold"`
	Price       float32 `json:"bidprice"`
	Type        string  `json:"type"`
}

// Transaction transaction settings
type Transaction struct {
	Bid *Order `json:"bid"`
	Ask *Order `json:"ask"`
}

// Config config entry
type Config struct {
	API         `json:"api"`
	Transaction `json:"transaction"`
}
