package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB_URL 			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var cfg Config

	homeDir := "~"
	fileName := "gatorconfig.json"
	fullPath := homeDir + "/" + fileName 
	dat, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Couldn't read configuration file")
		return Config{}, err
	}

	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		fmt.Println("Couldn't unmarshal json data")
		return Config{}, err
	}

	fmt.Printf("%s", cfg)

	return cfg, nil
}