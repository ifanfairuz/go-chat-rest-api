package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// One by id
func One(id string) (Models.Chat, error) {
	var chat Models.Chat
	result := Connection.Get().First(&chat, id)

	Connection.Close()

	return chat, result.Error
}
