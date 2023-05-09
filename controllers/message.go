package controllers

import (
	final_project "github.com/ZhanserikKalmukhambet/Go_Final_Project"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/middleware"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfMessages(c *gin.Context) {
	var messages []models.Message

	initializers.DB.Find(&messages)
	if len(messages) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": messages})
}

func CreateMessage(c *gin.Context) {
	var message models.Message

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func GetMessageByID(c *gin.Context) {
	var message models.Message

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": message})
}

func UpdateMessageByID(c *gin.Context) {
	var message models.Message

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User unauthorized"})
		return
	}

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var chat models.Chat
	err := initializers.DB.Where("id = ?", message.ChatID).Find(&chat)
	if err != nil {
		panic(err)
		return
	}

	ownerID := int(middleware.GetUserDetailsFromToken(c)["userID"].(float64))
	isAdmin := final_project.IsAdmin(c)

	if !isAdmin && ownerID != chat.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not admin or owner"})
		return
	}

	initializers.DB.Save(&message)
	c.JSON(http.StatusOK, gin.H{"data": message})
}

func DeleteMessageByID(c *gin.Context) {
	var message models.Message

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User unauthorized"})
		return
	}

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var chat models.Chat
	err := initializers.DB.Where("id = ?", message.ChatID).Find(&chat)
	if err != nil {
		panic(err)
		return
	}

	ownerID := int(middleware.GetUserDetailsFromToken(c)["userID"].(float64))

	if ownerID != chat.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not admin or owner"})
		return
	}

	initializers.DB.Delete(&message)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
