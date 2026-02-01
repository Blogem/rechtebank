package gemini

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"testing"
)

// createTestJPEGWithDimensions creates a test JPEG image with the specified dimensions
func createTestJPEGWithDimensions(width, height int) []byte {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Fill with a pattern to ensure compression has something to work with
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8((x * 255) / width),
				G: uint8((y * 255) / height),
				B: 128,
				A: 255,
			})
		}
	}

	var buf bytes.Buffer
	jpeg.Encode(&buf, img, &jpeg.Options{Quality: 95}) // High quality for testing compression
	return buf.Bytes()
}

// createTestPNG creates a test PNG image with the specified dimensions
func createTestPNG(width, height int) []byte {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Fill with a pattern
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8((x * 255) / width),
				G: uint8((y * 255) / height),
				B: 128,
				A: 255,
			})
		}
	}

	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf.Bytes()
}

// TestCompressJPEG tests that JPEG compression produces smaller output at quality 75
func TestCompressJPEG(t *testing.T) {
	// Create a test JPEG image (high quality)
	originalJPEG := createTestJPEGWithDimensions(800, 600)
	originalSize := len(originalJPEG)

	// Compress the JPEG
	compressed, err := compressJPEG(originalJPEG)
	if err != nil {
		t.Fatalf("compressJPEG failed: %v", err)
	}

	compressedSize := len(compressed)

	// Verify compression reduced size
	if compressedSize >= originalSize {
		t.Errorf("Expected compressed size (%d) to be smaller than original (%d)", compressedSize, originalSize)
	}

	// Verify output is still valid JPEG
	if detectMIMEType(compressed) != "jpeg" {
		t.Error("Compressed output is not a valid JPEG")
	}
}

// TestCompressJPEG_FallbackWhenLarger tests fallback to original when compressed is larger
func TestCompressJPEG_FallbackWhenLarger(t *testing.T) {
	// Create a small, already optimized JPEG (low quality)
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, color.RGBA{R: 128, G: 128, B: 128, A: 255})
		}
	}

	var buf bytes.Buffer
	jpeg.Encode(&buf, img, &jpeg.Options{Quality: 50}) // Already low quality
	original := buf.Bytes()

	// Compress - should return original if compressed would be larger
	result, err := compressJPEG(original)
	if err != nil {
		t.Fatalf("compressJPEG failed: %v", err)
	}

	// Result should be <= original size (either compressed smaller or fallback to original)
	if len(result) > len(original) {
		t.Errorf("Result size (%d) should not be larger than original (%d)", len(result), len(original))
	}
}

// TestCompressPNG tests that PNG compression with BestSpeed level produces smaller output
func TestCompressPNG(t *testing.T) {
	// Create a test PNG image
	originalPNG := createTestPNG(800, 600)
	originalSize := len(originalPNG)

	// Compress the PNG
	compressed, err := compressPNG(originalPNG)
	if err != nil {
		t.Fatalf("compressPNG failed: %v", err)
	}

	compressedSize := len(compressed)

	// Verify compression reduced size (or at least didn't increase it significantly)
	// Note: PNG compression can be variable depending on content
	maxAllowedSize := int(float64(originalSize) * 1.1)
	if compressedSize > maxAllowedSize {
		t.Errorf("Compressed size (%d) should not be significantly larger than original (%d)", compressedSize, originalSize)
	}

	// Verify output is still valid PNG
	if detectMIMEType(compressed) != "png" {
		t.Error("Compressed output is not a valid PNG")
	}
}

// TestCompressPNG_FallbackWhenLarger tests fallback to original when compressed PNG is larger
func TestCompressPNG_FallbackWhenLarger(t *testing.T) {
	// Create a small PNG with solid color (already well compressed)
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, color.RGBA{R: 128, G: 128, B: 128, A: 255})
		}
	}

	var buf bytes.Buffer
	// Use BestCompression for original to make it hard to compress further
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	encoder.Encode(&buf, img)
	original := buf.Bytes()

	// Compress - should return original if compressed would be larger
	result, err := compressPNG(original)
	if err != nil {
		t.Fatalf("compressPNG failed: %v", err)
	}

	// Result should be <= original size
	if len(result) > len(original) {
		t.Errorf("Result size (%d) should not be larger than original (%d)", len(result), len(original))
	}
}

// createTestWebP creates a minimal valid WebP file for testing
func createTestWebP() []byte {
	// Minimal WebP file header (RIFF....WEBPVP8 ...)
	// This is a simplified WebP structure for testing purposes
	webp := []byte{
		0x52, 0x49, 0x46, 0x46, // "RIFF"
		0x1A, 0x00, 0x00, 0x00, // File size (26 bytes)
		0x57, 0x45, 0x42, 0x50, // "WEBP"
		0x56, 0x50, 0x38, 0x20, // "VP8 "
		0x0E, 0x00, 0x00, 0x00, // VP8 chunk size
		0x00, 0x00, 0x00, 0x9D, 0x01, 0x2A, // VP8 bitstream header (1x1 image)
		0x01, 0x00, 0x01, 0x00, 0x00, 0x00,
	}
	return webp
}

