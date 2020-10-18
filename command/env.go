package command

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv function
/**
* load from .env
**/
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv function
/**
* get value from env
**/
func GetEnv(key string) string {
	return os.Getenv(key)
}
