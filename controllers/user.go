package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListOfUser(c *gin.Context) {
	var user []models.User

	initializers.DB.Find(&user)
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUserByID(c *gin.Context) {
	var user models.User

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUserByID(c *gin.Context) {
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserByID(c *gin.Context) {
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}

func UserChats(c *gin.Context) {
	UserID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	var chats []models.Chat
	var myChats []models.Chat
	initializers.DB.Find(&chats)

	for _, val := range chats {
		if val.UserID == UserID {
			myChats = append(myChats, val)
		}
	}

	c.JSON(http.StatusOK, myChats)
}
