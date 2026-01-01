package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	registerRoutesForEvent(server)
	registerRoutesForUser(server)

}

func registerRoutesForEvent(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.DELETE("/events/:id", delteEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events", emptyTable)
}

func registerRoutesForUser(server *gin.Engine) {
	server.GET("/users", getUser)
	server.POST("/signup", createUser)
}
