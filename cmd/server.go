package main

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	//r.GET("/users", GetAllUsers)
	//r.POST("/users", CreateUser)
	//r.GET("/users/:id", GetUser)
	//r.PUT("/users/:id", UpdateUser)
	//r.DELETE("/users/:id", DeleteUser)
	//
	//r.GET("/chats", GetAllChats)
	//r.POST("/chats", CreateChat)
	//r.GET("/chats/:id", GetChat)
	//r.PUT("/chats/:id", UpdateChat)
	//r.DELETE("/chats:id", DeleteChat)
	//
	//r.GET("/messages", getAllMessages)
	//r.POST("/messages", createMessage)
	//r.GET("/messages/:id", getMessage)
	//r.PUT("/messages/:id", updateMessage)
	//r.DELETE("/messages/:id", deleteMessage)

	r.Run()
}
