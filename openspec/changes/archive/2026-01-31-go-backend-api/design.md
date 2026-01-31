## Context

The rechtbank.org project is an interactive joke website where users photograph furniture (particularly couches) to receive humorous legal judgments about their alignment. The frontend (Svelte) is complete with camera access, spirit level overlay, and photo upload capabilities. The backend service needs to receive photos, analyze them using AI, and return structured verdicts in Dutch legal jargon.

**Constraints:**
- Must run in Docker containers for easy deployment
- Requires HTTPS for camera API (handled by reverse proxy)
- Budget-conscious (using free tier of Gemini API)
- Target deployment: lightweight VM

**Current State:**
- Frontend: Complete (Svelte/TypeScript with hexagonal architecture)
- Backend: Does not exist
- Infrastructure: Docker Compose planned with Nginx/Traefik reverse proxy

## Goals / Non-Goals

**Goals:**
- Build a Go-based HTTP API service that receives photo uploads and returns structured verdicts
- Integrate with Google Gemini 2.5 Flash Lite API for multimodal image analysis
- Return type-safe, structured JSON responses with guaranteed format
- Handle errors gracefully (rate limits, timeouts, invalid uploads)
- Follow hexagonal architecture principles to match frontend patterns
- Deploy as a Docker container alongside frontend

**Non-Goals:**
- Client-side image processing or compression (frontend responsibility)
- User authentication or session management (public, stateless service)
- Photo storage or history (ephemeral analysis only)
- Custom ML model training (using Gemini API as-is)
- Real-time streaming or WebSocket connections

## Decisions

### 1. Use Gemini Structured Output API (JSON Schema)

**Decision:** Use Gemini's native JSON Schema support instead of prompt engineering for JSON format.

**Rationale:**
- Gemini 2.5 Flash Lite supports structured outputs with JSON Schema
- Guarantees valid JSON response format at API boundary
- Type-safe parsing into Go structs without manual validation
- Eliminates parsing errors from malformed JSON
- Clear contract between backend and Gemini API

**Alternative Considered:** Prompt engineering ("respond in JSON format")
- Rejected: Not guaranteed, requires fallback parsing logic, brittle

**Implementation:**
```go
type VerdictResponse struct {
    Admissible bool            `json:"admissible"`
    Score      int             `json:"score"`
    Verdict    VerdictDetails  `json:"verdict"`
    RequestID  string          `json:"requestId"`
    Timestamp  string          `json:"timestamp"`
}

type VerdictDetails struct {
    Crime     string `json:"crime"`      // What offense was committed
    Sentence  string `json:"sentence"`   // The punishment
    Reasoning string `json:"reasoning"`  // Legal justification
}
```

### 2. Structured Verdict Components

**Decision:** Split verdict into three fields (crime, sentence, reasoning) instead of single freeform text.

**Rationale:**
- Frontend can display components separately for better UX (e.g., progressive reveal)
- Easier to test individual aspects of AI output quality
- Allows for styling variations (bold sentence, italic reasoning, etc.)
- Maintains creative freedom while adding structure
- Still allows Gemini to generate humorous Dutch legal jargon in each field

**Alternative Considered:** Single `verdict` string field
- Rejected: Less flexible for frontend presentation, harder to test quality

**System Prompt Update:**
The prompt will instruct Gemini to fill each field:
- `crime`: Describe the furniture offense (e.g., "Rugleuning-afwijking van 5 graden")
- `sentence`: State the punishment (e.g., "Veroordeeld tot de brandstapel")
- `reasoning`: Provide legal reasoning (e.g., "Artikel 42 van de Meubilair-wet verbiedt...")

### 3. Go Standard Library HTTP + Gin Framework

**Decision:** Use Gin framework for HTTP routing and middleware.

**Rationale:**
- Lightweight and fast (suitable for VM deployment)
- Built-in middleware for CORS, logging, recovery
- Clean routing for `/v1/judge` endpoint
- Multipart form parsing well-supported
- Still uses standard `net/http` underneath (not reinventing)

**Alternative Considered:** Standard library only (`http.ServeMux`)
- Rejected: More boilerplate for CORS, logging, error handling

### 4. Hexagonal Architecture (Ports & Adapters)

**Decision:** Apply hexagonal architecture with clear port/adapter boundaries.

**Rationale:**
- Consistency with frontend architecture (team familiarity)
- Testability: Can mock Gemini API for unit tests
- Future flexibility: Could swap AI providers without changing core logic
- Clear separation of concerns

