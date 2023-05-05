package handlers

import (
	"fmt"
	"os"
	"strconv"
	"txqueue/database"
	"txqueue/models"

	"github.com/gin-gonic/gin"
)

// Health shows alive
func Push(c *gin.Context) {
	var payload models.RawTxStruct
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	size := database.Que.Size()
	maxSize, _ := strconv.Atoi(os.Getenv("MAX_QUEUE_SIZE"))
	if size > maxSize {
		c.JSON(400, gin.H{
			"error": "max size exceeded (" + os.Getenv("MAX_QUEUE_SIZE") + ")",
		})
		return
	}

	fmt.Println("received RawTX:", payload.RawTx)

	database.Que.Push([]byte(payload.RawTx))

	c.JSON(200, gin.H{
		"status": "pushed", "queueSize": size + 1,
	})
}
