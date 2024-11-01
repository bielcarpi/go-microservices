package server

import (
	"broker/internal/handlers"
	"broker/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	router := gin.Default()

	// Add CORS middleware
	router.Use(middleware.CORS())

	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Initialize handlers
	h := handlers.NewHandler()

	// Routes
	router.POST("/", h.Broker)

	return router
}
