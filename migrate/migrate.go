package main

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Car{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Message{})
	initializers.DB.AutoMigrate(&models.Advertisement{})
	initializers.DB.AutoMigrate(&models.AdvertisementImage{})
	initializers.DB.AutoMigrate(&models.Favorite{})
	initializers.DB.AutoMigrate(&models.Complaint{})
	initializers.DB.AutoMigrate(&models.Category{})
	initializers.DB.AutoMigrate(&models.Complaint{})
	initializers.DB.AutoMigrate(&models.Address{})
}
