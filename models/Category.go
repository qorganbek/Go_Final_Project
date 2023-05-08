package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"unique_index"`
	Cars []Car  `gorm:"foreignKey:CategoryID; constraint:OnDelete:CASCADE"`
}
