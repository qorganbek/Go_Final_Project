package models

import (
	"gorm.io/gorm"
)

type UserV1 struct {
	gorm.Model
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Chats []ChatV1 `json:"chats" gorm:"foreignKey:UserID; null=true"` // one-to-many relationship with Message
}

type ChatV1 struct {
	gorm.Model
	AdvertisementID uint        `json:"advertisementID"`
	UserID          uint        `json:"userID"`
	Messages        []MessageV1 `json:"messages" gorm:"foreignKey:ChatID; null=true"`
}

type MessageV1 struct {
	gorm.Model
	Text   string `json:"text"`
	ChatID uint   `json:"chatID"`
}
