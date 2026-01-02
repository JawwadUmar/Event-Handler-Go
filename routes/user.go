package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body",
		})

		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save to database",
		})

		return
	}

	context.JSON(http.StatusOK, user)
}

func getUser(context *gin.Context) {
	var user []models.User
	user, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch all users"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body",
		})

		return
	}

	err = user.ValidateCredential()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not validate credentials",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully login"})
}
