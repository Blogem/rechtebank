# Overlay Rotation Button

## Context

The current photo rotation UI places two buttons ("Links" and "Rechts") below the photo preview. This separates the rotation controls from the photo itself, requiring users to shift their visual focus away from the image to find and use the controls. An overlay button positioned directly on the photo keeps the user's attention focused on the image while providing a more streamlined, modern interaction pattern.

## Solution Approach

Place a single semi-transparent rotation button as an overlay on the photo preview that:
- Rotates the image 90° clockwise on each click
- Is visually distinct but not obstructive (semi-transparent design)
- Stays positioned consistently on the photo (e.g., bottom-right corner)
- Provides clear visual feedback on interaction

This simplifies the UI by removing the need for separate controls below the photo and the associated hint text.

## ADDED Requirements

### Requirement: Display semi-transparent rotation button overlay
The system SHALL display a semi-transparent rotation button positioned as an overlay on the photo preview.

#### Scenario: Button is visible on photo preview
- **WHEN** user views photo confirmation screen
- **THEN** a semi-transparent rotation button is displayed overlaid on the photo
- **AND** button is positioned in a consistent location (e.g., bottom-right corner)
- **AND** button has sufficient opacity to be clearly visible against photo content

#### Scenario: Button does not obstruct photo viewing
- **WHEN** rotation button is displayed on photo
- **THEN** button uses semi-transparent background (e.g., rgba with alpha 0.6-0.8)
- **AND** button size is large enough to be easily clickable but small enough not to dominate the photo
- **AND** primary photo content remains clearly visible

#### Scenario: Button has clear rotation affordance
- **WHEN** user views rotation button
- **THEN** button displays a clear rotation icon or symbol (e.g., ↻ or circular arrow)
- **AND** button appearance suggests it is interactive

### Requirement: Rotate image 90° clockwise on click
The system SHALL rotate the photo preview 90° clockwise each time the overlay button is clicked.

#### Scenario: Single click rotates 90° clockwise
- **GIVEN** photo rotation is 0°
- **WHEN** user clicks overlay rotation button
- **THEN** rotation changes to 90°
- **AND** photo visual updates immediately with smooth CSS transition

#### Scenario: Multiple clicks cycle through full rotation
- **GIVEN** photo rotation is 0°
- **WHEN** user clicks overlay rotation button four times
- **THEN** rotation cycles through 90°, 180°, 270°, then back to 0°
- **AND** each rotation change is smooth and immediate

#### Scenario: Rotation from 270° cycles back to 0°
- **GIVEN** photo rotation is 270°
- **WHEN** user clicks overlay rotation button
- **THEN** rotation changes to 0°

### Requirement: Provide touch-friendly click target
The system SHALL ensure the overlay button has a touch-friendly click target for mobile users.

#### Scenario: Button meets minimum touch target size
- **WHEN** rotation button is rendered
- **THEN** button has minimum dimensions of 44×44 pixels (or larger)
- **AND** button is easily tappable on mobile devices

#### Scenario: Button responds to touch and mouse events
- **WHEN** user interacts with rotation button via touch (mobile) or mouse (desktop)
- **THEN** button responds to click/tap events
- **AND** provides visual feedback on interaction (e.g., hover state, active state)

### Requirement: Maintain overlay positioning during rotation
The system SHALL keep the overlay button in a consistent position relative to the photo container as the photo rotates.

#### Scenario: Button remains in fixed position during rotation
- **GIVEN** rotation button is positioned in bottom-right corner of photo container
- **WHEN** photo is rotated to any angle (90°, 180°, 270°)
- **THEN** button remains in the same bottom-right corner position
- **AND** button does not rotate with the photo

#### Scenario: Button is accessible at all rotation angles
- **GIVEN** photo is at any rotation angle
- **WHEN** user views photo preview
- **THEN** rotation button remains visible and clickable
- **AND** button does not overlap with other UI controls (e.g., "Bevestig", "Opnieuw" buttons)

## Styling Guidelines

**Button Appearance:**
- Semi-transparent background: rgba(0, 0, 0, 0.6) or similar
- White or light-colored icon for contrast
- Circular or rounded rectangle shape
- Subtle shadow or border for definition
- Minimum size: 48×48 pixels (touch-friendly)

**Position:**
- Bottom-right corner of photo preview container
- Consistent spacing from edges (e.g., 16px margin)
- Above the photo layer (z-index)

**Interaction States:**
- Default: Semi-transparent
- Hover: Slightly more opaque or color change
- Active/Click: Brief scale or opacity animation
- Disabled: Reduced opacity (if needed during processing)

**Animation:**
- Photo rotation: 0.3s ease transition (existing)
- Button state changes: 0.2s ease transition

## Accessibility

**ARIA Labels:**
- Button MUST have `aria-label="Roteer foto 90 graden"` or similar
- Button role is implicitly "button"

**Keyboard Support:**
- Button SHOULD be keyboard-focusable (if user navigates to photo preview area)
- Support Enter/Space key to trigger rotation

**Visual Clarity:**
- Icon MUST have sufficient contrast against semi-transparent background
- Consider adding a small text label "Roteer" if icon alone is unclear

## Technical Notes

**Implementation Approach:**
- Position button using absolute positioning within photo preview container
- Preview container must have `position: relative`
- Button uses `position: absolute; bottom: 16px; right: 16px;`
- Z-index ensures button appears above photo

**Event Handling:**
- Single click handler increments rotation by 90°
- Rotation logic remains unchanged (cycles 0° → 90° → 180° → 270° → 0°)
- Canvas transformation logic unchanged (applied on upload)

**Responsive Considerations:**
- Button size scales appropriately on smaller screens
- Touch target remains minimum 44×44px
- Button position adjusts if photo preview dimensions change
