package chat

import (
	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Chat "github.com/ifanfairuz/go-chat-rest-api/database/query/chat/status"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// ReadMany handler
func ReadMany(c *fiber.Ctx) error {
	c.Locals("inputvalidations", []string{"chat_ids"})
	token := c.Locals("token").(Models.Token)
	var param Request
	done, res := Network.Parse(c, &param)
	if !done {
		return res()
	}

	err := Chat.ReadAll(param.ChatIDs, token.Email)

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
