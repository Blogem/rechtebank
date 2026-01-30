## Context

The Rechtbank voor Meubilair is an interactive joke website that judges furniture alignment through photos. The frontend needs to provide an engaging, mobile-first experience where users can capture or upload photos of furniture and receive humorous legal verdicts. The application must run in a Docker container with Nginx, interact with a Go backend API at `/v1/judge`, and support camera access on mobile devices (requiring HTTPS).

Current state: New greenfield frontend application with no existing code.

Key constraints:
- Must work on mobile browsers (primary use case)
- Camera access requires HTTPS (strict browser security policy)
- Should be lightweight for minimal VM footprint
- TypeScript and Svelte stack

## Goals / Non-Goals

**Goals:**
- Build a Svelte SPA that captures/uploads photos and displays verdicts
- Implement camera access with proper HTTPS and permission handling
- Create an interactive spirit level overlay using device orientation sensors
- Ensure mobile-first responsive design with comedic/legal theming
- Containerize with Nginx for production deployment
- Communicate with backend API via multipart/form-data photo upload

**Non-Goals:**
- Backend API implementation (separate change)
- HTTPS/SSL certificate management (handled by reverse proxy layer)
- Photo processing or AI integration (backend responsibility)
- User authentication or persistence (stateless joke app)
- Multi-language support (Dutch only for comedic effect)
- SvelteKit server-side features (client-side SPA only)

## Decisions

### 1. Svelte with TypeScript for UI framework
**Decision**: Use Vite with Svelte + TypeScript (SvelteKit adapter-static for SPA mode)
**Rationale**: 
- Svelte provides smaller bundle sizes and better performance than React for this use case
- TypeScript provides type safety for API contracts and browser APIs
- Svelte's reactive syntax is ideal for real-time spirit level updates
- SvelteKit with adapter-static compiles to pure SPA (no SSR complexity)
- Vite provides fast dev experience with HMR
- Alternative considered: React - rejected in favor of lighter bundle and simpler reactive patterns

### 2. MediaDevices API for camera access
**Decision**: Use `navigator.mediaDevices.getUserMedia()` with video constraints  
**Rationale**:
- Standard browser API with good mobile support
- No external dependencies needed
- Provides direct access to camera stream for preview
- Alternative: File input only - rejected because camera UX is more engaging for mobile users

### 3. DeviceOrientationEvent for spirit level
**Decision**: Use `DeviceOrientationEvent.beta` (front-to-back tilt) to create visual level indicator  
**Rationale**:
- Matches the "waterpas" (spirit level) concept from proposal
- Native browser API, no dependencies
- Enhances comedic experience by forcing user engagement
- Alternative: Gyroscope API - rejected as overkill for simple tilt detection

### 4. Conditional upload button activation
**Decision**: Disable photo capture/upload unless device is within ±5° of level (based on beta orientation)  
**Rationale**:
- Enforces the absurdist "legal requirement" that furniture must be photographed straight
- Provides immediate visual feedback via overlay indicator
- Threshold of ±5° balances strictness with usability

### 5. Monorepo structure with separate service directories
**Decision**: Organize as monorepo with service-based directories:
- `frontend/` - Svelte application (this change)
- `backend/` - Go API (future change)
- `docker-compose.yml` - Service orchestration at root
- `docs/` or `infrastructure/` - Shared documentation and reverse proxy configs

**Rationale**:
- Keeps all Rechtbank services in single repository for easier development
- Clear separation of concerns between frontend and backend
- Docker Compose can reference `frontend/Dockerfile` and `backend/Dockerfile`
- Shared config files (docker-compose.yml) live at root
- Alternative: Separate repos per service - rejected for joke project scale, simpler deployment
- Alternative: Monolith with frontend in `static/` - rejected for cleaner architecture

### 6. Docker with Nginx for serving frontend
**Decision**: Multi-stage Docker build at `frontend/Dockerfile` - build stage with Node.js, production stage with Nginx Alpine  
**Rationale**:
- Separates build-time dependencies from runtime
- Nginx Alpine minimizes image size (~20MB vs 1GB+ with Node)
- Nginx efficiently serves static assets with proper caching headers
- Dockerfile located in frontend/ directory for Docker Compose context
- Alternative: Serve directly from Node.js - rejected for larger footprint

### 6. Hexagonal Architecture adaptation
**Decision**: Organize Svelte app with ports/adapters pattern:
- **Core domain**: UI state, business logic (spirit level thresholds, form validation)
- **Adapters**: API client (photo upload), browser APIs (camera, orientation)
- **Ports**: Interfaces for API and device capabilities (enables testing/mocking)

**Rationale**:
- Maintains consistency with project's Hexagonal Architecture principle
- Isolates browser API dependencies for testability
- Enables TDD by mocking camera and orientation adapters
- Alternative: Standard SvelteKit folder structure - rejected to align with architecture principles

### 7. Component structure within frontend/
**Decision**: Feature-based organization with:
- `frontend/src/routes/` - SvelteKit page route (single +page.svelte)
- `frontend/src/lib/features/camera/` - camera capture, preview, controls (.svelte components)
- `frontend/src/lib/features/upload/` - photo upload form and state
- `frontend/src/lib/features/spirit-level/` - orientation overlay component
- `frontend/src/lib/features/verdict/` - verdict display and styling
- `frontend/src/lib/adapters/` - API client, device APIs (TypeScript modules)
- `frontend/src/lib/shared/` - common UI components, types, stores
- `frontend/Dockerfile` - Multi-stage build definition
- `frontend/nginx.conf` - Nginx configuration for SPA routing

**Rationale**:
- Follows SvelteKit conventions (src/lib for reusable code, src/routes for pages)
- Co-locates related components, logic, and tests
- Supports incremental development and testing per feature
- Clear separation between features and technical adapters
- Svelte stores in shared/ for cross-component reactive state
- All frontend code isolated in frontend/ directory for monorepo compatibility

## Risks / Trade-offs

**[Risk]** Browser API compatibility on older mobile devices → **Mitigation**: Feature detection with graceful fallback (disable spirit level if DeviceOrientationEvent unavailable, show file upload if camera denied)

**[Risk]** HTTPS requirement for camera may complicate local development → **Mitigation**: Use `localhost` (exempt from HTTPS requirement) or mkcert for local SSL certificates

**[Risk]** Device orientation permission requires user gesture on iOS → **Mitigation**: Request permission explicitly via button click before enabling spirit level feature

**[Risk]** Camera permission denial breaks core functionality → **Mitigation**: Always provide file upload fallback option, clear permission instructions

**[Trade-off]** Spirit level enforcement (disabled upload) may frustrate users → Acceptable for comedic/absurdist purpose; consider "skip level check" escape hatch for accessibility

**[Trade-off]** SvelteKit + adapter-static + Nginx adds build complexity vs simpler static host → Acceptable for production-quality deployment and alignment with Docker architecture
