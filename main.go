package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/utility"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)
	server.DELETE("/events", emptyTable)

	server.Run(":8081")
}

func getEvent(context *gin.Context) {
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

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event, Try again later",
		})
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created :)",
		"event":   event,
	})
}

func emptyTable(context *gin.Context) {
	err := utility.TruncateTable("events")

	if err != nil {
		fmt.Print(err)
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
