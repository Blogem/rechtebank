## ADDED Requirements

### Requirement: Accept multipart photo uploads
The API SHALL accept HTTP POST requests with multipart/form-data containing a photo file at the `/v1/judge` endpoint.

#### Scenario: Valid photo upload
- **WHEN** client sends POST request to `/v1/judge` with multipart/form-data containing an image file
- **THEN** system accepts the request and returns HTTP 200 status

#### Scenario: Missing photo data
- **WHEN** client sends POST request without photo data
- **THEN** system returns HTTP 400 with error message "Photo file is required"

#### Scenario: Invalid content type
- **WHEN** client sends POST request with non-multipart content type
- **THEN** system returns HTTP 400 with error message "Content-Type must be multipart/form-data"

### Requirement: Support common image formats
The API SHALL accept JPEG, PNG, and WebP image formats.

#### Scenario: JPEG upload
- **WHEN** client uploads a JPEG image
- **THEN** system processes the image successfully

#### Scenario: PNG upload
- **WHEN** client uploads a PNG image
- **THEN** system processes the image successfully

#### Scenario: WebP upload
- **WHEN** client uploads a WebP image
- **THEN** system processes the image successfully

#### Scenario: Unsupported format
- **WHEN** client uploads an unsupported format (e.g., GIF, BMP)
- **THEN** system returns HTTP 400 with error message "Unsupported image format. Use JPEG, PNG, or WebP"

### Requirement: Enforce file size limits
The API SHALL reject photo uploads larger than 10MB.

#### Scenario: File within size limit
- **WHEN** client uploads a photo smaller than 10MB
- **THEN** system accepts and processes the photo

#### Scenario: File exceeds size limit
- **WHEN** client uploads a photo larger than 10MB
- **THEN** system returns HTTP 413 with error message "Photo file size must not exceed 10MB"

### Requirement: Handle concurrent uploads
The API SHALL handle multiple concurrent photo upload requests without data corruption or request interference.

#### Scenario: Concurrent uploads from different clients
- **WHEN** multiple clients upload photos simultaneously
- **THEN** each request is processed independently and returns correct verdict for respective photo
