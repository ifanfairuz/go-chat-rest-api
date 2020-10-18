package models

import (
	"time"
)

// Token Model
type Token struct {
	Token     string    `gorm:"primaryKey" json:"token" xml:"token" form:"token"`
	Email     string    `gorm:"index" json:"email" xml:"email" form:"email"`
	ExpiredAt int64     `json:"expired_at" xml:"expired_at" form:"expired_at"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" form:"updated_at"`
}
