## Why

The frontend is complete but needs a backend service to analyze furniture photos and return verdicts. The backend must receive photo uploads from the frontend, integrate with Google's Gemini 2.5 Flash Lite API for multimodal image analysis, and return humorous legal judgments about furniture alignment.

## What Changes

- New Go-based HTTP API service with RESTful endpoints
- Docker container for the Go backend service
- Integration with Google Gemini 2.5 Flash Lite API for image analysis
- Multipart form-data handling for photo uploads
- JSON response formatting for verdicts (score, legal text, admissibility)
- Docker Compose integration for orchestration with frontend
- Environment variable configuration for API keys and service settings

## Capabilities

### New Capabilities
- `photo-upload-endpoint`: REST API endpoint that accepts multipart/form-data photo uploads
- `gemini-integration`: Integration with Google Gemini API for multimodal image analysis with custom legal system prompt
- `verdict-response`: Structured JSON response containing furniture alignment score (1-10), legal verdict text, and case admissibility

### Modified Capabilities

## Impact

- **New Service**: Go backend service in new `/backend` directory
- **Docker Compose**: Updated to include backend container alongside frontend
- **Frontend Integration**: Frontend API adapter will connect to new backend endpoints
- **Dependencies**: Google Generative AI Go SDK, Docker, environment configuration for API keys
- **Infrastructure**: Requires Gemini API key configuration and network bridge between frontend/backend containers
