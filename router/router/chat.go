package router

import (
	"github.com/gofiber/fiber/v2"

	Chat "github.com/ifanfairuz/go-chat-rest-api/handler/chat"
	Middleware "github.com/ifanfairuz/go-chat-rest-api/router/middleware"
)

func createChat(app *fiber.App) {
	chat := app.Group("/chat", Middleware.CheckToken)
	withTarget := app.Group("/chat/with", Middleware.CheckToken)

	chat.Get("/history/:limit?", Chat.GetHistory)
	chat.Get("/get/:id", Chat.Get)

	chat.Post("/readall", Chat.ReadMany)
	chat.Post("/read", Chat.Read)
	chat.Post("/received", Chat.Received)
	chat.Post("/send", Chat.Send)

	withTarget.Get("/:target/last/:limit?", Chat.GetLastWith)
	withTarget.Get("/:target/between/:start-:end", Chat.GetBetweenWith)
	withTarget.Get("/:target/delay", Chat.GetDelayWith)
	withTarget.Get("/:target/unread", Chat.GetUnreadWith)

	withTarget.Post("/:target/readall", Chat.ReadAllWith)
}
