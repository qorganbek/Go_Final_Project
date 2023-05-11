package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/middleware"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListOfChats(c *gin.Context) {
	var chats []models.Chat

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Admin."})
		return
	}

	initializers.DB.Find(&chats)
	if len(chats) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chats})
}

func CreateChat(c *gin.Context) {
	var chat models.Chat

	isAuth := final_project.IsAuthorizedOrReadOnly(c)
	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedUser := middleware.GetPayloadFromToken(c)
	chat.UserID = int(loggedUser["userID"].(float64))

	if err := initializers.DB.Create(&chat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chat})
}

func GetChatByID(c *gin.Context) {
	var chat models.Chat

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

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	isAdmin := final_project.IsAdmin(c)
	ownerID := int(middleware.GetPayloadFromToken(c)["userID"].(float64))

	if !isAdmin && ownerID != chat.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Owner or Admin."})
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

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	isAdmin := final_project.IsAdmin(c)
	ownerID := int(middleware.GetPayloadFromToken(c)["userID"].(float64))

	if !isAdmin && ownerID != chat.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are nor Admin or Owner"})
		return
	}

	initializers.DB.Delete(&chat)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}

func ChatMessages(c *gin.Context) {
	if final_project.IsAuthorizedOrReadOnly(c) == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user."})
		return
	}

	var messages []models.Message

	chatID, _ := strconv.Atoi(c.Param("id"))

	if err := initializers.DB.Where("chat_id = ?", chatID).First(&messages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Chat messages": messages})
}
