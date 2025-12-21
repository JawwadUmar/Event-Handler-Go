package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)
	server.DELETE("/events", emptyTable)
}
