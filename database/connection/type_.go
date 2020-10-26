package connection

// Config structure
/**
* Database config structure
**/
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Charset  string
	Timezone string
}
