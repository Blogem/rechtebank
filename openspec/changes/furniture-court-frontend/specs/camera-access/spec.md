## ADDED Requirements

### Requirement: Camera permission request
The system SHALL request camera access permission from the user's browser.

#### Scenario: First-time camera access
- **WHEN** user opens the application for the first time
- **THEN** browser displays a permission prompt for camera access

#### Scenario: Permission denied
- **WHEN** user denies camera permission
- **THEN** system displays a fallback file upload option with instructions on how to enable camera access

#### Scenario: Permission granted
- **WHEN** user grants camera permission
- **THEN** system activates the camera and displays a live preview

### Requirement: Camera stream preview
The system SHALL display a live camera preview when camera access is granted.

#### Scenario: Preview display
- **WHEN** camera access is granted
- **THEN** system displays the camera feed in a preview area with appropriate aspect ratio for furniture photography

#### Scenario: Preview stops on photo capture
- **WHEN** user captures a photo
- **THEN** system freezes the preview and displays the captured image for confirmation

### Requirement: HTTPS enforcement for camera access
The system SHALL only request camera access when served over HTTPS or localhost.

#### Scenario: HTTPS connection
- **WHEN** application is accessed via HTTPS
- **THEN** system enables camera access functionality

#### Scenario: HTTP connection (non-localhost)
- **WHEN** application is accessed via HTTP on non-localhost domain
- **THEN** system displays an error message that camera requires HTTPS and shows file upload fallback

#### Scenario: Localhost development
- **WHEN** application is accessed via http://localhost
- **THEN** system enables camera access functionality (browser exemption)

### Requirement: Mobile camera optimization
The system SHALL optimize camera settings for mobile device photography.

#### Scenario: Mobile rear camera preference
- **WHEN** user is on a mobile device
- **THEN** system requests the rear-facing camera by default for better furniture photography

#### Scenario: Camera switching capability
- **WHEN** user is on a device with multiple cameras
- **THEN** system provides a control to switch between front and rear cameras

### Requirement: Photo capture
The system SHALL allow users to capture a photo from the camera stream.

#### Scenario: Successful photo capture
- **WHEN** user clicks the capture button
- **THEN** system captures the current camera frame as an image file

#### Scenario: Captured photo confirmation
- **WHEN** photo is captured
- **THEN** system displays the captured image with options to retake or proceed to upload

#### Scenario: Retake photo
- **WHEN** user chooses to retake the photo
- **THEN** system resumes the live camera preview and allows another capture
