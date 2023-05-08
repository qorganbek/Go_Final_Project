package models

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `json:"password"`

	Role   enums.Role   `json:"role"`
	Gender enums.Gender `json:"gender"`

	Favorite       Favorite        `json:"favorite"`
	Chats          []Chat          `json:"chats" gorm:"foreignKey:UserID"`
	Advertisements []Advertisement `json:"advertisements" gorm:"foreignKey:UserID"`
}

type CreateUserInput struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `json:"password"`

	Role   enums.Role   `json:"role"`
	Gender enums.Gender `json:"gender"`
}

type LoginUserInput struct {
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `json:"password"`
}
