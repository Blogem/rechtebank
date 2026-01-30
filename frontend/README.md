# Rechtbank voor Meubilair - Frontend

Interactive web application for the "Furniture Court" - submit photos of your furniture for comedic legal judgment.

## Features

- 📸 **Camera Access**: Capture photos directly from your device camera
- ⚖️ **Spirit Level Overlay**: Real-time device orientation guidance
- 📤 **Photo Upload**: Submit furniture photos for AI-powered judgment
- 🏛️ **Verdict Display**: Receive comedic legal verdicts with Dutch legal styling
- ♿ **Accessibility**: Optional bypass for level requirements
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

**Note**: Camera access requires HTTPS or localhost. For local development, `localhost` works fine. For network access on mobile devices, see the HTTPS setup guide below.

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

## HTTPS Setup for Local Development

Camera access on mobile devices requires HTTPS. For local development with mobile testing:

### Option 1: mkcert (Recommended)

1. Install mkcert:
   ```bash
   brew install mkcert  # macOS
   ```

2. Create local CA:
   ```bash
   mkcert -install
   ```

3. Generate certificate:
   ```bash
   mkcert localhost 192.168.1.x  # Replace with your local IP
   ```

4. Configure Vite to use HTTPS:
   ```typescript
   // vite.config.ts
   export default defineConfig({
     server: {
       https: {
         key: fs.readFileSync('./localhost-key.pem'),
         cert: fs.readFileSync('./localhost.pem')
       }
     }
   });
   ```

### Option 2: Reverse Proxy

Use a reverse proxy (Nginx, Caddy) with SSL termination in production.

## Project Structure

```
frontend/
├── src/
│   ├── lib/
│   │   ├── adapters/
│   │   │   ├── api/           # Backend API communication
│   │   │   ├── camera/        # Camera access via MediaDevices API
│   │   │   ├── orientation/   # Device orientation via DeviceOrientationEvent
│   │   │   └── ports/         # Port interfaces for adapters
│   │   ├── features/          # UI components
│   │   └── shared/
│   │       ├── stores/        # Svelte stores for state management
│   │       └── types/         # TypeScript type definitions
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

The application follows **Hexagonal Architecture** (Ports & Adapters):

### Core Domain
- Application state machine (app stores)
- Business logic (spirit level thresholds, validation)

### Adapters
- **CameraAdapter**: Wraps `navigator.mediaDevices` API
- **OrientationAdapter**: Wraps `DeviceOrientationEvent` API
- **ApiAdapter**: HTTP client for backend communication

### Ports (Interfaces)
- **ICameraPort**: Camera access interface
- **IOrientationPort**: Device orientation interface
- **IApiPort**: Backend API interface

### Benefits
- Testability: Adapters can be mocked in tests
- Flexibility: Easy to swap implementations
- Separation: Browser APIs isolated from business logic

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
- **Safari**: ✅ Full support (iOS 13+ requires orientation permission)
- **Mobile browsers**: ✅ Optimized for iOS Safari and Chrome Android

### Required Browser APIs
- MediaDevices API (camera access)
- DeviceOrientationEvent API (spirit level)
- Fetch API (upload)
- Canvas API (photo capture)

## Accessibility Features

- **Level Check Toggle**: Disable spirit level requirement for users who cannot hold device steady
- **Keyboard Navigation**: Full keyboard support for controls
- **Screen Reader Support**: ARIA labels and semantic HTML
- **File Upload Fallback**: Alternative to camera for users without camera access

## Troubleshooting

### Camera not working

1. **Check HTTPS**: Camera requires HTTPS or localhost
2. **Check permissions**: Ensure browser has camera permission
3. **Check device**: Some browsers block camera on certain devices
4. **Fallback**: Use file upload option if camera doesn't work

### Spirit level not working

1. **iOS 13+**: Requires explicit permission - click "Allow" when prompted
2. **Desktop**: Spirit level may not work on devices without orientation sensors
3. **Accessibility**: Use toggle to disable level requirement

### Upload failing

1. **Check API URL**: Verify `PUBLIC_API_URL` in `.env`
2. **Check network**: Ensure backend is running and accessible
3. **Check file size**: Photos must be under 10MB
4. **Check CORS**: Backend must allow requests from frontend origin

## License

© 2026 Rechtbank voor Meubilair | Satirical Project | No Real Legal Judgments

