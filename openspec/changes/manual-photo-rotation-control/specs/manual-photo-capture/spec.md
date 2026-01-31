## ADDED Requirements

### Requirement: Photo capture using file input
The component SHALL provide a native file input interface that triggers the device camera on mobile browsers.

#### Scenario: User initiates photo capture on mobile device
- **WHEN** user taps the capture button
- **THEN** the mobile device's native camera interface opens

#### Scenario: User selects photo from file input
- **WHEN** user selects an image file via the file input
- **THEN** the component loads the image and displays a preview

#### Scenario: File input restricted to images only
- **WHEN** the file input dialog opens
- **THEN** only image file types are selectable (accept="image/*")

### Requirement: Photo preview display
The component SHALL display a preview of the captured or selected photo before confirmation.

#### Scenario: Preview shows selected image
- **WHEN** an image is loaded into the component
- **THEN** the preview displays the image with current rotation applied visually

#### Scenario: Preview is responsive
- **WHEN** the preview is displayed
- **THEN** the image fits within the viewport without horizontal scrolling or distortion

### Requirement: Manual rotation controls
The component SHALL provide rotation controls allowing users to rotate the photo in 90-degree increments.

#### Scenario: Rotate left button rotates counter-clockwise
- **WHEN** user clicks the rotate-left button
- **THEN** the preview rotates 90 degrees counter-clockwise

#### Scenario: Rotate right button rotates clockwise
- **WHEN** user clicks the rotate-right button
- **THEN** the preview rotates 90 degrees clockwise

#### Scenario: Rotation wraps at 360 degrees
- **WHEN** rotation reaches 360 degrees in either direction
- **THEN** rotation value wraps to 0 degrees

#### Scenario: Rotation is tracked in state
- **WHEN** user rotates the image multiple times
- **THEN** the component tracks the cumulative rotation (0, 90, 180, or 270 degrees)

### Requirement: Initial rotation heuristic
The component SHALL attempt to determine an initial rotation value using the screen orientation API.

#### Scenario: Screen orientation API available
- **WHEN** screen.orientation.angle is available
- **THEN** initial rotation is set to screen.orientation.angle value

#### Scenario: Fallback to window.orientation
- **WHEN** screen.orientation.angle is unavailable but window.orientation exists
- **THEN** initial rotation is derived from window.orientation value

#### Scenario: No orientation API available
- **WHEN** neither screen.orientation nor window.orientation is available
- **THEN** initial rotation defaults to 0 degrees

### Requirement: Photo confirmation
The component SHALL provide a confirm action that exports the final rotated image.

#### Scenario: User confirms photo
- **WHEN** user clicks the confirm button
- **THEN** the component invokes the onPhotoConfirmed callback with the rotated image blob and rotation value

#### Scenario: Confirmed image includes baked rotation
- **WHEN** user confirms a rotated photo
- **THEN** the exported blob contains the rotation applied permanently (not just metadata)

### Requirement: Photo retake
The component SHALL provide a retake action that allows users to select a new photo.

#### Scenario: User retakes photo
- **WHEN** user clicks the retake button
- **THEN** the current photo is discarded and the file input is re-triggered

#### Scenario: Memory cleanup on retake
- **WHEN** a new photo is selected after retake
- **THEN** the previous photo's object URL is revoked to prevent memory leaks

### Requirement: Memory management
The component SHALL properly manage object URLs to prevent memory leaks.

#### Scenario: Object URL created for preview
- **WHEN** an image file is selected
- **THEN** a blob URL is created using URL.createObjectURL()

#### Scenario: Object URL revoked on new photo
- **WHEN** a new photo replaces the current one
- **THEN** the previous object URL is revoked

#### Scenario: Object URL revoked on component unmount
- **WHEN** the component is destroyed
- **THEN** any active object URL is revoked

### Requirement: Browser compatibility
The component SHALL function correctly on modern mobile browsers.

#### Scenario: iOS Safari support
- **WHEN** component is used on iOS Safari 13 or later
- **THEN** all features work correctly including camera access and rotation

#### Scenario: Android Chrome support
- **WHEN** component is used on Android Chrome 90 or later
- **THEN** all features work correctly including camera access and rotation

### Requirement: Component API
The component SHALL expose a clear and minimal prop-based API.

#### Scenario: onPhotoConfirmed callback required
- **WHEN** component is instantiated
- **THEN** onPhotoConfirmed prop is required and receives (blob: Blob, rotation: number)

#### Scenario: onCancelled callback optional
- **WHEN** component is instantiated
- **THEN** onCancelled prop is optional and called when user cancels

### Requirement: Error handling
The component SHALL handle error cases gracefully.

#### Scenario: Non-image file selected
- **WHEN** user somehow selects a non-image file type
- **THEN** component displays an error message and does not process the file

#### Scenario: Image load failure
- **WHEN** selected image fails to load
- **THEN** component displays an error message and allows retry
