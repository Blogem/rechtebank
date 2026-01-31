package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerdictResponse_JSONMarshaling(t *testing.T) {
	verdict := VerdictResponse{
		Admissible: true,
		Score:      8,
		Verdict: VerdictDetails{
			Crime:     "Rugleuning-afwijking van 5 graden",
			Sentence:  "Veroordeeld tot lichte berisping",
			Reasoning: "Artikel 42 van de Meubilair-wet",
		},
		RequestID: "test-request-123",
		Timestamp: "2026-01-30T10:00:00Z",
	}

	// Test marshaling
	data, err := json.Marshal(verdict)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"admissible":true`)
	assert.Contains(t, string(data), `"score":8`)
	assert.Contains(t, string(data), `"requestId":"test-request-123"`)

	// Test unmarshaling
	var decoded VerdictResponse
	err = json.Unmarshal(data, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, verdict.Admissible, decoded.Admissible)
	assert.Equal(t, verdict.Score, decoded.Score)
	assert.Equal(t, verdict.RequestID, decoded.RequestID)
	assert.Equal(t, verdict.Timestamp, decoded.Timestamp)
}

func TestVerdictDetails_AllFields(t *testing.T) {
	details := VerdictDetails{
		Crime:     "Ernstige horizontale schending",
		Sentence:  "Veroordeeld tot de brandstapel",
		Reasoning: "Het wetboek vereist perfecte rechtheid",
	}

	data, err := json.Marshal(details)
	assert.NoError(t, err)

	var decoded VerdictDetails
	err = json.Unmarshal(data, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, details.Crime, decoded.Crime)
	assert.Equal(t, details.Sentence, decoded.Sentence)
	assert.Equal(t, details.Reasoning, decoded.Reasoning)
}

func TestPhotoMetadata_Marshaling(t *testing.T) {
	metadata := PhotoMetadata{
		Filename:    "couch.jpg",
		ContentType: "image/jpeg",
		Size:        1024000,
	}

	data, err := json.Marshal(metadata)
	assert.NoError(t, err)

	var decoded PhotoMetadata
	err = json.Unmarshal(data, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, metadata.Filename, decoded.Filename)
	assert.Equal(t, metadata.ContentType, decoded.ContentType)
	assert.Equal(t, metadata.Size, decoded.Size)
}
