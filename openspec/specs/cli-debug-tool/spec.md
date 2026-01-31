## ADDED Requirements

### Requirement: CLI accepts image file path
The CLI tool SHALL accept an image file path as a command-line argument and load the image from disk.

#### Scenario: Valid image file provided
- **WHEN** user runs `debug-gemini /path/to/image.jpg`
- **THEN** tool loads the image file and proceeds with analysis

#### Scenario: Image file does not exist
- **WHEN** user runs `debug-gemini /path/to/nonexistent.jpg`
- **THEN** tool exits with error message "Image file not found: /path/to/nonexistent.jpg"

#### Scenario: File is not a valid image
- **WHEN** user runs `debug-gemini /path/to/document.pdf`
- **THEN** tool exits with error message "Unsupported image format (must be JPEG, PNG, or WebP)"

### Requirement: CLI displays image metadata
The CLI tool SHALL display metadata about the loaded image before sending to Gemini API.

#### Scenario: Image metadata shown
- **WHEN** tool loads an image successfully
- **THEN** tool displays file path, file size, MIME type, and dimensions

### Requirement: CLI shows system prompt
The CLI tool SHALL display the complete system prompt that will be sent to the Gemini API.

#### Scenario: System prompt displayed before API call
- **WHEN** tool prepares to call Gemini API
- **THEN** tool prints the full system prompt text under a "SYSTEM PROMPT" header

### Requirement: CLI shows user prompt
The CLI tool SHALL display the user prompt sent alongside the image.

#### Scenario: User prompt displayed
- **WHEN** tool prepares to call Gemini API
- **THEN** tool prints the user prompt text under a "USER PROMPT" header

### Requirement: CLI displays request metadata
The CLI tool SHALL display configuration details for the API request.

#### Scenario: Request metadata shown
- **WHEN** tool prepares to call Gemini API
- **THEN** tool displays model name, timeout duration, and max retry count

### Requirement: CLI shows raw API response
The CLI tool SHALL display the raw JSON response received from the Gemini API.

#### Scenario: Raw response displayed on success
- **WHEN** Gemini API returns a successful response
- **THEN** tool prints the complete JSON response under a "RAW RESPONSE" header

#### Scenario: Raw response shown on error
- **WHEN** Gemini API returns an error
- **THEN** tool prints the error response and stack trace

### Requirement: CLI displays parsed verdict
The CLI tool SHALL display the parsed verdict in a human-readable format.

#### Scenario: Parsed verdict shown after successful response
- **WHEN** tool successfully parses the Gemini response
- **THEN** tool displays admissible status, score, crime, sentence, and reasoning in a formatted layout

### Requirement: CLI exits with appropriate code
The CLI tool SHALL exit with status code 0 on success and non-zero on failure.

#### Scenario: Successful analysis
- **WHEN** tool completes analysis without errors
- **THEN** tool exits with status code 0

#### Scenario: Analysis fails
- **WHEN** tool encounters an error (missing file, API error, parse error)
- **THEN** tool exits with status code 1

### Requirement: CLI uses production analyzer code
The CLI tool SHALL use the existing GeminiAnalyzer from the adapters package to ensure behavior matches production.

#### Scenario: Same analyzer used
- **WHEN** tool analyzes an image
- **THEN** tool imports and uses the GeminiAnalyzer from `backend/internal/adapters/gemini`
