package models

import (
	"time"
)

// APIToken Model
type APIToken struct {
	Token     string    `gorm:"primaryKey" json:"token" xml:"token" form:"token"`
	Origin    string    `gorm:"size:255" json:"origin" xml:"origin" form:"origin"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" form:"updated_at"`
}
