package router

import (
	"github.com/gofiber/fiber/v2"

	Token "github.com/ifanfairuz/go-chat-rest-api/handler/token"
	Middleware "github.com/ifanfairuz/go-chat-rest-api/router/middleware"
)

func createToken(app *fiber.App) {
	app.Post("/token/get", Token.Get)

	routeToken := app.Group("/token", Middleware.CheckToken)
	routeToken.Post("/purge", Token.Purge)
}
