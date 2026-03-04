package main

import (
	"internal-office-backend/config"
	"internal-office-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	r := gin.Default()

	api := r.Group("/api")
	{
		handler.RegisterRoutes(api)
	}

	r.Run(":8080")
}