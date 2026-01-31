# Rechtbank voor Meubilair - Frontend

Interactive web application for the "Furniture Court" - submit photos of your furniture for comedic legal judgment.

## Features

- 📸 **Photo Capture**: Native camera integration via file input
- 🔄 **Manual Rotation**: Rotate photos before submission with visual preview
- 📤 **Photo Upload**: Submit furniture photos for AI-powered judgment
- 🏛️ **Verdict Display**: Receive comedic legal verdicts with Dutch legal styling
- 📱 **Mobile-First**: Optimized for smartphone use

## Technology Stack

- **Framework**: SvelteKit 2 with TypeScript
- **Build Tool**: Vite
- **Testing**: Vitest + Testing Library
- **Architecture**: Hexagonal (Ports & Adapters)
- **Deployment**: Docker (multi-stage) + Nginx Alpine
- **Styling**: Component-scoped CSS

## Prerequisites

- Node.js 20+
- npm or yarn
- Docker (for containerized deployment)

## Development Setup

### 1. Install Dependencies

```bash
npm install
```

### 2. Environment Configuration

Copy the example environment file:

```bash
cp .env.example .env
```

Edit `.env` to configure the API endpoint:

```env
PUBLIC_API_URL=http://localhost:8080
```

### 3. Run Development Server

```bash
npm run dev
```

The app will be available at `http://localhost:5173`

**Note**: Camera access via file input works on both HTTP and HTTPS. For local development, `localhost` is sufficient.

### 4. Run Tests

```bash
# Run tests in watch mode
npm run test

# Run tests with UI
npm run test:ui
```

### 5. Lint and Format

```bash
# Lint code
npm run lint

# Format code
npm run format
```

## Production Build

### Build Static Files

```bash
npm run build
```

Output will be in `build/` directory.

### Preview Production Build

```bash
npm run preview
```

## Docker Deployment

### Build Docker Image

```bash
docker build -t rechtbank-frontend .
```

### Run Container

```bash
docker run -p 8080:80 rechtbank-frontend
```

The application will be available at `http://localhost:8080`

### Environment Variables in Docker

Pass environment variables at runtime:

```bash
docker run -p 8080:80 -e PUBLIC_API_URL=https://api.example.com rechtbank-frontend
```

## Project Structure

```
frontend/
├── src/
│   ├── lib/
│   │   ├── adapters/
│   │   │   ├── api/           # Backend API communication
│   │   │   └── ports/         # Port interfaces for adapters
│   │   ├── features/          # UI components
│   │   └── shared/
│   │       ├── stores/        # Svelte stores for state management
│   │       ├── types/         # TypeScript type definitions
│   │       └── utils/         # Utility functions (rotation, etc.)
│   ├── routes/
│   │   └── +page.svelte       # Main application page
│   └── test-setup.ts          # Vitest configuration
├── Dockerfile                 # Multi-stage Docker build
├── nginx.conf                 # Nginx configuration for SPA
├── svelte.config.js           # SvelteKit configuration
├── vite.config.ts             # Vite configuration
└── package.json
```

## Architecture

The application follows **Hexagonal Architecture** (Ports & Adapters) for external integrations:

### Core Domain
- Application state machine (Svelte stores)
- Photo capture and rotation logic

### Adapters
- **ApiAdapter**: HTTP client for backend communication

### Ports (Interfaces)
- **IApiPort**: Backend API interface

### Benefits
- Testability: Adapters can be mocked in tests
- Flexibility: Easy to swap implementations
- Separation: External APIs isolated from business logic

## API Contract

### POST /v1/judge

Upload a furniture photo for judgment.

**Request**:
- Method: `POST`
- Content-Type: `multipart/form-data`
- Fields:
  - `photo` (file): Image file (JPEG, PNG, WEBP, GIF)
  - `userAgent` (string): Browser user agent
  - `timestamp` (string): ISO 8601 timestamp
  - `captureMethod` (string): `"camera"` or `"file"`
  - `rotation` (number): Applied rotation in degrees (0, 90, 180, 270)

**Response**:
```json
{
  "type": "guilty" | "acquittal" | "niet-ontvankelijk",
  "score": 5,
  "verdictText": "Na grondige bestudering...",
  "sentence": "Veroordeling tot plaatsing tegen de muur",
  "angleDeviation": 3.2,
  "isFurniture": true
}
```

## Browser Compatibility

- **Chrome/Edge**: ✅ Full support
- **Firefox**: ✅ Full support
- **Safari**: ✅ Full support (iOS 13+)
- **Mobile browsers**: ✅ Optimized for iOS Safari and Chrome Android

### Required Browser APIs
- File API (photo capture via file input)
- Canvas API (rotation processing)
- Fetch API (upload)

## Troubleshooting

### Photo capture not working

1. **Check permissions**: Ensure browser has camera permission
2. **Check device**: Verify device has a camera
3. **Try different browser**: Some older browsers may not support file input camera

### Photo rotation issues

1. **Manual controls**: Use rotate left/right buttons to correct orientation
2. **Preview**: Check photo preview before submitting

### Upload failing

1. **Check API URL**: Verify `PUBLIC_API_URL` in `.env`
2. **Check network**: Ensure backend is running and accessible
3. **Check file size**: Photos must be under 10MB
4. **Check CORS**: Backend must allow requests from frontend origin

## License

© 2026 Rechtbank voor Meubilair | Satirical Project | No Real Legal Judgments

