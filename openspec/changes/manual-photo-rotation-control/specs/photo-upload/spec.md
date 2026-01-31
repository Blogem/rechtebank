## ADDED Requirements

### Requirement: Photo upload with rotation metadata
The upload flow SHALL accept a photo blob and rotation value from the PhotoCapture component.

#### Scenario: Upload receives rotated blob
- **WHEN** user confirms a photo in the PhotoCapture component
- **THEN** the upload handler receives a blob with rotation already baked in

#### Scenario: Rotation value passed to backend
- **WHEN** photo is uploaded
- **THEN** the final rotation value (0, 90, 180, or 270) is included in the upload metadata

### Requirement: Simplified metadata payload
The upload flow SHALL no longer include device orientation sensor data.

#### Scenario: No beta/gamma sensor data in upload
- **WHEN** photo metadata is constructed for upload
- **THEN** beta and gamma orientation sensor values are not included

#### Scenario: Upload metadata contains essential fields only
- **WHEN** photo metadata is constructed
- **THEN** metadata includes userAgent, timestamp, captureMethod, and rotation only

### Requirement: Upload error handling
The upload flow SHALL handle errors gracefully without dependency on orientation data.

#### Scenario: Upload succeeds with rotated photo
- **WHEN** a rotated photo is uploaded successfully
- **THEN** the upload response includes the verdict based on the correctly oriented image

#### Scenario: Upload fails gracefully
- **WHEN** upload encounters an error
- **THEN** error is displayed to user without exposing orientation sensor issues

### Requirement: Backend compatibility
The upload flow SHALL remain compatible with existing backend API expectations.

#### Scenario: Photo blob format accepted by backend
- **WHEN** rotated photo blob is uploaded
- **THEN** backend accepts the JPEG/PNG blob format without issues

#### Scenario: Metadata schema compatible
- **WHEN** upload metadata is sent to backend
- **THEN** backend processes the simplified metadata without errors

### Requirement: Upload state management
The upload flow SHALL track upload progress and state correctly.

#### Scenario: Upload shows progress indicator
- **WHEN** photo upload is in progress
- **THEN** user sees a loading/progress indicator

#### Scenario: Upload success transitions to verdict display
- **WHEN** photo upload completes successfully
- **THEN** application state transitions to display the verdict

#### Scenario: Upload failure allows retry
- **WHEN** photo upload fails
- **THEN** user can retry the upload or retake the photo
