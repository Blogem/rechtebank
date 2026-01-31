package domain

// VerdictResponse represents the full verdict response from the API
type VerdictResponse struct {
	Admissible bool           `json:"admissible"`
	Score      int            `json:"score"`
	Verdict    VerdictDetails `json:"verdict"`
	RequestID  string         `json:"requestId"`
	Timestamp  string         `json:"timestamp"`
}

// VerdictDetails contains the structured components of the legal verdict
type VerdictDetails struct {
	Crime     string `json:"crime"`     // The furniture offense
	Sentence  string `json:"sentence"`  // The punishment
	Reasoning string `json:"reasoning"` // Legal justification
}

// PhotoMetadata contains information about an uploaded photo
type PhotoMetadata struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Size        int64  `json:"size"`
}
