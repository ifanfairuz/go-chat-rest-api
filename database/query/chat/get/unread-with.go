package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// UnreadWith get chat with target status is delay
func UnreadWith(email string, target string) ([]Models.Chat, error) {
	query := Connection.Get()
	query = query.Where("status < ?", 3)
	query = query.Order("send_at")

	var chats []Models.Chat
	result := query.Find(&chats)

	Connection.Close()

	return chats, result.Error
}
