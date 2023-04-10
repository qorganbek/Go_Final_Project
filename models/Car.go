package models

import (
	"gorm.io/gorm"
	"time"
)

type Car struct {
	gorm.Model
	Mark           string     `json:"mark"`
	Modell         string     `json:"modell"`
	Year           time.Time  `json:"year"`
	Category       string     `json:"category"`
	EngineCapacity float64    `json:"engineCapacity"`
	Color          string     `json:"color"`
	Description    string     `json:"description" gorm:"default:No Description"`
	Transmisson    string     `json:"transmisson"`
	Price          float64    `json:"price"`
	Milage         float64    `json:"milage"`
	Images         []CarImage `gorm:"foreignKey:CarID"`
}

type CarImage struct {
	gorm.Model
	CarID uint
	Image string `json:"image"`
}
