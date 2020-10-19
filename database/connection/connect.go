package connection

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect return bool
/**
* connect to database
* return boolean
**/
func Connect() bool {
	var config *Config = getConfig()
	var dsn string = getDsn(config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return false
	}

	database = db
	return true
}
