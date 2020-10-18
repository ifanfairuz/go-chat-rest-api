package migration

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
)

/**
* migrate via model
**/
func migrateModel(model interface{}) {
	if !Connection.Get().Migrator().HasTable(model) {
		Connection.Get().Migrator().CreateTable(model)
	}
}
