package session

import (
	"github.com/gofiber/fiber/v2"

	Insert "github.com/ifanfairuz/go-chat-rest-api/database/query/session/insert"
	UserInsert "github.com/ifanfairuz/go-chat-rest-api/database/query/users/insert"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Connect handler
func Connect(c *fiber.Ctx) error {
	c.Locals("inputvalidations", []string{"email", "socket", "image"})
	var param Request
	done, res := Network.Parse(c, &param)
	if !done {
		return res()
	}

	session, erro := Insert.One(param.Email, param.Socket)
	if erro == nil {
		UserInsert.OneOrUpdate(param.Email, true, param.Image)
	}

	if erro != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetToken",
			Message: erro.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Session: session,
		Status:  true,
		Message: "Success",
	})
}
