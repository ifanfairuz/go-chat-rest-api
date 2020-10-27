package main

import (
	"os"

	Command "github.com/ifanfairuz/go-chat-rest-api/command"
)

func init() {
	// Command.LoadEnv()
}

func main() {
	Command.Run(os.Getenv("PORT"))
}
