package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rechtebank/backend/internal/core/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestVerdictHandler_GetByID_Success(t *testing.T) {
	// Create temporary storage directory
	tmpDir := t.TempDir()

	// Create test data
	dateDir := "2026-02-01"
	filename := "153045_abc123"
	fullDir := filepath.Join(tmpDir, dateDir)
	os.MkdirAll(fullDir, 0755)

	// Write test photo
	photoData := []byte{0xFF, 0xD8, 0xFF, 0xE0} // JPEG magic bytes
	photoPath := filepath.Join(fullDir, filename+".jpg")
	os.WriteFile(photoPath, photoData, 0644)

	// Write test verdict JSON
	verdictData := domain.VerdictResponse{
		Admissible: true,
		Score:      8,
		Verdict: domain.VerdictDetails{
			Crime:     "Rugleuning-afwijking",
			Sentence:  "Lichte berisping",
			Reasoning: "Artikel 42",
		},
		RequestID: "abc123",
		Timestamp: "2026-02-01T15:30:45Z",
	}
	verdictJSON, _ := json.Marshal(verdictData)
	jsonPath := filepath.Join(fullDir, filename+".json")
	os.WriteFile(jsonPath, verdictJSON, 0644)

	// Create handler
	handler := NewVerdictHandler(tmpDir)

	// Encode ID
	encodedID := domain.EncodeVerdictID(dateDir + "/" + filename)

	// Create request
	router := gin.New()
	router.GET("/v1/verdict/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/v1/verdict/"+encodedID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)

	var response VerdictWithImageResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Check verdict data
	assert.Equal(t, verdictData.Admissible, response.Verdict.Admissible)
	assert.Equal(t, verdictData.Score, response.Verdict.Score)
	assert.Equal(t, verdictData.RequestID, response.Verdict.RequestID)

	// Check image data (should be base64 data URL)
	assert.Contains(t, response.Image, "data:image/jpeg;base64,")

	// Decode and verify image
	base64Data := response.Image[len("data:image/jpeg;base64,"):]
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	assert.NoError(t, err)
	assert.Equal(t, photoData, decoded)
}

func TestVerdictHandler_GetByID_InvalidID(t *testing.T) {
	tmpDir := t.TempDir()
	handler := NewVerdictHandler(tmpDir)

	router := gin.New()
	router.GET("/v1/verdict/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/v1/verdict/invalid-base64!!!", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["error"], "Invalid verdict ID")
}

func TestVerdictHandler_GetByID_MissingVerdictJSON(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test data with only photo (no JSON)
	dateDir := "2026-02-01"
	filename := "153045_abc123"
	fullDir := filepath.Join(tmpDir, dateDir)
	os.MkdirAll(fullDir, 0755)

	photoData := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	photoPath := filepath.Join(fullDir, filename+".jpg")
	os.WriteFile(photoPath, photoData, 0644)
	// Don't create JSON file

	handler := NewVerdictHandler(tmpDir)
	encodedID := domain.EncodeVerdictID(dateDir + "/" + filename)

	router := gin.New()
	router.GET("/v1/verdict/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/v1/verdict/"+encodedID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["error"], "Verdict data not found")
}

func TestVerdictHandler_GetByID_MissingPhoto(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test data with only JSON (no photo)
	dateDir := "2026-02-01"
	filename := "153045_abc123"
	fullDir := filepath.Join(tmpDir, dateDir)
	os.MkdirAll(fullDir, 0755)

	verdictData := domain.VerdictResponse{
		Admissible: true,
		Score:      8,
	}
	verdictJSON, _ := json.Marshal(verdictData)
	jsonPath := filepath.Join(fullDir, filename+".json")
	os.WriteFile(jsonPath, verdictJSON, 0644)
	// Don't create photo file

	handler := NewVerdictHandler(tmpDir)
	encodedID := domain.EncodeVerdictID(dateDir + "/" + filename)

	router := gin.New()
	router.GET("/v1/verdict/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/v1/verdict/"+encodedID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["error"], "Photo file not found")
}

func TestVerdictHandler_GetByID_FileReadError(t *testing.T) {
	tmpDir := t.TempDir()

	// Create directory without write permissions to simulate read errors
	dateDir := "2026-02-01"
	filename := "153045_abc123"
	fullDir := filepath.Join(tmpDir, dateDir)
	os.MkdirAll(fullDir, 0755)

	// Create files
	photoPath := filepath.Join(fullDir, filename+".jpg")
	os.WriteFile(photoPath, []byte{0xFF, 0xD8}, 0644)

	verdictData := domain.VerdictResponse{Admissible: true}
	verdictJSON, _ := json.Marshal(verdictData)
	jsonPath := filepath.Join(fullDir, filename+".json")
	os.WriteFile(jsonPath, verdictJSON, 0644)

	// Make directory unreadable
	os.Chmod(fullDir, 0000)
	defer os.Chmod(fullDir, 0755) // Restore for cleanup

	handler := NewVerdictHandler(tmpDir)
	encodedID := domain.EncodeVerdictID(dateDir + "/" + filename)

	router := gin.New()
	router.GET("/v1/verdict/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/v1/verdict/"+encodedID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return 500 for file read errors
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Tests for CreateShareURL handler

func TestVerdictHandler_CreateShareURL_Success(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test data
	dateDir := "2026-02-01"
	timestamp := "153045"
	requestID := "abc123"
	filename := timestamp + "_" + requestID
	fullDir := filepath.Join(tmpDir, dateDir)
	os.MkdirAll(fullDir, 0755)

	// Write test files
	photoPath := filepath.Join(fullDir, filename+".jpg")
	os.WriteFile(photoPath, []byte{0xFF, 0xD8}, 0644)

	verdictData := domain.VerdictResponse{Admissible: true}
	verdictJSON, _ := json.Marshal(verdictData)
	jsonPath := filepath.Join(fullDir, filename+".json")
	os.WriteFile(jsonPath, verdictJSON, 0644)

	handler := NewVerdictHandler(tmpDir)

	router := gin.New()
	router.POST("/v1/verdict/share", handler.CreateShareURL)

	// Create request body
	reqBody := `{"timestamp":"2026-02-01T15:30:45Z","requestId":"abc123"}`
	req := httptest.NewRequest(http.MethodPost, "/v1/verdict/share", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response ShareResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.ID)

	// Verify the ID can be decoded back
	decoded, err := domain.DecodeVerdictID(response.ID)
	assert.NoError(t, err)
	assert.Equal(t, dateDir+"/"+filename, decoded)
}

func TestVerdictHandler_CreateShareURL_MissingFiles(t *testing.T) {
	tmpDir := t.TempDir()
	handler := NewVerdictHandler(tmpDir)

	router := gin.New()
	router.POST("/v1/verdict/share", handler.CreateShareURL)

	reqBody := `{"timestamp":"2026-02-01T15:30:45Z","requestId":"nonexistent"}`
	req := httptest.NewRequest(http.MethodPost, "/v1/verdict/share", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["error"], "Verdict not found")
}

func TestVerdictHandler_CreateShareURL_InvalidRequest(t *testing.T) {
	tmpDir := t.TempDir()
	handler := NewVerdictHandler(tmpDir)

	router := gin.New()
	router.POST("/v1/verdict/share", handler.CreateShareURL)

	// Missing required fields
	reqBody := `{"timestamp":"2026-02-01T15:30:45Z"}`
	req := httptest.NewRequest(http.MethodPost, "/v1/verdict/share", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
