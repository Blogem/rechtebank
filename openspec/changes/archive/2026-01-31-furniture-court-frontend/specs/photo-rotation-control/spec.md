# Photo Rotation Control

## Context

When users capture photos with their mobile device camera or upload existing photos, the image orientation can be incorrect from Gemini's perspective. This happens because:

1. **Camera captures**: Video stream orientation doesn't always match device physical orientation
2. **File uploads**: EXIF orientation metadata is stripped when converting to JPEG via canvas
3. **Result**: Gemini receives sideways or upside-down images, reducing analysis accuracy

The spirit level feature ensures photos are taken level (not tilted), but doesn't address rotational orientation (portrait vs landscape alignment).

## Solution Approach

Rather than trying to auto-detect orientation (which is device-specific and unreliable), we:
- Default to 0° rotation (no rotation)
- Show the captured/selected photo to the user
- Provide manual rotation controls (rotate left/right)
- Apply the rotation before uploading to backend

This gives users direct control and visual confirmation.

## ADDED Requirements

### Requirement: Display photo with rotation preview
The system SHALL display the captured or selected photo with the current rotation applied as a CSS transform.

#### Scenario: Photo displayed with no rotation
- **GIVEN** user has captured a photo
- **WHEN** confirmation screen is shown
- **THEN** photo is displayed with 0° rotation by default

#### Scenario: Photo displayed with rotation applied
- **GIVEN** user has set rotation to 90°
- **WHEN** confirmation screen is shown
- **THEN** photo is displayed rotated 90° clockwise
- **AND** rotation is applied smoothly with CSS transition

### Requirement: Provide rotation controls
The system SHALL provide buttons to rotate the photo left (counter-clockwise) or right (clockwise).

#### Scenario: Rotate right button increases rotation
- **GIVEN** photo rotation is 0°
- **WHEN** user clicks "Rechts" (rotate right) button
- **THEN** rotation changes to 90°
- **AND** photo visual updates immediately

#### Scenario: Rotate left button decreases rotation
- **GIVEN** photo rotation is 90°
- **WHEN** user clicks "Links" (rotate left) button
- **THEN** rotation changes to 0°
- **AND** photo visual updates immediately

#### Scenario: Rotation cycles through 360°
- **GIVEN** photo rotation is 270°
- **WHEN** user clicks "Rechts" (rotate right) button
- **THEN** rotation changes to 0°

#### Scenario: Rotation cycles backwards through 360°
- **GIVEN** photo rotation is 0°
- **WHEN** user clicks "Links" (rotate left) button
- **THEN** rotation changes to 270°

### Requirement: Store rotation state
The system SHALL maintain the photo rotation state (0°, 90°, 180°, or 270°) throughout the photo confirmation flow.

#### Scenario: Smart initial rotation for camera captures
- **GIVEN** user captures photo with camera
- **WHEN** beta angle > 45° at moment of capture
- **THEN** rotation is set to 0° (portrait photo assumed)

#### Scenario: Smart initial rotation for landscape photos
- **GIVEN** user captures photo with camera
- **WHEN** beta angle ≤ 45° at moment of capture
- **THEN** rotation is set to 90° (landscape or floor photo assumed)

#### Scenario: Default rotation for file uploads
- **GIVEN** user selects photo from file system
- **WHEN** confirmation screen is shown
- **THEN** rotation defaults to 0° (no sensor data available)

#### Scenario: Rotation persists during confirmation
- **GIVEN** rotation has been set (smart or manually)
- **WHEN** user views confirmation screen
- **THEN** rotation remains unchanged until user adjusts it

#### Scenario: Rotation resets for new photo
- **GIVEN** user has adjusted rotation for previous photo
- **WHEN** user captures a new photo
- **THEN** rotation is recalculated from new beta reading

#### Scenario: Rotation resets on retake
- **GIVEN** user has rotated photo
- **WHEN** user clicks "Opnieuw" (retake) button
- **THEN** rotation will be recalculated when new photo is captured

