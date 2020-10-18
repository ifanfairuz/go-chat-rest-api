package status

import (
	"time"

	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
)

// ReadAllWith chats
func ReadAllWith(email string, target string) error {
	query := Connection.Get()
	query = query.Where("(`from` = ? OR `to` = ?)", email, email)
	query = query.Where("(`from` = ? OR `to` = ?)", target, target)
	result := query.Update("status", 3).Update("read_at", time.Now().Unix())

	return result.Error
}
