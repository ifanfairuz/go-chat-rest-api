package session

import (
	"github.com/gofiber/fiber/v2"

	Delete "github.com/ifanfairuz/go-chat-rest-api/database/query/session/delete"
	GetSession "github.com/ifanfairuz/go-chat-rest-api/database/query/session/get"
	UserSet "github.com/ifanfairuz/go-chat-rest-api/database/query/users/set"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Diconnect handler
func Diconnect(c *fiber.Ctx) error {
	c.Locals("inputvalidations", []string{"email", "socket"})
	var param Request
	done, res := Network.Parse(c, &param)
	if !done {
		return res()
	}

	erro := Delete.OneByEmailSocket(param.Email, param.Socket)
	if erro == nil {
		sessions, getErr := GetSession.AllByEmail(param.Email)
		if getErr == nil && len(sessions) <= 0 {
			UserSet.Online(param.Email, false)
		}
	}

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