**Structure:**
```
/backend
  /cmd/server         - Entry point, dependency injection
  /internal
    /core
      /domain         - VerdictResponse, PhotoMetadata (entities)
      /ports          - IPhotoAnalyzer, IPhotoValidator (interfaces)
      /services       - VerdictService (business logic)
    /adapters
      /http           - Gin handlers, request/response DTOs
      /gemini         - GeminiAnalyzer (implements IPhotoAnalyzer)
```

### 5. Error Handling Strategy

**Decision:** Map Gemini API errors to appropriate HTTP status codes with user-friendly messages.

**Error Mapping:**
- `400 Bad Request`: Invalid file format, missing file, file too large
- `413 Payload Too Large`: File exceeds 10MB
- `429 Too Many Requests`: Rate limit (with Retry-After header)
- `500 Internal Server Error`: Gemini initialization failed
- `502 Bad Gateway`: Gemini API error response
- `503 Service Unavailable`: Retry exhausted
- `504 Gateway Timeout`: Gemini timeout (30s)

**Rationale:**
- Clear distinction between client errors (4xx) and server errors (5xx)
- Retry-After header guides client behavior during rate limiting
- User-friendly error messages in Dutch for frontend display

### 6. Retry Logic with Exponential Backoff

**Decision:** Implement exponential backoff for Gemini API rate limits (429 errors).

**Configuration:**
- Max retries: 3
- Initial delay: 1s
- Backoff multiplier: 2x
- Max delay: 8s

**Rationale:**
- Gemini free tier has rate limits
- Exponential backoff is industry best practice
- Prevents thundering herd
- 3 retries balance reliability with timeout budget (30s total)

**Alternative Considered:** No retries
- Rejected: Poor UX during rate limiting

### 7. Environment Configuration

**Decision:** Use environment variables for all configuration (no config files).

**Required Variables:**
- `GEMINI_API_KEY`: Google API key (required)
- `PORT`: HTTP server port (default: 8080)
- `CORS_ORIGIN`: Allowed frontend origin (default: *)
- `MAX_FILE_SIZE`: Max upload size in bytes (default: 10MB)
- `GEMINI_TIMEOUT`: API timeout in seconds (default: 30)

**Rationale:**
- 12-factor app principles
- Easy Docker/Docker Compose integration
- No secrets in code or files
- Environment-specific configuration (dev/prod)

### 8. Testing Strategy

**Decision:** Follow TDD with three test levels:

1. **Unit Tests:** Core services with mocked adapters
2. **Integration Tests:** Gemini adapter with real API (skipped in CI if no key)
3. **HTTP Tests:** Gin handlers with mocked services

**Rationale:**
- TDD requirement from project context
- Mocking enables fast, reliable unit tests
- Integration tests validate real Gemini behavior
- HTTP tests ensure correct status codes and response format

## Risks / Trade-offs

### Risk: Gemini API Cost Escalation
**Mitigation:** Monitor usage, set budget alerts in Google Cloud, implement rate limiting at backend level if needed

### Risk: Gemini Response Quality
- Gemini might not follow Dutch legal jargon style consistently
- Gemini might struggle with furniture detection in edge cases
**Mitigation:** Comprehensive system prompt testing, fallback to generic verdicts, monitor feedback

### Risk: Gemini API Availability
- Gemini API downtime impacts entire service
- No local fallback
**Mitigation:** Clear error messages, graceful degradation (return friendly error), SLA monitoring

### Risk: File Upload Abuse
- No authentication means public endpoint
- Could receive spam/large files
**Mitigation:** 10MB file size limit, rate limiting (consider adding IP-based limits post-MVP)

### Trade-off: Structured Verdict vs Creative Freedom
- Structured fields might constrain Gemini's creativity
**Mitigation:** System prompt emphasizes creativity within structure, fields are freeform text

### Trade-off: Docker Container Size
- Go compiles to small binary (~10-20MB)
- But multi-stage Docker build needed to minimize image size
**Mitigation:** Use `golang:alpine` for build, `alpine` for runtime (final image ~30MB)

## Migration Plan

**Phase 1: Development**
1. Initialize Go module in `/backend`
2. Implement core domain types and ports
3. Build Gemini adapter with integration tests
4. Build HTTP adapter with unit tests
5. Wire dependencies in `main.go`

**Phase 2: Dockerization**
1. Create multi-stage Dockerfile
2. Add backend service to `docker-compose.yml`
3. Configure network bridge between frontend/backend
4. Test local Docker deployment

**Phase 3: Deployment**
1. Deploy Docker Compose stack to VM
2. Configure reverse proxy for HTTPS
3. Set environment variables (API key)
4. Smoke test with frontend

**Rollback Strategy:**
- If backend fails: Remove from Docker Compose, frontend shows "service unavailable"
- Git tag for each deployment
- Keep previous Docker image tagged

## Open Questions

None - ready for implementation.