// TestPassThroughWebP tests that WebP images pass through unchanged
func TestPassThroughWebP(t *testing.T) {
	// Create a test WebP
	originalWebP := createTestWebP()
	originalSize := len(originalWebP)

	// Pass through WebP (should return unchanged)
	result := passThroughWebP(originalWebP)

	// Verify size is unchanged
	if len(result) != originalSize {
		t.Errorf("WebP pass-through changed size from %d to %d", originalSize, len(result))
	}

	// Verify content is identical
	for i := 0; i < len(originalWebP); i++ {
		if result[i] != originalWebP[i] {
			t.Errorf("WebP pass-through changed content at byte %d", i)
			break
		}
	}

	// Verify output is still valid WebP
	if detectMIMEType(result) != "webp" {
		t.Error("Pass-through output is not a valid WebP")
	}
}

// TestResizeIfNeeded_Oversized tests resizing images over 1600px
func TestResizeIfNeeded_Oversized(t *testing.T) {
	// Create a large image (2000x1500)
	img := image.NewRGBA(image.Rect(0, 0, 2000, 1500))

	// Resize
	resized := resizeIfNeeded(img)

	bounds := resized.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Should be resized to fit within 1600x1600
	if width > MaxDimension || height > MaxDimension {
		t.Errorf("Resized image (%dx%d) exceeds maximum dimension %d", width, height, MaxDimension)
	}

	// Verify it was actually resized (not passed through)
	if width == 2000 && height == 1500 {
		t.Error("Image was not resized when it should have been")
	}
}

// TestResizeIfNeeded_MaintainsAspectRatio tests that resize maintains aspect ratio
func TestResizeIfNeeded_MaintainsAspectRatio(t *testing.T) {
	// Create a wide image (3200x1200) - aspect ratio 8:3
	img := image.NewRGBA(image.Rect(0, 0, 3200, 1200))

	// Resize
	resized := resizeIfNeeded(img)

	bounds := resized.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Original aspect ratio: 3200/1200 = 2.666...
	originalAspect := 3200.0 / 1200.0
	newAspect := float64(width) / float64(height)

	// Aspect ratio should be maintained (within small tolerance for rounding)
	if newAspect < originalAspect-0.01 || newAspect > originalAspect+0.01 {
		t.Errorf("Aspect ratio not maintained: original=%.3f, new=%.3f", originalAspect, newAspect)
	}

	// Should be limited by width (1600px wide, proportionally shorter height)
	if width != MaxDimension {
		t.Errorf("Expected width to be %d, got %d", MaxDimension, width)
	}
}

// TestResizeIfNeeded_WithinLimit tests images within size limit remain unchanged
func TestResizeIfNeeded_WithinLimit(t *testing.T) {
	// Create an image within limits (1200x800)
	img := image.NewRGBA(image.Rect(0, 0, 1200, 800))

	// Resize
	resized := resizeIfNeeded(img)

	bounds := resized.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Should remain unchanged
	if width != 1200 || height != 800 {
		t.Errorf("Image size changed from 1200x800 to %dx%d when it should have stayed the same", width, height)
	}
}

// TestCompressImage_JPEG tests that compressImage routes JPEG to correct handler
func TestCompressImage_JPEG(t *testing.T) {
	// Create a test JPEG
	jpegData := createTestJPEGWithDimensions(800, 600)
	originalSize := len(jpegData)

	// Compress
	compressed, err := compressImage(jpegData)
	if err != nil {
		t.Fatalf("compressImage failed: %v", err)
	}

	// Should be compressed (or at least not larger)
	maxExpected := int(float64(originalSize) * 1.1)
	if len(compressed) > maxExpected {
		t.Errorf("Compressed JPEG size (%d) is larger than expected based on original (%d)", len(compressed), originalSize)
	}

	// Should still be a JPEG
	if detectMIMEType(compressed) != "jpeg" {
		t.Error("Compressed output is not a JPEG")
	}
}

// TestCompressImage_PNG tests that compressImage routes PNG to correct handler
func TestCompressImage_PNG(t *testing.T) {
	// Create a test PNG
	pngData := createTestPNG(800, 600)

	// Compress
	compressed, err := compressImage(pngData)
	if err != nil {
		t.Fatalf("compressImage failed: %v", err)
	}

	// Should still be a PNG
	if detectMIMEType(compressed) != "png" {
		t.Error("Compressed output is not a PNG")
	}
}

