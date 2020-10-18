package chat

import (
	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Chat "github.com/ifanfairuz/go-chat-rest-api/database/query/chat/get"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// GetBetweenWith handler
func GetBetweenWith(c *fiber.Ctx) error {
	token := c.Locals("token").(Models.Token)
	target := c.Params("target", "")
	start := c.Params("start", "")
	end := c.Params("end", "")

	if target == "" || start == "" || end == "" {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "InvalidParam",
			Message: "CannotGetChats",
		})
	}

	chats, err := Chat.BetweenWith(token.Email, target, start, end)
	if err != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetChats",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Datas:   chats,
		Status:  true,
		Message: "Success",
	})
}
