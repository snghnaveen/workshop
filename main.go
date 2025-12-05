package main

import (
	"github.com/snghnaveen/workshop/config"
	"github.com/snghnaveen/workshop/models"
	"github.com/snghnaveen/workshop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect database
	config.ConnectDatabase()

	// Auto-migrate models
	config.DB.AutoMigrate(&models.Urls{})

	router := gin.Default()

	// Routes
	routes.AppRoutes(router)

	// Run server
	router.Run(":8080")
}
