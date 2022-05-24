package main

import (
	"restful_api/controllers"
	"restful_api/database"

	"github.com/gin-gonic/gin"
)



func main() {

	router := gin.Default()

	database.ConnectDB()

	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.GET("/messages", controllers.GetAllMessages)
	router.GET("/messages/:id", controllers.GetMessage)
	router.POST("/users", controllers.CreateUser)
	router.POST("/messages", controllers.CreateMessage)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.DELETE("/messages/:id", controllers.DeleteMessage)
	router.PATCH("/users/:id", controllers.UpdateUser)
	router.PATCH("/messages/:id", controllers.UpdateMessage)

	router.Run(":4040")
}

// password to the postgres 12345
