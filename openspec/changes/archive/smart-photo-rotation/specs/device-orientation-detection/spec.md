## ADDED Requirements

### Requirement: Capture device orientation at photo moment
The system SHALL capture the DeviceOrientationEvent beta reading at the exact moment a photo is taken.

#### Scenario: Portrait photo capture (normal)
- **WHEN** user captures a photo with phone held vertically with camera at top (beta > 45°)
- **THEN** system records the beta value for orientation calculation

#### Scenario: Portrait photo capture (upside-down)
- **WHEN** user captures a photo with phone held vertically with camera at bottom (beta < -45°)
- **THEN** system records the beta value for orientation calculation

#### Scenario: Landscape photo capture
- **WHEN** user captures a photo with phone held horizontally (-45° ≤ beta ≤ 45°)
- **THEN** system records the beta value for orientation calculation

#### Scenario: Sensor unavailable
- **WHEN** user captures a photo and DeviceOrientationEvent is unavailable
- **THEN** system defaults beta to null and continues without error

### Requirement: Calculate initial rotation from beta reading
The system SHALL calculate initial photo rotation angle based on captured beta reading, detecting portrait normal, portrait upside-down, and landscape orientations using beta angle and sign.

#### Scenario: Portrait normal orientation detected
- **WHEN** captured beta reading is greater than 45° (positive)
- **THEN** system sets initial rotation to 0° (no rotation needed)

#### Scenario: Portrait upside-down orientation detected
- **WHEN** captured beta reading is less than -45° (negative)
- **THEN** system sets initial rotation to 180° (flip correction)

#### Scenario: Landscape orientation detected
- **WHEN** captured beta reading is between -45° and 45° (inclusive)
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
- **WHEN** user retakes photoand calculated rotation in development mode for debugging and threshold tuning.

#### Scenario: Development mode orientation logging
- **WHEN** photo is captured in development environment
- **THEN** system logs beta value, calculated rotation, and detected orientation type to console
- **EXAMPLE**: "[Rotation] Beta: -87° → Initial rotation: 180° (portrait upside-down)"

#### Scenario: All orientation types logged correctly
- **WHEN** beta = 90° → logs "Beta: 90° → Initial rotation: 0° (portrait normal)"
- **WHEN** beta = -90° → logs "Beta: -90° → Initial rotation: 180° (portrait upside-down)"
- **WHEN** beta = 10° → logs "Beta: 10° → Initial rotation: 90° (landscape)"de for debugging and threshold tuning.

#### Scenario: Development mode beta logging
- **WHEN** photo is captured in development environment
- **THEN** system logs beta value and calculated rotation to console
