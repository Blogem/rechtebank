## 1. Project Setup and Infrastructure

- [x] 1.1 Initialize SvelteKit project with TypeScript in `frontend/` directory
- [x] 1.2 Configure adapter-static for SPA mode in svelte.config.js
- [x] 1.3 Set up Vite configuration for TypeScript and environment variables
- [x] 1.4 Install base dependencies (svelte, typescript, vite)
- [x] 1.5 Create frontend/Dockerfile with multi-stage build (Node.js build + Nginx Alpine)
- [x] 1.6 Create frontend/nginx.conf for SPA routing and caching headers
- [x] 1.7 Set up ESLint and Prettier for code quality
- [x] 1.8 Configure Vitest for unit and component testing

## 2. Project Structure and Hexagonal Architecture

- [x] 2.1 Create folder structure: src/lib/features/, src/lib/adapters/, src/lib/shared/
- [x] 2.2 Define port interfaces in src/lib/adapters/ports/ (ICameraPort, IOrientationPort, IApiPort)
- [x] 2.3 Create TypeScript types in src/lib/shared/types/ (Verdict, PhotoMetadata, OrientationData)
- [x] 2.4 Set up Svelte stores in src/lib/shared/stores/ for app state
- [x] 2.5 Create main route at src/routes/+page.svelte

## 3. Camera Access Feature (camera-access capability)

- [x] 3.1 Write unit tests for CameraAdapter interface (permission request, stream access)
- [x] 3.2 Implement CameraAdapter in src/lib/adapters/camera/ wrapping MediaDevices API
- [x] 3.3 Write tests for HTTPS detection logic with localhost exemption
- [x] 3.4 Add HTTPS detection logic with localhost exemption check
- [x] 3.5 Write component tests for CameraPermission (granted, denied scenarios)
- [x] 3.6 Create CameraPermission component for permission request flow
- [x] 3.7 Write component tests for CameraPreview (stream display, capture)
- [x] 3.8 Create CameraPreview component displaying live video stream
- [x] 3.9 Write tests for mobile camera constraints (rear camera preference)
- [x] 3.10 Implement mobile rear camera preference in getUserMedia constraints
- [x] 3.11 Write tests for camera switching functionality
- [x] 3.12 Add camera switching control for front/rear camera toggle
- [x] 3.13 Write tests for photo capture from video stream
- [x] 3.14 Create photo capture function to extract frame from video stream
- [x] 3.15 Write component tests for captured photo confirmation UI
- [x] 3.16 Implement captured photo confirmation UI with retake option
- [x] 3.17 Write component tests for FileUploadFallback
- [x] 3.18 Create FileUploadFallback component for camera permission denial

## 4. Spirit Level Overlay Feature (spirit-level-overlay capability)

- [x] 4.1 Write unit tests for OrientationAdapter (beta angle detection, threshold logic)
- [x] 4.2 Implement OrientationAdapter in src/lib/adapters/orientation/ wrapping DeviceOrientationEvent
- [x] 4.3 Write tests for ±5° threshold detection logic
- [x] 4.4 Implement ±5° threshold logic for level detection
- [x] 4.5 Write tests for iOS 13+ permission request handling
- [x] 4.6 Add iOS 13+ permission request handling for orientation access
- [x] 4.7 Write tests for orientation reactive store updates
- [x] 4.8 Implement orientation reactive store for component updates
- [x] 4.9 Write component tests for SpiritLevel visual states (green/red feedback)
- [x] 4.10 Create SpiritLevel component with visual bubble/level indicator
- [x] 4.11 Implement real-time tilt angle calculation from beta axis
- [x] 4.12 Create visual feedback states (green for level, red for tilted)
- [x] 4.13 Write tests for accessibility toggle functionality
- [x] 4.14 Add accessibility toggle to disable level requirement
- [x] 4.15 Write component tests for help text display when off-level
- [x] 4.16 Create help text showing tilt direction when off-level

## 5. Photo Upload Feature (photo-upload capability)

- [x] 5.1 Write unit tests for ApiAdapter (upload function, multipart/form-data format)
- [x] 5.2 Implement ApiAdapter in src/lib/adapters/api/ for backend communication
- [x] 5.3 Write tests for photo format conversion to JPEG
- [x] 5.4 Add file format conversion to JPEG for camera captures
- [x] 5.5 Write tests for file size validation (10MB max) and format validation
- [x] 5.6 Implement file size validation (10MB max) and format validation
- [x] 5.7 Write tests for metadata addition to form data
- [x] 5.8 Add metadata fields (user agent, timestamp) to form data
- [x] 5.9 Write tests for upload retry logic on network errors
- [x] 5.10 Implement retry logic for network errors
- [x] 5.11 Write tests for server error handling
- [x] 5.12 Add error handling for server errors with user-friendly messages
- [x] 5.13 Write integration tests for complete upload flow with mocked API
- [x] 5.14 Create photo upload function sending multipart/form-data to /v1/judge
- [x] 5.15 Write component tests for file picker component
- [x] 5.16 Implement file picker component for fallback upload
- [x] 5.17 Write component tests for upload progress indicator
- [x] 5.18 Create upload progress indicator component
- [x] 5.19 Write component tests for file preview
- [x] 5.20 Create file preview component for selected/uploaded files

