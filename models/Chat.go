package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Messages []Message `json:"messages" gorm:"foreignKey:ChatId"`
}

type Message struct {
	gorm.Model
	UserID uint   `json:"userID"`
	ChatId uint   `json:"chatID"`
	Text   string `json:"text"`
	IsRead bool   `json:"IsRead"`
}
