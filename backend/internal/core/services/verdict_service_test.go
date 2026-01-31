package services

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"rechtebank/backend/internal/core/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAnalyzer mocks the IPhotoAnalyzer interface
type MockAnalyzer struct {
	mock.Mock
}

func (m *MockAnalyzer) AnalyzePhoto(ctx context.Context, imageData []byte) (*domain.VerdictResponse, error) {
	args := m.Called(ctx, imageData)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.VerdictResponse), args.Error(1)
}

// MockValidator mocks the IPhotoValidator interface
type MockValidator struct {
	mock.Mock
}

func (m *MockValidator) ValidatePhoto(imageData []byte, metadata domain.PhotoMetadata) error {
	args := m.Called(imageData, metadata)
	return args.Error(0)
}

func TestVerdictService_JudgePhoto_Success(t *testing.T) {
	mockAnalyzer := new(MockAnalyzer)
	mockValidator := new(MockValidator)
	service := NewVerdictService(mockAnalyzer, mockValidator)

	imageData := []byte{0xFF, 0xD8, 0xFF}
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        1024,
	}

	analyzerResponse := &domain.VerdictResponse{
		Admissible: true,
		Score:      8,
		Verdict: domain.VerdictDetails{
			Crime:       "Rugleuning-afwijking",
			Sentence:    "Lichte berisping",
			Reasoning:   "Artikel 42",
			VerdictType: "waarschuwing",
		},
	}

	mockValidator.On("ValidatePhoto", imageData, metadata).Return(nil)
	mockAnalyzer.On("AnalyzePhoto", mock.Anything, imageData).Return(analyzerResponse, nil)

	result, err := service.JudgePhoto(context.Background(), imageData, metadata)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Admissible)
	assert.Equal(t, 8, result.Score)
	assert.Equal(t, "Rugleuning-afwijking", result.Verdict.Crime)
	assert.Equal(t, "waarschuwing", result.Verdict.VerdictType)
	assert.NotEmpty(t, result.RequestID)
	assert.NotEmpty(t, result.Timestamp)
	mockValidator.AssertExpectations(t)
	mockAnalyzer.AssertExpectations(t)
}

func TestVerdictService_JudgePhoto_ValidationError(t *testing.T) {
	mockAnalyzer := new(MockAnalyzer)
	mockValidator := new(MockValidator)
	service := NewVerdictService(mockAnalyzer, mockValidator)

	imageData := []byte{0x00, 0x00, 0x00}
	metadata := domain.PhotoMetadata{
		Filename:    "test.gif",
		ContentType: "image/gif",
		Size:        1024,
	}

	validationErr := errors.New("Unsupported image format. Use JPEG, PNG, or WebP")
	mockValidator.On("ValidatePhoto", imageData, metadata).Return(validationErr)

	result, err := service.JudgePhoto(context.Background(), imageData, metadata)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Unsupported image format. Use JPEG, PNG, or WebP", err.Error())
	mockValidator.AssertExpectations(t)
	mockAnalyzer.AssertNotCalled(t, "AnalyzePhoto")
}

func TestVerdictService_JudgePhoto_AnalysisError(t *testing.T) {
	mockAnalyzer := new(MockAnalyzer)
	mockValidator := new(MockValidator)
	service := NewVerdictService(mockAnalyzer, mockValidator)

	imageData := []byte{0xFF, 0xD8, 0xFF}
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        1024,
	}

	analysisErr := errors.New("AI analysis failed")
	mockValidator.On("ValidatePhoto", imageData, metadata).Return(nil)
	mockAnalyzer.On("AnalyzePhoto", mock.Anything, imageData).Return(nil, analysisErr)

	result, err := service.JudgePhoto(context.Background(), imageData, metadata)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "AI analysis failed", err.Error())
	mockValidator.AssertExpectations(t)
	mockAnalyzer.AssertExpectations(t)
}

