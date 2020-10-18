package models

import (
	"time"
)

// Session model
type Session struct {
	ID        uint      `gorm:"primaryKey, autoIncrement" json:"id" xml:"id" form:"id"`
	Email     string    `gorm:"index" json:"email" xml:"email" form:"email"`
	SocketID  string    `json:"socket_id" xml:"socket_id" form:"socket_id"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" form:"updated_at"`
}
