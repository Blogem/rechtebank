## ADDED Requirements

### Requirement: Display rotation controls in confirmation screen
The system SHALL display rotation controls (left and right buttons) in the photo confirmation screen.

#### Scenario: Rotation controls visible
- **WHEN** user views photo confirmation screen
- **THEN** system displays "↶ Links" and "↷ Rechts" rotation buttons

#### Scenario: Hint text displayed
- **WHEN** user views photo confirmation screen
- **THEN** system displays hint text "Staat de foto niet goed? Roteer hem eerst:"

### Requirement: Rotate photo left
The system SHALL rotate the photo 90° counter-clockwise when user clicks the left rotation button.

#### Scenario: Rotate left from 0°
- **WHEN** photo rotation is 0° and user clicks "↶ Links"
- **THEN** system sets rotation to 270°

#### Scenario: Rotate left from 90°
- **WHEN** photo rotation is 90° and user clicks "↶ Links"
- **THEN** system sets rotation to 0°

#### Scenario: Rotation cycles correctly
- **WHEN** user clicks "↶ Links" four times from any starting rotation
- **THEN** photo returns to original rotation angle

### Requirement: Rotate photo right
The system SHALL rotate the photo 90° clockwise when user clicks the right rotation button.

#### Scenario: Rotate right from 0°
- **WHEN** photo rotation is 0° and user clicks "↷ Rechts"
- **THEN** system sets rotation to 90°

#### Scenario: Rotate right from 270°
- **WHEN** photo rotation is 270° and user clicks "↷ Rechts"
- **THEN** system sets rotation to 0°

#### Scenario: Rotation cycles correctly
- **WHEN** user clicks "↷ Rechts" four times from any starting rotation
- **THEN** photo returns to original rotation angle

### Requirement: Show rotated preview
The system SHALL display a visually rotated preview of the photo in real-time as user adjusts rotation.

#### Scenario: Preview updates on rotation
- **WHEN** user clicks rotation button
- **THEN** photo preview updates instantly to show new rotation angle

#### Scenario: CSS transform applied
- **WHEN** rotation is set to 90°
- **THEN** photo preview has CSS transform "rotate(90deg)" applied

#### Scenario: Preview shows final orientation
- **WHEN** user views rotated preview
- **THEN** preview matches exactly how photo will appear after upload to Gemini

### Requirement: Maintain rotation state across interactions
The system SHALL maintain the current rotation angle as user interacts with other confirmation screen controls.

#### Scenario: Rotation preserved during other actions
- **WHEN** user rotates photo to 180°
- **WHEN** user clicks accessibility toggle or other UI element
- **THEN** photo rotation remains at 180°

#### Scenario: Rotation state available to upload
- **WHEN** user clicks submit button
- **THEN** current rotation angle is passed to upload function
