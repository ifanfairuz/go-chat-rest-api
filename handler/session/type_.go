package session

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Functions "github.com/ifanfairuz/go-chat-rest-api/lib/functions"
)

// Request type
type Request struct {
	Email  string `json:"email" xml:"email" form:"email"`
	Socket string `json:"socket" xml:"socket" form:"socket"`
	Image  string `json:"image" xml:"image" form:"image"`
}

// Validate Request
func (param Request) Validate(c *fiber.Ctx) error {
	validations := c.Locals("inputvalidations").([]string)

	if Functions.InArrayString("email", validations) && param.Email == "" {
		return errors.New("missing-email")
	}

	if Functions.InArrayString("socket", validations) && param.Socket == "" {
		return errors.New("missing-socket")
	}

	if Functions.InArrayString("image", validations) && param.Socket == "" {
		return errors.New("missing-image")
	}

	return nil
}

// Response type
type Response struct {
	Datas   []Models.Session `json:"datas" xml:"datas" form:"datas"`
	Session interface{}      `json:"session,omitempty" xml:"session" form:"session"`
	Status  bool             `json:"status" xml:"status" form:"status"`
	Message string           `json:"message" xml:"message" form:"message"`
}
