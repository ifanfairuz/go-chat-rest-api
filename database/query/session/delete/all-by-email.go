package delete

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// AllByEmail session
func AllByEmail(email string) error {
	session := Models.Session{}
	result := Connection.Get().Where("email = ?", email).Delete(&session)

	return result.Error
}
