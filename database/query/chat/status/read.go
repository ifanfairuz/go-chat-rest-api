package status

import (
	"errors"
	"time"

	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// Read chats
func Read(id int, email string) error {
	query := Connection.Get()

	var chat Models.Chat
	result := query.First(&chat, id)
	if result.Error != nil {
		return result.Error
	}

	if chat.To != email {
		return errors.New("ChatNotYours")
	}

	result = result.Update("status", 3).Update("read_at", time.Now().Unix())

	Connection.Close()

	return result.Error
}
