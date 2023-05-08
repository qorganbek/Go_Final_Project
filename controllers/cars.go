package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfCars(c *gin.Context) {
	var cars []models.Car

	initializers.DB.Find(&cars)
	if len(cars) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

func CreateCar(c *gin.Context) {
	var car models.Car

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": car})
}

func GetCarByID(c *gin.Context) {
	var car models.Car

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})
}

func UpdateCarByID(c *gin.Context) {
	var car models.Car
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&car)
	c.JSON(http.StatusOK, gin.H{"data": car})
}

func DeleteCarByID(c *gin.Context) {
	var car models.Car
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&car)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
