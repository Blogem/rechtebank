package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"rechtebank/backend/internal/core/domain"
	"rechtebank/backend/internal/core/ports"
)

// VerdictService orchestrates photo validation and analysis
type VerdictService struct {
	analyzer  ports.IPhotoAnalyzer
	validator ports.IPhotoValidator
}

// NewVerdictService creates a new VerdictService with the given dependencies
func NewVerdictService(analyzer ports.IPhotoAnalyzer, validator ports.IPhotoValidator) *VerdictService {
	return &VerdictService{
		analyzer:  analyzer,
		validator: validator,
	}
}

// JudgePhoto validates and analyzes a photo, returning a verdict
func (s *VerdictService) JudgePhoto(ctx context.Context, imageData []byte, metadata domain.PhotoMetadata) (*domain.VerdictResponse, error) {
	// Step 1: Validate the photo
	if err := s.validator.ValidatePhoto(imageData, metadata); err != nil {
		return nil, err
	}

	// Step 2: Analyze the photo with AI
	result, err := s.analyzer.AnalyzePhoto(ctx, imageData)
	if err != nil {
		return nil, err
	}

	// Step 3: Add request metadata
	result.RequestID = uuid.New().String()
	result.Timestamp = time.Now().UTC().Format(time.RFC3339)

	return result, nil
}
