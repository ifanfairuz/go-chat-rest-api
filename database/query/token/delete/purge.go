package delete

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// Purge Token
func Purge(token Models.Token) error {
	result := Connection.Get().Delete(&token)

	Connection.Close()

	return result.Error
}
