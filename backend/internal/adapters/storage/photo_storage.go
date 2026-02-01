package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// PhotoStorage handles saving and managing photo files
type PhotoStorage struct {
	basePath string
}

// NewPhotoStorage creates a new PhotoStorage instance
func NewPhotoStorage(basePath string) (*PhotoStorage, error) {
	// Create base directory if it doesn't exist
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &PhotoStorage{
		basePath: basePath,
	}, nil
}

// SavePhoto saves a photo to disk with timestamp-based naming
func (s *PhotoStorage) SavePhoto(imageData []byte, llmResponse []byte, requestID string, timestampISO string) (string, error) {
	// Create subdirectory based on current date (YYYY-MM-DD)
	now := time.Now()
	dateDir := now.Format("2006-01-02")
	fullDir := filepath.Join(s.basePath, dateDir)

	if err := os.MkdirAll(fullDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create date directory: %w", err)
	}

	// Generate filename: timestamp_requestID.jpg
	timestamp := now.Format("150405") // HHMMSS
	filename := fmt.Sprintf("%s_%s.jpg", timestamp, requestID)
	filePath := filepath.Join(fullDir, filename)

	// Write photo
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to write photo: %w", err)
	}
	// Write LLM response as JSON with added metadata
	filenameJSON := fmt.Sprintf("%s_%s.json", now.Format("150405"), requestID)
	filePathJSON := filepath.Join(fullDir, filenameJSON)

	// Parse the raw JSON and add requestId and timestamp
	var jsonData map[string]interface{}
	if err := json.Unmarshal(llmResponse, &jsonData); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}
	jsonData["requestId"] = requestID
	jsonData["timestamp"] = timestampISO

	// Marshal back to JSON
	completeJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(filePathJSON, completeJSON, 0644); err != nil {
		return "", fmt.Errorf("failed to write JSON: %w", err)
	}

	return filePath, nil
}

// CleanupOldPhotos removes photos older than the specified number of days
func (s *PhotoStorage) CleanupOldPhotos(retentionDays int) error {
	cutoffDate := time.Now().AddDate(0, 0, -retentionDays)

	// Walk through all date directories
	entries, err := os.ReadDir(s.basePath)
	if err != nil {
		return fmt.Errorf("failed to read storage directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		// Parse directory name as date (YYYY-MM-DD)
		dirDate, err := time.Parse("2006-01-02", entry.Name())
		if err != nil {
			// Skip directories that don't match the date format
			continue
		}

		// Remove directory if older than retention period
		if dirDate.Before(cutoffDate) {
			dirPath := filepath.Join(s.basePath, entry.Name())
			if err := os.RemoveAll(dirPath); err != nil {
				return fmt.Errorf("failed to remove old directory %s: %w", dirPath, err)
			}
		}
	}

	return nil
}
