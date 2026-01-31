## 1. Project Setup

- [x] 1.1 Create `backend/cmd/debug-gemini/` directory
- [x] 1.2 Create `backend/cmd/debug-gemini/main.go` with basic CLI structure
- [x] 1.3 Add build configuration to compile the debug tool

## 2. CLI Argument Handling

- [x] 2.1 Implement command-line argument parsing for image file path
- [x] 2.2 Add validation to check if file exists
- [x] 2.3 Add validation to check if file is a valid image format (JPEG, PNG, WebP)
- [x] 2.4 Implement error messages for missing or invalid files

## 3. Image Loading and Metadata

- [x] 3.1 Implement function to load image file from disk
- [x] 3.2 Implement function to detect image MIME type from file bytes
- [x] 3.3 Implement function to get image dimensions
- [x] 3.4 Implement function to format file size in human-readable format (KB/MB)

## 4. Verbose Output Implementation

- [x] 4.1 Create output formatter with section headers (=== format)
- [x] 4.2 Implement function to display image metadata section
- [x] 4.3 Add function to display system prompt from analyzer
- [x] 4.4 Add function to display user prompt
- [x] 4.5 Add function to display request configuration (model, timeout, retries)

## 5. Analyzer Integration

- [x] 5.1 Import GeminiAnalyzer from `backend/internal/adapters/gemini`
- [x] 5.2 Load GEMINI_API_KEY from environment variables
- [x] 5.3 Initialize GeminiAnalyzer with API key and timeout
- [x] 5.4 Create wrapper to expose system prompt from analyzer (may need to export systemPrompt constant)
- [x] 5.5 Create wrapper to expose request metadata (model name, timeout, retry config)

## 6. API Call and Response Logging

- [x] 6.1 Implement function to call analyzer with loaded image
- [x] 6.2 Capture and display raw JSON response from Gemini API
- [x] 6.3 Format raw JSON for readability (pretty-print)
- [x] 6.4 Display parsed verdict in human-readable format
- [x] 6.5 Handle and display API errors (timeout, rate limit, invalid response)

## 7. Error Handling

- [x] 7.1 Implement exit code 0 for successful analysis
- [x] 7.2 Implement exit code 1 for all error cases
- [x] 7.3 Add error handling for missing API key
- [x] 7.4 Add error handling for API timeout errors
- [x] 7.5 Add error handling for rate limit errors
- [x] 7.6 Add error handling for JSON parsing errors
- [x] 7.7 Display full error context without automatic retry

## 8. Testing

- [x] 8.1 Create test with valid JPEG image
- [x] 8.2 Create test with valid PNG image
- [x] 8.3 Create test with invalid file (non-image)
- [x] 8.4 Create test with missing file
- [x] 8.5 Test output formatting matches specification
- [x] 8.6 Verify all section headers are displayed correctly
- [x] 8.7 Compare debug tool output with production analyzer behavior

## 9. Documentation

- [x] 9.1 Add README to `backend/cmd/debug-gemini/` with usage examples
- [x] 9.2 Document required environment variables (GEMINI_API_KEY)
- [x] 9.3 Add example output to documentation
- [x] 9.4 Document build and run commands
