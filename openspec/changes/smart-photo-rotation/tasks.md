## 1. State Management Setup (TDD)

- [ ] 1.1 Write test: photoRotation store initializes to 0°
- [ ] 1.2 Run test, verify it fails (store doesn't exist yet)
- [ ] 1.3 Create `photoRotation` writable store with default value 0°
- [ ] 1.4 Run test, verify it passes
- [ ] 1.5 Write test: photoRotation store can be updated to 90°, 180°, 270°
- [ ] 1.6 Run test, verify it fails
- [ ] 1.7 Implement rotation value updates in store
- [ ] 1.8 Run test, verify it passes
- [ ] 1.9 Write test: reset() function clears rotation to 0°
- [ ] 1.10 Run test, verify it fails
- [ ] 1.11 Implement reset function
- [ ] 1.12 Run test, verify it passes

## 2. Device Orientation Detection (TDD)

- [ ] 2.1 Write test: getBetaAtCapture() returns current beta value
- [ ] 2.2 Run test, verify it fails (method doesn't exist)
- [ ] 2.3 Extend orientation adapter to add getBetaAtCapture() method
- [ ] 2.4 Run test, verify it passes
- [ ] 2.5 Write test: calculateInitialRotation(beta > 45°) returns 0°
- [ ] 2.6 Run test, verify it fails
- [ ] 2.7 Implement calculateInitialRotation with 45° threshold logic
- [ ] 2.8 Run test, verify it passes
- [ ] 2.9 Write test: calculateInitialRotation(beta ≤ 45°) returns 90°
- [ ] 2.10 Run test, verify it passes (should work with implementation from 2.7)
- [ ] 2.11 Write test: calculateInitialRotation(null) returns 0° (sensor unavailable)
- [ ] 2.12 Run test, verify it fails
- [ ] 2.13 Add null handling to calculateInitialRotation
- [ ] 2.14 Run test, verify it passes
- [ ] 2.15 Write test: calculateInitialRotation(undefined) returns 0°
- [ ] 2.16 Run test, verify it passes (should work with null handling)
- [ ] 2.17 Add development mode logging for beta values and calculated rotation
- [ ] 2.18 Manually verify logging output in browser console

## 3. Photo Confirmation UI Controls (TDD)

- [ ] 3.1 Write test: PhotoConfirmation renders "↶ Links" button
- [ ] 3.2 Run test, verify it fails (button doesn't exist)
- [ ] 3.3 Add "↶ Links" button to PhotoConfirmation component
- [ ] 3.4 Run test, verify it passes
- [ ] 3.5 Write test: PhotoConfirmation renders "↷ Rechts" button
- [ ] 3.6 Run test, verify it passes (should work with button addition)
- [ ] 3.7 Write test: PhotoConfirmation renders hint text "Staat de foto niet goed? Roteer hem eerst:"
- [ ] 3.8 Run test, verify it fails
- [ ] 3.9 Add hint text to PhotoConfirmation component
- [ ] 3.10 Run test, verify it passes
- [ ] 3.11 Write test: clicking "↶ Links" calls rotateLeft handler
- [ ] 3.12 Run test, verify it fails (handler not connected)
- [ ] 3.13 Connect rotateLeft handler to button click event
- [ ] 3.14 Run test, verify it passes
- [ ] 3.15 Write test: clicking "↷ Rechts" calls rotateRight handler
- [ ] 3.16 Run test, verify it fails (handler not connected)
- [ ] 3.17 Connect rotateRight handler to button click event
- [ ] 3.18 Run test, verify it passes
- [ ] 3.19 Style rotation buttons (accessibility, hover states, touch targets)
- [ ] 3.20 Manually verify button appearance and touch targets

## 4. Rotation Logic Implementation (TDD)

- [ ] 4.1 Write test: rotateLeft from 0° returns 270°
- [ ] 4.2 Run test, verify it fails (function doesn't exist)
- [ ] 4.3 Implement rotateLeft function (subtract 90° mod 360)
- [ ] 4.4 Run test, verify it passes
- [ ] 4.5 Write test: rotateLeft from 90° returns 0°
- [ ] 4.6 Run test, verify it passes
- [ ] 4.7 Write test: rotateLeft cycles correctly (270° → 180° → 90° → 0° → 270°)
- [ ] 4.8 Run test, verify it passes
- [ ] 4.9 Write test: rotateRight from 0° returns 90°
- [ ] 4.10 Run test, verify it fails (function doesn't exist)
- [ ] 4.11 Implement rotateRight function (add 90° mod 360)
- [ ] 4.12 Run test, verify it passes
- [ ] 4.13 Write test: rotateRight from 270° returns 0°
- [ ] 4.14 Run test, verify it passes
- [ ] 4.15 Write test: rotateRight cycles correctly (0° → 90° → 180° → 270° → 0°)
- [ ] 4.16 Run test, verify it passes
- [ ] 4.17 Write test: rotateLeft updates photoRotation store
- [ ] 4.18 Run test, verify it fails (store not connected)
- [ ] 4.19 Connect rotateLeft to photoRotation store
- [ ] 4.20 Run test, verify it passes
- [ ] 4.21 Write test: rotateRight updates photoRotation store
- [ ] 4.22 Run test, verify it fails (store not connected)
- [ ] 4.23 Connect rotateRight to photoRotation store
- [ ] 4.24 Run test, verify it passes

## 5. Visual Preview Implementation (TDD)

- [ ] 5.1 Write test: preview image has transform style when rotation is 90°
- [ ] 5.2 Run test, verify it fails (transform not applied)
- [ ] 5.3 Apply CSS transform to photo preview based on rotation store value
- [ ] 5.4 Run test, verify it passes
- [ ] 5.5 Write test: transform updates when rotation changes from 0° to 180°
- [ ] 5.6 Run test, verify it fails (reactive update not working)
- [ ] 5.7 Ensure transform updates reactively when rotation store changes
- [ ] 5.8 Run test, verify it passes
- [ ] 5.9 Write test: transform is "rotate(0deg)" when rotation is 0°
- [ ] 5.10 Run test, verify it passes
- [ ] 5.11 Write test: transform is "rotate(270deg)" when rotation is 270°
- [ ] 5.12 Run test, verify it passes
- [ ] 5.13 Manually verify preview display at all rotation angles (0°, 90°, 180°, 270°)
- [ ] 5.14 Manually verify preview matches expected final uploaded orientation

## 6. Canvas Rotation Transformation (TDD)

- [ ] 6.1 Write test: canvas dimensions are not swapped for 0° rotation (1080×1920 → 1080×1920)
- [ ] 6.2 Run test, verify it fails (function doesn't exist)
- [ ] 6.3 Create applyRotation helper function in API adapter with dimension logic
- [ ] 6.4 Run test, verify it passes
- [ ] 6.5 Write test: canvas dimensions are not swapped for 180° rotation (1080×1920 → 1080×1920)
- [ ] 6.6 Run test, verify it passes
- [ ] 6.7 Write test: canvas dimensions are swapped for 90° rotation (1080×1920 → 1920×1080)
- [ ] 6.8 Run test, verify it fails (swap logic not implemented)
- [ ] 6.9 Implement dimension swap logic for 90°/270° rotations
- [ ] 6.10 Run test, verify it passes
- [ ] 6.11 Write test: canvas dimensions are swapped for 270° rotation (1080×1920 → 1920×1080)
- [ ] 6.12 Run test, verify it passes
- [ ] 6.13 Write test: 90° rotation applies correct transformation matrix
- [ ] 6.14 Run test, verify it fails (transformation not implemented)
- [ ] 6.15 Implement canvas transformation matrix (translate, rotate, draw)
- [ ] 6.16 Run test, verify it passes
- [ ] 6.17 Write test: 180° rotation applies correct transformation matrix
- [ ] 6.18 Run test, verify it passes
- [ ] 6.19 Write test: 270° rotation applies correct transformation matrix
- [ ] 6.20 Run test, verify it passes
- [ ] 6.21 Write test: image is centered in canvas after rotation
- [ ] 6.22 Run test, verify it passes (should work with transformation implementation)
- [ ] 6.23 Write test: rotated canvas exports as JPEG blob
- [ ] 6.24 Run test, verify it fails (export not implemented)
- [ ] 6.25 Export rotated canvas as JPEG blob with quality setting
- [ ] 6.26 Run test, verify it passes
- [ ] 6.27 Write test: canvas creation failure logs error and returns original blob
- [ ] 6.28 Run test, verify it fails (error handling not implemented)
- [ ] 6.29 Add error handling for canvas operations with fallback to original
- [ ] 6.30 Run test, verify it passes
- [ ] 6.31 Write test: image decode failure logs error and returns original blob
- [ ] 6.32 Run test, verify it passes (should work with error handling from 6.29)

## 7. Upload Integration (TDD)

- [ ] 7.1 Write test: uploadPhoto with rotation 0° uploads original blob without transformation
- [ ] 7.2 Run test, verify it fails (rotation parameter not implemented)
- [ ] 7.3 Modify uploadPhoto to accept rotation parameter
- [ ] 7.4 Implement skip transformation logic when rotation is 0°
- [ ] 7.5 Run test, verify it passes
- [ ] 7.6 Write test: uploadPhoto with rotation 90° calls applyRotation before upload
- [ ] 7.7 Run test, verify it fails (transformation not called)
- [ ] 7.8 Apply transformation when rotation is 90°, 180°, or 270°
- [ ] 7.9 Run test, verify it passes
- [ ] 7.10 Write test: uploadPhoto with rotation 180° uploads rotated blob
- [ ] 7.11 Run test, verify it passes
- [ ] 7.12 Write test: uploadPhoto with rotation 270° uploads rotated blob
- [ ] 7.13 Run test, verify it passes
- [ ] 7.14 Write test: rotation value from photoRotation store is passed to uploadPhoto
- [ ] 7.15 Run test, verify it fails (store not connected to upload)
- [ ] 7.16 Connect confirmation screen to pass rotation value to uploadPhoto
- [ ] 7.17 Run test, verify it passes

## 8. Capture Flow Integration (TDD)

- [ ] 8.1 Write test: capturing photo with beta 80° sets rotation to 0°
- [ ] 8.2 Run test, verify it fails (beta capture not integrated)
- [ ] 8.3 Capture beta reading when photo is taken via camera
- [ ] 8.4 Calculate and set initial rotation from beta reading
- [ ] 8.5 Run test, verify it passes
- [ ] 8.6 Write test: capturing photo with beta 20° sets rotation to 90°
- [ ] 8.7 Run test, verify it passes
- [ ] 8.8 Write test: retaking photo resets rotation state and recalculates from new beta
- [ ] 8.9 Run test, verify it fails (reset not called on retake)
- [ ] 8.10 Reset rotation state when user retakes photo
- [ ] 8.11 Run test, verify it passes
- [ ] 8.12 Write test: manual rotation is preserved until new photo is captured
- [ ] 8.13 Run test, verify it fails (manual changes reset prematurely)
- [ ] 8.14 Ensure manual rotation persists across UI interactions (not on retake)
- [ ] 8.15 Run test, verify it passes
- [ ] 8.16 Write test: file upload defaults rotation to 0° (no beta reading)
- [ ] 8.17 Run test, verify it fails (file upload path not handled)
- [ ] 8.18 Set default rotation 0° for file upload path
- [ ] 8.19 Run test, verify it passes
- [ ] 8.20 Manually test full flow: camera capture → orientation detection → confirmation → rotation → upload

## 9. End-to-End Testing and Validation

- [ ] 9.1 Manual test: portrait photo capture (beta > 45°) → verify rotation 0° → preview correct → upload
- [ ] 9.2 Manual test: landscape photo capture (beta ≤ 45°) → verify rotation 90° → preview correct → upload
- [ ] 9.3 Manual test: portrait photo → manually rotate right → verify rotation 90° → preview updates → upload
- [ ] 9.4 Manual test: landscape photo → manually rotate left → verify rotation 0° → preview updates → upload
- [ ] 9.5 Manual test: capture photo → rotate to 180° → retake photo → verify rotation resets and recalculates
- [ ] 9.6 Manual test: file upload → verify default rotation 0° → manual rotation works → upload
- [ ] 9.7 Manual test: sensor unavailable → verify default rotation 0° → manual rotation works
- [ ] 9.8 Visual verification: upload rotated image → check Gemini receives correctly oriented image
- [ ] 9.9 Visual verification: compare canvas transformation output with expected orientation
- [ ] 9.10 Visual verification: ensure no image distortion or cropping at any rotation angle

## 10. Documentation and Cleanup

- [ ] 10.1 Add JSDoc comments to rotation functions
- [ ] 10.2 Update component documentation with rotation feature
- [ ] 10.3 Remove unused exifOrientation.ts file (if confirmed unused elsewhere)
- [ ] 10.4 Add user-facing documentation for rotation controls
