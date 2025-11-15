package handlers

import (
	"net/http"
	"url-shortener/database"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	Original string `json:"original"`
}

func ShortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.BindJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	shortCode := utils.GenerateShortCode(6)
	url := models.URL{Original: req.Original, Short: shortCode}

	if err := database.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save url"})
		return
	}

	c.JSON(200, gin.H{
		"original": req.Original,
		"short":    shortCode,
	})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	var url models.URL
	if err := database.DB.Where("short = ?", code).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "url not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Original)
}
