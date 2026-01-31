## ADDED Requirements

### Requirement: Device orientation detection
The system SHALL detect the device's tilt angle using the DeviceOrientationEvent API.

#### Scenario: Request orientation permission on iOS
- **WHEN** user is on an iOS device (iOS 13+)
- **THEN** system displays a button to request device orientation permission

#### Scenario: Permission granted for orientation
- **WHEN** user grants orientation permission
- **THEN** system begins monitoring device tilt angle (beta axis for front-to-back tilt)

#### Scenario: Permission denied for orientation
- **WHEN** user denies orientation permission
- **THEN** system disables spirit level feature and allows photo capture without level check

#### Scenario: Automatic orientation on non-iOS
- **WHEN** user is on a non-iOS device
- **THEN** system automatically begins monitoring device orientation without permission prompt

### Requirement: Spirit level visual indicator
The system SHALL display a visual spirit level overlay showing the device's tilt status.

#### Scenario: Level indicator overlay
- **WHEN** camera preview is active
- **THEN** system displays a spirit level graphic overlaid on the camera preview

#### Scenario: Tilt angle visualization
- **WHEN** device tilts
- **THEN** spirit level indicator reflects the current tilt angle in real-time

#### Scenario: Level achievement visual feedback
- **WHEN** device reaches level position (within ±5° threshold)
- **THEN** spirit level indicator changes to green/success state

#### Scenario: Off-level visual feedback
- **WHEN** device is tilted beyond ±5° threshold
- **THEN** spirit level indicator shows red/warning state with tilt direction

### Requirement: Photo capture conditional on level state
The system SHALL only enable photo capture when device is held level.

#### Scenario: Capture disabled when tilted
- **WHEN** device tilt exceeds ±5° from level
- **THEN** photo capture button is disabled with visual indication

#### Scenario: Capture enabled when level
- **WHEN** device is within ±5° of level position
- **THEN** photo capture button is enabled and highlighted

#### Scenario: Level status message
- **WHEN** capture button is disabled due to tilt
- **THEN** system displays a message instructing user to hold device level

### Requirement: Level threshold configuration
The system SHALL use a ±5° threshold for determining level state.

#### Scenario: Within threshold
- **WHEN** device beta angle is between -5° and +5°
- **THEN** system considers device to be level

#### Scenario: Beyond threshold
- **WHEN** device beta angle is less than -5° or greater than +5°
- **THEN** system considers device to be off-level

### Requirement: Accessibility escape hatch
The system SHALL provide an option to bypass the level requirement for accessibility.

#### Scenario: Disable level check option
- **WHEN** user cannot hold device steady due to accessibility needs
- **THEN** system provides a clearly labeled option to disable the level check requirement

#### Scenario: Photo capture without level check
- **WHEN** user has disabled the level check
- **THEN** photo capture button is always enabled regardless of device orientation

#### Scenario: Re-enable level check
- **WHEN** user wants to re-enable the level requirement
- **THEN** system provides an option to turn the level check back on
