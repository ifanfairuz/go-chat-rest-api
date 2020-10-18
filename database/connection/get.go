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
	return database
}
