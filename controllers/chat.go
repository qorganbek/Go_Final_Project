package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfChats(c *gin.Context) {
	var chats []models.Chat

	initializers.DB.Find(&chats)
	if len(chats) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chats})
}

func CreateChat(c *gin.Context) {
	var chat models.Chat

	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&chat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chat})
}

func GetChatByID(c *gin.Context) {
	var chat models.Chat

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&chat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": chat})
}

func UpdateChatByID(c *gin.Context) {
	var chat models.Chat
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&chat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&chat)
	c.JSON(http.StatusOK, gin.H{"data": chat})
}

func DeleteChatByID(c *gin.Context) {
	var chat models.Chat
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&chat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&chat)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
