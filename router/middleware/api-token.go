package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"

	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// CheckAPIToken middleware
func CheckAPIToken(c *fiber.Ctx) error {
	token := c.Get("Chat-App-Token", "")

	if token == "" {
		return c.Status(403).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "APITokenNotFound",
			Message: "App Not Verified",
		})
	}

	var apiToken Models.APIToken
	result := Connection.Get().Where("token = ?", token).First(&apiToken)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(403).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "APITokenInvalid",
			Message: "App Not Verified",
		})
	}
	Connection.Close()

	c.Set("Accept", apiToken.Origin)
	return cors.New(cors.Config{
		AllowOrigins: apiToken.Origin,
		AllowHeaders: "Origin, Content-Type, Accept, Chat-App-Token, Chat-User-Token",
	})(c)
}
