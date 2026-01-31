## ADDED Requirements

### Requirement: Canvas dimension handling for rotation
The system SHALL correctly set canvas dimensions based on rotation angle to prevent image clipping.

#### Scenario: No rotation (0 degrees)
- **WHEN** rotation is 0 degrees
- **THEN** canvas width equals image width and canvas height equals image height

#### Scenario: 180-degree rotation
- **WHEN** rotation is 180 degrees
- **THEN** canvas width equals image width and canvas height equals image height

#### Scenario: 90-degree rotation (portrait)
- **WHEN** rotation is 90 degrees
- **THEN** canvas width equals image height and canvas height equals image width

#### Scenario: 270-degree rotation (portrait)
- **WHEN** rotation is 270 degrees
- **THEN** canvas width equals image height and canvas height equals image width

### Requirement: Canvas coordinate transformation
The system SHALL apply proper coordinate transformations to rotate images without clipping or displacement.

#### Scenario: Translate to center before rotation
- **WHEN** applying rotation transformation
- **THEN** the canvas context is translated to the center point before rotation

#### Scenario: Rotate around center point
- **WHEN** rotation is applied
- **THEN** the context is rotated by (rotation × π / 180) radians

#### Scenario: Draw image centered after rotation
- **WHEN** drawing the rotated image
- **THEN** the image is drawn with its center at the origin (-imageWidth/2, -imageHeight/2)

### Requirement: Image quality preservation
The system SHALL maintain acceptable image quality during canvas rotation processing.

#### Scenario: Default canvas smoothing enabled
- **WHEN** drawing image to canvas
- **THEN** image smoothing is enabled (default behavior)

#### Scenario: High-quality JPEG export
- **WHEN** exporting canvas to blob
- **THEN** JPEG quality parameter is set between 0.85 and 0.95

### Requirement: Blob export functionality
The system SHALL convert the rotated canvas to a blob for upload.

#### Scenario: Export as JPEG blob
- **WHEN** user confirms the rotated photo
- **THEN** canvas is converted to a JPEG blob using toBlob()

#### Scenario: Export as PNG blob for transparency
- **WHEN** the original image has transparency (PNG, WebP with alpha)
- **THEN** canvas is converted to a PNG blob to preserve transparency

### Requirement: Rotation angle validation
The system SHALL only accept valid rotation angles.

#### Scenario: Valid rotation angles accepted
- **WHEN** rotation angle is 0, 90, 180, or 270
- **THEN** the transformation is applied correctly

#### Scenario: Invalid rotation angles normalized
- **WHEN** rotation angle is outside standard values (e.g., 360, -90, 450)
- **THEN** the angle is normalized to 0, 90, 180, or 270 using modulo arithmetic

### Requirement: Canvas transformation correctness
The system SHALL ensure rotated images remain fully visible within canvas bounds.

#### Scenario: 90-degree clockwise rotation visual correctness
- **WHEN** an image with text "TOP" at the top is rotated 90 degrees clockwise
- **THEN** the text "TOP" appears on the right edge of the rotated canvas

#### Scenario: 180-degree rotation visual correctness
- **WHEN** an image with text "TOP" at the top is rotated 180 degrees
- **THEN** the text "TOP" appears at the bottom edge upside-down

#### Scenario: 270-degree rotation visual correctness
- **WHEN** an image with text "TOP" at the top is rotated 270 degrees clockwise
- **THEN** the text "TOP" appears on the left edge of the rotated canvas

#### Scenario: No image clipping occurs
- **WHEN** any valid rotation is applied
- **THEN** all pixels of the original image are visible in the rotated canvas

### Requirement: Performance considerations
The system SHALL handle canvas operations efficiently for typical mobile image sizes.

#### Scenario: Rotation completes in reasonable time
- **WHEN** rotating an image up to 4000x3000 pixels
- **THEN** the canvas transformation completes within 500ms on modern mobile devices

#### Scenario: Memory usage is acceptable
- **WHEN** processing images on mobile devices
- **THEN** canvas operations do not cause memory warnings or crashes
