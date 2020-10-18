package connection

import (
	Command "github.com/ifanfairuz/go-chat-rest-api/command"
)

/**
* get configuration database
* return Config type
**/
func getConfig() *Config {
	return &Config{
		Host:     Command.GetEnv("DB_HOST"),
		Port:     Command.GetEnv("DB_PORT"),
		User:     Command.GetEnv("DB_USER"),
		Password: Command.GetEnv("DB_PASSWORD"),
		Dbname:   Command.GetEnv("DB_DB_NAME"),
		Charset:  Command.GetEnv("DB_CHARSET"),
	}
}
