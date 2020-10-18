package connection

/**
* get configuration database
* return Config type
**/
func getConfig() *Config {
	return &Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "1q1q2w1q",
		Dbname:   "go_chat",
		Charset:  "utf8mb4",
	}
}
