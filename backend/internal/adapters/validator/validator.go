package validator

import (
	"bytes"
	"errors"

	"rechtebank/backend/internal/core/domain"
)

const maxFileSize = 10 * 1024 * 1024 // 10MB

type PhotoValidator struct{}

func NewPhotoValidator() *PhotoValidator {
	return &PhotoValidator{}
}

func (v *PhotoValidator) ValidatePhoto(imageData []byte, metadata domain.PhotoMetadata) error {
	// Validate file size
	if metadata.Size > maxFileSize {
		return errors.New("Photo file size must not exceed 10MB")
	}

	// Validate format based on magic bytes
	if !v.isSupportedFormat(imageData) {
		return errors.New("Unsupported image format. Use JPEG, PNG, or WebP")
	}

	return nil
}

func (v *PhotoValidator) isSupportedFormat(data []byte) bool {
	if len(data) < 12 {
		return false
	}

	// JPEG: FF D8 FF
	if bytes.HasPrefix(data, []byte{0xFF, 0xD8, 0xFF}) {
		return true
	}

	// PNG: 89 50 4E 47 0D 0A 1A 0A
	if bytes.HasPrefix(data, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return true
	}

	// WebP: RIFF....WEBP
	if bytes.HasPrefix(data, []byte{0x52, 0x49, 0x46, 0x46}) && 
	   bytes.HasPrefix(data[8:], []byte{0x57, 0x45, 0x42, 0x50}) {
		return true
	}

	return false
}
