package routes

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/controllers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/filters"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/middleware"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(routes *gin.Engine) {
	routes.GET("/categories", controllers.ListOfCategories)
	routes.GET("/categories/:id", controllers.GetCategoryByID)
	routes.POST("/categories", controllers.CreateCategory)
	routes.PATCH("/categories/:id", controllers.UpdateCategoryByID)
	routes.DELETE("/categories/:id", controllers.DeleteCategoryByID)
}

func CarRoutes(routes *gin.Engine) {
	routes.GET("/cars", controllers.ListOfCars)
	routes.GET("/cars/:id", controllers.GetCarByID)
	routes.POST("/cars", controllers.CreateCar)
	routes.PATCH("/cars/:id", controllers.UpdateCarByID)
	routes.DELETE("/cars/:id", controllers.DeleteCarByID)
}

func AuthRoutes(routes *gin.Engine) {
	routes.POST("users/signup", controllers.SignUp)
	routes.POST("users/login", controllers.SignIn)
	routes.GET("users/validate", middleware.RequireAuth, controllers.Validate)
	routes.POST("users/logout", controllers.Logout)
}

func UserRoutes(routes *gin.Engine) {
	routes.GET("/users", controllers.ListOfUser)
	routes.GET("/users/:id", controllers.GetUserByID)
	routes.POST("/users", controllers.CreateUser)
	routes.PATCH("/users/:id", controllers.UpdateUserByID)
	routes.DELETE("/users/:id", controllers.DeleteUserByID)

	routes.GET("/users/loggedUserDetail", controllers.GetUserDetailsFromToken)
}

func ChatRoutes(routes *gin.Engine) {
	routes.GET("/chats", controllers.ListOfChats)
	routes.GET("/chats/:id", controllers.GetChatByID)
	routes.POST("/chats", controllers.CreateChat)
	routes.PATCH("/chats/:id", controllers.UpdateChatByID)
	routes.DELETE("/chats/:id", controllers.DeleteChatByID)
}

func MessageRoutes(routes *gin.Engine) {
	routes.GET("/messages", controllers.ListOfMessages)
	routes.GET("/messages/:id", controllers.GetMessageByID)
	routes.POST("/messages", controllers.CreateMessage)
	routes.PATCH("/messages/:id", controllers.UpdateMessageByID)
	routes.DELETE("/messages/:id", controllers.DeleteMessageByID)
}

func AdvertisementRoutes(routes *gin.Engine) {
	routes.GET("/advertisements", controllers.ListOfAdvertisements)
	routes.GET("/advertisements/:id", controllers.GetAdvertisementByID)
	routes.POST("/advertisements", controllers.CreateAdvertisement)
	routes.PATCH("/advertisements/:id", controllers.UpdateAdvertisementByID)
	routes.DELETE("/advertisements/:id", controllers.DeleteAdvertisementByID)

	routes.GET("advertisements/filterByPrice/:min/:max", filters.ListOfAdvertisementsByPrice)
	routes.GET("advertisements/filterByYear/:min/:max", filters.ListOfAdvertisementsByYears)
}

func AdvertisementImageRoutes(routes *gin.Engine) {
	routes.GET("/advertisement_images", controllers.ListOfAdvertisementImages)
	routes.GET("/advertisement_images/:id", controllers.GetAdvertisementImageByID)
	routes.POST("/advertisement_images", controllers.CreateAdvertisementImage)
	routes.PATCH("/advertisement_images/:id", controllers.UpdateAdvertisementImageByID)
	routes.DELETE("/advertisement_images/:id", controllers.DeleteAdvertisementImageByID)
}

func AddressRoutes(routes *gin.Engine) {
	routes.POST("/addresses", controllers.CreateAddress)
}

func ComplaintRoutes(routes *gin.Engine) {
	routes.GET("/complaints", controllers.ListOfComplaints)
	routes.GET("/complaints/:id", controllers.GetComplaintByID)
	routes.POST("/complaints", controllers.CreateComplaint)
	routes.DELETE("/complaints/:id", controllers.DeleteComplaintByID)
}
