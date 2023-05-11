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
	routes.POST("users/signUp", middleware.SignUp)
	routes.POST("users/signIn", middleware.SignIn)
	routes.GET("users/validate", middleware.RequireAuth, middleware.ValidateUser)
	routes.POST("users/signOut", middleware.SignOut)
}

func UserRoutes(routes *gin.Engine) {
	routes.GET("/users", controllers.ListOfUsers)
	routes.GET("/users/:id", controllers.GetUserByID)
	routes.DELETE("/users/:id", controllers.DeleteUserByID)

	routes.GET("/myChats", controllers.ListOfUserChats)
	routes.GET("/myFavorites", controllers.ListOfUserFavoriteItems)
	routes.GET("/myAdvertisements", controllers.ListOfUserAdvertisements)

	//routes.GET("/users/loggedUserDetail", controllers.GetUserDetailsFromToken)
}

func ChatRoutes(routes *gin.Engine) {
	routes.GET("/chats", controllers.ListOfChats)
	routes.GET("/chats/:id", controllers.GetChatByID)
	routes.POST("/chats", controllers.CreateChat)
	routes.PATCH("/chats/:id", controllers.UpdateChatByID)
	routes.DELETE("/chats/:id", controllers.DeleteChatByID)
	routes.GET("/chats/:id/messages", controllers.ChatMessages)
}

func MessageRoutes(routes *gin.Engine) {
	routes.GET("/messages", controllers.ListOfMessages)
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
	routes.GET("advertisements/filterByYear/:min/:max", filters.ListOfAdvertisementsByYear)
}

func ComplaintRoutes(routes *gin.Engine) {
	routes.GET("/complaints", controllers.ListOfComplaints)
	routes.GET("/complaints/:id", controllers.GetComplaintByID)
	routes.POST("/complaints", controllers.CreateComplaint)
}

func FavoriteItemRoutes(routes *gin.Engine) {
	routes.POST("/favorites", controllers.CreateFavoriteItem)
	routes.DELETE("/favorites/:id", controllers.DeleteFavoriteItemByID)
}
