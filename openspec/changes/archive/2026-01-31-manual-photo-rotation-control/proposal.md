## Why

The current automatic photo rotation logic uses device orientation sensors (beta/gamma angles) to detect phone orientation and automatically rotate photos. This approach has become complex, unreliable across different mobile browsers, and doesn't work consistently on iOS Safari and Android Chrome due to varying sensor implementations and permissions. Users need a simpler, more reliable way to capture and orient their photos correctly.

## What Changes

- **Remove automatic rotation logic**: Delete `calculateInitialRotation()` function and all device orientation sensor-based rotation detection from `rotation.ts`
- **Remove OrientationAdapter integration**: Remove orientation monitoring, permission requests, and beta/gamma tracking used solely for rotation detection from the capture flow
- **Build new PhotoCapture component**: Create a reusable Svelte component that uses native `<input type="file" capture="camera">` for photo capture with manual rotation controls
- **Implement canvas-based rotation**: Build rotation logic using HTML Canvas API with proper coordinate translation to "bake" rotation into the final image blob
- **Simple initial heuristic**: Use basic `screen.orientation.angle` or `window.orientation` for initial rotation guess (best-effort, no complex logic)
- **Keep manual rotation buttons**: Preserve the ability for users to manually correct photo orientation with rotate-left/rotate-right controls
- **Memory management**: Add explicit cleanup with `URL.revokeObjectURL()` to prevent memory leaks

The existing `rotateLeft()` and `rotateRight()` helper functions will be retained and reused.

## Capabilities

### New Capabilities
- `manual-photo-capture`: Mobile-friendly photo capture component with preview, manual rotation controls, and canvas-based rotation processing
- `canvas-rotation-transform`: Canvas coordinate math for rotating images (handling dimension swaps, translation, and centering)

### Modified Capabilities
- `photo-upload`: Simplified photo upload flow that no longer relies on complex orientation sensor data, only on user-corrected rotation state

## Impact

**Code affected:**
- `frontend/src/lib/shared/utils/rotation.ts` - Remove `calculateInitialRotation()`, keep `rotateLeft()` and `rotateRight()`
- `frontend/src/lib/adapters/orientation/OrientationAdapter.ts` - Remove beta/gamma capture methods if only used for rotation
- `frontend/src/lib/shared/stores/appStore.ts` - Simplify or remove `orientationData` store if only used for rotation
- `frontend/src/routes/+page.svelte` - Replace orientation monitoring and capture logic with new PhotoCapture component
- `frontend/src/lib/adapters/camera/CameraAdapter.ts` - May be fully replaced by new component using file input API

**Dependencies:**
- No new external libraries (uses native Web APIs: Canvas, File API)

**Browser compatibility:**
- Must work on iOS Safari 13+ and Android Chrome (current support targets)

**UX impact:**
- Users will have simpler, more predictable photo capture experience
- No permission prompts for orientation sensors (only camera)
- Users must manually rotate if initial heuristic is wrong (acceptable tradeoff for reliability)

## Implementation Outcome

**Actual changes exceeded original scope:**

During implementation, it became clear that the spirit level feature no longer worked with the native file input approach (users are in the OS camera app, not viewing our SPA during capture). This led to removing the entire orientation system:

- ✅ OrientationAdapter - **FULLY REMOVED** (not just beta/gamma methods)
- ✅ SpiritLevel.svelte - **REMOVED** (didn't display during native camera)
- ✅ AccessibilityToggle.svelte - **REMOVED** (no longer needed)
- ✅ CameraAdapter - **FULLY REMOVED** (replaced by file input)
- ✅ CameraPreview.svelte - **REMOVED** (replaced by PhotoCapture)
- ✅ PhotoConfirmation.svelte - **REMOVED** (merged into PhotoCapture)
- ✅ FileUploadFallback.svelte - **REMOVED** (unused)
- ✅ All orientation stores - **REMOVED** from appStore.ts
- ✅ OrientationData type - **REMOVED**
- ✅ IOrientationPort, ICameraPort - **REMOVED**

**Result:** ~1,500+ lines of code removed, significantly simpler architecture.
