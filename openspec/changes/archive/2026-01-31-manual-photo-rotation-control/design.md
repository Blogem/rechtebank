## Context

The current photo capture system uses a complex automatic rotation detection system that relies on:
- DeviceOrientationEvent API (beta/gamma sensor readings)
- Permission management for iOS 13+ devices
- Complex heuristics in `calculateInitialRotation()` to map sensor angles to rotation values
- OrientationAdapter for monitoring and capturing sensor data at photo capture time

This has proven unreliable due to:
- Inconsistent sensor availability across browsers
- Permission friction on iOS
- Complex edge cases in angle-to-rotation mapping
- Maintenance burden of sensor-based logic

The new approach simplifies to:
- Native file input with camera capture
- Simple screen orientation API for initial guess
- User-driven manual rotation controls
- Canvas-based "baking" of rotation into final image

## Goals / Non-Goals

**Goals:**
- Replace automatic rotation detection with simple heuristic + manual control
- Remove OrientationAdapter dependency for rotation purposes
- Build reusable PhotoCapture Svelte component with canvas rotation
- Maintain manual rotation buttons (rotate left/right)
- Properly handle canvas coordinate transformations (no "disappearing" images)
- Implement memory-safe image handling with URL cleanup
- Support iOS Safari 13+ and Android Chrome

