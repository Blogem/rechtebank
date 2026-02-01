package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"rechtebank/backend/internal/core/domain"

	"github.com/gin-gonic/gin"
)

// VerdictWithImageResponse combines verdict data with base64-encoded image
type VerdictWithImageResponse struct {
	Verdict domain.VerdictResponse `json:"verdict"`
	Image   string                 `json:"image"` // data URL format: "data:image/jpeg;base64,..."
}

// VerdictHandler handles verdict retrieval requests
type VerdictHandler struct {
	storagePath string
}

// NewVerdictHandler creates a new VerdictHandler
func NewVerdictHandler(storagePath string) *VerdictHandler {
	return &VerdictHandler{
		storagePath: storagePath,
	}
}

// GetByID handles GET /v1/verdict/:id requests
func (h *VerdictHandler) GetByID(c *gin.Context) {
	// Get encoded ID from URL parameter
	encodedID := c.Param("id")

	// Decode base64url ID to get file path
	filePath, err := domain.DecodeVerdictID(encodedID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verdict ID"})
		return
	}

	// Construct full file paths
	baseFilePath := filepath.Join(h.storagePath, filePath)
	jsonPath := baseFilePath + ".json"
	photoPath := baseFilePath + ".jpg"

	// Read verdict JSON file
	verdictData, err := os.ReadFile(jsonPath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Verdict data not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read verdict data"})
		}
		return
	}

	// Parse verdict JSON (stored in flat format on disk)
	var flatVerdict struct {
		Admissible  bool   `json:"admissible"`
		Score       int    `json:"score"`
		Crime       string `json:"crime"`
		Sentence    string `json:"sentence"`
		Reasoning   string `json:"reasoning"`
		Observation string `json:"observation"`
		VerdictType string `json:"verdictType"`
		RequestID   string `json:"requestId"`
		Timestamp   string `json:"timestamp"`
	}
	if err := json.Unmarshal(verdictData, &flatVerdict); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse verdict data"})
		return
	}

	// Convert to VerdictResponse structure
	verdict := domain.VerdictResponse{
		Admissible: flatVerdict.Admissible,
		Score:      flatVerdict.Score,
		Verdict: domain.VerdictDetails{
			Crime:       flatVerdict.Crime,
			Sentence:    flatVerdict.Sentence,
			Reasoning:   flatVerdict.Reasoning,
			Observation: flatVerdict.Observation,
			VerdictType: flatVerdict.VerdictType,
		},
		RequestID: flatVerdict.RequestID,
		Timestamp: flatVerdict.Timestamp,
	}

	// Read photo file
	photoData, err := os.ReadFile(photoPath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Photo file not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read photo file"})
		}
		return
	}

	// Encode photo as base64 data URL
	base64Photo := base64.StdEncoding.EncodeToString(photoData)
	imageDataURL := fmt.Sprintf("data:image/jpeg;base64,%s", base64Photo)

	// Return combined response
	response := VerdictWithImageResponse{
		Verdict: verdict,
		Image:   imageDataURL,
	}

	c.JSON(http.StatusOK, response)
}

// ShareRequest represents the request body for POST /v1/verdict/share
type ShareRequest struct {
	Timestamp string `json:"timestamp" binding:"required"`
	RequestID string `json:"requestId" binding:"required"`
}

// ShareResponse represents the response for POST /v1/verdict/share
type ShareResponse struct {
	ID string `json:"id"`
}

// CreateShareURL handles POST /v1/verdict/share requests
func (h *VerdictHandler) CreateShareURL(c *gin.Context) {
	var req ShareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Parse timestamp to extract date (YYYY-MM-DD) and time (HHMMSS)
	// Expected format: "2026-02-01T15:30:45Z" or similar ISO 8601
	// Extract YYYY-MM-DD for directory
	dateDir := ""
	timeStr := ""

	if len(req.Timestamp) >= 10 {
		dateDir = req.Timestamp[:10] // YYYY-MM-DD
	}

	// Extract time component (HHMMSS)
	// If timestamp is "2026-02-01T15:30:45Z", we need to extract "153045"
	if strings.Contains(req.Timestamp, "T") {
		parts := strings.Split(req.Timestamp, "T")
		if len(parts) > 1 {
			timePart := strings.Split(parts[1], ":") // ["15", "30", "45Z"]
			if len(timePart) >= 3 {
				// Clean seconds part (remove Z or other timezone info)
				seconds := strings.TrimRight(timePart[2], "Z")
				seconds = strings.Split(seconds, ".")[0]      // Remove milliseconds if present
				timeStr = timePart[0] + timePart[1] + seconds // HHMMSS
			}
		}
	}

	if dateDir == "" || timeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timestamp format"})
		return
	}

	// Construct file path
	filename := fmt.Sprintf("%s_%s", timeStr, req.RequestID)
	filePath := fmt.Sprintf("%s/%s", dateDir, filename)

	// Verify both files exist
	baseFilePath := filepath.Join(h.storagePath, filePath)
	jsonPath := baseFilePath + ".json"
	photoPath := baseFilePath + ".jpg"

	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Verdict not found"})
		return
	}

	if _, err := os.Stat(photoPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Verdict not found"})
		return
	}

	// Generate base64url ID
	encodedID := domain.EncodeVerdictID(filePath)

	response := ShareResponse{
		ID: encodedID,
	}

	c.JSON(http.StatusOK, response)
}
