package domain

import (
	"encoding/base64"
	"errors"
)

// EncodeVerdictID encodes a file path (date directory + filename without extension)
// into a base64url-encoded identifier for use in shareable URLs.
// Example: "2026-02-01/153045_abc123" -> "MjAyNi0wMi0wMS8xNTMwNDVfYWJjMTIz"
func EncodeVerdictID(filePath string) string {
	if filePath == "" {
		return ""
	}
	return base64.URLEncoding.EncodeToString([]byte(filePath))
}

// DecodeVerdictID decodes a base64url-encoded verdict ID back into the original file path.
// Returns an error if the ID is empty, malformed, or cannot be decoded.
func DecodeVerdictID(encodedID string) (string, error) {
	if encodedID == "" {
		return "", errors.New("verdict ID cannot be empty")
	}
	
	decoded, err := base64.URLEncoding.DecodeString(encodedID)
	if err != nil {
		return "", errors.New("invalid verdict ID: malformed base64")
	}
	
	return string(decoded), nil
}
