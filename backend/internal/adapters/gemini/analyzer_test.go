package gemini

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockGeminiClient mocks the Gemini API client for testing
type MockGeminiClient struct {
	mock.Mock
}

func (m *MockGeminiClient) GenerateContent(ctx context.Context, imageData []byte) (*GeminiResponse, error) {
	args := m.Called(ctx, imageData)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*GeminiResponse), args.Error(1)
}

// Test Gemini client initialization
func TestNewGeminiAnalyzer_WithAPIKey(t *testing.T) {
	// Skip this test in CI as it requires actual API connection
	t.Skip("Skipping test that requires actual Gemini API connection")
	analyzer, err := NewGeminiAnalyzer("test-api-key", 30*time.Second)
	assert.NoError(t, err)
	assert.NotNil(t, analyzer)
}

func TestNewGeminiAnalyzer_WithoutAPIKey(t *testing.T) {
	analyzer, err := NewGeminiAnalyzer("", 30*time.Second)
	assert.Error(t, err)
	assert.Nil(t, analyzer)
	assert.Equal(t, "GEMINI_API_KEY environment variable is required", err.Error())
}

// Test successful photo analysis
func TestGeminiAnalyzer_AnalyzePhoto_Success(t *testing.T) {
	mockClient := new(MockGeminiClient)
	analyzer := &GeminiAnalyzer{
		client:  mockClient,
		timeout: 30 * time.Second,
	}

	imageData := []byte{0xFF, 0xD8, 0xFF} // JPEG header

	expectedResponse := &GeminiResponse{
		Admissible: true,
		Score:      8,
		Crime:      "Rugleuning-afwijking van 5 graden",
		Sentence:   "Veroordeeld tot lichte berisping",
		Reasoning:  "Artikel 42 van de Meubilair-wet",
	}

	mockClient.On("GenerateContent", mock.Anything, imageData).Return(expectedResponse, nil)

	result, err := analyzer.AnalyzePhoto(context.Background(), imageData)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Admissible)
	assert.Equal(t, 8, result.Score)
	assert.Equal(t, "Rugleuning-afwijking van 5 graden", result.Verdict.Crime)
	assert.Equal(t, "Veroordeeld tot lichte berisping", result.Verdict.Sentence)
	assert.Equal(t, "Artikel 42 van de Meubilair-wet", result.Verdict.Reasoning)
	mockClient.AssertExpectations(t)
}

// Test non-furniture detection
func TestGeminiAnalyzer_AnalyzePhoto_NonFurniture(t *testing.T) {
	mockClient := new(MockGeminiClient)
	analyzer := &GeminiAnalyzer{
		client:  mockClient,
		timeout: 30 * time.Second,
	}

	imageData := []byte{0xFF, 0xD8, 0xFF}

	expectedResponse := &GeminiResponse{
		Admissible: false,
		Score:      0,
		Crime:      "Geen meubilair gedetecteerd",
		Sentence:   "Zaak niet-ontvankelijk",
		Reasoning:  "Alleen meubilair kan worden berecht",
	}

	mockClient.On("GenerateContent", mock.Anything, imageData).Return(expectedResponse, nil)

	result, err := analyzer.AnalyzePhoto(context.Background(), imageData)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Admissible)
	assert.Equal(t, 0, result.Score)
	mockClient.AssertExpectations(t)
}

// Test rate limit handling (429 errors)
func TestGeminiAnalyzer_AnalyzePhoto_RateLimit_RetrySuccess(t *testing.T) {
	mockClient := new(MockGeminiClient)
	analyzer := &GeminiAnalyzer{
		client:     mockClient,
		timeout:    30 * time.Second,
		maxRetries: 3,
	}

	imageData := []byte{0xFF, 0xD8, 0xFF}

	expectedResponse := &GeminiResponse{
		Admissible: true,
		Score:      7,
		Crime:      "Minor offense",
		Sentence:   "Warning",
		Reasoning:  "First time",
	}

	// First call fails with rate limit, second succeeds
	mockClient.On("GenerateContent", mock.Anything, imageData).
		Return(nil, &RateLimitError{RetryAfter: 10 * time.Millisecond}).Once()
	mockClient.On("GenerateContent", mock.Anything, imageData).
		Return(expectedResponse, nil).Once()

	result, err := analyzer.AnalyzePhoto(context.Background(), imageData)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 7, result.Score)
	mockClient.AssertExpectations(t)
}

func TestGeminiAnalyzer_AnalyzePhoto_RateLimit_RetryExhausted(t *testing.T) {
	mockClient := new(MockGeminiClient)
	analyzer := &GeminiAnalyzer{
		client:     mockClient,
		timeout:    30 * time.Second,
		maxRetries: 3,
	}

	imageData := []byte{0xFF, 0xD8, 0xFF}

	// All retries fail with rate limit
	rateLimitErr := &RateLimitError{RetryAfter: 10 * time.Millisecond}
	mockClient.On("GenerateContent", mock.Anything, imageData).
		Return(nil, rateLimitErr).Times(4) // Initial + 3 retries

	result, err := analyzer.AnalyzePhoto(context.Background(), imageData)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "AI analysis service temporarily unavailable")
	mockClient.AssertExpectations(t)
}

// Test timeout scenarios
func TestGeminiAnalyzer_AnalyzePhoto_Timeout(t *testing.T) {
	mockClient := new(MockGeminiClient)
	analyzer := &GeminiAnalyzer{
		client:  mockClient,
		timeout: 30 * time.Second,
	}

	imageData := []byte{0xFF, 0xD8, 0xFF}

	mockClient.On("GenerateContent", mock.Anything, imageData).
		Return(nil, context.DeadlineExceeded)

	result, err := analyzer.AnalyzePhoto(context.Background(), imageData)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "AI analysis timeout")
	mockClient.AssertExpectations(t)
}

// Test API error scenarios
func TestGeminiAnalyzer_AnalyzePhoto_APIError(t *testing.T) {
	mockClient := new(MockGeminiClient)
	analyzer := &GeminiAnalyzer{
		client:  mockClient,
		timeout: 30 * time.Second,
	}

	imageData := []byte{0xFF, 0xD8, 0xFF}

	mockClient.On("GenerateContent", mock.Anything, imageData).
		Return(nil, errors.New("API error"))

	result, err := analyzer.AnalyzePhoto(context.Background(), imageData)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "AI analysis failed")
	mockClient.AssertExpectations(t)
}

func TestGeminiAnalyzer_AnalyzePhoto_InvalidResponse(t *testing.T) {
	mockClient := new(MockGeminiClient)
	analyzer := &GeminiAnalyzer{
		client:  mockClient,
		timeout: 30 * time.Second,
	}

	imageData := []byte{0xFF, 0xD8, 0xFF}

	mockClient.On("GenerateContent", mock.Anything, imageData).
		Return(nil, &InvalidResponseError{Message: "invalid JSON schema"})

	result, err := analyzer.AnalyzePhoto(context.Background(), imageData)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Invalid AI response format")
	mockClient.AssertExpectations(t)
}
