package models

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phoneNumber" gorm:"unique_index"`
	Password    string `json:"password"`

	Role   enums.Role   `json:"role"`
	Gender enums.Gender `json:"gender"`

	Favorite       Favorite        `json:"favorite"`
	Advertisements []Advertisement `json:"advertisements" gorm:"foreignKey:UserID"`

	Messages       []Message       `json:"messages" gorm:"foreignKey:UserID"`
	Advertisements []Advertisement `json:"advertisements" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
}
