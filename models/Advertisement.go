package models

import "gorm.io/gorm"

type Advertisement struct {
	gorm.Model
	UserID int `json:"userID"`
	CarID  int `json:"carID"`

	Year           int     `json:"year"`
	EngineCapacity float64 `json:"engineCapacity" gorm:"check:engine_capacity > 0"`
	Millage        int     `json:"millage" gorm:"check:millage > 0"`
	Description    string  `json:"description" gorm:"default:No Description"`
	Transmisson    string  `json:"transmisson"`
	Color          string  `json:"color"`
	Price          float64 `json:"price" gorm:"check:price > 0"`
	IsTop          bool    `json:"isTop" gorm:"default:false"`
	Raiting        int     `json:"raiting" gorm:"default:0"`
	Address        string  `json:"address"`

	Chats []Chat `json:"chats" gorm:"foreignKey:AdvertisementID"`

	Images      []AdvertisementImage `json:"images" gorm:"AdvertisementID; constraint:OnDelete:CASCADE"`
	Complaintes []Complaint          `json:"complaintes" gorm:"foreignKey:AdvertisementID; constraint:OnDelete:CASCADE"`
}

type AdvertisementImage struct {
	gorm.Model
	AdvertisementID int
	Image           string `json:"image"`
}
