package connection

import (
	"os"
)

/**
* get configuration database
* return Config type
**/
func getConfig() *Config {
	return &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_DB_NAME"),
		Charset:  os.Getenv("DB_CHARSET"),
	}
}
