package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// AllByEmail sessions
func AllByEmail(email string) ([]Models.Session, error) {
	var sessions []Models.Session
	result := Connection.Get().Where("email = ?", email).Find(&sessions)

	return sessions, result.Error
}
