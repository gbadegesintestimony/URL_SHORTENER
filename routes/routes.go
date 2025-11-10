package routes

import (
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:code", handlers.RedirectURL)
}
