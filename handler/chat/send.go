package chat

import (
	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Insert "github.com/ifanfairuz/go-chat-rest-api/database/query/chat/insert"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Send handler
func Send(c *fiber.Ctx) error {
	c.Locals("inputvalidations", []string{"to", "text", "send_at"})
	token := c.Locals("token").(Models.Token)
	var param Request
	done, res := Network.Parse(c, &param)
	if !done {
		return res()
	}

	chat, err := Insert.One(token.Email, param.To, param.Text, param.SendAt)

	if err != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotSendChat",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Chat:    chat,
		Status:  true,
		Message: "Success",
	})
}
