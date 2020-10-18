package chat

import (
	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Chat "github.com/ifanfairuz/go-chat-rest-api/database/query/chat/status"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// ReadAllWith handler
func ReadAllWith(c *fiber.Ctx) error {
	token := c.Locals("token").(Models.Token)
	target := c.Params("target", "")

	err := Chat.ReadAllWith(token.Email, target)

	if err != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotReadChats",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Status:  true,
		Message: "Success",
	})
}
