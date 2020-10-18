package token

import (
	"github.com/gofiber/fiber/v2"

	Insert "github.com/ifanfairuz/go-chat-rest-api/database/query/token/insert"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// Get handler
func Get(c *fiber.Ctx) error {
	var param Request
	done, res := Network.Parse(c, &param)
	if !done {
		return res()
	}

	token, err := Insert.Generate(param.Email)
	if err != nil {
		return c.Status(500).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetToken",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		Token:   token,
		Status:  true,
		Message: "Success",
	})
}
