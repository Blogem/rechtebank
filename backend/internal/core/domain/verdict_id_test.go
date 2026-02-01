package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeVerdictID(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "simple path with date and filename",
			filePath: "2026-02-01/153045_abc123",
			expected: "MjAyNi0wMi0wMS8xNTMwNDVfYWJjMTIz",
		},
		{
			name:     "path with different date",
			filePath: "2026-01-15/093022_xyz789",
			expected: "MjAyNi0wMS0xNS8wOTMwMjJfeHl6Nzg5",
		},
		{
			name:     "empty path",
			filePath: "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EncodeVerdictID(tt.filePath)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDecodeVerdictID(t *testing.T) {
	tests := []struct {
		name        string
		encodedID   string
		expected    string
		expectError bool
	}{
		{
			name:        "valid encoded ID",
			encodedID:   "MjAyNi0wMi0wMS8xNTMwNDVfYWJjMTIz",
			expected:    "2026-02-01/153045_abc123",
			expectError: false,
		},
		{
			name:        "another valid encoded ID",
			encodedID:   "MjAyNi0wMS0xNS8wOTMwMjJfeHl6Nzg5",
			expected:    "2026-01-15/093022_xyz789",
			expectError: false,
		},
		{
			name:        "empty ID",
			encodedID:   "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid base64",
			encodedID:   "not-valid-base64!!!",
			expected:    "",
			expectError: true,
		},
		{
			name:        "malformed base64",
			encodedID:   "@@@@",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DecodeVerdictID(tt.encodedID)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestVerdictID_RoundTrip(t *testing.T) {
	originalPath := "2026-02-01/153045_abc123"

	// Encode
	encoded := EncodeVerdictID(originalPath)
	assert.NotEmpty(t, encoded)

	// Decode
	decoded, err := DecodeVerdictID(encoded)
	assert.NoError(t, err)
	assert.Equal(t, originalPath, decoded)
}
