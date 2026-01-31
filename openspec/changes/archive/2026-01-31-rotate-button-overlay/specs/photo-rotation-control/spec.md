# Photo Rotation Control

## MODIFIED Requirements

### Requirement: Provide rotation controls
The system SHALL provide a single overlay button to rotate the photo clockwise only.

#### Scenario: Rotate button increases rotation
- **GIVEN** photo rotation is 0°
- **WHEN** user clicks overlay rotation button
- **THEN** rotation changes to 90°
- **AND** photo visual updates immediately

#### Scenario: Rotation cycles through 360°
- **GIVEN** photo rotation is 270°
- **WHEN** user clicks overlay rotation button
- **THEN** rotation changes to 0°

## REMOVED Requirements

### Requirement: Rotate left button decreases rotation
**Reason**: Simplified to single clockwise rotation button overlay. Counter-clockwise rotation removed to streamline UI and reduce cognitive load.
**Migration**: Users can achieve 270° rotation by clicking the rotation button three times (instead of one counter-clockwise click). This is an acceptable tradeoff for a cleaner, more focused UI.

### Requirement: Rotation cycles backwards through 360°
**Reason**: Counter-clockwise rotation capability removed with the "Links" button.
**Migration**: Use clockwise rotation button multiple times to achieve desired angle.

### Requirement: Rotation controls hint text
**Reason**: Overlay button is self-explanatory and positioned directly on the photo, eliminating the need for separate hint text ("Staat de foto niet goed? Roteer hem eerst:").
**Migration**: Visual affordance of overlay button (rotation icon) replaces textual hint.
