package migration

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"

	"gorm.io/gorm"
)

// Migrate DB
func Migrate() {
	var connection *gorm.DB = Connection.Get()

	migrateModel(connection, &Models.APIToken{})
	migrateModel(connection, &Models.Token{})
	migrateModel(connection, &Models.Session{})
	migrateModel(connection, &Models.Chat{})

	Connection.Close()
}
