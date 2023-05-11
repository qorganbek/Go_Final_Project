package controllers

import (
	final_project "github.com/ZhanserikKalmukhambet/Go_Final_Project"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/middleware"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfAdvertisements(c *gin.Context) {
	var advertisement []models.Advertisement

	initializers.DB.Find(&advertisement)
	if len(advertisement) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data."})
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

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are unauthorized."})
		return
	}

	userID := int(middleware.GetPayloadFromToken(c)["userID"].(float64))
	advertisement.UserID = userID

	if err := initializers.DB.Create(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

func GetAdvertisementByID(c *gin.Context) {
	var advertisement models.Advertisement

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

func UpdateAdvertisementByID(c *gin.Context) {
	var advertisement models.Advertisement
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found."})
		return
	}

	if err := c.ShouldBindJSON(&advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	ownerID := int(middleware.GetPayloadFromToken(c)["userID"].(float64))
	isAdmin := final_project.IsAdmin(c)

	if ownerID != advertisement.UserID && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not owner or admin."})
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

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User unauthorized."})
		return
	}

	ownerID := int(middleware.GetPayloadFromToken(c)["userID"].(float64))
	isAdmin := final_project.IsAdmin(c)

	if ownerID != advertisement.UserID && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Owner or Admin."})
		return
	}

	initializers.DB.Delete(&advertisement)
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted."})
}
