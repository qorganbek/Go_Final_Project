package models

import "gorm.io/gorm"

type Advertisement struct {
	gorm.Model
	UserID uint `json:"userID"`
	CarID  uint `json:"carID"`

	Year           int     `json:"year"`
	EngineCapacity float64 `json:"engineCapacity"`
	Millage        int     `json:"millage"`
	Description    string  `json:"description" gorm:"default:No Description"`
	Transmisson    string  `json:"transmisson"`
	Color          string  `json:"color"`
	Price          float64 `json:"price"`
	IsTop          bool    `json:"isTop"`
	Raiting        int     `json:"raiting"`
	Address        Address `json:"address"`

	Chats       []ChatV1             `json:"chats" gorm:"AdvertisementID"`
	Images      []AdvertisementImage `json:"images" gorm:"AdvertisementID"`
	Complaintes []Complaint          `json:"complaintes" gorm:"foreignKey:AdvertisementID"`
}

type AdvertisementImage struct {
	gorm.Model
	AdvertisementID uint
	Image           string `json:"image"`
}
