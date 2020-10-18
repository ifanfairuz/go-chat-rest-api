package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// OneByEmailSocket sessions
func OneByEmailSocket(email string, socket string) (Models.Session, error) {
	var sessions Models.Session
	result := Connection.Get().Where("email = ? AND socket_id = ?", email, socket).First(&sessions)

	return sessions, result.Error
}
