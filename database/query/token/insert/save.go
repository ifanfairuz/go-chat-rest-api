package insert

import (
	"time"

	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

func save(email string, token string) bool {
	// epired in one day
	expired := time.Now().AddDate(0, 0, 1)
	result := Connection.Get().Create(&Models.Token{
		Email:     email,
		Token:     token,
		ExpiredAt: expired.Unix(),
	})

	return (result.Error == nil)
}
