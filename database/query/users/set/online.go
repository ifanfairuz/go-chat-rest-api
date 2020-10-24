package set

import (
	"time"

	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// Online chats
func Online(email string, status bool) error {
	query := Connection.Get()

	var user Models.User
	result := query.First(&user, email)
	if result.Error != nil {
		return result.Error
	}

	result = result.Update("status_online", status).Update("last_online", time.Now().Unix())

	Connection.Close()

	return result.Error
}
