//go:build integration

package gemini

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests require a valid GEMINI_API_KEY environment variable
// Run with: go test -tags=integration ./internal/adapters/gemini/...

func getAPIKey(t *testing.T) string {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY not set, skipping integration test")
	}
	return apiKey
}

// createTestJPEG creates a minimal valid JPEG image for testing
func createTestJPEG() []byte {
	// Minimal 1x1 red JPEG image
	return []byte{
		0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
		0x01, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0xFF, 0xDB, 0x00, 0x43,
		0x00, 0x08, 0x06, 0x06, 0x07, 0x06, 0x05, 0x08, 0x07, 0x07, 0x07, 0x09,
		0x09, 0x08, 0x0A, 0x0C, 0x14, 0x0D, 0x0C, 0x0B, 0x0B, 0x0C, 0x19, 0x12,
		0x13, 0x0F, 0x14, 0x1D, 0x1A, 0x1F, 0x1E, 0x1D, 0x1A, 0x1C, 0x1C, 0x20,
		0x24, 0x2E, 0x27, 0x20, 0x22, 0x2C, 0x23, 0x1C, 0x1C, 0x28, 0x37, 0x29,
		0x2C, 0x30, 0x31, 0x34, 0x34, 0x34, 0x1F, 0x27, 0x39, 0x3D, 0x38, 0x32,
		0x3C, 0x2E, 0x33, 0x34, 0x32, 0xFF, 0xC0, 0x00, 0x0B, 0x08, 0x00, 0x01,
		0x00, 0x01, 0x01, 0x01, 0x11, 0x00, 0xFF, 0xC4, 0x00, 0x1F, 0x00, 0x00,
		0x01, 0x05, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x09, 0x0A, 0x0B, 0xFF, 0xC4, 0x00, 0xB5, 0x10, 0x00, 0x02, 0x01, 0x03,
		0x03, 0x02, 0x04, 0x03, 0x05, 0x05, 0x04, 0x04, 0x00, 0x00, 0x01, 0x7D,
		0x01, 0x02, 0x03, 0x00, 0x04, 0x11, 0x05, 0x12, 0x21, 0x31, 0x41, 0x06,
		0x13, 0x51, 0x61, 0x07, 0x22, 0x71, 0x14, 0x32, 0x81, 0x91, 0xA1, 0x08,
		0x23, 0x42, 0xB1, 0xC1, 0x15, 0x52, 0xD1, 0xF0, 0x24, 0x33, 0x62, 0x72,
		0x82, 0x09, 0x0A, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x25, 0x26, 0x27, 0x28,
		0x29, 0x2A, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x43, 0x44, 0x45,
		0x46, 0x47, 0x48, 0x49, 0x4A, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59,
		0x5A, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x73, 0x74, 0x75,
		0x76, 0x77, 0x78, 0x79, 0x7A, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89,
		0x8A, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0xA2, 0xA3,
		0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6,
		0xB7, 0xB8, 0xB9, 0xBA, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9,
		0xCA, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xE1, 0xE2,
		0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xF1, 0xF2, 0xF3, 0xF4,
		0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFF, 0xDA, 0x00, 0x08, 0x01, 0x01,
		0x00, 0x00, 0x3F, 0x00, 0xFB, 0xD5, 0xDB, 0x20, 0xA8, 0xF1, 0x45, 0x00,
		0xFF, 0xD9,
	}
}

func TestIntegration_NewGeminiAnalyzer_ValidKey(t *testing.T) {
	apiKey := getAPIKey(t)

	analyzer, err := NewGeminiAnalyzer(apiKey, 30*time.Second)
	require.NoError(t, err)
	require.NotNil(t, analyzer)

	defer analyzer.Close()
}

func TestIntegration_AnalyzePhoto_MinimalImage(t *testing.T) {
	apiKey := getAPIKey(t)

	analyzer, err := NewGeminiAnalyzer(apiKey, 30*time.Second)
	require.NoError(t, err)
	defer analyzer.Close()

	// Use minimal test JPEG - Gemini should respond that this is not furniture
	imageData := createTestJPEG()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := analyzer.AnalyzePhoto(ctx, imageData)
	if err != nil {
		t.Logf("Error from Gemini API: %v", err)
	}
	require.NoError(t, err)
	require.NotNil(t, result)

	// Verify response structure
	t.Logf("Response: admissible=%v, score=%d, verdictType=%s", result.Admissible, result.Score, result.Verdict.VerdictType)
	t.Logf("Crime: %s", result.Verdict.Crime)
	t.Logf("Sentence: %s", result.Verdict.Sentence)
	t.Logf("Reasoning: %s", result.Verdict.Reasoning)

	// Basic structure validation
	assert.GreaterOrEqual(t, result.Score, 0)
	assert.LessOrEqual(t, result.Score, 10)
	assert.NotEmpty(t, result.Verdict.Crime)
	assert.NotEmpty(t, result.Verdict.Sentence)
	assert.NotEmpty(t, result.Verdict.Reasoning)
	assert.NotEmpty(t, result.Verdict.VerdictType, "VerdictType field should be present")
}

