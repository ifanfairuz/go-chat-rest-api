package command

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Migration "github.com/ifanfairuz/go-chat-rest-api/database/migration"
)

func migration() {
	if Connection.Connect() {
		Migration.Migrate()
		Connection.Close()
	}
}
