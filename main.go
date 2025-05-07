package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // GET POST PUT DELETE PATCH

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusAccepted, gin.H{"message": "Hello World"})
}
