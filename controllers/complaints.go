package controllers

import (
	final_project "github.com/ZhanserikKalmukhambet/Go_Final_Project"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/middleware"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfComplaints(c *gin.Context) {
	if final_project.IsAdmin(c) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can not see complaints! You are not admin!"})
		return
	}

	var complaints []models.Complaint

	initializers.DB.Find(&complaints)
	if len(complaints) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": complaints})

}

func CreateComplaint(c *gin.Context) {
	if final_project.IsAuthorizedOrReadOnly(c) == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please, log in to write complaint on advertisement!"})
		return
	}

	var complaint models.Complaint

	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedUser := middleware.GetUserDetailsFromToken(c)
	complaint.UserID = int(loggedUser["userID"].(float64))

	if err := initializers.DB.Create(&complaint).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": complaint})
}

func GetComplaintByID(c *gin.Context) {
	if final_project.IsAdmin(c) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can not see specific complaint! You are not admin!"})
		return
	}

	var complaint models.Complaint

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&complaint).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": complaint})
}
