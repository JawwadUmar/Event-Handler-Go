package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utility"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot get the requested data, try again later",
		})

		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request data",
		})

		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event, Try again later",
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created :)",
		"event":   event,
	})
}

func emptyTable(context *gin.Context) {
	err := utility.TruncateTable("events")

	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not truncate events table, Try again later",
			"err":     err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Events table emptied successfully",
	})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parses event id",
			"err":     err,
		})

		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch this event with this id",
			"err":     err,
		})

		return
	}

	context.JSON(http.StatusOK, event)

}

func delteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not fetch this event with this id",
			"err":     err,
		})

		return
	}

	eventModel, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch this event with this id",
			"err":     err,
		})

		return
	}

	userId := context.GetInt64("userId")

	if eventModel.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "This user is not allowed to delete this event"})
		return
	}

	err = models.DeleteEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch this event with this id",
			"err":     err,
		})
	}
}

func updateEvent(context *gin.Context) {
	//first extract the id --> integer
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}

	eventModel, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})

		return
	}

	userId := context.GetInt64("userId")

	if eventModel.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "This user is not allowed to update",
		})

		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	//updated event is updated :)
	updatedEvent.Id = id

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}
