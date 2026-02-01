package http

import (
	"net/http"
	"time"

	"rechtebank/backend/internal/adapters/http/handlers"

	"github.com/gin-gonic/gin"
)

// RouterConfig holds configuration for the router
type RouterConfig struct {
	CORSOrigin string
}

// NewRouter creates a new Gin router with all middleware and routes configured
func NewRouter(judgeHandler *handlers.JudgeHandler, verdictHandler *handlers.VerdictHandler, config RouterConfig) *gin.Engine {
	router := gin.New()

	// Add middleware
	router.Use(gin.Recovery())
	router.Use(loggingMiddleware())
	router.Use(corsMiddleware(config.CORSOrigin))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	})

	// API v1 routes
	v1 := router.Group("/v1")
	{
		v1.POST("/judge", judgeHandler.Handle)
		v1.GET("/verdict/:id", verdictHandler.GetByID)
		v1.POST("/verdict/share", verdictHandler.CreateShareURL)
	}

	return router
}

// loggingMiddleware logs request information
func loggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/health"},
	})
}

// corsMiddleware handles CORS headers
func corsMiddleware(allowedOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := allowedOrigin
		if origin == "" {
			origin = "*"
		}

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
