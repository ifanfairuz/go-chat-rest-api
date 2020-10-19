package migration

import (
	"gorm.io/gorm"
)

/**
* migrate via model
**/
func migrateModel(connection *gorm.DB, model interface{}) {
	if !connection.Migrator().HasTable(model) {
		connection.Migrator().CreateTable(model)
	}
}
