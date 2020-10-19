package middleware

import (
	"github.com/gofiber/fiber/v2"

	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// CheckParams middleware
func CheckParams(params []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		for i := 0; i < len(params); i++ {
			param := params[i]
			value := c.Params(param, "")
			if value == "" {
				return c.Status(400).JSON(&Network.ErrResponse{
					Status:  false,
					Error:   "TargetNotFound",
					Message: "Error",
				})
			}
		}

		return c.Next()
	}
}