**Non-Goals:**
- EXIF parsing or preservation (explicitly avoiding)
- AI/ML-based orientation detection
- Maintaining the complex automatic rotation logic
- Supporting older browsers (focusing on modern mobile browsers)
- Removing the spirit level feature (that's separate from rotation logic)

## Decisions

### Decision 1: Use `<input type="file" capture="camera">` instead of MediaStream API

**Rationale:**
- Simpler API with built-in mobile camera integration
- No need to manage MediaStream lifecycle
- Native browser handling of camera permissions
- Better mobile UX (native camera app integration on some devices)

**Alternatives considered:**
- Keep MediaStream API (`getUserMedia`): More control but adds complexity we don't need
- Hybrid approach: Too confusing, introduces conditional logic

**Trade-off:** Less control over camera settings (resolution, facing mode), but acceptable for this use case.

---

### Decision 2: Simple heuristic using `screen.orientation.angle` or `window.orientation`

**Rationale:**
- Much simpler than beta/gamma sensor calculations
- Directly reflects screen orientation (0, 90, 180, 270)
- No permissions required
- Fallback to 0° if unavailable (user can correct manually)

**Alternatives considered:**
- Parse EXIF data: Adds dependency, not reliable (many phones don't set it)
- No initial rotation: Could work, but slight UX improvement for free
- Keep DeviceOrientationEvent: Rejected due to complexity and unreliability

**Implementation:**
```typescript
function getInitialRotation(): number {
  if (screen?.orientation?.angle !== undefined) {
    return screen.orientation.angle;
  }
  if (window.orientation !== undefined) {
    // iOS fallback: window.orientation returns -90, 0, 90, 180
    const angle = window.orientation;
    return angle < 0 ? 360 + angle : angle;
  }
  return 0; // Fallback: no rotation
}
```

---

### Decision 3: Canvas coordinate transformation strategy

**Problem:** When rotating 90° or 270°, canvas width/height must swap, and image must be translated to stay visible.

**Solution:**
```typescript
function rotateImageOnCanvas(
  img: HTMLImageElement, 
  rotation: number
): HTMLCanvasElement {
  const canvas = document.createElement('canvas');
  const ctx = canvas.getContext('2d')!;
  
  const isPortrait = rotation === 90 || rotation === 270;
  
  // Swap dimensions for portrait rotations
  canvas.width = isPortrait ? img.height : img.width;
  canvas.height = isPortrait ? img.width : img.height;
  
  // Move origin to center, rotate, move back
  ctx.translate(canvas.width / 2, canvas.height / 2);
  ctx.rotate((rotation * Math.PI) / 180);
  ctx.drawImage(img, -img.width / 2, -img.height / 2);
  
  return canvas;
}
```

**Rationale:**
- Translate to center before rotation prevents image from disappearing
- Dimension swap ensures canvas matches final image orientation
- Drawing from center handles all rotation cases uniformly

**Alternatives considered:**
- Rotate without translation: Image disappears off-canvas
- CSS transforms only: Doesn't modify the actual image data
- Pre-calculate offsets per rotation: More complex, error-prone

---

### Decision 4: Component architecture - self-contained PhotoCapture component

**Structure:**
```
PhotoCapture.svelte
├── File input (hidden, triggered by button)
├── Image preview (<img> element)
├── Rotation controls (left/right buttons)
└── Confirm/Retake actions
```

**Rationale:**
- Encapsulates all photo capture + rotation logic
- Reusable across the app
- Clear separation of concerns
- Testable in isolation

**Props:**
```typescript
{
  onPhotoConfirmed: (blob: Blob, rotation: number) => void;
  onCancelled?: () => void;
}
```

**Alternatives considered:**
- Split into PhotoCapture + PhotoRotate: Over-engineering for this use case
- Keep logic in +page.svelte: Violates reusability and testability

---

### Decision 5: Memory management with URL lifecycle tracking

**Strategy:**
- Create object URL when file selected: `URL.createObjectURL(file)`
- Revoke on new photo selected
- Revoke on component unmount
- Revoke after blob exported (optional, depends on usage)

**Implementation pattern:**
```typescript
let currentObjectURL: string | null = null;

function handleFileSelected(file: File) {
  // Clean up previous URL
  if (currentObjectURL) {
    URL.revokeObjectURL(currentObjectURL);
  }
  currentObjectURL = URL.createObjectURL(file);
}

onDestroy(() => {
  if (currentObjectURL) {
    URL.revokeObjectURL(currentObjectURL);
  }
});
```

**Rationale:**
- Prevents memory leaks from orphaned Blob URLs
- Explicit lifecycle management
- Browser can garbage collect image data

---

### Decision 6: Remove OrientationAdapter entirely

**Decision (UPDATED):** Remove OrientationAdapter AND spirit level feature completely.

**Original plan:** Keep OrientationAdapter for spirit level, remove only beta/gamma methods.

**Actual implementation:** During development, discovered that spirit level doesn't display during native camera capture (user is in OS camera app, not viewing SPA). Since the spirit level couldn't function with the new approach, the entire orientation system became dead code.

**What was removed:**
- OrientationAdapter.ts - **DELETED** (entire adapter)
- SpiritLevel.svelte - **DELETED** (component)
- AccessibilityToggle.svelte - **DELETED** (no longer needed)
- CameraAdapter.ts - **DELETED** (replaced by file input)
- CameraPreview.svelte - **DELETED** (replaced by PhotoCapture)
- PhotoConfirmation.svelte - **DELETED** (merged into PhotoCapture)
- FileUploadFallback.svelte - **DELETED** (unused)
- All orientation stores - **REMOVED** from appStore.ts
- OrientationData type - **DELETED**
- IOrientationPort, ICameraPort - **DELETED**
- DeviceOrientationEvent integration - **REMOVED**
- Empty adapter directories - **DELETED**

**Rationale for full removal:**
- Spirit level invisible during native file input camera capture
- No other features used orientation data
- Simpler architecture without sensor permissions
- Reduced code complexity (~1,500+ lines removed)
- No other features used orientation data
- Simpler architecture without sensor permissions
- Reduced code complexity (~1,500+ lines removed)

## Risks / Trade-offs

**[Risk] Initial rotation heuristic may be wrong** → Mitigation: Manual rotation buttons always available, user can correct in 1-2 taps

**[Risk] Canvas rotation quality on high-res images** → Mitigation: Use default canvas smoothing, monitor performance, may need to add image scaling if issues arise

**[Risk] Large file sizes from high-resolution cameras** → Mitigation: Canvas `toBlob()` with quality parameter (0.85-0.9), consider max dimension cap if needed

**[Risk] Browser compatibility with screen.orientation API** → Mitigation: Fallback to window.orientation, then fallback to 0° (user corrects manually)

**[Risk] File input doesn't work on older browsers** → Mitigation: Acceptable, targeting modern mobile browsers (iOS 13+, Chrome 90+)

**[Trade-off] Less control over camera** → Acceptable: Native file input is simpler, reliability > fine-grained control

**[Trade-off] Users must manually rotate if heuristic wrong** → Acceptable: Predictable UX, user has full control

## Migration Plan

**Phase 1: Build new component** (parallel to existing code)
1. Create PhotoCapture.svelte with canvas rotation
2. Add tests for canvas transformation logic
3. Test on iOS Safari and Android Chrome

**Phase 2: Integrate and replace**
1. Update +page.svelte to use PhotoCapture component
2. Remove OrientationAdapter integration from capture flow
3. Remove `calculateInitialRotation()` from rotation.ts
4. Keep `rotateLeft()` and `rotateRight()` helpers

**Phase 3: Cleanup**
1. Remove unused `orientationData` store if only used for rotation
2. Remove beta/gamma capture methods from OrientationAdapter
3. Update tests to reflect new behavior

**Rollback strategy:**
- New component is isolated, can be feature-flagged if needed
- Old code remains until new component fully integrated
- Can revert by switching component reference in +page.svelte

**Testing checklist:**
- [ ] Rotation transformations correct for all angles (0, 90, 180, 270)
- [ ] Image stays centered and visible after rotation
- [ ] Memory cleanup verified (no Blob URL leaks)
- [ ] Works on iOS Safari (test on real device)
- [ ] Works on Android Chrome (test on real device)
- [ ] Manual rotation buttons respond correctly
- [ ] File input triggers camera on mobile

## Open Questions

- **Q: Should we cap maximum image resolution before canvas processing?**
  - If yes, what dimensions? (e.g., 2048x2048 max)
  - Decision deferred: Implement first, add cap only if performance issues arise

- **Q: Should we preserve original file if rotation is 0°?**
  - Avoids re-encoding if no rotation needed
  - Decision deferred: Start with always processing for consistency

- **Q: Do we need to handle non-image file types gracefully?**
  - File input has `accept="image/*"` but users can bypass
  - Decision: Add MIME type validation before processing
