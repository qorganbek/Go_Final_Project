package controllers

import (
	final_project "github.com/ZhanserikKalmukhambet/Go_Final_Project"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfCategories(c *gin.Context) {
	var categories []models.Category

	initializers.DB.Find(&categories)
	if len(categories) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "user not admin"})
		return
	}

	if err := initializers.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func GetCategoryByID(c *gin.Context) {
	var category models.Category

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func UpdateCategoryByID(c *gin.Context) {
	var category models.Category
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "user not admin"})
		return
	}

	initializers.DB.Save(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategoryByID(c *gin.Context) {
	var category models.Category
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "user not admin"})
		return
	}

	initializers.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
