package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	//userId, eventId
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Some problem while converting event id to int"})
		return
	}

	userId := context.GetInt64("userId")

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The event with this id does not exist"})
		return
	}

	registration := models.Register{
		UserId:  userId,
		EventId: eventId,
	}

	err = registration.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save this"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Added to the registeration Table"})

}

func getRegistrations(context *gin.Context) {
	registerations, err := models.GetAllRegisterations()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot get the requested data, try again later",
		})

		return
	}
	context.JSON(http.StatusOK, registerations)
}
