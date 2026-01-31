# Rechtbank voor Meubilair (Furniture Court)

A comedic web application that puts your furniture on trial. Upload a photo of your furniture and receive an official legal verdict in Dutch from the Honorable Judge of the Furniture Court.

## Overview

The Furniture Court ("Rechtbank voor Meubilair") analyzes furniture photos using Google's Gemini AI to deliver humorous legal judgments. If your furniture doesn't meet the strict standards of the Furniture Law, expect a sentence ranging from "re-orientation therapy" to "immediate scrapping."

**Verdict System:**
The court delivers verdicts in three categories:
- **Vrijspraak** (Acquittal): For well-aligned furniture (typically scores 8-10)
- **Waarschuwing** (Warning): For borderline cases with minor violations (typically scores 6-7)
- **Schuldig** (Guilty): For serious alignment violations (typically scores 1-5)

**Example Verdict:**
- **Verdict Type**: Waarschuwing
- **Score**: 7/10
- **Crime**: Slight backrest deviation of 3 degrees
- **Sentence**: Warning with mandatory observation
- **Reasoning**: While Article 12 of the Furniture Act requires strict straightness, the court considers this minimal deviation excusable given the otherwise impeccable condition of the piece.

## Architecture

This is a full-stack application built with:

- **Frontend**: SvelteKit + TypeScript (mobile-first web app)
- **Backend**: Go HTTP API with hexagonal architecture
- **AI**: Google Gemini 2.5 Flash Lite for image analysis
- **Deployment**: Docker Compose with multi-stage builds

```
┌─────────────┐      ┌──────────────┐      ┌──────────────┐
│   Browser   │──────│   Frontend   │──────│   Backend    │
│  (Camera)   │      │  (SvelteKit) │      │  (Go API)    │
└─────────────┘      └──────────────┘      └──────────────┘
                                                   │
                                                   ▼
                                           ┌──────────────┐
                                           │  Gemini AI   │
                                           │  (Analysis)  │
                                           └──────────────┘
```

## Project Structure

```
.
├── backend/                 # Go API server
│   ├── cmd/
│   │   ├── server/         # Main HTTP server
│   │   └── debug-gemini/   # CLI debugging tool
│   ├── internal/
│   │   ├── adapters/       # External integrations (Gemini, HTTP, storage)
│   │   ├── config/         # Configuration management
│   │   └── core/           # Business logic & domain models
│   └── README.md           # Backend documentation
│
├── frontend/               # SvelteKit web application
│   ├── src/
│   │   ├── lib/
│   │   │   ├── features/   # UI components (camera, upload, verdict display)
│   │   │   └── adapters/   # API adapters
│   │   └── routes/         # SvelteKit pages
│   └── README.md           # Frontend documentation
│
├── openspec/               # OpenSpec framework for structured development
│   └── changes/            # Feature proposals and designs
│
├── docs/                   # Project documentation
├── docker-compose.yml      # Production deployment configuration
└── .env.example           # Environment variables template
```

## Quick Start

### Prerequisites

- Docker & Docker Compose
- Google Gemini API key ([get one here](https://ai.google.dev/))

### 1. Environment Setup

```bash
# Copy environment template
cp .env.example .env

# Edit .env and add your Gemini API key
# GEMINI_API_KEY=your-api-key-here
```

### 2. Run with Docker Compose

```bash
docker-compose up --build
```

This will start:
- **Backend**: http://localhost:8080
- **Frontend**: http://localhost:5173

### 3. Use the Application

1. Open http://localhost:5173 in your browser
2. Allow camera access (or choose a file)
3. Take a photo of furniture (or upload one)
4. Optionally rotate the photo
5. Submit and receive your verdict!

## Development

### Backend Development

See [backend/README.md](backend/README.md) for detailed instructions.

```bash
cd backend
cp .env.example .env
# Add your GEMINI_API_KEY to .env
go run ./cmd/server
```

**Debug Tool:**
```bash
cd backend
go build -o debug-gemini ./cmd/debug-gemini
export GEMINI_API_KEY="your-key"
./debug-gemini path/to/furniture.jpg
```

### Frontend Development

See [frontend/README.md](frontend/README.md) for detailed instructions.

```bash
cd frontend
npm install
npm run dev
```

**Run Tests:**
```bash
npm run test        # Watch mode
npm run test:ui     # Interactive UI
```

## API Documentation

### POST /v1/judge

Submit a furniture photo for judgment.

**Request:**
- Content-Type: `multipart/form-data`
- Field: `photo` (JPEG, PNG, or WebP, max 10MB)

**Response:**
```json
{
  "admissible": true,
  "score": 8,
  "verdict": {
    "crime": "Lichte rugleuning-afwijking van 3 graden",
    "sentence": "Vrijgesproken met een waarschuwing",
    "reasoning": "Hoewel artikel 12..."
  },
  "requestId": "550e8400-e29b-41d4-a716-446655440000",
  "timestamp": "2026-01-31T10:30:00Z"
}
```

### GET /health

Health check endpoint.

## Environment Variables

### Backend
| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `GEMINI_API_KEY` | Yes | - | Google Gemini API key |
| `PORT` | No | `8080` | HTTP server port |
| `CORS_ORIGIN` | No | `*` | Allowed CORS origin |
| `GEMINI_TIMEOUT` | No | `30` | Gemini API timeout (seconds) |
| `MAX_FILE_SIZE` | No | `10485760` | Max upload size (bytes) |

### Frontend
| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `PUBLIC_API_URL` | Yes | - | Backend API URL |

## Testing

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Tests
```bash
cd frontend
npm run test
```

## Deployment

The application is designed for containerized deployment using Docker Compose. Both frontend and backend use multi-stage Docker builds for optimized production images.

### Production Build
```bash
docker-compose up -d --build
```

### Health Checks
The backend includes health checks for container orchestration:
- Endpoint: `GET /health`
- Interval: 30s
- Timeout: 3s

## License

This is a personal project for entertainment purposes.

## Credits

Built with Google Gemini AI for image analysis and comedic verdict generation.
