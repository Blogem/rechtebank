## ADDED Requirements

### Requirement: Upload furniture photo to backend
The system SHALL upload the captured or selected photo to the backend API for judgment.

#### Scenario: Successful photo upload
- **WHEN** user confirms a photo for submission
- **THEN** system sends the photo to the `/v1/judge` endpoint as multipart/form-data

#### Scenario: Upload progress indication
- **WHEN** photo upload is in progress
- **THEN** system displays a loading indicator to the user

#### Scenario: Upload failure with network error
- **WHEN** upload fails due to network connectivity
- **THEN** system displays an error message and provides a retry button

#### Scenario: Upload failure with server error
- **WHEN** backend returns an error response
- **THEN** system displays the error message from the backend to the user

### Requirement: File upload fallback
The system SHALL provide file upload functionality as a fallback when camera access is unavailable.

#### Scenario: File picker for upload
- **WHEN** user chooses to upload from files or camera is denied
- **THEN** system displays a file picker that accepts image formats (JPEG, PNG, WEBP)

#### Scenario: Selected file preview
- **WHEN** user selects a file from their device
- **THEN** system displays a preview of the selected image before upload

#### Scenario: File size validation
- **WHEN** user selects a file larger than 10MB
- **THEN** system displays an error message and prevents upload

### Requirement: Photo format handling
The system SHALL ensure photos are in an acceptable format for the backend API.

#### Scenario: Camera capture format
- **WHEN** photo is captured from camera
- **THEN** system converts the image to JPEG format for upload

#### Scenario: File upload format validation
- **WHEN** user selects a file to upload
- **THEN** system validates the file is an image format (JPEG, PNG, WEBP, GIF)

#### Scenario: Unsupported format rejection
- **WHEN** user attempts to upload a non-image file
- **THEN** system displays an error message and prevents upload

### Requirement: Photo metadata preparation
The system SHALL prepare photo metadata for submission to the backend.

#### Scenario: Multipart form data structure
- **WHEN** system uploads a photo
- **THEN** photo is sent as multipart/form-data with field name "photo"

#### Scenario: Include client metadata
- **WHEN** system uploads a photo
- **THEN** request includes user agent and timestamp in form fields for backend logging
