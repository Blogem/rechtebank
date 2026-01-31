## Context

Photos sent to Gemini for furniture analysis are sometimes oriented incorrectly (sideways or upside-down), reducing analysis accuracy. The existing spirit level feature ensures photos are level (not tilted) but doesn't address rotational orientation.

**Current State:**
- Camera captures via MediaDevices API have no EXIF metadata
- File uploads may have EXIF metadata, but it's stripped during canvas re-encoding
- No mechanism to detect or correct photo orientation before upload
- Gemini receives incorrectly oriented images, reducing analysis quality

**Constraints:**
- Must work for both camera captures AND file uploads
- Cannot rely on EXIF metadata for camera captures
- Should leverage existing DeviceOrientationEvent infrastructure from spirit level
- Must maintain simple, predictable UX

## Goals / Non-Goals

**Goals:**
- Automatically guess correct orientation for most photos using device orientation data
- Provide manual controls for users to adjust rotation when auto-detection is wrong
- Ensure Gemini receives correctly oriented images for optimal analysis
- Maintain instant visual feedback for rotation changes

**Non-Goals:**
- Perfect auto-detection for all edge cases (floor photos, unusual angles)
- EXIF metadata parsing or preservation
- Backend image processing or rotation
- Gesture-based rotation (pinch, swipe)

## Decisions

### Decision 1: Smart Initial Rotation Using Beta Reading

**Choice:** Use DeviceOrientationEvent beta value at moment of capture to guess orientation

**Rationale:**
- Beta reading indicates phone angle: ~90° for vertical (portrait), ~0° for horizontal (landscape)
- Heuristic: If beta > 45° → rotation = 0° (portrait), else rotation = 90° (landscape)
- Works for majority of use cases without user intervention
- Leverages existing spirit level sensor infrastructure

**Alternatives Considered:**
- Screen orientation API: Unreliable, inconsistent browser support, camera sensor ≠ screen orientation
- EXIF parsing: Doesn't help camera captures, adds complexity for partial solution
- Always default to 0°: Forces manual adjustment for all landscape/floor photos

**Trade-off:** Floor photos (phone vertical, camera pointing down) will be guessed as landscape and require manual correction. This is acceptable because floor photos are uncommon.

### Decision 2: Manual Rotation Controls with Visual Preview

**Choice:** Provide rotate left/right buttons in photo confirmation screen with CSS-rotated preview

**Rationale:**
- Handles all edge cases the heuristic misses
- Users see exactly what Gemini will receive
- Simple, predictable UX with instant visual feedback
- No device-specific bugs or unreliable auto-detection

**Implementation:**
- Store rotation state (0°, 90°, 180°, 270°) in writable store
- Apply CSS `transform: rotate()` to preview image for instant feedback
- Rotate left: -90° mod 360, Rotate right: +90° mod 360
- Buttons labeled "↶ Links" and "↷ Rechts" for accessibility

**Alternatives Considered:**
- Slider/dial control: More complex, less intuitive for 90° increments
- Gesture-based rotation: Complex, harder to discover, requires touch events
- No manual controls: Unacceptable, leaves some photos incorrectly oriented

### Decision 3: Canvas Transformation Before Upload

**Choice:** Apply rotation transformation to canvas, then export as JPEG for upload

**Rationale:**
- Gemini receives pixel-perfect oriented image
- No backend changes needed
- Works identically for camera captures and file uploads
- Preserves image quality

**Implementation:**
```javascript
const needsSwap = rotation === 90 || rotation === 270;
canvas.width = needsSwap ? height : width;
canvas.height = needsSwap ? width : height;
ctx.translate(canvas.width / 2, canvas.height / 2);
ctx.rotate((rotation * Math.PI) / 180);
ctx.drawImage(img, -width / 2, -height / 2, width, height);
```

**Key Details:**
- Swap canvas dimensions for 90°/270° rotations
- Translate to center, rotate, then draw centered image
- Export as JPEG blob (no EXIF metadata needed)

**Alternatives Considered:**
- Backend rotation: Adds backend complexity, doesn't leverage frontend preview
- Send rotation metadata: Requires backend changes, Gemini API changes
- No transformation: Unacceptable, defeats the purpose

### Decision 4: State Management Approach

**Choice:** Add `photoRotation` writable store, reset on new capture

**Rationale:**
- Centralized state accessible to confirmation component and API adapter
- Reset on new capture ensures fresh heuristic calculation
- Reactive store updates preview instantly

**State Flow:**
1. Photo captured → Calculate initial rotation from beta reading
2. User views confirmation screen → Store value applied to CSS preview
3. User rotates → Store updated, preview updates reactively
4. User submits → Store value passed to canvas transformation
5. New photo captured → Store reset, recalculate from new beta reading

### Decision 5: Beta Threshold Selection

**Choice:** Use 45° as the threshold (beta > 45° = portrait, else landscape)

**Rationale:**
- Midpoint between typical portrait (90°) and landscape (0°) readings
- Real-world testing confirms clear separation:
  - Portrait photos: beta ≈ 90°
  - Landscape photos: beta ≈ 0°
  - Floor photos: beta ≈ 0° (edge case)
- Provides margin for natural hand wobble

**Future Refinement:**
- Could be adjusted based on usage data if pattern emerges
- Could add additional signals (screen orientation API when available)

## Risks / Trade-offs

**Risk:** Floor photos (phone vertical, camera down) misidentified as landscape  
**Mitigation:** Manual controls allow user to correct. Floor photos are uncommon edge case.

**Risk:** Device orientation sensor unavailable or disabled  
**Mitigation:** Default rotation = 0°, user can still manually adjust.

**Risk:** Beta reading fluctuates during capture  
**Mitigation:** Single reading at moment of capture provides stable value. Spirit level already requires stable reading.

**Risk:** User confusion about rotation controls  
**Mitigation:** Clear hint text "Staat de foto niet goed? Roteer hem eerst:", arrow symbols, Dutch labels.

**Trade-off:** Adds one extra UI element to confirmation screen  
**Acceptance:** Visual confirmation and control outweighs slight UI complexity.

**Trade-off:** Canvas transformation adds processing time  
**Acceptance:** Only applied when rotation ≠ 0°, typically <100ms, imperceptible to user.
