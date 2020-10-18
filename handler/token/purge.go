package token

import (
	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Delete "github.com/ifanfairuz/go-chat-rest-api/database/query/token/delete"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Purge handler
func Purge(c *fiber.Ctx) error {
	token := c.Locals("token").(Models.Token)
	err := Delete.Purge(token)

	if err != nil {
		return c.Status(500).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetToken",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Status:  true,
		Message: "Success",
	})
}
