## Why

Debugging the Gemini API integration requires visibility into the exact prompts sent and responses received. Currently, the `AnalyzePhoto` function in the analyzer obscures this interaction, making it difficult to troubleshoot issues with prompt engineering, response parsing, or API behavior.

## What Changes

- Add a CLI debug tool for testing Gemini interaction outside the main application
- Implement verbose logging to show system prompts, user prompts, and raw API responses
- Support loading test images from disk to replay scenarios
- Display structured JSON output alongside raw responses for comparison

## Capabilities

### New Capabilities
- `cli-debug-tool`: Command-line interface that accepts an image path and calls the Gemini API with full request/response logging
- `verbose-logging`: Detailed logging that shows the system prompt, image metadata, request payload, and raw JSON response from Gemini

### Modified Capabilities

(No existing capabilities are being modified)

## Impact

- **New files**: `backend/cmd/debug-gemini/main.go` (new CLI entry point)
- **Modified files**: `backend/internal/adapters/gemini/analyzer.go` (potential refactoring to expose internal details for debugging)
- **Dependencies**: No new dependencies; uses existing `genai` SDK and `analyzer` package
- **Testing**: Helps validate the Gemini integration and prompt effectiveness
