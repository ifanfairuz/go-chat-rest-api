package token

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Request type
type Request struct {
	Email string `json:"email" xml:"email" form:"email"`
}

// Validate Request
func (tr Request) Validate(c *fiber.Ctx) error {
	if tr.Email == "" {
		return errors.New("Email Not Provided")
	}

	return nil
}

// Response type
type Response struct {
	Token   string `json:"token" xml:"token" form:"token"`
	Status  bool   `json:"status" xml:"status" form:"status"`
	Message string `json:"message" xml:"message" form:"message"`
}
