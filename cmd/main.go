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
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	routes.CarRoutes(router)
	routes.CategoryRoutes(router)
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.ChatRoutes(router)
	routes.AdvertisementRoutes(router)
	routes.AdvertisementImageRoutes(router)
	routes.MessageRoutes(router)
	routes.AddressRoutes(router)

	log.Fatal(router.Run(":8000"))
}