### Requirement: Apply rotation before upload
The system SHALL apply the rotation transformation to the image before uploading to the backend.

#### Scenario: Upload photo with no rotation
- **GIVEN** photo rotation is 0°
- **WHEN** user confirms photo for upload
- **THEN** original image orientation is sent to backend

#### Scenario: Upload photo with 90° rotation
- **GIVEN** photo rotation is 90°
- **WHEN** user confirms photo for upload
- **THEN** image pixels are rotated 90° clockwise using canvas transformation
- **AND** canvas dimensions are swapped (width ↔ height)
- **AND** rotated JPEG is sent to backend

#### Scenario: Upload photo with 180° rotation
- **GIVEN** photo rotation is 180°
- **WHEN** user confirms photo for upload
- **THEN** image pixels are rotated 180° using canvas transformation
- **AND** canvas dimensions remain the same
- **AND** rotated JPEG is sent to backend

#### Scenario: Upload photo with 270° rotation
- **GIVEN** photo rotation is 270°
- **WHEN** user confirms photo for upload
- **THEN** image pixels are rotated 270° clockwise using canvas transformation
- **AND** canvas dimensions are swapped (width ↔ height)
- **AND** rotated JPEG is sent to backend

### Requirement: Canvas rotation transformation
The system SHALL use HTML5 canvas transformations to apply rotation to image pixels before JPEG encoding.

#### Scenario: 90° and 270° rotations swap dimensions
- **GIVEN** original image is 1080×1920 pixels (portrait)
- **WHEN** rotation is 90° or 270°
- **THEN** canvas is created as 1920×1080 pixels (landscape)
- **AND** image is drawn with rotation transform

#### Scenario: 0° and 180° rotations keep dimensions
- **GIVEN** original image is 1080×1920 pixels
- **WHEN** rotation is 0° or 180°
- **THEN** canvas is created as 1080×1920 pixels
- **AND** image is drawn with rotation transform

#### Scenario: Rotation transform calculation
- **GIVEN** rotation angle in degrees
- **WHEN** applying canvas transformation
- **THEN** canvas is translated to center point
- **AND** rotation is applied in radians (degrees × π / 180)
- **AND** image is drawn centered at origin

## User Interface

### PhotoConfirmation Component Updates

**Visual Elements:**
- Photo preview with CSS `transform: rotate({rotation}deg)`
- Rotation controls section with hint text
- Two rotation buttons (left/right)

**Button Labels:**
- "↶ Links" - Rotate counter-clockwise (-90°)
- "↷ Rechts" - Rotate clockwise (+90°)

**Hint Text:**
- "Staat de foto niet goed? Roteer hem eerst:"

