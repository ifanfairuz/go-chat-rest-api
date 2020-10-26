package insert

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// OneOrUpdate chats
func OneOrUpdate(email string, statusOnline bool, image string) (Models.User, error) {
	var exist Models.User
	connection := Connection.Get()
	query := connection.Where("email = ?", email).First(&exist)
	if query.RowsAffected > 0 {
		query.Update("status_online", true)
		Connection.Close()
		return exist, nil
	}

	chat := Models.User{
		Email:        email,
		StatusOnline: true,
		LastOnline:   0,
		Image:        image,
	}
	result := connection.Create(&chat)

	Connection.Close()

	return chat, result.Error
}
