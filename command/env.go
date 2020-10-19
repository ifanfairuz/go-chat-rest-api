package command

import (
	"log"

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
