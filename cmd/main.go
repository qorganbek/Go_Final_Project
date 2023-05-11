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

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	routes.CategoryRoutes(router)
	routes.CarRoutes(router)

	routes.AdvertisementRoutes(router)
	routes.FavoriteItemRoutes(router)

	routes.ChatRoutes(router)
	routes.MessageRoutes(router)
	routes.ComplaintRoutes(router)

	log.Fatal(router.Run(":8000"))
}
