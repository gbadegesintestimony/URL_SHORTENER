package main

import (
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/models"
	"url-shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.ConnectDB()
	database.DB.AutoMigrate(&models.URL{})

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
