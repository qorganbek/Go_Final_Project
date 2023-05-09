package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfFavorite(c *gin.Context) {
	var favorite []models.Favorite

	initializers.DB.Find(&favorite)
	if len(favorite) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": favorite})
}

func GetFavoriteByID(c *gin.Context) {
	var favorite models.Favorite

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&favorite).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": favorite})
}

func UpdateFavoriteByID(c *gin.Context) {
	var favorite models.Favorite

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&favorite).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&favorite)
	c.JSON(http.StatusOK, gin.H{"data": favorite})
}
