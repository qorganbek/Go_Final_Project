package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/middleware"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFavoriteItem(c *gin.Context) {
	var favoriteItem models.FavoriteItem

	if err := c.ShouldBindJSON(&favoriteItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !final_project.IsAuthorizedOrReadOnly(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
		return
	}

	loggedUser := middleware.GetPayloadFromToken(c)
	if loggedUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't load data from token!"})
		return
	}

	favoriteItem.UserID = int(loggedUser["userID"].(float64))

	if err := initializers.DB.Create(&favoriteItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": favoriteItem})
}

func DeleteFavoriteItemByID(c *gin.Context) {
	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
		return
	}

	var favoriteItem models.FavoriteItem
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&favoriteItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&favoriteItem)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
