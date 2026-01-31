## Why

We need an interactive frontend for the "Rechtbank voor Meubilair" (Furniture Court) joke website that allows users to submit photos of their furniture for comedic legal judgment. The frontend must support camera access and photo upload functionality while adding an interactive spirit level overlay to enhance the absurdist experience.

## What Changes

- Create a React-based single-page application (SPA) for the Furniture Court
- Implement camera access and photo upload functionality for furniture submissions
- Add an interactive spirit level overlay using device orientation sensors
- Display AI-generated court verdicts from the backend API
- Serve the application via Nginx in a Docker container
- Ensure HTTPS compatibility for camera access on mobile devices

## Capabilities

### New Capabilities
- `camera-access`: Access device camera for taking photos of furniture, with HTTPS requirement for mobile compatibility
- `photo-upload`: Upload furniture photos to the backend API endpoint for judgment
- `spirit-level-overlay`: Real-time spirit level visualization using DeviceOrientationEvent API to guide users in taking straight photos
- `verdict-display`: Display the AI-generated legal verdict and score (1-10) from the backend with appropriate comedic styling

### Modified Capabilities
<!-- No existing capabilities are being modified -->

## Impact

- New frontend application requiring a Docker container with Nginx
- Requires HTTPS configuration (via reverse proxy) for camera access on smartphones
- Depends on backend API endpoint `/v1/judge` for photo processing
- Adds browser API dependencies: MediaDevices API (camera), DeviceOrientationEvent API (spirit level)
- Creates `frontend/` directory at repository root (prepares for `backend/` directory in future change)
- Docker Compose will orchestrate frontend container alongside future backend and reverse proxy containers
- No breaking changes to existing systems
