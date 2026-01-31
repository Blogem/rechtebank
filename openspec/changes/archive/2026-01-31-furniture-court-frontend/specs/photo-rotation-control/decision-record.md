# Decision Record: Photo Rotation Approach

**Date**: January 31, 2026  
**Status**: Accepted  
**Context**: Explore mode discussion on image orientation for Gemini AI analysis

## Problem Statement

Photos sent to Gemini for furniture analysis are sometimes oriented incorrectly (sideways or upside-down), reducing analysis accuracy. The spirit level feature ensures photos are taken level (not tilted), but doesn't address rotational orientation.

## Investigation

### Root Cause Analysis

We traced the image flow from capture to Gemini:

```
📱 Camera/File → 💻 Browser → 🔙 Backend → 🤖 Gemini
```

**Camera Capture Path:**
- Video stream captured via MediaDevices API
- Canvas renders frame from video
- No EXIF metadata present
- Orientation depends on device sensor, not screen orientation

**File Upload Path:**
- Existing photo may have EXIF Orientation tag (1-8)
- Browser displays correctly (auto-applies EXIF)
- Canvas re-encoding strips EXIF metadata
- Result: Pixels rotated, but metadata lost

**Bug in ApiAdapter.convertToJPEG():**
```typescript
// ❌ WRONG: naturalWidth/Height ignore EXIF rotation
canvas.width = img.naturalWidth;  // Returns pre-rotation dimensions!
canvas.height = img.naturalHeight;

// Browser draws ROTATED image into WRONG-sized canvas
ctx.drawImage(img, 0, 0);
```

Example:
- Photo: 1080×1920 pixels, EXIF Orientation: 6 (rotate 90° CW)
- Browser displays: ✅ Correct (auto-rotates)
- Canvas dimensions: 1080×1920 (should be 1920×1080 for rotated image)
- Result: Image squished/cropped incorrectly

### Attempted Solutions

**Attempt 1: EXIF Orientation Parsing**
- Created `exifOrientation.ts` with `getExifOrientation()` and `correctImageOrientation()`
- Problem: Never called! Code exists but unused
- Deeper problem: Still has the same `naturalWidth/Height` bug
- Further problem: Doesn't help camera captures (no EXIF)

**Attempt 2: Auto-Detect Screen Orientation**
- Consider using `screen.orientation.angle`
- Problem: API availability varies by browser
- Problem: Camera sensor orientation ≠ screen orientation
- Problem: Device-specific, no reliable heuristic

**Attempt 3: Use Spirit Level Data for Smart Initial Rotation** ✅
- Use DeviceOrientationEvent beta reading at moment of capture to guess orientation
- **Real-world testing shows**:
  - Phone VERTICAL (portrait), camera pointing forward → beta ≈ 90°
  - Phone HORIZONTAL (landscape), camera pointing forward → beta ≈ 0°
  - Phone pointing at FLOOR (unusual) → beta ≈ 0°
- **Smart heuristic**:
  - If beta > 45° when capturing → Assume portrait, set rotation = 0°
  - If beta ≤ 45° when capturing → Assume landscape OR floor photo, set rotation = 90°
  - User can manually adjust if guess is wrong (e.g., floor photos)
- **Advantages**:
  - Most photos will have correct rotation automatically
  - Reduces manual rotation needed
  - Still provides manual controls as fallback
- **Decision**: Use this as initial rotation guess, keep manual controls for edge cases

## Options Considered

### Option A: Fix EXIF Handling (Backend)
Apply rotation on the Go backend before sending to Gemini.

**Pros:**
- Frontend stays simple
- Handles all upload sources

**Cons:**
- Backend complexity
- Needs EXIF parsing library
- Doesn't help camera captures (no EXIF to parse)

**Verdict:** ❌ Rejected - Doesn't solve camera capture case

### Option B: Fix EXIF Handling (Frontend)
Read EXIF, calculate correct canvas dimensions, apply transforms.

**Pros:**
- Gemini receives correct pixels
- No backend changes needed

**Cons:**
- Complex transformation matrix for 8 EXIF orientations
- Still doesn't help camera captures
- Need to correctly swap width/height for 90°/270° rotations
- ~100 lines of error-prone code

