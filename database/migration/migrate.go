package migration

import (
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// Migrate DB
func Migrate() {
	migrateModel(&Models.APIToken{})
	migrateModel(&Models.Token{})
	migrateModel(&Models.Session{})
	migrateModel(&Models.Chat{})
}
