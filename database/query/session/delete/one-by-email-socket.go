package delete

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// OneByEmailSocket sessions
func OneByEmailSocket(email string, socket string) error {
	session := Models.Session{}
	result := Connection.Get().Where("email = ? AND socket_id = ?", email, socket).Delete(&session)

	Connection.Close()

	return result.Error
}
