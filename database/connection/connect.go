package connection

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect return bool
/**
* connect to database
* return boolean
**/
func Connect() bool {
	var config *Config = getConfig()
	var dbms string = os.Getenv("DBMS")
	var dsn string = getDsn(config, dbms)
	var err error
	var db *gorm.DB

	if dbms == "postgres" {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Print("Database Connected.")
	database = db
	return true
}
