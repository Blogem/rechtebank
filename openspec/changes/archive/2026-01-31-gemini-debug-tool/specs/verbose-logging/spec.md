## ADDED Requirements

### Requirement: Output sections clearly separated
The verbose logging output SHALL use visual separators to distinguish between different sections of information.

#### Scenario: Section headers displayed
- **WHEN** tool outputs information
- **THEN** each major section (image metadata, prompts, response) is preceded by a header with "===" markers

### Requirement: System prompt logged in full
The verbose logging SHALL output the complete system prompt text without truncation.

#### Scenario: Full system prompt shown
- **WHEN** tool displays the system prompt
- **THEN** entire prompt text is printed, including all instructions and guidelines

### Requirement: Image metadata logged
The verbose logging SHALL output detailed metadata about the input image.

#### Scenario: Complete image metadata shown
- **WHEN** tool loads an image
- **THEN** tool logs file path, file size in KB, MIME type, and pixel dimensions

### Requirement: Request configuration logged
The verbose logging SHALL output the API request configuration parameters.

#### Scenario: Request config displayed
- **WHEN** tool prepares API request
- **THEN** tool logs model name, timeout setting, and retry limit

### Requirement: Raw JSON response logged
The verbose logging SHALL output the complete raw JSON response from Gemini without modification.

#### Scenario: Raw JSON printed
- **WHEN** Gemini API returns a response
- **THEN** tool prints the exact JSON string received, formatted for readability

#### Scenario: Empty response logged
- **WHEN** Gemini API returns an empty response
- **THEN** tool logs this as "empty response from Gemini" with error context

### Requirement: Response parsing errors logged
The verbose logging SHALL output detailed error messages when response parsing fails.

#### Scenario: JSON parse error shown
- **WHEN** tool fails to parse Gemini JSON response
- **THEN** tool logs the parsing error message and the problematic JSON string

### Requirement: API errors logged with detail
The verbose logging SHALL output comprehensive error information for API failures.

#### Scenario: Timeout error logged
- **WHEN** Gemini API request times out
- **THEN** tool logs timeout duration and context about the request

#### Scenario: Rate limit error logged
- **WHEN** Gemini API returns rate limit error
- **THEN** tool logs the error and retry-after duration if available

#### Scenario: Invalid response error logged
- **WHEN** Gemini API returns unexpected response format
- **THEN** tool logs the complete response and format validation errors

### Requirement: Logging does not retry automatically
The verbose logging SHALL display errors without automatic retry to show actual API behavior.

#### Scenario: No automatic retry on error
- **WHEN** API call fails with recoverable error
- **THEN** tool logs the error and exits without attempting retry

### Requirement: Output remains machine-parseable
The verbose logging SHALL structure output so key sections can be extracted programmatically if needed.

#### Scenario: Consistent section markers
- **WHEN** tool outputs verbose logs
- **THEN** each section uses consistent header format (e.g., "=== SECTION NAME ===")
