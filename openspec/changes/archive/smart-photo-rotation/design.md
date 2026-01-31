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

### Decision 1: Smart Initial Rotation Using Beta + Gamma Readings

**Choice:** Use DeviceOrientationEvent beta AND gamma values at moment of capture to determine correct orientation

**Problem with Beta-Only Approach:**
The initial beta-only heuristic was **fundamentally flawed** because it measured phone orientation, not camera sensor orientation:
- Camera sensor is **landscape-native** in phone hardware
- When phone is vertical (beta ~90°), camera captures **sideways image**
- Browser auto-rotates `<video>` preview for display, but captured blob has no rotation metadata
- Result: Preview appears correct, but captured image is rotated incorrectly

**Root Cause Discovery:**
Testing revealed that landscape mode has **two orientations** that beta alone cannot distinguish:
- **Landscape-right** (rotated CW from portrait): Home button on left, camera top-left → Image **correct** ✓
- **Landscape-left** (rotated CCW from portrait): Home button on right, camera top-right → Image **upside-down** ✗

Both have beta ~0°, but gamma values differ:
- Landscape-right: gamma ~-90° (tilted left)
- Landscape-left: gamma ~+90° (tilted right)

**CRITICAL BUG DISCOVERED (2026-01-31):**

Implementation testing revealed the heuristic had **inverted rotation logic**. The confusion stemmed from conflating two different reference frames:

1. **Device Orientation** (what beta/gamma measure): Which way is the phone tilted?
2. **Image Orientation** (what Gemini sees): How are the pixels arranged in the captured blob?

**The Inversion:**
- Mobile camera sensors are landscape-native (horizontal)
- When phone is portrait (vertical), the camera ALREADY captures at 90° to user's view
- The `<video>` preview element auto-rotates for display, but the captured blob has NO rotation metadata
- The blob pixels are ALREADY in the correct orientation for portrait photos
- Applying ADDITIONAL 90° rotation makes images sideways (180° total error)

**Wrong Heuristic (Initial Implementation):**

```
Portrait Normal (beta > 45°):     return 90° ❌  → Makes portrait photos SIDEWAYS
Portrait Upside-Down (beta < -45°): return 270° ❌ → Makes inverted photos MORE wrong
Landscape-Right (gamma < -45°):   return 0° ❌   → Landscape photos need 90°, not 0°
Landscape-Left (gamma > 45°):     return 180° ❌ → Wrong correction angle
Ambiguous/Flat:                   return 90° ❌  → Default should be no rotation
```

**Corrected Heuristic:**

```
Portrait Normal (beta > 45°, gamma ~0°):
  Phone vertical, camera top
  → Captured blob pixels already correctly oriented
  → Rotation needed: 0° (no correction) ✅

Portrait Upside-Down (beta < -45°, gamma ~0°):
  Phone vertical inverted, camera bottom  
  → Captured blob is upside-down
  → Rotation needed: 180° ✅

Landscape-Right (|beta| ≤ 45°, gamma < -45°):
  Phone horizontal, home button left
  → Captured blob is rotated 90° from expected orientation
  → Rotation needed: 90° CW ✅

Landscape-Left (|beta| ≤ 45°, gamma > 45°):
  Phone horizontal, home button right
  → Captured blob is rotated 270° from expected orientation
  → Rotation needed: 270° CW (or -90°) ✅

Ambiguous (|beta| ≤ 45°, |gamma| ≤ 45°):
  Phone flat or nearly flat
  → Most likely portrait, pixels already correct
  → Rotation needed: 0° (no correction, user can adjust) ✅
```

**Why This Happens:**
The browser's MediaDevices API returns video frames that are ALREADY oriented for display. When you call `canvas.drawImage(videoElement, ...)`, the pixels are drawn in their display orientation, NOT their sensor-native orientation. This is different from file uploads with EXIF metadata, where `img.naturalWidth/Height` returns pre-rotation dimensions.

