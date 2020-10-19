package chat

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
	Chat "github.com/ifanfairuz/go-chat-rest-api/database/query/chat/get"
	Functions "github.com/ifanfairuz/go-chat-rest-api/lib/functions"
)

// Request type
type Request struct {
	ChatIDs []int  `json:"chat_ids" xml:"chat_ids" form:"chat_ids"`
	ChatID  int    `json:"chat_id" xml:"chat_id" form:"chat_id"`
	To      string `json:"to" xml:"to" form:"to"`
	Text    string `json:"text" xml:"text" form:"text"`
	SendAt  int64  `json:"send_at" xml:"send_at" form:"send_at"`
}

// Validate Request
func (param Request) Validate(c *fiber.Ctx) error {
	validations := c.Locals("inputvalidations").([]string)

	if Functions.InArrayString("chat_ids", validations) && len(param.ChatIDs) == 0 {
		return errors.New("missing-chat_id")
	}

	if Functions.InArrayString("chat_id", validations) && param.ChatID == 0 {
		return errors.New("missing-chat_id")
	}

	if Functions.InArrayString("to", validations) && param.To == "" {
		return errors.New("missing-to")
	}

	if Functions.InArrayString("text", validations) && param.Text == "" {
		return errors.New("missing-text")
	}

	if Functions.InArrayString("send_at", validations) && param.SendAt == 0 {
		return errors.New("missing-send_at")
	}

	return nil
}

// Response type
type Response struct {
	History []Chat.HistoryResult `json:"history,omitempty" xml:"history" form:"history"`
	Datas   []Models.Chat        `json:"datas" xml:"datas" form:"datas"`
	Chat    interface{}          `json:"chat,omitempty" xml:"chat" form:"chat"`
	Status  bool                 `json:"status" xml:"status" form:"status"`
	Message string               `json:"message" xml:"message" form:"message"`
}
