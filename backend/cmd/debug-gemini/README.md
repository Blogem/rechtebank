# Gemini Debug Tool

A standalone CLI tool for debugging Gemini API interactions. This tool provides complete visibility into the request/response cycle when analyzing furniture images with the Gemini AI model.

## Purpose

The debug tool helps developers:
- Inspect the exact system prompt sent to Gemini
- See the user prompt and image metadata
- View raw JSON responses from the Gemini API
- Compare parsed verdicts with raw responses
- Troubleshoot API errors, timeouts, and response parsing issues
- Test prompt engineering changes quickly

## Requirements

- Go 1.21 or later
- Valid `GEMINI_API_KEY` environment variable
- A furniture image file (JPEG, PNG, or WebP format)

## Building

From the `backend/` directory:

```bash
go build -o debug-gemini ./cmd/debug-gemini
```

This creates a `debug-gemini` executable in the current directory.

## Usage

```bash
export GEMINI_API_KEY="your-api-key-here"
./debug-gemini path/to/furniture-image.jpg
```

### Example

```bash
export GEMINI_API_KEY="AIza..."
./debug-gemini test-images/chair.jpg
```

## Output Format

The tool displays the following sections:

### 1. Header with Image Metadata
```
=== GEMINI DEBUG TOOL ===
Image: test-images/chair.jpg
Size: 245.3 KB
MIME Type: jpeg
Dimensions: 1920x1080
```

### 2. System Prompt
```
=== SYSTEM PROMPT ===
Je bent de Eerwaarde Rechter van de Meubilair-rechtbank...
[full system prompt text]
```

### 3. User Prompt
```
=== USER PROMPT ===
Analyseer dit meubelstuk en spreek je vonnis uit.
```

### 4. Request Configuration
```
=== REQUEST METADATA ===
Model: gemini-2.5-flash-lite
Timeout: 30s
Max Retries: 3
```

### 5. API Call Status
```
=== CALLING GEMINI API ===
Sending request...
```

### 6. Raw JSON Response
```
=== RAW RESPONSE ===
{
  "admissible": true,
  "score": 7,
  "crime": "Lichte scheefstand van 3 graden",
  "sentence": "Veroordeeld tot heroriëntatie",
  "reasoning": "Artikel 42 van de Meubilair-wet...",
  "observation": ""
}
```

### 7. Parsed Verdict
```
=== PARSED VERDICT ===
Admissible: true
Score: 7/10
Crime: Lichte scheefstand van 3 graden
Sentence: Veroordeeld tot heroriëntatie
Reasoning: Artikel 42 van de Meubilair-wet...
```

## Error Handling

The tool exits with code `0` on success and `1` on failure.

### Common Errors

**Missing file:**
```
Error: image file not found: /path/to/image.jpg
```

**Invalid image format:**
```
Error: unsupported image format (must be JPEG, PNG, or WebP)
```

**Missing API key:**
```
Error: GEMINI_API_KEY environment variable is required
```

**API timeout:**
```
Error: AI analysis timeout
```

**API errors** are displayed with full context:
```
=== ERROR ===
API call failed: <error details>

=== RAW RESPONSE (ERROR) ===
<error response if available>
```

## Supported Image Formats

- **JPEG** (`.jpg`, `.jpeg`)
- **PNG** (`.png`)
- **WebP** (`.webp`)

Images are detected by magic bytes, not file extension.

## Tips

### Quick Iteration
Use the tool to rapidly test prompt changes:
1. Modify the system prompt in `backend/internal/adapters/gemini/analyzer.go`
2. Rebuild: `go build ./cmd/debug-gemini`
3. Test: `./debug-gemini test-image.jpg`

### Comparing with Production
The debug tool uses the same `GeminiAnalyzer` code as the production API, ensuring identical behavior.

### Testing Different Scenarios
Create a collection of test images representing different cases:
- Perfectly straight furniture (expected score: 9-10)
- Slightly tilted furniture (expected score: 5-8)
- Very crooked furniture (expected score: 1-4)
- Non-furniture items (expected: `admissible: false`)

### Saving Output
Redirect output to a file for analysis:
```bash
./debug-gemini chair.jpg > debug-output.txt 2>&1
```

## Development

### Running Tests
```bash
go test ./cmd/debug-gemini -v
```

### Test Coverage
```bash
go test ./cmd/debug-gemini -cover
```

Tests cover:
- MIME type detection (JPEG, PNG, WebP, invalid)
- File size formatting
- Section header formatting
- Missing file handling
- Invalid image handling
- Missing API key handling
- Command-line argument validation

## Architecture

The tool is designed to reuse production code:
- Imports `GeminiAnalyzer` from `backend/internal/adapters/gemini`
- Uses the same system prompt via `GetSystemPrompt()` function
- Calls the same `AnalyzePhoto()` method as the HTTP API
- No code duplication or test-only implementations

## Limitations

- Cannot capture the truly raw response from Gemini (reconstructs JSON from parsed response)
- Does not support batch processing (one image at a time)
- No retry visualization (errors are displayed immediately)
- Requires rebuilding after prompt changes

## Security Note

⚠️ **Never commit tool output containing API keys or sensitive information to version control.**

The tool may expose:
- API keys in error messages
- Image file paths
- Full API responses

Use caution when sharing debug output.
