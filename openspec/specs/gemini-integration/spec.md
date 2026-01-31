## ADDED Requirements

### Requirement: Initialize Gemini API client
The system SHALL initialize the Google Gemini 2.5 Flash Lite API client with valid API credentials on service startup.

#### Scenario: Valid API key
- **WHEN** service starts with valid GEMINI_API_KEY environment variable
- **THEN** Gemini client initializes successfully

#### Scenario: Missing API key
- **WHEN** service starts without GEMINI_API_KEY environment variable
- **THEN** service fails to start with error "GEMINI_API_KEY environment variable is required"

#### Scenario: Invalid API key
- **WHEN** service makes first API call with invalid credentials
- **THEN** system returns HTTP 500 with error message "AI analysis service unavailable"

### Requirement: Send photo with legal system prompt
The system SHALL send the uploaded photo to Gemini API with a Dutch legal-themed system instruction prompt that requests structured verdict components.

#### Scenario: Photo analysis request with structured output
- **WHEN** system sends photo to Gemini API
- **THEN** request includes system prompt instructing Gemini to analyze furniture and return structured JSON with crime, sentence, and reasoning fields in Dutch legal jargon

#### Scenario: JSON schema definition
- **WHEN** system configures Gemini request
- **THEN** request includes JSON schema defining required fields: admissible (boolean), score (integer 0-10), verdict object with crime, sentence, and reasoning (strings)

### Requirement: Parse Gemini response
The system SHALL parse the Gemini API structured JSON response to extract furniture type, alignment score, and verdict components (crime, sentence, reasoning).

#### Scenario: Valid furniture detected
- **WHEN** Gemini responds with furniture analysis
- **THEN** system extracts score (1-10), crime description, sentence text, reasoning text, and determines case is admissible

#### Scenario: Non-furniture detected
- **WHEN** Gemini identifies object is not furniture
- **THEN** system marks case as non-admissible with score 0 and appropriate verdict components

#### Scenario: Gemini API error
- **WHEN** Gemini API returns error response
- **THEN** system returns HTTP 502 with error message "AI analysis failed"

#### Scenario: Invalid JSON schema response
- **WHEN** Gemini response does not match expected JSON schema
- **THEN** system returns HTTP 502 with error message "Invalid AI response format"

### Requirement: Handle API rate limits
The system SHALL handle Gemini API rate limits gracefully with exponential backoff retry logic.

#### Scenario: Rate limit exceeded
- **WHEN** Gemini API returns 429 rate limit error
- **THEN** system retries request with exponential backoff up to 3 attempts

#### Scenario: Retry exhausted
- **WHEN** all retry attempts fail due to rate limiting
- **THEN** system returns HTTP 503 with error message "AI analysis service temporarily unavailable"

### Requirement: Set request timeout
The system SHALL timeout Gemini API requests after 30 seconds to prevent indefinite waiting.

#### Scenario: Gemini responds within timeout
- **WHEN** Gemini API responds within 30 seconds
- **THEN** system processes response normally

#### Scenario: Gemini timeout
- **WHEN** Gemini API does not respond within 30 seconds
- **THEN** system returns HTTP 504 with error message "AI analysis timeout"
