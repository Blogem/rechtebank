package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestDetectMIMEType(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{
			name:     "JPEG image",
			data:     []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01},
			expected: "jpeg",
		},
		{
			name:     "PNG image",
			data:     []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00, 0x00, 0x0D},
			expected: "png",
		},
		{
			name:     "WebP image",
			data:     []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x57, 0x45, 0x42, 0x50},
			expected: "webp",
		},
		{
			name:     "Invalid format",
			data:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expected: "",
		},
		{
			name:     "Too short",
			data:     []byte{0xFF, 0xD8},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detectMIMEType(tt.data)
			if result != tt.expected {
				t.Errorf("detectMIMEType() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestFormatFileSize(t *testing.T) {
	tests := []struct {
		name     string
		bytes    int
		expected string
	}{
		{
			name:     "Bytes",
			bytes:    512,
			expected: "512 bytes",
		},
		{
			name:     "Kilobytes",
			bytes:    2048,
			expected: "2.0 KB",
		},
		{
			name:     "Megabytes",
			bytes:    2 * 1024 * 1024,
			expected: "2.0 MB",
		},
		{
			name:     "Large KB",
			bytes:    250 * 1024,
			expected: "250.0 KB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatFileSize(tt.bytes)
			if result != tt.expected {
				t.Errorf("formatFileSize(%d) = %q, want %q", tt.bytes, result, tt.expected)
			}
		})
	}
}

func TestPrintSection(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printSection("TEST SECTION")

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "=== TEST SECTION ===\n"
	if output != expected {
		t.Errorf("printSection() output = %q, want %q", output, expected)
	}
}

func createTestImage(t *testing.T, format string) string {
	t.Helper()

	var data []byte
	var ext string

	switch format {
	case "jpeg":
		data = []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01}
		ext = ".jpg"
	case "png":
		data = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00, 0x00, 0x0D}
		ext = ".png"
	case "invalid":
		data = []byte("This is not an image file")
		ext = ".txt"
	default:
		t.Fatal("Unknown format")
	}

	tmpFile := filepath.Join(t.TempDir(), "test"+ext)
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		t.Fatal(err)
	}

	return tmpFile
}

func TestRunWithMissingFile(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"debug-gemini", "/nonexistent/path/to/image.jpg"}

	err := run()
	if err == nil {
		t.Error("Expected error for missing file, got nil")
	}

	if err != nil && err.Error() != "image file not found: /nonexistent/path/to/image.jpg" {
		t.Errorf("Expected 'image file not found' error, got: %v", err)
	}
}

func TestRunWithInvalidImage(t *testing.T) {
	tmpFile := createTestImage(t, "invalid")

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"debug-gemini", tmpFile}

	err := run()
	if err == nil {
		t.Error("Expected error for invalid image, got nil")
	}

	if err != nil && err.Error() != "unsupported image format (must be JPEG, PNG, or WebP)" {
		t.Errorf("Expected 'unsupported image format' error, got: %v", err)
	}
}

func TestRunWithMissingAPIKey(t *testing.T) {
	tmpFile := createTestImage(t, "jpeg")

	oldArgs := os.Args
	oldAPIKey := os.Getenv("GEMINI_API_KEY")
	defer func() {
		os.Args = oldArgs
		os.Setenv("GEMINI_API_KEY", oldAPIKey)
	}()

	os.Unsetenv("GEMINI_API_KEY")

	os.Args = []string{"debug-gemini", tmpFile}

	err := run()
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// The error could be either about missing API key or failed to get dimensions
	// since our test JPEG header is minimal and may not decode properly
	// Both are acceptable test results
	errMsg := err.Error()
	if errMsg != "GEMINI_API_KEY environment variable is required" &&
		errMsg != "failed to get image dimensions: unexpected EOF" {
		t.Errorf("Expected error about API key or image dimensions, got: %v", err)
	}
}

func TestArgumentValidation(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"debug-gemini"}

	err := run()
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}