**Impact:**
- Portrait photos (most common) were being stored sideways
- Landscape photos were stored in various incorrect orientations
- Manual rotation controls could fix it, but default experience was broken
- Testing with real devices revealed 100% of auto-detected orientations were wrong

**Rationale:**
- Corrects the fundamental reference frame confusion
- Distinguishes between two landscape orientations using gamma
- Handles upside-down portrait cases
- Leverages existing spirit level sensor infrastructure (beta + gamma already available)
- Works for majority of real-world use cases when corrected

**Alternatives Considered:**
- Screen orientation API: Unreliable when screen rotation is locked, doesn't reflect camera sensor orientation
- Video dimension comparison: Complex, requires comparing video display vs sensor dimensions
- EXIF parsing: Doesn't help camera captures, adds complexity for partial solution
- Always default to 0°: Forces manual adjustment for all photos

**Trade-offs:** 
- Flat/floor photos (phone horizontal but level) are ambiguous and may require manual correction
- Acceptable because: (a) uncommon use case, (b) manual controls available as fallback
- Gamma sensor adds minimal complexity since it's already tracked for spirit level

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

### Decision 5: Beta + Gamma Threshold Selection

**Choice:** Use 45° as the threshold for both beta and gamma

**Beta Threshold (portrait vs landscape):**
- `|beta| > 45°` = Portrait mode (phone vertical)
- `|beta| ≤ 45°` = Landscape mode (phone horizontal)
- Midpoint between typical portrait (90°) and landscape (0°) readings
- Provides margin for natural hand wobble

**Gamma Threshold (landscape orientation):**
- `gamma > 45°` = Landscape-left (home button right, image upside-down)
- `gamma < -45°` = Landscape-right (home button left, image correct)
- `|gamma| ≤ 45°` = Ambiguous (phone flat/level)
- Ensures clear tilt detection, avoids false positives when phone is level

**Rationale:**
- Real-world testing on iOS and Android confirms clear separation
- Portrait: beta ≈ 90°, gamma ≈ 0°
- Landscape-right: beta ≈ 0°, gamma ≈ -90°
- Landscape-left: beta ≈ 0°, gamma ≈ +90°
- 45° threshold provides sufficient margin for natural hand wobble
- Consistent threshold value (45°) simplifies logic and testing

**Edge Cases:**
- Phone flat on table (beta ~0°, gamma ~0°): Defaults to 90° (portrait guess)
- Phone at exactly 45° angle: Boundary cases handled consistently
- Sensor unavailable: Defaults to 90° (most common use case)

**Future Refinement:**
- Could collect telemetry on rotation corrections to validate thresholds
- Could add hysteresis (different thresholds for entering vs exiting states)
- Could combine with screen orientation API when available and reliable

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

## Implementation Fix Required

**File:** `frontend/src/lib/shared/utils/rotation.ts`  
**Function:** `calculateInitialRotation()`

**Changes Required:**

| Line/Case | Current (Wrong) | Corrected |
|-----------|----------------|-----------|
| Portrait normal (beta > 45°) | `return 90` | `return 0` |
| Portrait upside-down (beta < -45°) | `return 270` | `return 180` |
| Landscape-left (gamma > 45°) | `return 180` | `return 270` |
| Landscape-right (gamma < -45°) | `return 0` | `return 90` |
| Ambiguous/flat | `return 90` | `return 0` |

**Explanation:**
The rotation values need to be inverted because the browser's `drawImage()` from video element already applies display orientation. The heuristic was adding ADDITIONAL rotation on top of already-correct pixels, resulting in sideways images.

**Testing Strategy:**
After fix, test with real mobile device in each orientation:
1. Portrait normal (phone vertical, camera forward) → Should store upright ✅
2. Portrait upside-down (phone inverted) → Should store upright after 180° rotation ✅
3. Landscape-right (home button left) → Should store upright after 90° rotation ✅
4. Landscape-left (home button right) → Should store upright after 270° rotation ✅
5. Flat/ambiguous → Should store upright with no rotation ✅
