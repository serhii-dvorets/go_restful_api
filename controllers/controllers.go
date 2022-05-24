package controllers

import (
	"restful_api/database"
	"restful_api/models"

	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUser(context *gin.Context) {
	var users models.User
	id := context.Param("id")

	getUser := database.ConnectDB().First(&users, id)
	err := getUser.Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func GetMessage(context *gin.Context) {
	var messages models.Message
	id := context.Param("id")

	getMessage := database.ConnectDB().First(&messages, id)
	err := getMessage.Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"messages": messages})
	}
}

func CreateUser(context *gin.Context) {
	var input models.User
	err := context.ShouldBindJSON(&input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		user := models.User{Name: input.Name}
		database.ConnectDB().Create(&user)

		context.JSON(http.StatusOK, gin.H{"newUser": user})
	}
}

func CreateMessage(context *gin.Context) {
	var input models.Message
	err := context.ShouldBindJSON(&input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		message := models.Message{
			Text:   input.Text,
			UserID: input.UserID,
		}
		database.ConnectDB().Create(&message)
	}

}

func GetAllUsers(context *gin.Context) {
	var users []models.User

	getAllUsers := database.ConnectDB().Find(&users)
	err := getAllUsers.Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"users": users})
	}

}

func GetAllMessages(context *gin.Context) {
	var messages []models.Message

	getAllMessages := database.ConnectDB().Find(&messages)
	err := getAllMessages.Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	}
}

func DeleteUser(context *gin.Context) {
	var users models.User
	id := context.Param("id")

	deleteUser := database.ConnectDB().Delete(&users, id)
	err := deleteUser.Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func DeleteMessage(context *gin.Context) {
	var messages models.Message
	id := context.Param("id")

	deleteMessage := database.ConnectDB().Delete(&messages, id)
	err := deleteMessage.Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"messages": messages})
	}
}

func UpdateUser(context *gin.Context) {

	var user models.User
	if err := database.ConnectDB().Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Such user doesn't exist"})
		return
	}

	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.ConnectDB().Model(&user).Updates(input)
}

func UpdateMessage(context *gin.Context) {

	var message models.Message
	if err := database.ConnectDB().Where("id = ?", context.Param("id")).First(&message).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Such message doesn't exist"})
		return
	}

	var input models.Message
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.ConnectDB().Model(&message).Updates(input)
}
