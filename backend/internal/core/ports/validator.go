package ports

import "rechtebank/backend/internal/core/domain"

// IPhotoValidator defines the interface for validating photo uploads
type IPhotoValidator interface {
	// ValidatePhoto validates image data and metadata
	// Returns nil if valid, error with user-friendly message if invalid
	ValidatePhoto(imageData []byte, metadata domain.PhotoMetadata) error
}