func TestVerdictService_JudgePhoto_RequestIDFormat(t *testing.T) {
	mockAnalyzer := new(MockAnalyzer)
	mockValidator := new(MockValidator)
	service := NewVerdictService(mockAnalyzer, mockValidator)

	imageData := []byte{0xFF, 0xD8, 0xFF}
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        1024,
	}

	analyzerResponse := &domain.VerdictResponse{
		Admissible: true,
		Score:      7,
		Verdict: domain.VerdictDetails{
			Crime:       "Test crime",
			Sentence:    "Test sentence",
			Reasoning:   "Test reasoning",
			VerdictType: "vrijspraak",
		},
	}

	mockValidator.On("ValidatePhoto", imageData, metadata).Return(nil)
	mockAnalyzer.On("AnalyzePhoto", mock.Anything, imageData).Return(analyzerResponse, nil)

	result, err := service.JudgePhoto(context.Background(), imageData, metadata)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify UUID format (8-4-4-4-12 hex pattern)
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	assert.True(t, uuidRegex.MatchString(result.RequestID), "RequestID should be a valid UUID format, got: %s", result.RequestID)
}

func TestVerdictService_JudgePhoto_TimestampFormat(t *testing.T) {
	mockAnalyzer := new(MockAnalyzer)
	mockValidator := new(MockValidator)
	service := NewVerdictService(mockAnalyzer, mockValidator)

	imageData := []byte{0xFF, 0xD8, 0xFF}
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        1024,
	}

	analyzerResponse := &domain.VerdictResponse{
		Admissible: true,
		Score:      7,
		Verdict: domain.VerdictDetails{
			Crime:       "Test crime",
			Sentence:    "Test sentence",
			Reasoning:   "Test reasoning",
			VerdictType: "vrijspraak",
		},
	}

	mockValidator.On("ValidatePhoto", imageData, metadata).Return(nil)
	mockAnalyzer.On("AnalyzePhoto", mock.Anything, imageData).Return(analyzerResponse, nil)

	beforeTime := time.Now().UTC().Add(-1 * time.Second) // Allow 1 second buffer
	result, err := service.JudgePhoto(context.Background(), imageData, metadata)
	afterTime := time.Now().UTC().Add(1 * time.Second) // Allow 1 second buffer

	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify ISO 8601 format
	parsedTime, parseErr := time.Parse(time.RFC3339, result.Timestamp)
	assert.NoError(t, parseErr, "Timestamp should be valid ISO 8601 format")

	// Verify timestamp is within expected range (with buffer for timing)
	assert.True(t, !parsedTime.Before(beforeTime), "Timestamp should be after test started, got: %s, expected after: %s", parsedTime, beforeTime)
	assert.True(t, !parsedTime.After(afterTime), "Timestamp should be before test ended, got: %s, expected before: %s", parsedTime, afterTime)
}

func TestVerdictService_JudgePhoto_UniqueRequestIDs(t *testing.T) {
	mockAnalyzer := new(MockAnalyzer)
	mockValidator := new(MockValidator)
	service := NewVerdictService(mockAnalyzer, mockValidator)

	imageData := []byte{0xFF, 0xD8, 0xFF}
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        1024,
	}

	analyzerResponse := &domain.VerdictResponse{
		Admissible: true,
		Score:      7,
		Verdict: domain.VerdictDetails{
			Crime:       "Test crime",
			Sentence:    "Test sentence",
			Reasoning:   "Test reasoning",
			VerdictType: "vrijspraak",
		},
	}

	mockValidator.On("ValidatePhoto", imageData, metadata).Return(nil)
	mockAnalyzer.On("AnalyzePhoto", mock.Anything, imageData).Return(analyzerResponse, nil)

	// Make multiple calls and verify unique IDs
	requestIDs := make(map[string]bool)
	for i := 0; i < 10; i++ {
		result, err := service.JudgePhoto(context.Background(), imageData, metadata)
		assert.NoError(t, err)
		assert.False(t, requestIDs[result.RequestID], "RequestID should be unique")
		requestIDs[result.RequestID] = true
	}
}
