package network

import (
	"github.com/gofiber/fiber/v2"
)

// Request interface
type Request interface {
	Validate(c *fiber.Ctx) error
}

// ErrResponse type
type ErrResponse struct {
	Status  bool   `json:"status" xml:"status" form:"status"`
	Error   string `json:"error" xml:"error" form:"error"`
	Message string `json:"message" xml:"message" form:"message"`
}
