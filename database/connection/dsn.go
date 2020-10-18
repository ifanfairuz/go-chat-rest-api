package connection

import (
	"fmt"
)

/**
* get dsn for connect to database
**/
func getDsn(config *Config) string {
	format := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
	return fmt.Sprintf(format, config.User, config.Password, config.Host, config.Port, config.Dbname, config.Charset)
}
