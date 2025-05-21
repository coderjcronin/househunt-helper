package config

import "time"

type Config struct {
	dbUrl       string
	dbToken     string
	secret      string
	platform    string
	port        string
	assets_root string
	timeout     time.Duration
}
