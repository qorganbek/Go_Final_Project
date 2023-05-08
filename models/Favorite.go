package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID uint `json:"userID"`

	Advertisements []Advertisement `json:"advertisements" gorm:"many2many:favorite_advertisements; constraint:OnDelete:CASCADE"`
}
