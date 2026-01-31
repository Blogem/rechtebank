## 1. Write tests for canvas rotation utilities (TDD - RED)

- [x] 1.1 Create `rotationUtils.test.ts` test file
- [x] 1.2 Write test for `getInitialRotation()` with screen.orientation.angle available
- [x] 1.3 Write test for `getInitialRotation()` fallback to window.orientation
- [x] 1.4 Write test for `getInitialRotation()` default to 0 when no API available
- [x] 1.5 Write test for rotation angle normalization (0, 90, 180, 270)
- [x] 1.6 Write test for canvas dimension swapping at 0° rotation
- [x] 1.7 Write test for canvas dimension swapping at 90° rotation
- [x] 1.8 Write test for canvas dimension swapping at 180° rotation
- [x] 1.9 Write test for canvas dimension swapping at 270° rotation
- [x] 1.10 Write test for `canvasToBlob()` with JPEG quality parameter
- [x] 1.11 Run tests to confirm they fail (RED)

## 2. Implement canvas rotation utilities (TDD - GREEN)

- [x] 2.1 Create `rotationUtils.ts` with canvas rotation helper functions
- [x] 2.2 Implement `getInitialRotation()` function using screen.orientation.angle with fallbacks
- [x] 2.3 Implement `rotateImageOnCanvas()` function with dimension swapping logic
- [x] 2.4 Add canvas coordinate transformation (translate, rotate, draw)
- [x] 2.5 Implement `canvasToBlob()` function with JPEG quality parameter (0.9)
- [x] 2.6 Run tests to confirm they pass (GREEN)
- [x] 2.7 Refactor if needed while keeping tests green

## 3. Write tests for PhotoCapture component (TDD - RED)

- [x] 3.1 Create `PhotoCapture.test.ts` test file
- [x] 3.2 Write test for file input triggering on capture button click
- [x] 3.3 Write test for image preview display after file selection
- [x] 3.4 Write test for rotate-left decreasing rotation by 90 degrees
- [x] 3.5 Write test for rotate-right increasing rotation by 90 degrees
- [x] 3.6 Write test for rotation wrapping at 0 and 360 degrees
- [x] 3.7 Write test for confirm callback receiving blob and rotation value
- [x] 3.8 Write test for retake resetting state and revoking object URL
- [x] 3.9 Write test for object URL cleanup on component unmount
- [x] 3.10 Write test for non-image file rejection with error display
- [x] 3.11 Run tests to confirm they fail (RED)

## 4. Implement PhotoCapture component (TDD - GREEN)

- [x] 4.1 Create `PhotoCapture.svelte` component file
- [x] 4.2 Add hidden file input with `accept="image/*"` and `capture="camera"`
- [x] 4.3 Implement capture button that triggers file input
- [x] 4.4 Add state for current photo (File), object URL, and rotation angle
- [x] 4.5 Implement file selection handler with MIME type validation
- [x] 4.6 Create image preview with responsive styling
- [x] 4.7 Add rotate-left and rotate-right buttons using existing rotateLeft/rotateRight utilities
- [x] 4.8 Implement visual rotation preview using CSS transform
- [x] 4.9 Add confirm button that calls canvas rotation and exports blob
- [x] 4.10 Add retake button that cleans up URLs and resets state
- [x] 4.11 Implement URL.createObjectURL() on file selection
- [x] 4.12 Implement URL.revokeObjectURL() on retake, new photo, and component destroy
- [x] 4.13 Add error handling for non-image files and load failures
- [x] 4.14 Define component props: onPhotoConfirmed (required), onCancelled (optional)
- [x] 4.15 Run tests to confirm they pass (GREEN)
- [x] 4.16 Refactor if needed while keeping tests green

## 5. Write tests for automatic rotation removal (TDD - RED)

