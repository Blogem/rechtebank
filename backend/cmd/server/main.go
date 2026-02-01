package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rechtebank/backend/internal/adapters/gemini"
	httpAdapter "rechtebank/backend/internal/adapters/http"
	"rechtebank/backend/internal/adapters/http/handlers"
	"rechtebank/backend/internal/adapters/storage"
	"rechtebank/backend/internal/adapters/validator"
	"rechtebank/backend/internal/config"
	"rechtebank/backend/internal/core/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set Gin mode based on environment
	if !cfg.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Log startup information
	log.Printf("Starting Rechtebank Backend API")
	log.Printf("  Port: %s", cfg.Port)
	log.Printf("  Environment: %s", cfg.Environment)
	log.Printf("  CORS Origin: %s", cfg.CORSOrigin)
	log.Printf("  Gemini Timeout: %s", cfg.GeminiTimeout)
	log.Printf("  Photo Storage: %s", cfg.PhotoStoragePath)
	log.Printf("  Photo Retention: %d days", cfg.PhotoRetentionDays)

	// Initialize dependencies
	// 1. Validator
	photoValidator := validator.NewPhotoValidator()

	// 2. Gemini Analyzer
	geminiAnalyzer, err := gemini.NewGeminiAnalyzer(cfg.GeminiAPIKey, cfg.GeminiTimeout)
	if err != nil {
		log.Fatalf("Failed to initialize Gemini analyzer: %v", err)
	}
	defer geminiAnalyzer.Close()

	// 3. Photo Storage
	photoStorage, err := storage.NewPhotoStorage(cfg.PhotoStoragePath)
	if err != nil {
		log.Fatalf("Failed to initialize photo storage: %v", err)
	}

	// 4. Verdict Service
	verdictService := services.NewVerdictService(geminiAnalyzer, photoValidator)

	// 5. HTTP Handlers
	judgeHandler := handlers.NewJudgeHandler(verdictService, photoStorage)
	verdictHandler := handlers.NewVerdictHandler(cfg.PhotoStoragePath)

	// 6. Router
	router := httpAdapter.NewRouter(judgeHandler, verdictHandler, httpAdapter.RouterConfig{
		CORSOrigin: cfg.CORSOrigin,
	})

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Create context for cleanup goroutine
	cleanupCtx, cleanupCancel := context.WithCancel(context.Background())
	defer cleanupCancel()

	// Start cleanup job in background
	go func() {
		ticker := time.NewTicker(24 * time.Hour) // Run daily
		defer ticker.Stop()

		// Run once on startup
		if err := photoStorage.CleanupOldPhotos(cfg.PhotoRetentionDays); err != nil {
			log.Printf("Warning: Failed to cleanup old photos: %v", err)
		} else {
			log.Printf("Photo cleanup completed successfully")
		}

		// Then run daily
		for {
			select {
			case <-cleanupCtx.Done():
				log.Println("Cleanup job stopped")
				return
			case <-ticker.C:
				if err := photoStorage.CleanupOldPhotos(cfg.PhotoRetentionDays); err != nil {
					log.Printf("Warning: Failed to cleanup old photos: %v", err)
				} else {
					log.Printf("Photo cleanup completed successfully")
				}
			}
		}
	}()

	// Start server in goroutine
	go func() {
		log.Printf("Server listening on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Cancel cleanup goroutine
	cleanupCancel()

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
