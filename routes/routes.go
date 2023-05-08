package routes

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/controllers"
	"github.com/gin-gonic/gin"
)

func AdminCategoryRoutes(routes *gin.Engine) {
	routes.GET("/categories", controllers.ListOfCategories)
	routes.GET("/categories/:id", controllers.GetCategoryByID)
	routes.POST("/categories", controllers.CreateCategory)
	routes.PATCH("/categories/:id", controllers.UpdateCategoryByID)
	routes.DELETE("/categories/:id", controllers.DeleteCategoryByID)
}

func AdminCarRoutes(routes *gin.Engine) {
	routes.GET("/cars", controllers.ListOfCars)
	routes.GET("/cars/:id", controllers.GetCarByID)
	routes.POST("/cars", controllers.CreateCar)
	routes.PATCH("/cars/:id", controllers.UpdateCarByID)
	routes.DELETE("/cars/:id", controllers.DeleteCarByID)
}