- [x] 5.1 Update `rotation.test.ts` to remove `calculateInitialRotation()` tests
- [x] 5.2 Verify `rotateLeft()` and `rotateRight()` tests remain
- [x] 5.3 Update OrientationAdapter tests to remove beta/gamma capture tests
- [x] 5.4 Run tests to confirm updated test suite runs

## 6. Remove automatic rotation logic (TDD - GREEN)

- [x] 6.1 Remove `calculateInitialRotation()` function from `rotation.ts`
- [x] 6.2 Verify `rotateLeft()` and `rotateRight()` functions remain in `rotation.ts`
- [x] 6.3 Remove `getBetaAtCapture()` method from OrientationAdapter.ts
- [x] 6.4 Remove `getGammaAtCapture()` method from OrientationAdapter.ts
- [x] 6.5 Run tests to confirm all tests pass

## 7. Integrate PhotoCapture into main page

- [x] 7.1 Import PhotoCapture component in `+page.svelte`
- [x] 7.2 Remove CameraAdapter instantiation and camera stream logic
- [x] 7.3 Remove orientation monitoring setup for photo capture flow
- [x] 7.4 Replace camera preview and capture logic with PhotoCapture component
- [x] 7.5 Update handleConfirm to receive blob and rotation from PhotoCapture
- [x] 7.6 Remove handleCapture function (now handled by PhotoCapture)
- [x] 7.7 Update handleRetake to work with PhotoCapture retake callback
- [x] 7.8 Verify spirit level functionality still works independently *(COMPLETED BUT LATER REMOVED - spirit level didn't work with native camera)*
- [x] 7.9 Run existing integration tests to verify no regressions

## 8. Update upload flow

- [x] 8.1 Update ApiAdapter.uploadPhoto to accept rotation parameter
- [x] 8.2 Simplify metadata payload to exclude beta/gamma sensor data
- [x] 8.3 Include rotation value in upload metadata
- [x] 8.4 Update upload state management to work with new component
- [x] 8.5 Test upload success flow with rotated photo
- [x] 8.6 Test upload failure and retry flow

## 9. Clean up unused code

- [x] 9.1 Evaluate if `orientationData` store is still needed (only for spirit level) *(DECISION: Removed entirely - spirit level didn't work with native camera)*
- [x] 9.2 Remove `orientationData` from appStore.ts if only used for rotation *(COMPLETED: All orientation stores removed)*
- [x] 9.3 Remove CameraAdapter.ts if fully replaced by file input approach *(COMPLETED: Fully removed)*
- [x] 9.4 Update imports and remove unused orientation-related code *(COMPLETED: Removed OrientationAdapter, SpiritLevel, AccessibilityToggle, CameraPreview, PhotoConfirmation, FileUploadFallback)*
- [x] 9.5 Remove unused state variables related to camera stream management
- [x] 9.6 Run full test suite to verify no broken dependencies

## 10. Manual testing on devices

- [~] 10.1 Test photo capture on iOS Safari (iPhone)
- [~] 10.2 Test rotation controls on iOS Safari
- [~] 10.3 Test canvas rotation output on iOS Safari
- [x] 10.4 Test photo capture on Android Chrome
- [x] 10.5 Test rotation controls on Android Chrome
- [x] 10.6 Test canvas rotation output on Android Chrome
- [ ] 10.7 Verify no memory leaks by capturing multiple photos
- [ ] 10.8 Test initial rotation heuristic accuracy on both platforms
- [x] 10.9 Verify spirit level feature still works *(N/A - Spirit level removed, doesn't work with native camera)*
- [ ] 10.10 Test full upload flow with backend integration

## 11. Documentation and polish

- [x] 11.1 Add JSDoc comments to rotation utility functions
- [x] 11.2 Add component documentation for PhotoCapture props
- [x] 11.3 Update README if camera capture flow documentation exists
- [x] 11.4 Ensure all tests pass with `npm test`
- [x] 11.5 Verify linting passes with no errors
- [ ] 11.6 Create commit following conventional commit format
