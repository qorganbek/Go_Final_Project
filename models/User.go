package models

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname   string `json:"firstname"`
	Secondname  string `json:"secondname"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`

	Role   enums.Role   `json:"role"`
	Gender enums.Gender `json:"gender"`

	Favorite       Favorite        `json:"favorite"`
	Advertisements []Advertisement `json:"advertisements" gorm:"foreignKey:UserID"`
}
