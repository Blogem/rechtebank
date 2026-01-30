## 1. Project Setup

- [ ] 1.1 Initialize Go module in `/backend` directory with `go mod init`
- [ ] 1.2 Create hexagonal architecture folder structure (`cmd/server`, `internal/core`, `internal/adapters`)
- [ ] 1.3 Add Gin dependency (`github.com/gin-gonic/gin`)
- [ ] 1.4 Add Google Generative AI Go SDK dependency
- [ ] 1.5 Add testify for testing utilities (`github.com/stretchr/testify`)
- [ ] 1.6 Create `.env.example` file with required environment variables
- [ ] 1.7 Create `.gitignore` for Go projects

## 2. Core Domain Layer

- [ ] 2.1 Write unit tests for `VerdictResponse` struct (JSON marshaling/unmarshaling)
- [ ] 2.2 Write unit tests for `VerdictDetails` struct with all three fields
- [ ] 2.3 Write unit tests for `PhotoMetadata` struct
- [ ] 2.4 Run tests to verify they fail (structs don't exist yet)
- [ ] 2.5 Define `VerdictResponse` struct in `internal/core/domain/verdict.go`
- [ ] 2.6 Define `VerdictDetails` struct with crime, sentence, reasoning fields
- [ ] 2.7 Define `PhotoMetadata` struct for upload metadata
- [ ] 2.8 Run tests to verify they pass
- [ ] 2.9 Create `IPhotoAnalyzer` port interface in `internal/core/ports/analyzer.go`
- [ ] 2.10 Create `IPhotoValidator` port interface in `internal/core/ports/validator.go`

## 3. Photo Validator

- [ ] 3.1 Write unit tests for file format validation (JPEG, PNG, WebP detection)
- [ ] 3.2 Write unit tests for unsupported format rejection
- [ ] 3.3 Write unit tests for file size validation (10MB limit)
- [ ] 3.4 Write unit tests for size limit exceeded scenarios
- [ ] 3.5 Run tests to verify they fail (validator doesn't exist yet)
- [ ] 3.6 Implement `PhotoValidator` in `internal/adapters/validator/validator.go`
- [ ] 3.7 Add file format validation logic
- [ ] 3.8 Add file size validation logic
- [ ] 3.9 Run tests to verify they pass

## 4. Gemini Adapter

- [ ] 4.1 Write unit tests for Gemini client initialization (with/without API key)
- [ ] 4.2 Write unit tests for successful photo analysis with mocked Gemini client
- [ ] 4.3 Write unit tests for non-furniture detection scenario
- [ ] 4.4 Write unit tests for exponential backoff retry logic (429 errors)
- [ ] 4.5 Write unit tests for timeout scenarios (30s)
- [ ] 4.6 Write unit tests for all Gemini API error scenarios
- [ ] 4.7 Run tests to verify they fail (analyzer doesn't exist yet)
- [ ] 4.8 Create `GeminiAnalyzer` struct in `internal/adapters/gemini/analyzer.go`
- [ ] 4.9 Implement Gemini client initialization with API key validation
- [ ] 4.10 Define JSON schema for structured output (admissible, score, verdict components)
- [ ] 4.11 Create Dutch legal system prompt with structured field instructions
- [ ] 4.12 Implement `AnalyzePhoto` method with Gemini API call
- [ ] 4.13 Implement response parsing from JSON schema output
- [ ] 4.14 Implement exponential backoff retry logic for 429 errors (max 3 retries)
- [ ] 4.15 Implement 30-second timeout for Gemini API requests
- [ ] 4.16 Add error handling for all Gemini API error scenarios
- [ ] 4.17 Run unit tests to verify they pass
- [ ] 4.18 Write integration tests with real Gemini API (conditional on API key)
- [ ] 4.19 Run integration tests to verify real API behavior

## 5. Verdict Service

- [ ] 5.1 Write unit tests for `JudgePhoto` method with mocked analyzer and validator
- [ ] 5.2 Write unit tests for validation error propagation
- [ ] 5.3 Write unit tests for analysis error propagation
- [ ] 5.4 Write unit tests for request ID generation (UUID format)
- [ ] 5.5 Write unit tests for timestamp generation (ISO 8601 format)
- [ ] 5.6 Run tests to verify they fail (service doesn't exist yet)
- [ ] 5.7 Create `VerdictService` in `internal/core/services/verdict_service.go`
- [ ] 5.8 Implement constructor with dependency injection (analyzer, validator)
- [ ] 5.9 Implement `JudgePhoto` method orchestrating validation and analysis
- [ ] 5.10 Add request ID generation (UUID)
- [ ] 5.11 Add timestamp generation (ISO 8601 format)
- [ ] 5.12 Run tests to verify they pass

## 6. HTTP Adapter

- [ ] 6.1 Write HTTP handler tests for successful photo upload with mocked service
- [ ] 6.2 Write tests for multipart form-data parsing
- [ ] 6.3 Write tests for missing file scenario (400)
- [ ] 6.4 Write tests for invalid content type (400)
- [ ] 6.5 Write tests for file too large scenario (413)
- [ ] 6.6 Write tests for rate limit scenario (429 with Retry-After header)
- [ ] 6.7 Write tests for internal errors (500, 502, 503, 504)
- [ ] 6.8 Write tests for CORS headers in responses
- [ ] 6.9 Run tests to verify they fail (handlers don't exist yet)
- [ ] 6.10 Create Gin router setup in `internal/adapters/http/router.go`
- [ ] 6.11 Configure CORS middleware with environment-based origin
- [ ] 6.12 Configure logging middleware
- [ ] 6.13 Configure recovery middleware for panic handling
- [ ] 6.14 Create `JudgeHandler` in `internal/adapters/http/handlers/judge.go`
- [ ] 6.15 Implement multipart form-data parsing
- [ ] 6.16 Implement HTTP error mapping (400, 413, 429, 500, 502, 503, 504)
- [ ] 6.17 Add Retry-After header for 429 responses
- [ ] 6.18 Create request/response DTOs if needed
- [ ] 6.19 Run tests to verify they pass

## 7. Configuration & Main

- [ ] 7.1 Create `Config` struct in `internal/config/config.go`
- [ ] 7.2 Implement environment variable loading (GEMINI_API_KEY, PORT, etc.)
- [ ] 7.3 Add config validation (required fields, defaults)
- [ ] 7.4 Create `main.go` in `cmd/server/main.go`
- [ ] 7.5 Wire dependencies (validator → analyzer → service → handler → router)
- [ ] 7.6 Implement graceful shutdown
- [ ] 7.7 Add startup logging (port, configuration summary)

## 8. Docker Integration

- [ ] 8.1 Create multi-stage Dockerfile (golang:alpine for build, alpine for runtime)
- [ ] 8.2 Configure WORKDIR and dependency caching
- [ ] 8.3 Add health check endpoint (`/health`) in router
- [ ] 8.4 Update root `docker-compose.yml` to include backend service
- [ ] 8.5 Configure backend service (build context, ports, environment)
- [ ] 8.6 Configure network bridge between frontend and backend containers
- [ ] 8.7 Add backend service to depends_on for frontend (if needed)
- [ ] 8.8 Test local Docker build
- [ ] 8.9 Test local Docker Compose deployment

## 9. Testing & Validation

- [ ] 9.1 Run all unit tests (`go test ./...`)
- [ ] 9.2 Run integration tests with Gemini API key
- [ ] 9.3 Test end-to-end with frontend in Docker Compose
- [ ] 9.4 Test CORS from frontend origin
- [ ] 9.5 Test all error scenarios (missing file, wrong format, too large)
- [ ] 9.6 Test rate limiting behavior (if possible)
- [ ] 9.7 Validate JSON response structure matches spec
- [ ] 9.8 Validate Dutch legal jargon quality in verdicts

## 10. Documentation

- [ ] 10.1 Create `backend/README.md` with setup instructions
- [ ] 10.2 Document environment variables and defaults
- [ ] 10.3 Document API endpoint (`POST /v1/judge`)
- [ ] 10.4 Add example request/response
- [ ] 10.5 Document local development setup
- [ ] 10.6 Document Docker deployment
- [ ] 10.7 Add troubleshooting section (common errors)
