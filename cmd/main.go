package main

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := gin.New()

	routes.AdminCarRoutes(router)
	routes.AdminCategoryRoutes(router)

	log.Fatal(router.Run(":8080"))
}
