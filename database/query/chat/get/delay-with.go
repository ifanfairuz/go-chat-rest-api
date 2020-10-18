package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// DelayWith get chat with target status is delay
func DelayWith(email string, target string) ([]Models.Chat, error) {
	query := Connection.Get()
	query = query.Where("status < ?", 2)
	query = query.Order("send_at")

	var chats []Models.Chat
	result := query.Find(&chats)

	return chats, result.Error
}
