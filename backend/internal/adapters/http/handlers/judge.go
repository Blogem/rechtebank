package handlers

import (
	"context"
	"errors"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"

	"rechtebank/backend/internal/core/domain"

	"github.com/gin-gonic/gin"
)

// VerdictServiceInterface defines the interface for the verdict service
type VerdictServiceInterface interface {
	JudgePhoto(ctx context.Context, imageData []byte, metadata domain.PhotoMetadata) (*domain.VerdictResponse, error)
}

// PhotoStorageInterface defines the interface for photo storage
type PhotoStorageInterface interface {
	SavePhoto(imageData []byte, llmResponse []byte, requestID string) (string, error)
}

// JudgeHandler handles POST /v1/judge requests
type JudgeHandler struct {
	service VerdictServiceInterface
	storage PhotoStorageInterface
}

// NewJudgeHandler creates a new JudgeHandler
func NewJudgeHandler(service VerdictServiceInterface, storage PhotoStorageInterface) *JudgeHandler {
	return &JudgeHandler{
		service: service,
		storage: storage,
	}
}

// Handle processes the photo upload request
func (h *JudgeHandler) Handle(c *gin.Context) {
	// Check content type
	contentType := c.ContentType()
	if contentType == "" || !isMultipartFormData(contentType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be multipart/form-data"})
		return
	}

	// Get the file from form
	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Photo file is required"})
		return
	}
	defer file.Close()

	// Read file data
	imageData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// Create metadata
	metadata := domain.PhotoMetadata{
		Filename:    header.Filename,
		ContentType: header.Header.Get("Content-Type"),
		Size:        int64(len(imageData)),
	}

	// Log incoming photo details
	log.Printf("[JUDGE] Incoming photo: filename=%s, size=%d bytes, content-type=%s",
		metadata.Filename, metadata.Size, metadata.ContentType)

	// Call service
	result, err := h.service.JudgePhoto(c.Request.Context(), imageData, metadata)
	if err != nil {
		h.handleError(c, err)
		return
	}

	// Log the Gemini response
	log.Printf("[JUDGE] Gemini response: admissible=%v, score=%d, requestID=%s",
		result.Admissible, result.Score, result.RequestID)
	log.Printf("[JUDGE] Gemini raw JSON: %s", result.RawJSON)

	// Save photo to disk (async, don't fail request if this fails)
	if h.storage != nil && result.RequestID != "" {
		go func() {
			if _, err := h.storage.SavePhoto(imageData, []byte(result.RawJSON), result.RequestID); err != nil {
				// Log error but don't fail the request
				fmt.Printf("Failed to save photo: %v\n", err)
			}
		}()
	}

	c.JSON(http.StatusOK, result)
}

func (h *JudgeHandler) handleError(c *gin.Context, err error) {
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		c.JSON(validationErr.StatusCode, gin.H{"error": validationErr.Message})
		return
	}

	var rateLimitErr *RateLimitError
	if errors.As(err, &rateLimitErr) {
		c.Header("Retry-After", fmt.Sprintf("%d", rateLimitErr.RetryAfter))
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
		return
	}

	var apiErr *APIError
	if errors.As(err, &apiErr) {
		c.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Message})
		return
	}

	// Default to internal server error
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func isMultipartFormData(contentType string) bool {
	return len(contentType) >= 19 && contentType[:19] == "multipart/form-data"
}

// Error types for HTTP responses

// ValidationError represents a validation error with specific status code
type ValidationError struct {
	Message    string
	StatusCode int
}

func (e *ValidationError) Error() string {
	return e.Message
}

// RateLimitError indicates rate limiting with retry information
type RateLimitError struct {
	RetryAfter int
}

func (e *RateLimitError) Error() string {
	return "rate limit exceeded"
}

// APIError represents an API error with specific status code
type APIError struct {
	Message    string
	StatusCode int
}

func (e *APIError) Error() string {
	return e.Message
}
