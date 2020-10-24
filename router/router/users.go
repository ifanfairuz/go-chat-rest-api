package router

import (
	"github.com/gofiber/fiber/v2"

	Users "github.com/ifanfairuz/go-chat-rest-api/handler/users"
	Middleware "github.com/ifanfairuz/go-chat-rest-api/router/middleware"
)

func createUsers(app *fiber.App) {
	users := app.Group("/users", Middleware.CheckToken)

	users.Get("/all", Users.All)
	users.Get("/get/:email", Users.Get)
}
