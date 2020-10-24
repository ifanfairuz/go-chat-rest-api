package users

import (
	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// Request type
type Request struct{}

// Validate Request
func (param Request) Validate(c *fiber.Ctx) error {
	return nil
}

// Response type
type Response struct {
	Datas   []Models.User `json:"datas" xml:"datas" form:"datas"`
	User    interface{}    `json:"user,omitempty" xml:"user" form:"user"`
	Status  bool           `json:"status" xml:"status" form:"status"`
	Message string         `json:"message" xml:"message" form:"message"`
}
