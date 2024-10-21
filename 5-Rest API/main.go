package main

import (
	"example.com/rest/db"
	"example.com/rest/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	server := gin.Default()
	handlers.AllRoutes(server)
	server.Run(":8080") 
}