// TestCompressImage_WebP tests that compressImage passes through WebP
func TestCompressImage_WebP(t *testing.T) {
	// Create a test WebP
	webpData := createTestWebP()

	// Compress (should pass through)
	result, err := compressImage(webpData)
	if err != nil {
		t.Fatalf("compressImage failed: %v", err)
	}

	// Should be unchanged
	if len(result) != len(webpData) {
		t.Errorf("WebP size changed from %d to %d", len(webpData), len(result))
	}

	// Should still be a WebP
	if detectMIMEType(result) != "webp" {
		t.Error("Result is not a WebP")
	}
}

// TestCompressImage_InvalidFormat tests error handling for invalid image
func TestCompressImage_InvalidFormat(t *testing.T) {
	// Create invalid image data
	invalidData := []byte{0x00, 0x01, 0x02, 0x03}

	// Should fail gracefully and return original
	result, err := compressImage(invalidData)
	if err != nil {
		t.Fatalf("compressImage should not error on invalid format, got: %v", err)
	}

	// Should return original data when compression fails
	if len(result) != len(invalidData) {
		t.Error("Should return original data when format is invalid")
	}
}

// TestCompressImage_WithResize tests integration of resizing before compression
func TestCompressImage_WithResize(t *testing.T) {
	// Create a large JPEG (2400x1800)
	largeJPEG := createTestJPEGWithDimensions(2400, 1800)

	// Compress (should resize and then compress)
	compressed, err := compressImage(largeJPEG)
	if err != nil {
		t.Fatalf("compressImage failed: %v", err)
	}

	// Decode to verify dimensions
	img, err := jpeg.Decode(bytes.NewReader(compressed))
	if err != nil {
		t.Fatalf("Failed to decode compressed image: %v", err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Should be resized to max 1600px
	if width > MaxDimension || height > MaxDimension {
		t.Errorf("Compressed image (%dx%d) exceeds maximum dimension %d", width, height, MaxDimension)
	}

	// Should be significantly smaller than original
	if len(compressed) >= len(largeJPEG) {
		t.Errorf("Compressed large image (%d bytes) should be smaller than original (%d bytes)", len(compressed), len(largeJPEG))
	}
}

// captureLogOutput captures log output during function execution
func captureLogOutput(fn func()) string {
	// Create a pipe to capture log output
	r, w, _ := os.Pipe()
	oldOutput := log.Writer()
	log.SetOutput(w)

	// Execute function
	fn()

	// Restore original output
	log.SetOutput(oldOutput)
	w.Close()

	// Read captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

// TestCompressionLogging_Success tests that successful compression logs metrics
func TestCompressionLogging_Success(t *testing.T) {
	// Create a test JPEG
	jpegData := createTestJPEGWithDimensions(800, 600)

	// Capture logs
	logOutput := captureLogOutput(func() {
		compressImage(jpegData)
	})

	// Verify log contains expected fields
	if !strings.Contains(logOutput, "[COMPRESSION]") {
		t.Error("Log should contain [COMPRESSION] tag")
	}
	if !strings.Contains(logOutput, "originalSize=") {
		t.Error("Log should contain originalSize")
	}
	if !strings.Contains(logOutput, "compressedSize=") {
		t.Error("Log should contain compressedSize")
	}
	if !strings.Contains(logOutput, "compressionRatio=") {
		t.Error("Log should contain compressionRatio")
	}
	if !strings.Contains(logOutput, "imageFormat=") {
		t.Error("Log should contain imageFormat")
	}
}

// TestCompressionLogging_WebPPassthrough tests that WebP pass-through is logged
func TestCompressionLogging_WebPPassthrough(t *testing.T) {
	// Create a test WebP
	webpData := createTestWebP()

	// Capture logs
	logOutput := captureLogOutput(func() {
		compressImage(webpData)
	})

	// Verify log indicates pass-through
	if !strings.Contains(logOutput, "[COMPRESSION]") {
		t.Error("Log should contain [COMPRESSION] tag")
	}
	if !strings.Contains(logOutput, "WebP pass-through") {
		t.Error("Log should indicate WebP pass-through")
	}
	if !strings.Contains(logOutput, "imageFormat=webp") {
		t.Error("Log should contain imageFormat=webp")
	}
}

// TestCompressionLogging_Skipped tests that skipped compression is logged
func TestCompressionLogging_Skipped(t *testing.T) {
	// Create invalid data
	invalidData := []byte{0x00, 0x01, 0x02, 0x03}

	// Capture logs
	logOutput := captureLogOutput(func() {
		compressImage(invalidData)
	})

	// Verify log indicates skipped compression
	if !strings.Contains(logOutput, "[COMPRESSION]") {
		t.Error("Log should contain [COMPRESSION] tag")
	}
	if !strings.Contains(logOutput, "Skipped") {
		t.Error("Log should indicate compression was skipped")
	}
}
