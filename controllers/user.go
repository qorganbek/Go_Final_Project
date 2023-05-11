package controllers

import (
	final_project "github.com/ZhanserikKalmukhambet/Go_Final_Project"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/middleware"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfUsers(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)
	if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Admin."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByID(c *gin.Context) {
	var user models.User

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Admin."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserByID(c *gin.Context) {
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Admin."})
		return
	}

	initializers.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}

func ListOfUserChats(c *gin.Context) {
	UserID := int(middleware.GetPayloadFromToken(c)["userID"].(float64))

	var chats []models.Chat

	if err := initializers.DB.Where("user_id = ?", UserID).First(&chats).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, chats)
}

func ListOfUserFavoriteItems(c *gin.Context) {
	if final_project.IsAuthorizedOrReadOnly(c) == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		return
	}

	loggedUser := middleware.GetPayloadFromToken(c)
	if loggedUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't load data from token!"})
		return
	}

	var favoriteItems []models.FavoriteItem

	if err := initializers.DB.Where("user_id = ?", int(loggedUser["userID"].(float64))).First(&favoriteItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if len(favoriteItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": favoriteItems})
}

func ListOfUserAdvertisements(c *gin.Context) {
	if final_project.IsAuthorizedOrReadOnly(c) == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		return
	}

	loggedUser := middleware.GetPayloadFromToken(c)
	if loggedUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't load data from token!"})
		return
	}

	var advertisements []models.Advertisement

	if err := initializers.DB.Where("user_id = ?", int(loggedUser["userID"].(float64))).First(&advertisements).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if len(advertisements) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisements})
}
