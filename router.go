package main

import (
	"net/http"
	"txqueue/handlers"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// InitializeRouter initialises all the routes in the app
func InitializeRouter() {
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// ----- PUBLIC APIs -----
	router.GET("/health", handlers.Health)
	router.POST("/push", handlers.Push)

	// ----- Other -----
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "api endpoint not found"})
	})
}