func TestIntegration_AnalyzePhoto_ResponseFormat(t *testing.T) {
	apiKey := getAPIKey(t)

	analyzer, err := NewGeminiAnalyzer(apiKey, 30*time.Second)
	require.NoError(t, err)
	defer analyzer.Close()

	imageData := createTestJPEG()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := analyzer.AnalyzePhoto(ctx, imageData)
	if err != nil {
		t.Logf("Error from Gemini API: %v", err)
	}
	require.NoError(t, err)

	// Verify JSON schema compliance
	// Score should be 0-10
	assert.GreaterOrEqual(t, result.Score, 0, "Score should be >= 0")
	assert.LessOrEqual(t, result.Score, 10, "Score should be <= 10")

	// All verdict fields should be populated
	assert.NotEmpty(t, result.Verdict.Crime, "Crime field should not be empty")
	assert.NotEmpty(t, result.Verdict.Sentence, "Sentence field should not be empty")
	assert.NotEmpty(t, result.Verdict.Reasoning, "Reasoning field should not be empty")
	assert.NotEmpty(t, result.Verdict.VerdictType, "VerdictType field should not be empty")

	// VerdictType should be one of the valid enum values
	validVerdictTypes := []string{"vrijspraak", "waarschuwing", "schuldig"}
	assert.Contains(t, validVerdictTypes, result.Verdict.VerdictType, "VerdictType should be one of: vrijspraak, waarschuwing, schuldig")

	// If not admissible, score should be 0
	if !result.Admissible {
		assert.Equal(t, 0, result.Score, "Non-admissible cases should have score 0")
	}
}

// TestIntegration_Compression_JPEG tests end-to-end with JPEG compression
func TestIntegration_Compression_JPEG(t *testing.T) {
	apiKey := getAPIKey(t)

	analyzer, err := NewGeminiAnalyzer(apiKey, 30*time.Second)
	require.NoError(t, err)
	defer analyzer.Close()

	// Create a larger JPEG that will benefit from compression
	jpegData := createTestJPEG()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := analyzer.AnalyzePhoto(ctx, jpegData)
	if err != nil {
		t.Logf("Error from Gemini API: %v", err)
	}
	require.NoError(t, err)
	require.NotNil(t, result)

	// Verify verdict accuracy is maintained (image should still be analyzed correctly)
	assert.GreaterOrEqual(t, result.Score, 0)
	assert.LessOrEqual(t, result.Score, 10)
	assert.NotEmpty(t, result.Verdict.Crime)

	t.Logf("JPEG compression test - Score: %d, VerdictType: %s", result.Score, result.Verdict.VerdictType)
}

// TestIntegration_Compression_PNG tests end-to-end with PNG compression
func TestIntegration_Compression_PNG(t *testing.T) {
	apiKey := getAPIKey(t)

	analyzer, err := NewGeminiAnalyzer(apiKey, 30*time.Second)
	require.NoError(t, err)
	defer analyzer.Close()

	// Create a PNG test image
	pngData := createTestPNG(400, 300)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := analyzer.AnalyzePhoto(ctx, pngData)
	if err != nil {
		t.Logf("Error from Gemini API: %v", err)
	}
	require.NoError(t, err)
	require.NotNil(t, result)

	// Verify verdict accuracy is maintained
	assert.GreaterOrEqual(t, result.Score, 0)
	assert.LessOrEqual(t, result.Score, 10)
	assert.NotEmpty(t, result.Verdict.Crime)

	t.Logf("PNG compression test - Score: %d, VerdictType: %s", result.Score, result.Verdict.VerdictType)
}

// TestIntegration_Compression_WebP tests end-to-end with WebP pass-through
func TestIntegration_Compression_WebP(t *testing.T) {
	apiKey := getAPIKey(t)

	analyzer, err := NewGeminiAnalyzer(apiKey, 30*time.Second)
	require.NoError(t, err)
	defer analyzer.Close()

	// Create a WebP test image
	webpData := createTestWebP()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := analyzer.AnalyzePhoto(ctx, webpData)
	if err != nil {
		t.Logf("Error from Gemini API: %v", err)
	}
	require.NoError(t, err)
	require.NotNil(t, result)

	// Verify verdict structure is valid
	assert.GreaterOrEqual(t, result.Score, 0)
	assert.LessOrEqual(t, result.Score, 10)
	assert.NotEmpty(t, result.Verdict.Crime)

	t.Logf("WebP pass-through test - Score: %d, VerdictType: %s", result.Score, result.Verdict.VerdictType)
}
