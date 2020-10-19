package command

import (
	"log"

	"github.com/gofiber/fiber/v2"

	Router "github.com/ifanfairuz/go-chat-rest-api/router/router"
)

// Run function
/**
* Running the app
* address: string for address listening app
*          example: "localhost:3000" or ":3000"
**/
func Run(address string) {
	var app *fiber.App

	app = fiber.New()
	migration()
	Router.Create(app)

	if app != nil {
		app.Listen(address)
	} else {
		log.Fatal("Error.")
	}
}
