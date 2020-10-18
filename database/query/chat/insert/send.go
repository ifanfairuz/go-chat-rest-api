package insert

import (
	"time"

	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// Send chats
func Send(email string, target string, text string, sendAt int64) (Models.Chat, error) {
	chat := Models.Chat{
		From:   email,
		To:     target,
		Text:   text,
		Status: 1,
		SendAt: sendAt,
		SentAt: time.Now().Unix(),
	}
	result := Connection.Get().Create(&chat)

	return chat, result.Error
}
