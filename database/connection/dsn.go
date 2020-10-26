package connection

import (
	"fmt"
)

/**
* get dsn for connect to database
**/
func getDsn(config *Config, dbms string) string {
	if dbms == "postgres" {
		return getDsnPosgres(config)
	}

	return getDsnMysql(config)
}

/**
* get dsn for connect to database mysql
**/
func getDsnMysql(config *Config) string {
	format := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
	return fmt.Sprintf(format, config.User, config.Password, config.Host, config.Port, config.Dbname, config.Charset)
}

/**
* get dsn for connect to database postgresql
**/
func getDsnPosgres(config *Config) string {
	format := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s"
	return fmt.Sprintf(format, config.Host, config.Port, config.User, config.Password, config.Dbname, config.Timezone)
}
