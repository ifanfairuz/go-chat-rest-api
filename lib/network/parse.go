package network

import (
	"github.com/gofiber/fiber/v2"
)

// Parse request
func Parse(c *fiber.Ctx, param Request) (bool, func() error) {
	if err := c.BodyParser(param); err != nil {
		return false, func() error {
			return c.Status(400).JSON(&ErrResponse{
				Status:  false,
				Error:   "CannotParseBody",
				Message: err.Error(),
			})
		}
	}

	if err := param.Validate(c); err != nil {
		return false, func() error {
			return c.Status(400).JSON(&ErrResponse{
				Status:  false,
				Error:   "InputError",
				Message: err.Error(),
			})
		}
	}

	return true, func() error { return nil }
}
