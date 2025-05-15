package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rest-api-event/db"
	"github.com/rest-api-event/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
