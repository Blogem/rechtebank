package gemini

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg" // Register JPEG decoder
	"image/png"
	_ "image/png" // Register PNG decoder
	"log"

	"golang.org/x/image/draw"
)

// Compression constants
// These values are chosen to balance file size reduction with image quality retention
// for furniture recognition by the Gemini LLM API.
const (
	// JPEGQuality sets the JPEG compression quality (0-100).
	// Quality 75 provides ~50-60% file size reduction while maintaining sufficient
	// detail for furniture edge detection and shape recognition.
	// Lower values (50-60) risk degrading fine details; higher values (85-90)
	// provide minimal additional compression benefit.
	JPEGQuality = 75

	// MaxDimension is the maximum width or height in pixels before resizing.
	// Images larger than 1600x1600 are resized proportionally to fit within this boundary.
	// Rationale:
	// - Typical phone photos are 3000-4000px, so this saves significant bandwidth
	// - Gemini API likely downsamples large images internally anyway
	// - Furniture can be recognized clearly at 1600px resolution
	// - Reduces token consumption without affecting recognition accuracy
	MaxDimension = 1600
)

// compressJPEG compresses a JPEG image to quality 75
// Returns the original if compressed version would be larger
func compressJPEG(data []byte) ([]byte, error) {
	// Decode the image
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode JPEG: %w", err)
	}

	// Encode with quality 75
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: JPEGQuality})
	if err != nil {
		return nil, fmt.Errorf("failed to encode JPEG: %w", err)
	}

	compressed := buf.Bytes()

	// Fallback to original if compressed is larger
	if len(compressed) >= len(data) {
		return data, nil
	}

	return compressed, nil
}

// compressPNG compresses a PNG image using BestSpeed compression level
// Returns the original if compressed version would be larger
func compressPNG(data []byte) ([]byte, error) {
	// Decode the image
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode PNG: %w", err)
	}

	// Encode with BestSpeed compression
	var buf bytes.Buffer
	encoder := png.Encoder{CompressionLevel: png.BestSpeed}
	err = encoder.Encode(&buf, img)
	if err != nil {
		return nil, fmt.Errorf("failed to encode PNG: %w", err)
	}

	compressed := buf.Bytes()

	// Fallback to original if compressed is larger
	if len(compressed) >= len(data) {
		return data, nil
	}

	return compressed, nil
}

// passThroughWebP returns WebP images unchanged as they are already efficiently compressed
func passThroughWebP(data []byte) []byte {
	return data
}

// resizeIfNeeded resizes an image if either dimension exceeds MaxDimension
// Maintains aspect ratio and returns the original if within limits
func resizeIfNeeded(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Check if resize is needed
	if width <= MaxDimension && height <= MaxDimension {
		return img
	}

	// Calculate new dimensions maintaining aspect ratio
	var newWidth, newHeight int
	if width > height {
		// Width is the limiting dimension
		newWidth = MaxDimension
		newHeight = (height * MaxDimension) / width
	} else {
		// Height is the limiting dimension
		newHeight = MaxDimension
		newWidth = (width * MaxDimension) / height
	}

	// Create new image with calculated dimensions
	resized := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Use high-quality scaling
	draw.BiLinear.Scale(resized, resized.Bounds(), img, img.Bounds(), draw.Over, nil)

	return resized
}

// compressImage compresses an image based on its MIME type to reduce token consumption
// for Gemini LLM API calls while maintaining sufficient quality for furniture analysis.
//
// Compression Strategy:
// 1. JPEG: Re-encode at quality 75 (typically 50-60% size reduction)
// 2. PNG: Re-encode with BestSpeed compression level (fast, moderate compression)
// 3. WebP: Pass through unchanged (already efficiently compressed)
// 4. Resize: If either dimension > 1600px, resize proportionally before compression
// 5. Fallback: Return original image if compression fails or produces larger output
//
// Error Handling:
// - Invalid/unknown formats return original data
// - Decode/encode errors return original data
// - All errors are logged but don't fail the request
//
// Resizes if needed, then applies format-specific compression
// Falls back to original image on any error
func compressImage(imageData []byte) ([]byte, error) {
	originalSize := len(imageData)

	// Detect MIME type
	mimeType := detectMIMEType(imageData)
	if mimeType == "" {
		// Unknown format, return original
		log.Printf("[COMPRESSION] Skipped: unknown image format, originalSize=%d", originalSize)
		return imageData, nil
	}

	// WebP: pass through unchanged
	if mimeType == "webp" {
		log.Printf("[COMPRESSION] WebP pass-through: originalSize=%d, imageFormat=%s", originalSize, mimeType)
		return passThroughWebP(imageData), nil
	}

	// For JPEG and PNG: decode, resize if needed, then compress
	var img image.Image
	var decodeErr error

	switch mimeType {
	case "jpeg":
		img, decodeErr = jpeg.Decode(bytes.NewReader(imageData))
	case "png":
		img, decodeErr = png.Decode(bytes.NewReader(imageData))
	default:
		// Unknown format, return original
		log.Printf("[COMPRESSION] Skipped: unsupported format %s, originalSize=%d", mimeType, originalSize)
		return imageData, nil
	}

	if decodeErr != nil {
		// Failed to decode, return original
		log.Printf("[COMPRESSION] Decode failed: %v, using original, originalSize=%d, imageFormat=%s", decodeErr, originalSize, mimeType)
		return imageData, nil
	}

	// Resize if needed
	img = resizeIfNeeded(img)

	// Compress based on format
	var compressed []byte
	var compressErr error

	switch mimeType {
	case "jpeg":
		// Encode with quality 75
		var buf bytes.Buffer
		compressErr = jpeg.Encode(&buf, img, &jpeg.Options{Quality: JPEGQuality})
		if compressErr == nil {
			compressed = buf.Bytes()
		}
	case "png":
		// Encode with BestSpeed compression
		var buf bytes.Buffer
		encoder := png.Encoder{CompressionLevel: png.BestSpeed}
		compressErr = encoder.Encode(&buf, img)
		if compressErr == nil {
			compressed = buf.Bytes()
		}
	}

	if compressErr != nil {
		// Compression failed, return original
		log.Printf("[COMPRESSION] Encode failed: %v, using original, originalSize=%d, imageFormat=%s", compressErr, originalSize, mimeType)
		return imageData, nil
	}

	// Return compressed if smaller, otherwise return original
	compressedSize := len(compressed)
	if compressedSize < originalSize {
		compressionRatio := float64(originalSize) / float64(compressedSize)
		log.Printf("[COMPRESSION] Success: originalSize=%d, compressedSize=%d, compressionRatio=%.2fx, imageFormat=%s",
			originalSize, compressedSize, compressionRatio, mimeType)
		return compressed, nil
	}

	// Compressed is larger, use original
	log.Printf("[COMPRESSION] Skipped: compressed larger than original, originalSize=%d, compressedSize=%d, imageFormat=%s",
		originalSize, compressedSize, mimeType)
	return imageData, nil
}