**Styling:**
- Rotation controls in light gray background (#f8f9fa)
- Smooth transition (0.3s ease) for rotation visual feedback
- Preview area with flex centering to handle dimension changes

## Technical Architecture

### State Management

**New Store:**
```typescript
export const photoRotation = writable<number>(0); // 0, 90, 180, 270
```

**Reset Points:**
- On new photo capture
- On photo retake
- On app state reset

### API Changes

**IApiPort Interface:**
```typescript
uploadPhoto(photo: Blob, metadata: PhotoMetadata, rotation?: number): Promise<Verdict>
```

**ApiAdapter Methods:**
- `uploadPhoto()` - accepts optional rotation parameter
- `applyRotation()` - private method for canvas transformation
- `convertToJPEG()` - existing method, used when rotation is 0°

### Canvas Transformation Logic

```
For rotation R:
1. Determine needsSwap = (R === 90 || R === 270)
2. Set canvas.width = needsSwap ? height : width
3. Set canvas.height = needsSwap ? width : height
4. ctx.translate(canvas.width / 2, canvas.height / 2)
5. ctx.rotate((R * Math.PI) / 180)
6. ctx.drawImage(img, -width / 2, -height / 2, width, height)
```

## Implementation Notes

### Why Manual Rotation (Not Auto-Detection)?

**Auto-detection challenges:**
1. `screen.orientation.angle` availability varies by browser
2. Camera sensor orientation is device-specific
3. Video stream rotation differs across Android/iOS
4. EXIF orientation only applies to file uploads, not camera captures
5. No reliable cross-device heuristic

**Manual rotation benefits:**
1. User sees exactly what Gemini will see
2. Works for all capture methods
3. Simple, predictable behavior
4. No device-specific bugs

### EXIF Handling

The existing `exifOrientation.ts` file is not used because:
- Camera captures don't have EXIF metadata
- Canvas re-encoding strips EXIF anyway
- User rotation replaces the need for EXIF parsing

The file can remain for potential future use with file uploads.

### Performance Considerations

- CSS rotation is instant (GPU-accelerated)
- Canvas rotation only on upload (not preview)
- No image quality loss (90° rotations are lossless pixel operations)

## Testing Strategy

### Unit Tests

**PhotoConfirmation Component:**
- Rotation prop applies CSS transform
- Rotate left emits correct event
- Rotate right emits correct event
- Rotation buttons are visible

**ApiAdapter:**
- Rotation 0° skips transformation
- Rotation 90° swaps dimensions
- Rotation 180° keeps dimensions
- Rotation 270° swaps dimensions
- Canvas transformation math is correct

**App Store:**
- photoRotation initializes to 0°
- Reset clears rotation

### Integration Tests

**Full Flow:**
1. Capture photo → rotation = 0°
2. Click rotate right → rotation = 90°, visual updates
3. Click rotate right → rotation = 180°, visual updates
4. Click rotate left → rotation = 90°, visual updates
5. Confirm → rotated image uploaded
6. Retake → rotation resets to 0°

### Manual Testing Checklist

- [ ] Portrait photo appears correct when rotated 0°, 90°, 180°, 270°
- [ ] Landscape photo appears correct when rotated 0°, 90°, 180°, 270°
- [ ] File upload (existing photo) can be rotated
- [ ] Backend/Gemini receives correctly oriented image
- [ ] Rotation resets when capturing new photo
- [ ] Rotation UI is clear and intuitive
- [ ] Smooth visual feedback on rotation change
- [ ] Works on iOS Safari
- [ ] Works on Android Chrome

## Future Enhancements

### Potential Improvements (Out of Scope)

1. **Smart initial rotation**: Detect screen.orientation.angle when available, use as starting rotation
2. **EXIF auto-rotation**: For file uploads, read EXIF and auto-set rotation
3. **Rotation gesture**: Pinch-rotate gesture support on touch devices
4. **Crop/zoom**: Allow users to crop or zoom the image
5. **Remember preference**: Store user's device orientation preference

## Risks & Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Users don't notice rotation is wrong | Med | Clear hint text, prominent buttons |
| Canvas rotation fails on old browsers | Low | Feature detection, graceful degradation |
| Rotation affects file size | Low | JPEG quality at 0.9, acceptable tradeoff |
| Confusion about which direction to rotate | Med | Use arrow symbols (↶ ↷) for clarity |

## Dependencies

- Existing: `photoRotation` store in appStore.ts
- Existing: Canvas API support (already used)
- Existing: CSS transforms (already used)
- New: Rotation handler in +page.svelte
- New: Rotation UI in PhotoConfirmation.svelte
- New: Rotation transform in ApiAdapter.ts

## Success Criteria

- ✅ User can rotate photo preview in 90° increments
- ✅ Visual feedback is immediate and smooth
- ✅ Rotated images are correctly oriented when uploaded to Gemini
- ✅ No image quality degradation from rotation
- ✅ Works for both camera captures and file uploads
- ✅ Rotation state resets appropriately
- ✅ UI is intuitive without requiring explanation
