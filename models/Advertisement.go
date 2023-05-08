package models

import "gorm.io/gorm"

type Advertisement struct {
	gorm.Model
	UserID uint `json:"userID"`
	CarID  uint `json:"carID"`

	Year           int     `json:"year"`
	EngineCapacity float64 `json:"engineCapacity" gorm:"check:engineCapacity > 0"`
	Millage        int     `json:"millage" gorm:"check:millage > 0"`
	Description    string  `json:"description" gorm:"default:No Description"`
	Transmisson    string  `json:"transmisson"`
	Color          string  `json:"color"`
	Price          float64 `json:"price" gorm:"check:price > 0"`
	IsTop          bool    `json:"isTop" gorm:"default:false"`
	Raiting        int     `json:"raiting" gorm:"default:0"`
	Address        Address `json:"address"`

	Chats       []ChatV1             `json:"chats" gorm:"AdvertisementID"`
	Images      []AdvertisementImage `json:"images" gorm:"AdvertisementID"`
	Complaintes []Complaint          `json:"complaintes" gorm:"foreignKey:AdvertisementID"`
	Images      []AdvertisementImage `json:"images" gorm:"AdvertisementID; constraint:OnDelete:CASCADE"`
	Complaintes []Complaint          `json:"complaintes" gorm:"foreignKey:AdvertisementID; constraint:OnDelete:CASCADE"`
}

type AdvertisementImage struct {
	gorm.Model
	AdvertisementID uint
	Image           string `json:"image"`
}
