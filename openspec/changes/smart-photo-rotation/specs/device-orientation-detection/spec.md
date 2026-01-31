## ADDED Requirements

### Requirement: Capture device orientation at photo moment
The system SHALL capture the DeviceOrientationEvent beta reading at the exact moment a photo is taken.

#### Scenario: Portrait photo capture
- **WHEN** user captures a photo with phone held vertically (beta > 45°)
- **THEN** system records the beta value for orientation calculation

#### Scenario: Landscape photo capture
- **WHEN** user captures a photo with phone held horizontally (beta ≤ 45°)
- **THEN** system records the beta value for orientation calculation

#### Scenario: Sensor unavailable
- **WHEN** user captures a photo and DeviceOrientationEvent is unavailable
- **THEN** system defaults beta to null and continues without error

### Requirement: Calculate initial rotation from beta reading
The system SHALL calculate initial photo rotation angle based on captured beta reading using a 45° threshold.

#### Scenario: Portrait orientation detected
- **WHEN** captured beta reading is greater than 45°
- **THEN** system sets initial rotation to 0° (no rotation needed)

#### Scenario: Landscape orientation detected
- **WHEN** captured beta reading is less than or equal to 45°
- **THEN** system sets initial rotation to 90° (rotate clockwise for correction)

#### Scenario: No beta reading available
- **WHEN** beta reading is null (sensor unavailable)
- **THEN** system sets initial rotation to 0° (default)

### Requirement: Reset orientation detection on new capture
The system SHALL reset and recalculate orientation detection each time a new photo is captured.

#### Scenario: Multiple photo captures
- **WHEN** user captures a first photo with beta 90° (rotation set to 0°)
- **WHEN** user retakes photo with beta 10° (different orientation)
- **THEN** system recalculates and sets rotation to 90° for the new photo

#### Scenario: Manual rotation cleared on retake
- **WHEN** user manually adjusts rotation to 180° on first photo
- **WHEN** user retakes photo
- **THEN** system discards manual rotation and recalculates from new beta reading

### Requirement: Expose beta reading for debugging
The system SHOULD expose the captured beta reading in development mode for debugging and threshold tuning.

#### Scenario: Development mode beta logging
- **WHEN** photo is captured in development environment
- **THEN** system logs beta value and calculated rotation to console
