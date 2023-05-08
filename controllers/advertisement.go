package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfAdvertisements(c *gin.Context) {
	var advertisement []models.Advertisement

	initializers.DB.Find(&advertisement)
	if len(advertisement) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

func CreateAdvertisement(c *gin.Context) {
	var advertisement models.Advertisement

	if err := c.ShouldBindJSON(&advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

func GetAdvertisementByID(c *gin.Context) {
	var advertisement models.Advertisement

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

func UpdateAdvertisementByID(c *gin.Context) {
	var advertisement models.Advertisement
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&advertisement)
	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

func DeleteAdvertisementByID(c *gin.Context) {
	var advertisement models.Advertisement
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&advertisement)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
