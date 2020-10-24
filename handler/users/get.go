package users

import (
	"github.com/gofiber/fiber/v2"

	Users "github.com/ifanfairuz/go-chat-rest-api/database/query/users/get"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Get handler
func Get(c *fiber.Ctx) error {
	email := c.Params("email", "")

	if email == "" {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "InvalidEmail",
			Message: "CannotGetUser",
		})
	}

	user, err := Users.One(email)
	if err != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetChats",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		User:    user,
		Status:  true,
		Message: "Success",
	})
}
