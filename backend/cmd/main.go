package main

import (
	"internal-office-backend/config"
	"internal-office-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDatabase()

	r := gin.Default()

	api := r.Group("/api")
	{
		routes.RegisterRoutes(api, db)
	}

	r.Run(":8080")
}