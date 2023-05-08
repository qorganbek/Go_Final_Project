package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Mark           string          `json:"mark"`
	Modell         string          `json:"modell" gorm:"unique_index"`
	CategoryID     int             `json:"categoryID"`
	Advertisements []Advertisement `json:"advertisements" gorm:"foreignKey:CarID; constraint:OnDelete:CASCADE"`
}
