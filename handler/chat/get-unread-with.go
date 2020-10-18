package chat

import (
	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Chat "github.com/ifanfairuz/go-chat-rest-api/database/query/chat/get"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// GetUnreadWith handler
func GetUnreadWith(c *fiber.Ctx) error {
	token := c.Locals("token").(Models.Token)
	target := c.Params("target", "")

	chats, err := Chat.UnreadWith(token.Email, target)
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
