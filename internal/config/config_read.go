package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
Retrieves .env variables from config to store in memory.
*/
func configRead() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	pathDB := os.Getenv("DB_PATH")
	if pathDB == "" {
		return Config{}, fmt.Errorf("DB_PATH not set")
	}

	return Config{}, nil
}
