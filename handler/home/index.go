package home

import (
	"github.com/gofiber/fiber/v2"
)

// Index page
func Index(c *fiber.Ctx) error {
	return c.Send([]byte("REST API Chat v2.0-candidate"))
}