**Verdict:** ❌ Rejected - Partial solution, high complexity

### Option C: Use Library (blueimp-load-image)
Battle-tested library handles EXIF orientation automatically.

**Pros:**
- Reliable, handles all 8 EXIF cases
- Minimal code to integrate

**Cons:**
- Still doesn't help camera captures (no EXIF)
- Adds dependency (~6KB)
- Doesn't give user control

**Verdict:** ❌ Rejected - Incomplete solution for camera path

### Option D: Manual Rotation Controls ✅ CHOSEN
Let users preview and rotate photos before upload.

**Pros:**
- Works for BOTH camera captures AND file uploads
- User sees exactly what Gemini will see
- Simple, predictable UX
- No device-specific bugs
- No EXIF complexity needed

**Cons:**
- Requires user action (but provides visual confirmation)
- Adds UI controls

**Verdict:** ✅ **Selected** - Complete solution, best UX

## Decision

**We will implement smart initial rotation with manual controls** with the following design:

1. **Smart initial rotation** based on spirit level reading at capture:
   - If beta > 45° → rotation = 0° (portrait photo)
   - If beta ≤ 45° → rotation = 90° (landscape photo or floor photo)
2. **Show photo preview** in confirmation screen with applied rotation
3. **Provide rotation buttons**: "↶ Links" and "↷ Rechts" for manual adjustment
4. **CSS rotation** for instant visual feedback
5. **Canvas transformation** applied before upload
6. **Reset rotation** when capturing new photo (recalculate from beta)

### Rotation Implementation

**State:**
- Store: `photoRotation: writable<number>(0)` // 0, 90, 180, 270

**UI:**
```svelte
<img style="transform: rotate({rotation}deg);" />
<button onclick={rotateLeft}>↶ Links</button>
<button onclick={rotateRight}>↷ Rechts</button>
```

**Canvas Transform:**
```javascript
const needsSwap = rotation === 90 || rotation === 270;
canvas.width = needsSwap ? height : width;
canvas.height = needsSwap ? width : height;
ctx.translate(canvas.width / 2, canvas.height / 2);
ctx.rotate((rotation * Math.PI) / 180);
ctx.drawImage(img, -width / 2, -height / 2, width, height);
```

## Consequences

### Positive

✅ **Complete solution**: Works for camera captures AND file uploads  
✅ **User confidence**: Visual confirmation of what Gemini sees  
✅ **Device-agnostic**: No special handling per platform  
✅ **Simple code**: No EXIF parsing, no complex heuristics  
✅ **Accessible**: Clear UI with arrow symbols and text labels

### Negative

⚠️ **Requires user action**: Auto-detection would be "nicer" if it worked reliably  
⚠️ **Extra step**: Adds one more action to photo submission flow

### Mitigations

- Clear hint text: "Staat de foto niet goed? Roteer hem eerst:"
- Smart initial rotation using beta reading (most photos correct automatically)
- Rotation is quick (instant visual feedback)
- Users can skip if photo looks correct
- Manual controls handle edge cases (floor photos, unusual angles)

## Future Enhancements
Possible improvements based on usage patterns:

1. **Refine beta threshold**: Adjust 45° threshold based on real usage data
2. **EXIF auto-detection**: For file uploads, parse EXIF and pre-set rotation
3. **Screen orientation API**: Use `screen.orientation.angle` as additional signal when available
4. **Remember preference**: Store user's typical rotation per device
5. **Gesture support**: Pinch-rotate on touch devices
But start simple: default 0°, let user rotate if needed.

## Lessons Learned

1. **Auto-detection is hard**: Device orientation, screen orientation, camera sensor orientation, and EXIF orientation are all different things
2. **User control is reliable**: When automation is unreliable, give users the controls
3. **Partial solutions don't help**: EXIF handling only solves file uploads, not camera captures
4. **Existing code may be incomplete**: `exifOrientation.ts` existed but was never used

## References

- Spec: `specs/photo-rotation-control/spec.md`
- Design Decision: `design.md` - Decision #6
- Implementation: `specs/photo-rotation-control/implementation-summary.md`
- Original exploration: This conversation (explore mode)
