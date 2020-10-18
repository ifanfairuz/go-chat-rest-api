package migration

import (
	Models "github.com/ifanfairuz/go-chat-rest-api/command"
)

// Migrate DB
func Migrate() {
	migrateModel(&Models.APIToken{})
	migrateModel(&Models.Token{})
	migrateModel(&Models.Session{})
	migrateModel(&Models.Chat{})
}
