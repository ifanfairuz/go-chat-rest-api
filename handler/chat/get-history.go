package chat

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Chat "github.com/ifanfairuz/go-chat-rest-api/database/query/chat/get"
	Network "github.com/ifanfairuz/go-chat-rest-api/lib/network"
)

// GetHistory handler
func GetHistory(c *fiber.Ctx) error {
	token := c.Locals("token").(Models.Token)
	limit, errParse := strconv.Atoi(c.Params("limit", "0"))
	if errParse != nil {
		limit = 0
	}

	history, err := Chat.History(token.Email, limit)
	if err != nil {
		return c.Status(400).JSON(&Network.ErrResponse{
			Status:  false,
			Error:   "CannotGetHistory",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&Response{
		History: history,
		Status:  true,
		Message: "Success",
	})
}
