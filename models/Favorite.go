package models

import "gorm.io/gorm"

type FavoriteItem struct {
	gorm.Model
	UserID          int `json:"userID"`
	AdvertisementID int `json:"advertisementID" gorm:"constraint:OnDelete:CASCADE"`
}
