package command

import (
	"log"

	"github.com/gofiber/fiber/v2"
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

	if app != nil {
		app.Listen(address)
	} else {
		log.Fatal("Error.")
	}
}
