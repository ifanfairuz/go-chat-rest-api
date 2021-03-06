package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// All get user
func All() ([]Models.User, error) {
	var users []Models.User
	result := Connection.Get().Find(&users)

	Connection.Close()

	return users, result.Error
}

// AllExcept get user
func AllExcept(email string) ([]Models.User, error) {
	var users []Models.User
	result := Connection.Get().Where("email != ?", email).Find(&users)

	Connection.Close()

	return users, result.Error
}
