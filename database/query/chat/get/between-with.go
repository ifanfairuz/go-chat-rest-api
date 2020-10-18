package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// BetweenWith get chat with target between two chats
func BetweenWith(email string, target string, start string, end string) ([]Models.Chat, error) {
	query := Connection.Get()

	var chatStart Models.Chat
	result := query.First(&chatStart, start)
	if result.Error != nil {
		return []Models.Chat{}, result.Error
	}

	var chatEnd Models.Chat
	result = query.First(&chatEnd, end)
	if result.Error != nil {
		return []Models.Chat{}, result.Error
	}

	var chats []Models.Chat
	query = query.Where("send_at >= ? AND send_at <= ?", chatStart.SendAt, chatEnd.SendAt)
	query = query.Order("send_at")
	result = query.Find(&chats)

	Connection.Close()

	return chats, result.Error
}
