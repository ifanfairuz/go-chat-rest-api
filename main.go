package main

import (
	Command "github.com/ifanfairuz/go-chat-rest-api/command"
)

func init() {
	Command.LoadEnv()
}

func main() {
	Command.Run(":3000")
}
