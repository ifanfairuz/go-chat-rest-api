package session

import (
	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Session "github.com/ifanfairuz/go-chat-rest-api/database/query/session/get"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Get handler
func Get(c *fiber.Ctx) error {
	token := c.Locals("token").(Models.Token)
	socket := c.Params("socketid")
	session, err := Session.OneByEmailSocket(token.Email, socket)

	if err != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetToken",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Session: session,
		Status:  true,
		Message: "Success",
	})
}
