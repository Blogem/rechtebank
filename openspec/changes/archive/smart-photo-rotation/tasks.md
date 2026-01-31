## 1. State Management Setup (TDD)

- [x] 1.1 Write test: photoRotation store initializes to 0°
- [x] 1.2 Run test, verify it fails (store doesn't exist yet)
- [x] 1.3 Create `photoRotation` writable store with default value 0°
- [x] 1.4 Run test, verify it passes
- [x] 1.5 Write test: photoRotation store can be updated to 90°, 180°, 270°
- [x] 1.6 Run test, verify it fails
- [x] 1.7 Implement rotation value updates in store
- [x] 1.8 Run test, verify it passes
- [x] 1.9 Write test: reset() function clears rotation to 0°
- [x] 1.10 Run test, verify it fails
- [x] 1.11 Implement reset function
- [x] 1.12 Run test, verify it passes

## 2. Device Orientation Detection (TDD)

- [x] 2.1 Write test: getBetaAtCapture() returns current beta value
- [x] 2.2 Run test, verify it fails (method doesn't exist)
- [x] 2.3 Extend orientation adapter to add getBetaAtCapture() method
- [x] 2.4 Run test, verify it passes
- [x] 2.5 Write test: calculateInitialRotation(beta > 45°) returns 0°
- [x] 2.6 Run test, verify it fails
- [x] 2.7 Implement calculateInitialRotation with 45° threshold logic
- [x] 2.8 Run test, verify it passes
- [x] 2.9 Write test: calculateInitialRotation(beta ≤ 45°) returns 90°
- [x] 2.10 Run test, verify it passes (should work with implementation from 2.7)
- [x] 2.11 Write test: calculateInitialRotation(null) returns 0° (sensor unavailable)
- [x] 2.12 Run test, verify it fails
- [x] 2.13 Add null handling to calculateInitialRotation
- [x] 2.14 Run test, verify it passes
- [x] 2.15 Write test: calculateInitialRotation(undefined) returns 0°
- [x] 2.16 Run test, verify it passes (should work with null handling)
- [x] 2.17 Add development mode logging for beta values and calculated rotation
- [ ] 2.18 Manually verify logging output in browser console
- [x] 2.19 Write test: calculateInitialRotation(beta < -45°) returns 180° (upside-down)
- [x] 2.20 Run test, verify it fails (upside-down detection not implemented)
- [x] 2.21 Update calculateInitialRotation to detect portrait upside-down using beta sign
- [x] 2.22 Run test, verify it passes
- [x] 2.23 Write test: calculateInitialRotation(-90°) returns 180°
- [x] 2.24 Run test, verify it passes
- [x] 2.25 Write test: calculateInitialRotation(-80°) returns 180°
- [x] 2.26 Run test, verify it passes
- [x] 2.27 Write test: calculateInitialRotation(-45°) returns 90° (landscape boundary)
- [x] 2.28 Run test, verify it passes
- [x] 2.29 Update development mode logging to show orientation type (portrait normal/upside-down/landscape)
- [ ] 2.30 Manually verify enhanced logging in browser console shows orientation types

## 3. Photo Confirmation UI Controls (TDD)

- [x] 3.1 Write test: PhotoConfirmation renders "↶ Links" button
- [x] 3.2 Run test, verify it fails (button doesn't exist)
- [x] 3.3 Add "↶ Links" button to PhotoConfirmation component
- [x] 3.4 Run test, verify it passes
- [x] 3.5 Write test: PhotoConfirmation renders "↷ Rechts" button
- [x] 3.6 Run test, verify it passes (should work with button addition)
- [x] 3.7 Write test: PhotoConfirmation renders hint text "Staat de foto niet goed? Roteer hem eerst:"
- [x] 3.8 Run test, verify it fails
- [x] 3.9 Add hint text to PhotoConfirmation component
- [x] 3.10 Run test, verify it passes
- [x] 3.11 Write test: clicking "↶ Links" calls rotateLeft handler
- [x] 3.12 Run test, verify it fails (handler not connected)
- [x] 3.13 Connect rotateLeft handler to button click event
- [x] 3.14 Run test, verify it passes
- [x] 3.15 Write test: clicking "↷ Rechts" calls rotateRight handler
- [x] 3.16 Run test, verify it fails (handler not connected)
- [x] 3.17 Connect rotateRight handler to button click event
- [x] 3.18 Run test, verify it passes
- [x] 3.19 Style rotation buttons (accessibility, hover states, touch targets)
- [ ] 3.20 Manually verify button appearance and touch targets

## 4. Rotation Logic Implementation (TDD)

- [x] 4.1 Write test: rotateLeft from 0° returns 270°
- [x] 4.2 Run test, verify it fails (function doesn't exist)
- [x] 4.3 Implement rotateLeft function (subtract 90° mod 360)
- [x] 4.4 Run test, verify it passes
- [x] 4.5 Write test: rotateLeft from 90° returns 0°
- [x] 4.6 Run test, verify it passes
- [x] 4.7 Write test: rotateLeft cycles correctly (270° → 180° → 90° → 0° → 270°)
- [x] 4.8 Run test, verify it passes
- [x] 4.9 Write test: rotateRight from 0° returns 90°
- [x] 4.10 Run test, verify it fails (function doesn't exist)
- [x] 4.11 Implement rotateRight function (add 90° mod 360)
- [x] 4.12 Run test, verify it passes
- [x] 4.13 Write test: rotateRight from 270° returns 0°
- [x] 4.14 Run test, verify it passes
- [x] 4.15 Write test: rotateRight cycles correctly (0° → 90° → 180° → 270° → 0°)
- [x] 4.16 Run test, verify it passes
- [x] 4.17 Write test: rotateLeft updates photoRotation store
- [x] 4.18 Run test, verify it fails (store not connected)
- [x] 4.19 Connect rotateLeft to photoRotation store
- [x] 4.20 Run test, verify it passes
- [x] 4.21 Write test: rotateRight updates photoRotation store
- [x] 4.22 Run test, verify it fails (store not connected)
- [x] 4.23 Connect rotateRight to photoRotation store
- [x] 4.24 Run test, verify it passes

## 5. Visual Preview Implementation (TDD)

- [x] 5.1 Write test: preview image has transform style when rotation is 90°
- [x] 5.2 Run test, verify it fails (transform not applied)
- [x] 5.3 Apply CSS transform to photo preview based on rotation store value
- [x] 5.4 Run test, verify it passes
- [x] 5.5 Write test: transform updates when rotation changes from 0° to 180°
- [x] 5.6 Run test, verify it fails (reactive update not working)
- [x] 5.7 Ensure transform updates reactively when rotation store changes
- [x] 5.8 Run test, verify it passes
- [x] 5.9 Write test: transform is "rotate(0deg)" when rotation is 0°
- [x] 5.10 Run test, verify it passes
- [x] 5.11 Write test: transform is "rotate(270deg)" when rotation is 270°
- [x] 5.12 Run test, verify it passes
- [ ] 5.13 Manually verify preview display at all rotation angles (0°, 90°, 180°, 270°)
- [ ] 5.14 Manually verify preview matches expected final uploaded orientation

## 6. Canvas Rotation Transformation (TDD)

- [x] 6.1 Write test: canvas dimensions are not swapped for 0° rotation (1080×1920 → 1080×1920)
- [x] 6.2 Run test, verify it fails (function doesn't exist)
- [x] 6.3 Create applyRotation helper function in API adapter with dimension logic
- [x] 6.4 Run test, verify it passes
- [x] 6.5 Write test: canvas dimensions are not swapped for 180° rotation (1080×1920 → 1080×1920)
- [x] 6.6 Run test, verify it passes
- [x] 6.7 Write test: canvas dimensions are swapped for 90° rotation (1080×1920 → 1920×1080)
- [x] 6.8 Run test, verify it fails (swap logic not implemented)
- [x] 6.9 Implement dimension swap logic for 90°/270° rotations
- [x] 6.10 Run test, verify it passes
- [x] 6.11 Write test: canvas dimensions are swapped for 270° rotation (1080×1920 → 1920×1080)
- [x] 6.12 Run test, verify it passes
- [x] 6.13 Write test: 90° rotation applies correct transformation matrix
- [x] 6.14 Run test, verify it fails (transformation not implemented)
- [x] 6.15 Implement canvas transformation matrix (translate, rotate, draw)
- [x] 6.16 Run test, verify it passes
- [x] 6.17 Write test: 180° rotation applies correct transformation matrix
- [x] 6.18 Run test, verify it passes
- [x] 6.19 Write test: 270° rotation applies correct transformation matrix
- [x] 6.20 Run test, verify it passes
- [x] 6.21 Write test: image is centered in canvas after rotation
- [x] 6.22 Run test, verify it passes (should work with transformation implementation)
- [x] 6.23 Write test: rotated canvas exports as JPEG blob
- [x] 6.24 Run test, verify it fails (export not implemented)
- [x] 6.25 Export rotated canvas as JPEG blob with quality setting
- [x] 6.26 Run test, verify it passes
- [x] 6.27 Write test: canvas creation failure logs error and returns original blob
- [x] 6.28 Run test, verify it fails (error handling not implemented)
- [x] 6.29 Add error handling for canvas operations with fallback to original
- [x] 6.30 Run test, verify it passes
- [x] 6.31 Write test: image decode failure logs error and returns original blob
- [x] 6.32 Run test, verify it passes (should work with error handling from 6.29)

## 7. Upload Integration (TDD)

- [x] 7.1 Write test: uploadPhoto with rotation 0° uploads original blob without transformation
- [x] 7.2 Run test, verify it fails (rotation parameter not implemented)
- [x] 7.3 Modify uploadPhoto to accept rotation parameter
- [x] 7.4 Implement skip transformation logic when rotation is 0°
- [x] 7.5 Run test, verify it passes
- [x] 7.6 Write test: uploadPhoto with rotation 90° calls applyRotation before upload
- [x] 7.7 Run test, verify it fails (transformation not called)
- [x] 7.8 Apply transformation when rotation is 90°, 180°, or 270°
- [x] 7.9 Run test, verify it passes
- [x] 7.10 Write test: uploadPhoto with rotation 180° uploads rotated blob
- [x] 7.11 Run test, verify it passes
- [x] 7.12 Write test: uploadPhoto with rotation 270° uploads rotated blob
- [x] 7.13 Run test, verify it passes
- [x] 7.14 Write test: rotation value from photoRotation store is passed to uploadPhoto
- [x] 7.15 Run test, verify it fails (store not connected to upload)
- [x] 7.16 Connect confirmation screen to pass rotation value to uploadPhoto
- [x] 7.17 Run test, verify it passes

## 8. Capture Flow Integration (TDD)

- [x] 8.1 Write test: capturing photo with beta 80° sets rotation to 0°
- [x] 8.2 Run test, verify it fails (beta capture not integrated)
- [x] 8.3 Capture beta reading when photo is taken via camera
- [x] 8.4 Calculate and set initial rotation from beta reading
- [x] 8.5 Run test, verify it passes
- [x] 8.6 Write test: capturing photo with beta 20° sets rotation to 90°
- [x] 8.7 Run test, verify it passes
- [x] 8.8 Write test: retaking photo resets rotation state and recalculates from new beta
- [x] 8.9 Run test, verify it fails (reset not called on retake)
- [x] 8.10 Reset rotation state when user retakes photo
- [x] 8.11 Run test, verify it passes
- [x] 8.12 Write test: manual rotation is preserved until new photo is captured
- [x] 8.13 Run test, verify it fails (manual changes reset prematurely)
- [x] 8.14 Ensure manual rotation persists across UI interactions (not on retake)
- [x] 8.15 Run test, verify it passes
- [x] 8.16 Write test: file upload defaults rotation to 0° (no beta reading)
- [x] 8.17 Run test, verify it fails (file upload path not handled)
- [x] 8.18 Set default rotation 0° for file upload path
- [x] 8.19 Run test, verify it passes
- [x] 8.20 Manually test full flow: camera capture → orientation detection → confirmation → rotation → upload
- [x] 8.21 Write test: capturing photo with beta -80° sets rotation to 180° (upside-down)
- [x] 8.22 Run test, verify it fails (upside-down capture not tested)
- [x] 8.23 Run test, verify it passes (should work with updated calculateInitialRotation)
- [ ] 8.24 Manual test: capture upside-down portrait (beta < -45°) → verify rotation 180° → preview correct

## 9. End-to-End Testing and Validation

- [x] 9.1 Manual test: portrait photo capture (beta > 45°) → verify rotation 0° → preview correct → upload
- [x] 9.2 Manual test: landscape photo capture (beta ≤ 45°) → verify rotation 90° → preview correct → upload
- [x] 9.3 Manual test: portrait photo → manually rotate right → verify rotation 90° → preview updates → upload
- [x] 9.4 Manual test: landscape photo → manually rotate left → verify rotation 0° → preview updates → upload
- [x] 9.5 Manual test: capture photo → rotate to 180° → retake photo → verify rotation resets and recalculates
- [x] 9.6 Manual test: file upload → verify default rotation 0° → manual rotation works → upload
- [x] 9.7 Manual test: sensor unavailable → verify default rotation 0° → manual rotation works
- [x] 9.8 Visual verification: upload rotated image → check Gemini receives correctly oriented image
- [x] 9.9 Visual verification: compare canvas transformation output with expected orientation
- [x] 9.10 Visual verification: ensure no image distortion or cropping at any rotation angle
- [ ] 9.11 Manual test: upside-down portrait → verify beta < -45° detected → rotation 180° applied
- [ ] 9.12 Manual test: upside-down portrait → rotate right → verify 270° → preview correct
- [ ] 9.13 Manual test: upside-down portrait → retake as normal portrait → verify rotation resets to 0°
- [ ] 9.14 Visual verification: upside-down photos upload correctly oriented to Gemini

## 10. Documentation and Cleanup

- [x] 10.1 Add JSDoc comments to rotation functions
- [ ] 10.2 Update component documentation with rotation feature
- [x] 10.3 Remove unused exifOrientation.ts file (if confirmed unused elsewhere)
- [ ] 10.4 Add user-facing documentation for rotation controls

## 11. Fix Rotation Logic - Add Gamma Support (CRITICAL BUG FIX)

**Context**: Current implementation has inverted logic - portrait returns 0° when it should return 90°. Also, landscape mode has two orientations that beta alone cannot distinguish. Need gamma to detect landscape-left vs landscape-right.

- [x] 11.1 Update OrientationAdapter: Add `lastGamma` property to track gamma values
- [x] 11.2 Update OrientationAdapter: Store `this.lastGamma = gamma` in orientation event handler
- [x] 11.3 Update OrientationAdapter: Add `getGammaAtCapture(): number | null` method
- [x] 11.4 Write test: getGammaAtCapture() returns current gamma value
- [x] 11.5 Run test, verify it passes
- [x] 11.6 Update calculateInitialRotation signature: add `gamma: number | null | undefined` parameter
- [x] 11.7 Write test: calculateInitialRotation(beta=90, gamma=0) returns 90° (portrait normal, INVERTED from old logic)
- [x] 11.8 Run test, verify it fails (current implementation returns 0°)
- [x] 11.9 Update calculateInitialRotation: Fix portrait normal logic (beta > 45° → return 90°)
- [x] 11.10 Run test, verify it passes
- [x] 11.11 Write test: calculateInitialRotation(beta=-90, gamma=0) returns 270° (portrait upside-down, INVERTED from old logic)
- [x] 11.12 Run test, verify it fails (current implementation returns 180°)
- [x] 11.13 Update calculateInitialRotation: Fix portrait upside-down logic (beta < -45° → return 270°)
- [x] 11.14 Run test, verify it passes
- [x] 11.15 Write test: calculateInitialRotation(beta=0, gamma=-90) returns 0° (landscape-right, home button left)
- [x] 11.16 Run test, verify it fails (gamma not used yet)
- [x] 11.17 Update calculateInitialRotation: Add landscape-right detection (gamma < -45° → return 0°)
- [x] 11.18 Run test, verify it passes
- [x] 11.19 Write test: calculateInitialRotation(beta=0, gamma=90) returns 180° (landscape-left, home button right)
- [x] 11.20 Run test, verify it fails (landscape-left not detected)
- [x] 11.21 Update calculateInitialRotation: Add landscape-left detection (gamma > 45° → return 180°)
- [x] 11.22 Run test, verify it passes
- [x] 11.23 Write test: calculateInitialRotation(beta=0, gamma=0) returns 90° (ambiguous/flat, default to portrait)
- [x] 11.24 Run test, verify it fails
- [x] 11.25 Update calculateInitialRotation: Add ambiguous case handling (|beta| ≤ 45°, |gamma| ≤ 45° → return 90°)
- [x] 11.26 Run test, verify it passes
- [x] 11.27 Write test: calculateInitialRotation(null, null) returns 90° (sensor unavailable, default to portrait)
- [x] 11.28 Run test, verify it fails (current default is 0°)
- [x] 11.29 Update calculateInitialRotation: Change default from 0° to 90° when sensors unavailable
- [x] 11.30 Run test, verify it passes
- [x] 11.31 Update +page.svelte: Get gamma value at capture using `orientationAdapter.getGammaAtCapture()`
- [x] 11.32 Update +page.svelte: Pass both beta and gamma to `calculateInitialRotation(beta, gamma)`
- [x] 11.33 Update development logging: Show both beta AND gamma values with orientation type
- [x] 11.34 Update development logging: Show calculated rotation and reasoning (e.g., "portrait normal", "landscape-left")
- [x] 11.35 Run all existing tests, verify they fail with new logic
- [x] 11.36 Update existing test expectations to match corrected rotation values
- [x] 11.37 Run all tests, verify they pass
- [ ] 11.38 Manual test: Portrait normal (beta ~90°, gamma ~0°) → verify rotation 90° → preview correct
- [ ] 11.39 Manual test: Portrait upside-down (beta ~-90°, gamma ~0°) → verify rotation 270° → preview correct
- [ ] 11.40 Manual test: Landscape-right (beta ~0°, gamma ~-90°) → verify rotation 0° → preview correct
- [ ] 11.41 Manual test: Landscape-left (beta ~0°, gamma ~90°) → verify rotation 180° → preview correct
- [ ] 11.42 Manual test: Phone flat (beta ~0°, gamma ~0°) → verify rotation 90° → can manually adjust
- [ ] 11.43 Visual verification: All orientations show correctly in preview
- [ ] 11.44 Visual verification: Gemini receives correctly oriented images for all cases
