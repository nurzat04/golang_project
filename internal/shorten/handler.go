package shorten

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func shortenHandler(c *gin.Context) {
	type Request struct {
		URL string `json:"url" binding:"required,url"`
	}
	type Response struct {
		ShortURL string `json:"short_url"`
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	shortCode := generateShortCode(req.URL)
	if err := saveURL(shortCode, req.URL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	shortURL := "http://localhost:8080/" + shortCode
	c.JSON(http.StatusOK, Response{ShortURL: shortURL})
}