## 6. Verdict Display Feature (verdict-display capability)

- [x] 6.1 Write component tests for VerdictDisplay (all verdict types)
- [x] 6.2 Create VerdictDisplay component with legal/courtroom themed styling
- [x] 6.3 Write tests for score display (1-10) rendering
- [x] 6.4 Implement score display (1-10) with visual prominence
- [x] 6.5 Write tests for verdict type styling variants
- [x] 6.6 Add verdict type styling (niet-ontvankelijk, guilty, acquittal)
- [x] 6.7 Write tests for loading animation and timeout behavior
- [x] 6.8 Create loading animation with Dutch text "De rechter beraadslaagt..."
- [x] 6.9 Implement 30-second timeout for backend response
- [x] 6.10 Write tests for legal terminology formatting
- [x] 6.11 Add legal terminology formatting with Dutch language support
- [x] 6.12 Write tests for sentencing display
- [x] 6.13 Create sentencing display with dramatic styling
- [x] 6.14 Write component tests for reset flow button
- [x] 6.15 Implement "Try another judgment" button to reset flow
- [x] 6.16 Write tests for share verdict functionality
- [x] 6.17 Add share verdict option (copy link or download image)
- [x] 6.18 Write tests for error display in legal format
- [x] 6.19 Create error display in legal-styled "case dismissed" format

## 7. Main Application Flow Integration

- [x] 7.1 Write tests for state machine flow (camera → capture → upload → verdict)
- [x] 7.2 Implement state machine for app flow (camera → capture → upload → verdict)
- [x] 7.3 Write tests for main page component orchestration
- [x] 7.4 Implement main page component orchestrating all features
- [x] 7.5 Write tests for spirit level integration with camera preview
- [x] 7.6 Integrate spirit level with camera preview overlay
- [x] 7.7 Write tests for conditional capture button based on level state
- [x] 7.8 Add conditional capture button based on level state
- [x] 7.9 Write tests for photo capture to upload flow connection
- [x] 7.10 Connect photo capture to upload flow
- [x] 7.11 Write tests for upload completion to verdict display
- [x] 7.12 Wire upload completion to verdict display
- [x] 7.13 Write tests for reset flow from verdict back to camera
- [x] 7.14 Implement reset flow from verdict back to camera
- [x] 7.15 Add responsive mobile-first CSS layout
- [x] 7.16 Create comedic legal theme styling (fonts, colors, icons)

## 8. End-to-End Testing and Quality Assurance

- [~] 8.1 Write end-to-end tests for complete user flow (camera → verdict)
- [~] 8.2 Write end-to-end tests for file upload fallback flow
- [x] 8.3 Test camera access on mobile browsers (iOS Safari, Chrome Android)
- [x] 8.4 Test orientation permission flow on iOS 13+ devices
- [x] 8.5 Verify HTTPS enforcement and localhost exemption
- [x] 8.6 Test file upload fallback when camera denied
- [x] 8.7 Verify spirit level accuracy with device tilt
- [~] 8.8 Test upload with various photo formats and sizes
- [~] 8.9 Verify verdict display with different backend responses
- [~] 8.10 Test error handling and retry flows
- [~] 8.11 Accessibility audit (keyboard nav, screen readers, level bypass)

## 9. Docker and Deployment Configuration

- [x] 9.1 Test multi-stage Docker build locally
- [x] 9.2 Verify Nginx serves SPA with correct routing (all routes → index.html)
- [x] 9.3 Configure Nginx caching headers for static assets
- [x] 9.4 Add environment variable injection for API endpoint URL
- [x] 9.5 Create .dockerignore file
- [x] 9.6 Document build and run commands in frontend/README.md
- [x] 9.7 Test container startup and health check

## 10. Documentation and Finalization

- [x] 10.1 Document camera permission setup instructions for users
- [x] 10.2 Add development setup guide in frontend/README.md
- [x] 10.3 Document environment variables and configuration
- [x] 10.4 Add API contract documentation for /v1/judge endpoint
- [x] 10.5 Create HTTPS setup guide for local development (localhost/mkcert)
- [x] 10.6 Document accessibility features and escape hatches
- [x] 10.7 Add troubleshooting guide for common issues
