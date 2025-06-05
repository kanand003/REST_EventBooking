package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-event/models"
)

func RegisterforEvents(context *gin.Context) {
	userId := context.GetInt64("userid")                          // Get user ID from context set by middleware
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Convert string to int
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registration successful", "event": event})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userid")                          // Get user ID from context set by middleware
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Convert string to int

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled", "event": event})
}
