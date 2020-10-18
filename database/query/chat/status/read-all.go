package status

import (
	"time"

	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
)

// ReadAll chats
func ReadAll(ids []int, email string) error {
	query := Connection.Get()
	query = query.Where("id IN (?)", ids)
	query = query.Where("(`from` = ? OR `to` = ?)", email, email)

	result := query.Update("status", 3).Update("read_at", time.Now().Unix())

	Connection.Close()

	return result.Error
}
