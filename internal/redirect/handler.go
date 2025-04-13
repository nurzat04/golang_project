package redirect

import (
	"net/http"
	"shorty/pkg/logger"

	"github.com/gin-gonic/gin"
)

func redirectHandler(c *gin.Context) {
	code := c.Param("code")
	originalURL, err := getOriginalURL(code)
	if err != nil {
		logger.ErrorLogger.Println("Failed to get original URL:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Short link not found"})
		return
	}

	logger.InfoLogger.Println("Redirecting to:", originalURL)
	incrementClick(code)
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func incrementClick(code string) {
	key := "stats:" + code
	err := rdb.Incr(ctx, key).Err()
	if err != nil {
		logger.ErrorLogger.Println("Failed to increment stats:", err)
	}
}
