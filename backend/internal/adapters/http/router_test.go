package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"rechtebank/backend/internal/adapters/http/handlers"
	"rechtebank/backend/internal/core/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockVerdictService for router tests
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

func TestRouter_HealthEndpoint(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := handlers.NewJudgeHandler(mockService, nil)
	verdictHandler := handlers.NewVerdictHandler("")
	router := NewRouter(handler, verdictHandler, RouterConfig{CORSOrigin: "*"})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
	assert.Contains(t, w.Body.String(), "timestamp")
}

func TestRouter_CORS_PreflightRequest(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := handlers.NewJudgeHandler(mockService, nil)
	verdictHandler := handlers.NewVerdictHandler("")
	router := NewRouter(handler, verdictHandler, RouterConfig{CORSOrigin: "http://localhost:5173"})

	req := httptest.NewRequest(http.MethodOptions, "/v1/judge", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	req.Header.Set("Access-Control-Request-Method", "POST")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "http://localhost:5173", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Methods"), "POST")
}

func TestRouter_CORS_PostRequest(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := handlers.NewJudgeHandler(mockService, nil)
	verdictHandler := handlers.NewVerdictHandler("")
	router := NewRouter(handler, verdictHandler, RouterConfig{CORSOrigin: "http://localhost:5173"})

	req := httptest.NewRequest(http.MethodPost, "/v1/judge", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	req.Header.Set("Content-Type", "application/json") // Will fail validation but CORS should still be set
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Check CORS header is set regardless of response status
	assert.Equal(t, "http://localhost:5173", w.Header().Get("Access-Control-Allow-Origin"))
}

func TestRouter_CORS_DefaultOrigin(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := handlers.NewJudgeHandler(mockService, nil)
	verdictHandler := handlers.NewVerdictHandler("")
	router := NewRouter(handler, verdictHandler, RouterConfig{CORSOrigin: ""})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
}

func TestRouter_V1JudgeEndpoint(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := handlers.NewJudgeHandler(mockService, nil)
	verdictHandler := handlers.NewVerdictHandler("")
	router := NewRouter(handler, verdictHandler, RouterConfig{CORSOrigin: "*"})

	// Request without proper content type should fail with 400
	req := httptest.NewRequest(http.MethodPost, "/v1/judge", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRouter_NotFound(t *testing.T) {
	mockService := new(MockVerdictService)
	handler := handlers.NewJudgeHandler(mockService, nil)
	verdictHandler := handlers.NewVerdictHandler("")
	router := NewRouter(handler, verdictHandler, RouterConfig{CORSOrigin: "*"})

	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
