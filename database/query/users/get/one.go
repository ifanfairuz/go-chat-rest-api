package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// One get user
func One(email string) (Models.User, error) {
	var user Models.User
	result := Connection.Get().First(&user, email)

	Connection.Close()

	return user, result.Error
}
