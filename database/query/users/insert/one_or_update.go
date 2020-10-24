package insert

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// OneOrUpdate chats
func OneOrUpdate(email string, statusOnline bool) (Models.User, error) {
	var exist Models.User
	query := Connection.Get().First(&exist, email)
	if query.Error != nil {
		query.Update("status_online", true)
		return exist, nil
	}

	chat := Models.User{
		Email:        email,
		StatusOnline: true,
		LastOnline:   0,
	}
	result := Connection.Get().Create(&chat)

	Connection.Close()

	return chat, result.Error
}
