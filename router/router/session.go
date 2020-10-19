package router

import (
	"github.com/gofiber/fiber/v2"

	Session "github.com/ifanfairuz/go-chat-rest-api/handler/session"
	Middleware "github.com/ifanfairuz/go-chat-rest-api/router/middleware"
)

func createSession(app *fiber.App) {
	session := app.Group("/session", Middleware.CheckToken)

	session.Get("/all", Session.All)
	session.Get("/get/:socketid", Session.Get)

	session.Post("/connect", Session.Connect)
	session.Post("/disconnect", Session.Diconnect)
	session.Post("/reset", Session.Reset)
}
