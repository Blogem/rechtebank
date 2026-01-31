# Rechtebank Backend API

Go-based HTTP API service for the Rechtebank furniture court application. Receives photo uploads, analyzes them using Google Gemini AI, and returns humorous legal verdicts in Dutch.

## Quick Start

```bash
# Copy environment file and add your API key
cp .env.example .env
# Edit .env and set GEMINI_API_KEY

# Run the server
go run ./cmd/server

# Or with Docker
docker build -t rechtebank-backend .
docker run -p 8080:8080 --env-file .env rechtebank-backend
```

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `GEMINI_API_KEY` | Yes | - | Google Gemini API key |
| `PORT` | No | `8080` | HTTP server port |
| `CORS_ORIGIN` | No | `*` | Allowed CORS origin (e.g., `http://localhost:5173`) |
| `GEMINI_TIMEOUT` | No | `30` | Gemini API timeout in seconds |
| `MAX_FILE_SIZE` | No | `10485760` | Max upload size in bytes (default 10MB) |
| `ENV` | No | `development` | Environment (`development` or `production`) |

## API Endpoint

### POST /v1/judge

Analyze a furniture photo and receive a legal verdict.

**Request:**
- Content-Type: `multipart/form-data`
- Body: Form field `photo` containing image file (JPEG, PNG, or WebP)
- Max file size: 10MB

**Example using curl:**
```bash
curl -X POST http://localhost:8080/v1/judge \
  -F "photo=@/path/to/furniture.jpg"
```

**Success Response (200 OK):**
```json
{
  "admissible": true,
  "score": 8,
  "verdict": {
    "crime": "Lichte rugleuning-afwijking van 3 graden",
    "sentence": "Vrijgesproken met een waarschuwing",
    "reasoning": "Hoewel artikel 12 van de Meubilair-wet strikte rechtheid vereist, acht de rechtbank deze minimale afwijking verschoonbaar gezien de verder onberispelijke staat van het meubelstuk.",
    "observation": "Een houten zetel met groene ribstof, lichte afwijking van verticale norm geconstateerd.",
    "verdictType": "vrijspraak"
  },
  "requestId": "550e8400-e29b-41d4-a716-446655440000",
  "timestamp": "2026-01-31T10:30:00Z"
}
```

**Verdict Types:**
The `verdictType` field indicates the court's decision:
- `"vrijspraak"` - Acquittal (typically scores 8-10, or exceptional alignment)
- `"waarschuwing"` - Warning (typically scores 6-7, or minor violations)
- `"schuldig"` - Guilty (typically scores 1-5, or serious violations)

Note: The LLM determines the verdict type based on its analysis and may consider context beyond just the numeric score.

**Non-Furniture Response (200 OK):**
```json
{
  "admissible": false,
  "score": 0,
  "verdict": {
    "crime": "Geen meubilair gedetecteerd",
    "sentence": "Zaak niet-ontvankelijk",
    "reasoning": "Alleen meubilair kan worden berecht door de Meubilair-rechtbank."
  },
  "requestId": "550e8400-e29b-41d4-a716-446655440001",
  "timestamp": "2026-01-31T10:31:00Z"
}
```

**Error Responses:**

| Status | Description | Example |
|--------|-------------|---------|
| 400 | Missing file or invalid format | `{"error": "Photo file is required"}` |
| 400 | Unsupported image format | `{"error": "Unsupported image format. Use JPEG, PNG, or WebP"}` |
| 413 | File too large | `{"error": "Photo file size must not exceed 10MB"}` |
| 429 | Rate limited | `{"error": "rate limit exceeded"}` (includes `Retry-After` header) |
| 500 | Internal server error | `{"error": "AI analysis service unavailable"}` |
| 502 | Gemini API error | `{"error": "AI analysis failed"}` |
| 503 | Service unavailable | `{"error": "AI analysis service temporarily unavailable"}` |
| 504 | Timeout | `{"error": "AI analysis timeout"}` |

### GET /health

Health check endpoint for container orchestration.

**Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2026-01-31T10:30:00Z"
}
```

## Local Development

### Prerequisites
- Go 1.24+ (or 1.23 with GOTOOLCHAIN=auto)
- Google Gemini API key

### Setup

1. Clone the repository and navigate to backend:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create environment file:
   ```bash
   cp .env.example .env
   # Edit .env and add your GEMINI_API_KEY
   ```

4. Run the server:
   ```bash
   go run ./cmd/server
   ```

5. Run tests:
   ```bash
   go test ./...
   ```

### Project Structure

```
backend/
├── cmd/server/           # Application entry point
│   └── main.go
├── internal/
│   ├── adapters/         # External interfaces
│   │   ├── gemini/       # Gemini AI adapter
│   │   ├── http/         # HTTP handlers and router
│   │   └── validator/    # Photo validation
│   ├── config/           # Configuration loading
│   └── core/             # Business logic
│       ├── domain/       # Domain entities
│       ├── ports/        # Interface definitions
│       └── services/     # Business services
├── Dockerfile
├── go.mod
└── go.sum
```

## Docker Deployment

### Build and run standalone:
```bash
docker build -t rechtebank-backend .
docker run -p 8080:8080 \
  -e GEMINI_API_KEY=your-api-key \
  -e CORS_ORIGIN=http://localhost:5173 \
  rechtebank-backend
```

### With Docker Compose (from project root):
```bash
# Set environment variables
export GEMINI_API_KEY=your-api-key

# Start all services
docker-compose up -d

# View logs
docker-compose logs -f backend
```

## Troubleshooting

### Common Errors

**"GEMINI_API_KEY environment variable is required"**
- Ensure you've set the `GEMINI_API_KEY` environment variable
- Check your `.env` file exists and contains the key

**"AI analysis service unavailable" (500)**
- The Gemini API key may be invalid
- Check your API key in Google Cloud Console

**"AI analysis failed" (502)**
- Gemini API returned an error
- Check the server logs for details
- The image might not be processable

**"AI analysis service temporarily unavailable" (503)**
- Rate limit exceeded after retries
- Wait and try again, or check your API quota

**"AI analysis timeout" (504)**
- Gemini API took too long to respond
- Try with a smaller image or retry later

**"Unsupported image format"**
- Only JPEG, PNG, and WebP are supported
- Check the file is a valid image (not corrupted)

**"Photo file size must not exceed 10MB"**
- Compress or resize the image before uploading

### Debug Mode

Set `ENV=development` to enable verbose logging:
```bash
ENV=development go run ./cmd/server
```

### Health Check

Verify the server is running:
```bash
curl http://localhost:8080/health
```
