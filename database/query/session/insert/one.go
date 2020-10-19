package insert

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"
	Models "github.com/ifanfairuz/go-chat-rest-api/database/models"
)

// One session
func One(email string, socket string) (Models.Session, error) {
	session := Models.Session{
		Email:    email,
		SocketID: socket,
	}
	result := Connection.Get().Create(&session)

	Connection.Close()

	return session, result.Error
}
