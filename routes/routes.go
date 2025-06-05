package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rest-api-event/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	// Define your routes here
	server.GET("/events", getEvents) // GET POST PUT DELETE PATCH
	server.GET("/events/:id", getEvent)

	// Protected routes
	authenticated := server.Group("/api")
	authenticated.Use(middleware.Authenticate) // Apply authentication middleware to this group
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", RegisterforEvents)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
