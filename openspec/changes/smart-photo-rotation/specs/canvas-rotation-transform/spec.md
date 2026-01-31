## ADDED Requirements

### Requirement: Apply rotation transformation before upload
The system SHALL apply canvas rotation transformation to the photo before uploading when rotation is not 0°.

#### Scenario: No rotation needed
- **WHEN** photo rotation is 0°
- **THEN** system uploads original photo without transformation

#### Scenario: Rotation applied
- **WHEN** photo rotation is 90°, 180°, or 270°
- **THEN** system creates rotated canvas and uploads transformed image

### Requirement: Correct canvas dimensions for rotation
The system SHALL set canvas dimensions appropriately based on rotation angle to prevent cropping or distortion.

#### Scenario: No dimension swap for 0° or 180°
- **WHEN** rotation is 0° or 180°
- **WHEN** original image is 1080×1920 pixels
- **THEN** canvas dimensions are 1080×1920 pixels

#### Scenario: Dimension swap for 90° or 270°
- **WHEN** rotation is 90° or 270°
- **WHEN** original image is 1080×1920 pixels
- **THEN** canvas dimensions are 1920×1080 pixels (width and height swapped)

### Requirement: Apply rotation transformation matrix
The system SHALL apply the correct rotation transformation matrix to render the image at the specified angle.

#### Scenario: 90° clockwise rotation
- **WHEN** rotation is 90°
- **THEN** canvas transformation matrix rotates image 90° clockwise from center

#### Scenario: 180° rotation
- **WHEN** rotation is 180°
- **THEN** canvas transformation matrix rotates image 180° from center

#### Scenario: 270° clockwise rotation
- **WHEN** rotation is 270°
- **THEN** canvas transformation matrix rotates image 270° clockwise from center

### Requirement: Center image in rotated canvas
The system SHALL center the image in the canvas after applying rotation to prevent offset or clipping.

#### Scenario: Image centered after rotation
- **WHEN** applying any rotation transformation
- **THEN** image is translated to canvas center before rotation
- **THEN** image is drawn centered with correct dimensions

### Requirement: Export rotated image as JPEG
The system SHALL export the rotated canvas as a JPEG blob for upload.

#### Scenario: JPEG format preserved
- **WHEN** rotation transformation is applied
- **THEN** output is JPEG format blob

#### Scenario: No EXIF metadata needed
- **WHEN** rotation transformation is applied
- **THEN** output image contains correctly rotated pixels without EXIF orientation tag

#### Scenario: Image quality preserved
- **WHEN** rotation transformation is applied
- **THEN** JPEG quality setting matches original conversion quality

### Requirement: Handle rotation transformation errors
The system SHALL handle errors during rotation transformation gracefully without blocking upload.

#### Scenario: Canvas creation failure
- **WHEN** canvas rotation transformation fails
- **THEN** system logs error and uploads original photo unrotated

#### Scenario: Image decode failure
- **WHEN** image cannot be decoded for rotation
- **THEN** system logs error and uploads original blob
