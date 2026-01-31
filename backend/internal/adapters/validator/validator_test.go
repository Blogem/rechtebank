package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"rechtebank/backend/internal/core/domain"
)

// JPEG magic bytes: FF D8 FF
var jpegHeader = []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01}

// PNG magic bytes: 89 50 4E 47 0D 0A 1A 0A
var pngHeader = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00, 0x00, 0x0D}

// WebP magic bytes: RIFF....WEBP
var webpHeader = []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x57, 0x45, 0x42, 0x50}

// GIF magic bytes: GIF89a
var gifHeader = []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

// BMP magic bytes: BM
var bmpHeader = []byte{0x42, 0x4D, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func TestPhotoValidator_FileFormatValidation_JPEG(t *testing.T) {
	v := NewPhotoValidator()
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        int64(len(jpegHeader)),
	}

	err := v.ValidatePhoto(jpegHeader, metadata)
	assert.NoError(t, err)
}

func TestPhotoValidator_FileFormatValidation_PNG(t *testing.T) {
	v := NewPhotoValidator()
	metadata := domain.PhotoMetadata{
		Filename:    "test.png",
		ContentType: "image/png",
		Size:        int64(len(pngHeader)),
	}

	err := v.ValidatePhoto(pngHeader, metadata)
	assert.NoError(t, err)
}

func TestPhotoValidator_FileFormatValidation_WebP(t *testing.T) {
	v := NewPhotoValidator()
	metadata := domain.PhotoMetadata{
		Filename:    "test.webp",
		ContentType: "image/webp",
		Size:        int64(len(webpHeader)),
	}

	err := v.ValidatePhoto(webpHeader, metadata)
	assert.NoError(t, err)
}

func TestPhotoValidator_UnsupportedFormat_GIF(t *testing.T) {
	v := NewPhotoValidator()
	metadata := domain.PhotoMetadata{
		Filename:    "test.gif",
		ContentType: "image/gif",
		Size:        int64(len(gifHeader)),
	}

	err := v.ValidatePhoto(gifHeader, metadata)
	assert.Error(t, err)
	assert.Equal(t, "Unsupported image format. Use JPEG, PNG, or WebP", err.Error())
}

func TestPhotoValidator_UnsupportedFormat_BMP(t *testing.T) {
	v := NewPhotoValidator()
	metadata := domain.PhotoMetadata{
		Filename:    "test.bmp",
		ContentType: "image/bmp",
		Size:        int64(len(bmpHeader)),
	}

	err := v.ValidatePhoto(bmpHeader, metadata)
	assert.Error(t, err)
	assert.Equal(t, "Unsupported image format. Use JPEG, PNG, or WebP", err.Error())
}

func TestPhotoValidator_FileSizeValidation_WithinLimit(t *testing.T) {
	v := NewPhotoValidator()
	// Create data within limit (1MB)
	size := int64(1 * 1024 * 1024)
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        size,
	}

	err := v.ValidatePhoto(jpegHeader, metadata)
	assert.NoError(t, err)
}

func TestPhotoValidator_FileSizeValidation_AtLimit(t *testing.T) {
	v := NewPhotoValidator()
	// Exactly 10MB
	size := int64(10 * 1024 * 1024)
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        size,
	}

	err := v.ValidatePhoto(jpegHeader, metadata)
	assert.NoError(t, err)
}

func TestPhotoValidator_FileSizeValidation_ExceedsLimit(t *testing.T) {
	v := NewPhotoValidator()
	// 11MB - over limit
	size := int64(11 * 1024 * 1024)
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        size,
	}

	err := v.ValidatePhoto(jpegHeader, metadata)
	assert.Error(t, err)
	assert.Equal(t, "Photo file size must not exceed 10MB", err.Error())
}

func TestPhotoValidator_FileSizeValidation_VeryLarge(t *testing.T) {
	v := NewPhotoValidator()
	// 100MB - way over limit
	size := int64(100 * 1024 * 1024)
	metadata := domain.PhotoMetadata{
		Filename:    "large.jpg",
		ContentType: "image/jpeg",
		Size:        size,
	}

	err := v.ValidatePhoto(jpegHeader, metadata)
	assert.Error(t, err)
	assert.Equal(t, "Photo file size must not exceed 10MB", err.Error())
}

func TestPhotoValidator_InvalidData_TooShort(t *testing.T) {
	v := NewPhotoValidator()
	shortData := []byte{0xFF, 0xD8}
	metadata := domain.PhotoMetadata{
		Filename:    "test.jpg",
		ContentType: "image/jpeg",
		Size:        int64(len(shortData)),
	}

	err := v.ValidatePhoto(shortData, metadata)
	assert.Error(t, err)
	assert.Equal(t, "Unsupported image format. Use JPEG, PNG, or WebP", err.Error())
}
