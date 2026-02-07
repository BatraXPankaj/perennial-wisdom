package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health returns a simple health check response.
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"project": "Perennial Wisdom API",
		"motto":   "The obstacle is the way.",
	})
}
