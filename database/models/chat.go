package models

import (
	"time"
)

// Chat model
type Chat struct {
	ID        uint      `gorm:"primaryKey, autoIncrement" json:"id" xml:"id" form:"id"`
	From      string    `gorm:"index" json:"from" xml:"from" form:"from"`
	To        string    `gorm:"index" json:"to" xml:"to" form:"to"`
	Text      string    `gorm:"not null" json:"text" xml:"text" form:"text"`
	Status    int8      `gorm:"default:0" json:"status" xml:"status" form:"status"`
	SendAt    int64     `json:"send_at" xml:"send_at" form:"send_at"`
	SentAt    int64     `json:"sent_at" xml:"sent_at" form:"sent_at"`
	ReceiveAt int64     `json:"receive_at" xml:"receive_at" form:"receive_at"`
	ReadAt    int64     `json:"read_at" xml:"read_at" form:"read_at"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" form:"updated_at"`
}
