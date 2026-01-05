package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	registerRoutesForEvent(server)
	registerRoutesForUser(server)

}

func registerRoutesForEvent(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// authGroup := server.Group("/")
	// authGroup.
	server.POST("/events", middlewares.Authentication, createEvent)
	server.DELETE("/events/:id", delteEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events", emptyTable)
}

func registerRoutesForUser(server *gin.Engine) {
	server.GET("/users", getUser)
	server.POST("/signup", createUser)
	server.POST("/login", login)
}
