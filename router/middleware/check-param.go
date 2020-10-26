package middleware

import (
	"github.com/gofiber/fiber/v2"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// CheckParamTarget middleware
func CheckParamTarget(c *fiber.Ctx) error {
	target := c.Params("target")
	if target == "" {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "TargetNotFound",
			Message: "Error",
		})
	}

	return c.Next()
}
