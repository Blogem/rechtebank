package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

// Config holds all configuration for the application
type Config struct {
	// Server settings
	Port string

	// CORS settings
	CORSOrigin string

	// Gemini API settings
	GeminiAPIKey  string
	GeminiTimeout time.Duration

	// File upload settings
	MaxFileSize int64

	// Photo storage settings
	PhotoStoragePath   string
	PhotoRetentionDays int

	// Environment
	Environment string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	config := &Config{
		Port:               getEnvOrDefault("PORT", "8080"),
		CORSOrigin:         getEnvOrDefault("CORS_ORIGIN", "*"),
		GeminiAPIKey:       os.Getenv("GEMINI_API_KEY"),
		GeminiTimeout:      getDurationOrDefault("GEMINI_TIMEOUT", 30*time.Second),
		MaxFileSize:        getInt64OrDefault("MAX_FILE_SIZE", 10*1024*1024), // 10MB
		PhotoStoragePath:   getEnvOrDefault("PHOTO_STORAGE_PATH", "./photos"),
		PhotoRetentionDays: getIntOrDefault("PHOTO_RETENTION_DAYS", 90),
		Environment:        getEnvOrDefault("ENV", "development"),
	}

	// Validate required fields
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// Validate checks that all required configuration is present
func (c *Config) Validate() error {
	if c.GeminiAPIKey == "" {
		return errors.New("GEMINI_API_KEY environment variable is required")
	}
	return nil
}

// IsDevelopment returns true if running in development mode
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if seconds, err := strconv.Atoi(value); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return defaultValue
}

func getInt64OrDefault(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
