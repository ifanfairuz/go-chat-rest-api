package router

import (
	"github.com/gofiber/fiber/v2"

	Home "github.com/ifanfairuz/go-chat-rest-api/handler/home"
	Middleware "github.com/ifanfairuz/go-chat-rest-api/router/middleware"
)

// Create Router
func Create(app *fiber.App) {
	app.Get("/", Home.Index)

	app.Use(Middleware.CheckAPIToken)

	createToken(app)
	createSession(app)
	createChat(app)
}
