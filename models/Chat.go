package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	AdvertisementID int       `json:"advertisementID"`
	UserID          int       `json:"userID"`
	Messages        []Message `json:"messages" gorm:"foreignKey:ChatID"`
}

type Message struct {
	gorm.Model
	ChatID int    `json:"chatID"`
	Text   string `json:"text"`
	Read   bool   `json:"read"`
}
