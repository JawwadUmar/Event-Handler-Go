package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8081")
}

func getEvent(context *gin.Context) {
	events := models.GetAllEvents()
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

	event.Id = 1
	event.UserId = 1

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created :)",
		"event":   event,
	})

	event.Save()

}
