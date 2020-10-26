package models

import (
	"time"
)

// User model
type User struct {
	Email        string    `gorm:"primaryKey" json:"email" xml:"email" form:"email"`
	StatusOnline bool      `json:"status_online" xml:"status_online" form:"status_online"`
	LastOnline   int64     `json:"last_online" xml:"last_online" form:"last_online"`
	Image        string    `json:"image" xml:"image" form:"image"`
	CreatedAt    time.Time `json:"created_at" xml:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" xml:"updated_at" form:"updated_at"`
}
