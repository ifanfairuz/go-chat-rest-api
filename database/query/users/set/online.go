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
	result := query.Where("email = ?", email).First(&user)
	if result.RowsAffected < 1 {
		return result.Error
	}

	user.StatusOnline = status
	user.LastOnline = time.Now().Unix()
	result = result.Save(&user)

	Connection.Close()

	return result.Error
}
