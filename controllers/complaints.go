package controllers

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfComplaints(c *gin.Context) {
	var complaints []models.Complaint

	initializers.DB.Find(&complaints)
	if len(complaints) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": complaints})
}

func CreateComplaint(c *gin.Context) {
	var complaint models.Complaint

	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedUser := GetUserDetailsFromToken(c)
	complaint.UserID = int(loggedUser["userID"].(float64))

	if err := initializers.DB.Create(&complaint).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": complaint})
}

func GetComplaintByID(c *gin.Context) {
	var complaint models.Complaint

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&complaint).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": complaint})
}

func DeleteComplaintByID(c *gin.Context) {
	var complaint models.Complaint
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&complaint).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&complaint)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
