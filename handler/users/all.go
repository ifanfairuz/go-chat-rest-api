package users

import (
	"github.com/gofiber/fiber/v2"

	Users "github.com/ifanfairuz/go-chat-rest-api/database/query/users/get"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// All handler
func All(c *fiber.Ctx) error {
	users, err := Users.All()

	if err != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetUsers",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Datas:   users,
		Status:  true,
		Message: "Success",
	})
}
