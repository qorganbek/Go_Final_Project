package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfAdvertisementImages(c *gin.Context) {
	var advertisementImage []models.AdvertisementImage

	initializers.DB.Find(&advertisementImage)
	if len(advertisementImage) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisementImage})
}

func CreateAdvertisementImage(c *gin.Context) {
	var advertisementImage models.AdvertisementImage

	if err := c.ShouldBindJSON(&advertisementImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&advertisementImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisementImage})
}

func GetAdvertisementImageByID(c *gin.Context) {
	var advertisementImage models.AdvertisementImage

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&advertisementImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": advertisementImage})
}

func UpdateAdvertisementImageByID(c *gin.Context) {
	var advertisementImage models.AdvertisementImage
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&advertisementImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&advertisementImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&advertisementImage)
	c.JSON(http.StatusOK, gin.H{"data": advertisementImage})
}

func DeleteAdvertisementImageByID(c *gin.Context) {
	var advertisementImage models.AdvertisementImage
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&advertisementImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&advertisementImage)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
