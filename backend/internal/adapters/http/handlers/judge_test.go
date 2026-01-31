package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rechtebank/backend/internal/core/domain"
)

// MockVerdictService mocks the verdict service
type MockVerdictService struct {
	mock.Mock
}

func (m *MockVerdictService) JudgePhoto(ctx context.Context, imageData []byte, metadata domain.PhotoMetadata) (*domain.VerdictResponse, error) {
	args := m.Called(ctx, imageData, metadata)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.VerdictResponse), args.Error(1)
}

func init() {
	gin.SetMode(gin.TestMode)
}

func createMultipartRequest(t *testing.T, fieldName, filename string, content []byte) (*http.Request, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile(fieldName, filename)
	if err != nil {
		return nil, err
	}
	_, err = part.Write(content)
	if err != nil {
		return nil, err
	}
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/v1/judge", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

func TestJudgeHandler_SuccessfulUpload(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	// JPEG magic bytes + some data
	imageData := []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01}

	expectedResponse := &domain.VerdictResponse{
		Admissible: true,
		Score:      8,
		Verdict: domain.VerdictDetails{
			Crime:     "Rugleuning-afwijking",
			Sentence:  "Lichte berisping",
			Reasoning: "Artikel 42",
		},
		RequestID: "test-123",
		Timestamp: "2026-01-31T10:00:00Z",
	}

	mockService.On("JudgePhoto", mock.Anything, imageData, mock.MatchedBy(func(m domain.PhotoMetadata) bool {
		return m.Filename == "test.jpg" && m.Size == int64(len(imageData))
	})).Return(expectedResponse, nil)

	req, _ := createMultipartRequest(t, "photo", "test.jpg", imageData)
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response domain.VerdictResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response.Admissible)
	assert.Equal(t, 8, response.Score)
	mockService.AssertExpectations(t)
}

func TestJudgeHandler_MissingFile(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	// Request without file
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/v1/judge", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.Equal(t, "Photo file is required", errorResponse["error"])
}

func TestJudgeHandler_InvalidContentType(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	req := httptest.NewRequest(http.MethodPost, "/v1/judge", bytes.NewReader([]byte("not multipart")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.Equal(t, "Content-Type must be multipart/form-data", errorResponse["error"])
}

func TestJudgeHandler_FileTooLarge(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	imageData := []byte{0xFF, 0xD8, 0xFF}

	validationErr := &ValidationError{Message: "Photo file size must not exceed 10MB", StatusCode: http.StatusRequestEntityTooLarge}
	mockService.On("JudgePhoto", mock.Anything, imageData, mock.Anything).Return(nil, validationErr)

	req, _ := createMultipartRequest(t, "photo", "large.jpg", imageData)
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusRequestEntityTooLarge, w.Code)

	var errorResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.Equal(t, "Photo file size must not exceed 10MB", errorResponse["error"])
}

func TestJudgeHandler_RateLimitError(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	imageData := []byte{0xFF, 0xD8, 0xFF}

	rateLimitErr := &RateLimitError{RetryAfter: 30}
	mockService.On("JudgePhoto", mock.Anything, imageData, mock.Anything).Return(nil, rateLimitErr)

	req, _ := createMultipartRequest(t, "photo", "test.jpg", imageData)
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.Equal(t, "30", w.Header().Get("Retry-After"))

	var errorResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.Contains(t, errorResponse["error"], "rate limit")
}

func TestJudgeHandler_InternalServerError(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	imageData := []byte{0xFF, 0xD8, 0xFF}

	mockService.On("JudgePhoto", mock.Anything, imageData, mock.Anything).Return(nil, errors.New("AI analysis service unavailable"))

	req, _ := createMultipartRequest(t, "photo", "test.jpg", imageData)
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestJudgeHandler_BadGateway(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	imageData := []byte{0xFF, 0xD8, 0xFF}

	badGatewayErr := &APIError{Message: "AI analysis failed", StatusCode: http.StatusBadGateway}
	mockService.On("JudgePhoto", mock.Anything, imageData, mock.Anything).Return(nil, badGatewayErr)

	req, _ := createMultipartRequest(t, "photo", "test.jpg", imageData)
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}

func TestJudgeHandler_ServiceUnavailable(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	imageData := []byte{0xFF, 0xD8, 0xFF}

	serviceErr := &APIError{Message: "AI analysis service temporarily unavailable", StatusCode: http.StatusServiceUnavailable}
	mockService.On("JudgePhoto", mock.Anything, imageData, mock.Anything).Return(nil, serviceErr)

	req, _ := createMultipartRequest(t, "photo", "test.jpg", imageData)
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestJudgeHandler_GatewayTimeout(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	imageData := []byte{0xFF, 0xD8, 0xFF}

	timeoutErr := &APIError{Message: "AI analysis timeout", StatusCode: http.StatusGatewayTimeout}
	mockService.On("JudgePhoto", mock.Anything, imageData, mock.Anything).Return(nil, timeoutErr)

	req, _ := createMultipartRequest(t, "photo", "test.jpg", imageData)
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusGatewayTimeout, w.Code)
}

func TestJudgeHandler_MultipartFormDataParsing(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := NewJudgeHandler(mockService, nil)

	// PNG magic bytes
	imageData := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00, 0x00, 0x0D}

	expectedResponse := &domain.VerdictResponse{
		Admissible: true,
		Score:      7,
		Verdict: domain.VerdictDetails{
			Crime:     "Test",
			Sentence:  "Test",
			Reasoning: "Test",
		},
		RequestID: "test-456",
		Timestamp: "2026-01-31T10:00:00Z",
	}

	mockService.On("JudgePhoto", mock.Anything, imageData, mock.MatchedBy(func(m domain.PhotoMetadata) bool {
		return m.Filename == "image.png" && m.Size == int64(len(imageData))
	})).Return(expectedResponse, nil)

	req, _ := createMultipartRequest(t, "photo", "image.png", imageData)
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/v1/judge", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
