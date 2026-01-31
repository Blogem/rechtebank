package ports

import (
	"context"

	"rechtebank/backend/internal/core/domain"
)

// IPhotoAnalyzer defines the interface for analyzing photos and generating verdicts
type IPhotoAnalyzer interface {
	// AnalyzePhoto analyzes an image and returns verdict details
	// Returns admissible status, score (1-10), and verdict components
	AnalyzePhoto(ctx context.Context, imageData []byte) (*domain.VerdictResponse, error)
}
