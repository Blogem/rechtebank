## Why

The current photo rotation UI uses two separate buttons ("Links" and "Rechts") positioned below the photo preview with hint text, requiring users to look away from the photo to find and use rotation controls. This creates unnecessary cognitive friction and visual separation. An overlay button directly on the photo would streamline the interaction by keeping the user's attention focused on the image itself.

## What Changes

- Replace two rotation buttons (rotate left/right) below photo with a single semi-transparent rotate button overlaid on the photo
- Button rotates image 90° clockwise on each click
- Remove counter-clockwise rotation logic and UI
- Remove rotation hint text ("Staat de foto niet goed? Roteer hem eerst:")
- Button should be visually distinct but not obstructive (semi-transparent overlay style)
- Maintain smooth CSS transition for rotation preview
- Keep canvas-based rotation processing unchanged (only UI changes)

## Capabilities

### New Capabilities
- `overlay-rotation-button`: Semi-transparent rotation button positioned as an overlay on the photo preview that rotates image 90° clockwise per click

### Modified Capabilities
- `photo-rotation-control`: Update rotation control UI from two buttons below photo (left/right) to single overlay button (clockwise only), removing counter-clockwise rotation capability

## Impact

**Frontend Code:**
- [PhotoCapture.svelte](../../frontend/src/lib/features/PhotoCapture.svelte) - Remove "Links" button, convert "Rechts" button to overlay, update styling
- [rotation.ts](../../frontend/src/lib/shared/utils/rotation.ts) - Remove `rotateLeft()` function if no longer used
- Component tests for PhotoCapture - Update to test single clockwise rotation only

**APIs:**
- No backend API changes (rotation processing remains canvas-based before upload)

**Dependencies:**
- No new dependencies required

**User Experience:**
- Simpler, more focused UI with rotation control integrated into photo preview
- Fewer clicks needed (cycling through all rotations requires max 4 clicks vs potentially reversing with left button)
- **Breaking change for users**: No counter-clockwise rotation (users must click 3 times to achieve 270° rotation instead of 1 left click)
