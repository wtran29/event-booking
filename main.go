package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wtran29/event-booking/db"
	"github.com/wtran29/event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
