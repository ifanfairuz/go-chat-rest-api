package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// LastWith get chat with target order by last
func LastWith(email string, target string, limit int) ([]Models.Chat, error) {
	query := Connection.Get()
	query = query.Where("(`from` = ? OR `to` = ?)", email, email)
	query = query.Where("(`from` = ? OR `to` = ?)", target, target)
	query = query.Order("send_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}

	var chats []Models.Chat
	result := query.Find(&chats)

	reversed := make([]Models.Chat, 0, len(chats))
	for i := len(chats) - 1; i >= 0; i-- {
		reversed = append(reversed, chats[i])
	}

	Connection.Close()

	return reversed, result.Error
}
