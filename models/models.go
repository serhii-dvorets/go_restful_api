package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:text"`
	Messages []Message 
}

type Message struct {
	gorm.Model
	Text   string `json:"text" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
}
