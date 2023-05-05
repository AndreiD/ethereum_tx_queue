package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health shows alive
func Health(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}
