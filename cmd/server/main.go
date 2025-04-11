// cmd/server/main.go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go-run-reports/internal/api"
	"go-run-reports/internal/config"
)

func main() {
	// Load environment variables/config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Setup Gin router
	router := gin.Default()

	// Register API routes
	api.RegisterRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("🚀 Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}