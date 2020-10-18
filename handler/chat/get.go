package chat

import (
	"github.com/gofiber/fiber/v2"

	Chat "github.com/ifanfairuz/go-chat-rest-api/database/query/chat/get"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Get handler
func Get(c *fiber.Ctx) error {
	id := c.Params("id", "")

	if id == "" {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "InvalidId",
			Message: "CannotGetChats",
		})
	}

	chat, err := Chat.One(id)
	if err != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetChats",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Chat:    chat,
		Status:  true,
		Message: "Success",
	})
}
