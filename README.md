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

The application uses Docker Compose with environment-specific configurations:
- **Local development**: Uses [docker-compose.yml](docker-compose.yml) with exposed ports
- **Production**: Combines base + [docker-compose.prod.yml](docker-compose.prod.yml) for Traefik integration

### Local/Development Deployment

```bash
docker-compose up -d --build
```

This runs with ports exposed (backend: 8080, frontend: 5173) without requiring Traefik.

### Production Deployment with Traefik

The application integrates with Traefik reverse proxy for SSL termination and routing.

**Prerequisites:**
- Linux VM with Docker and Docker Compose installed
- Traefik reverse proxy running with network named `proxy`
- Domain name pointing to your VM
- SSH access to the VM

**Required GitHub Secrets:**

Configure the following secrets in your GitHub repository settings (Settings → Secrets and variables → Actions):

| Secret | Description | Example |
|--------|-------------|---------|
| `VM_IP` | IP address or hostname of your VM | `192.168.1.100` or `vm.example.com` |
| `VM_USER` | SSH username for deployment | `deploy` or `ubuntu` |
| `SSH_PRIVATE_KEY` | SSH private key for authentication | `-----BEGIN OPENSSH PRIVATE KEY-----...` |
| `GEMINI_API_KEY` | Google Gemini API key | `AIzaSyD...` |
| `DOMAIN` | Production domain name | `rechtbank.example.com` |

**VM Setup:**

1. Clone the repository to deployment directory:
```bash
# On your VM
sudo mkdir -p /opt/rechtbank
sudo chown $USER:$USER /opt/rechtbank
cd /opt/rechtbank
git clone <your-repo-url> .
```

2. Verify Traefik is running with proxy network:
```bash
docker network ls | grep proxy
```

3. Ensure SSH key authentication is enabled for the deployment user

**Deployment Process:**

The application automatically deploys when code is pushed to the `main` branch:

1. Push to main branch triggers GitHub Actions workflow
2. Workflow connects to VM via SSH
3. Pulls latest code from repository
4. Creates `.env` file with production secrets
5. Rebuilds and restarts containers with `docker compose up -d --build`

**Manual Deployment:**

If needed, you can deploy manually on the VM:

```bash
cd /opt/rechtbank
git pull origin main

# Create .env file with required variables
cat > .env << EOF
GEMINI_API_KEY=your-key-here
ENV=production
CORS_ORIGIN=https://rechtbank.example.com
PUBLIC_API_URL=https://rechtbank.example.com/api
DOMAIN=rechtbank.example.com
EOF

docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build
```

**Rollback Procedure:**

To rollback to a previous version:

```bash
# On your VM
cd /opt/rechtbank
git log --oneline  # Find the commit hash to rollback to
git reset --hard <commit-hash>
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build
```

Or push a revert commit to trigger automatic deployment:

```bash
# On your local machine
git revert <bad-commit-hash>
git push origin main
```

**Verification:**

After deployment, verify:
- Frontend accessible at `https://your-domain.com`
- HTTPS certificate is valid (Let's Encrypt)
- API endpoints work correctly
- Backend is not directly accessible from external network

### Health Checks

The backend includes health checks for container orchestration:
- Endpoint: `GET /health`
- Interval: 30s
- Timeout: 3s

## License

This is a personal project for entertainment purposes.

## Credits

Built with Google Gemini AI for image analysis and comedic verdict generation.
