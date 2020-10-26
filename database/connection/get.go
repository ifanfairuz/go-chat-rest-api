package connection

import (
	"gorm.io/gorm"
)

// Get return Database
/**
* get database connection
* return *gorm.DB
**/
func Get() *gorm.DB {
	sqlDB, err := database.DB()
	if err = sqlDB.Ping(); err != nil {
		Connect()
	}

	return database
}
