package middleware

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// CheckToken middleware
func CheckToken(c *fiber.Ctx) error {
	t := c.Get("Chat-User-Token", "")

	if t == "" {
		return c.Status(403).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "TokenNotFound",
			Message: "Not Auth",
		})
	}

	var token Models.Token
	result := Connection.Get().Where("token", t).First(&token)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(403).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "TokenInvalid",
			Message: "Not Auth",
		})
	}

	now := time.Now().Unix()
	if now > token.ExpiredAt {
		return c.Status(403).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "Expired",
			Message: "TokenExpired",
		})
	}

	c.Locals("token", token)
	return c.Next()
}
