package session

import (
	"github.com/gofiber/fiber/v2"

	Delete "github.com/ifanfairuz/go-chat-rest-api/database/query/session/delete"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Reset handler
func Reset(c *fiber.Ctx) error {
	c.Locals("inputvalidations", []string{"email"})
	var param Request
	done, res := Network.Parse(c, &param)
	if !done {
		return res()
	}

	erro := Delete.AllByEmail(param.Email)

	if erro != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetToken",
			Message: erro.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Status:  true,
		Message: "Success",
	})
}
