package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/event", getEvent)
	server.Run(":8081")
}

func getEvent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Fhd": "fzi"})
}
