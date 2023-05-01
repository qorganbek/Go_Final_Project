package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	City            string `json:"city"`
	Region          string `json:"region"`
	AdvertisementID uint   `json:"advertisementID"`
}
