## 1. Write Failing Tests (TDD Red Phase)

- [x] 1.1 Add test verifying overlay button is rendered inside `.preview` container (should fail)
- [x] 1.2 Add test verifying overlay button has ARIA label "Roteer foto 90 graden" (should fail)
- [x] 1.3 Add test verifying overlay button calls rotation handler on click (should fail)
- [x] 1.4 Add test verifying overlay button is disabled when `isProcessing` is true (should fail)
- [x] 1.5 Add test verifying overlay button has class `rotation-button-overlay` (should fail)
- [x] 1.6 Run tests and confirm they fail as expected

## 2. Implement Overlay Rotation Button (TDD Green Phase)

- [x] 2.1 Remove `rotateLeft` import from PhotoCapture.svelte
- [x] 2.2 Remove `handleRotateLeft()` function from PhotoCapture.svelte
- [x] 2.3 Remove `.rotation-hint` section from template (div with "Staat de foto niet goed?" text)
- [x] 2.4 Remove `.rotation-controls` section from template (div containing Links/Rechts buttons)
- [x] 2.5 Add overlay button element inside `.preview` div with class `rotation-button-overlay`
- [x] 2.6 Set button content to Unicode circular arrow symbol (↻)
- [x] 2.7 Add ARIA label `aria-label="Roteer foto 90 graden"` to button
- [x] 2.8 Wire button `onclick` to `handleRotateRight` handler
- [x] 2.9 Bind button `disabled` attribute to `isProcessing` state
- [x] 2.10 Run tests and confirm new tests pass

## 3. Style Overlay Button (TDD Green Phase)

- [x] 3.1 Add CSS rule for `.preview` with `position: relative`
- [x] 3.2 Add CSS rule for `.rotation-button-overlay` with `position: absolute; bottom: 16px; right: 16px; z-index: 10`
- [x] 3.3 Set button background to `rgba(0, 0, 0, 0.6)` (semi-transparent black)
- [x] 3.4 Set button color to white for icon contrast
- [x] 3.5 Set minimum button dimensions to 48×48px (touch-friendly)
- [x] 3.6 Add border-radius for circular or rounded rectangle appearance
- [x] 3.7 Add subtle box-shadow for visual definition
- [x] 3.8 Add hover state with slight opacity increase or color change
- [x] 3.9 Add transition for smooth interaction feedback (0.2s ease)
- [x] 3.10 Update `:disabled` style for reduced opacity during processing

## 4. Refactor and Clean Up (TDD Refactor Phase)

- [x] 4.1 Remove `.rotation-hint` CSS styles from PhotoCapture.svelte
- [x] 4.2 Remove `.rotation-hint p` CSS styles
- [x] 4.3 Remove `.rotation-controls` CSS styles
- [x] 4.4 Remove `.button-rotation` CSS styles
- [x] 4.5 Remove tests for "Links" button from PhotoCapture.test.ts
- [x] 4.6 Remove tests for rotation hint text from PhotoCapture.test.ts
- [x] 4.7 Update existing rotation tests to only test clockwise rotation (remove counter-clockwise scenarios)
- [x] 4.8 Run all tests and confirm everything passes

## 5. Manual Testing

- [x] 5.1 Test overlay button visibility on photo preview (verify semi-transparent background)
- [x] 5.2 Test button position remains in bottom-right corner across different photo orientations
- [x] 5.3 Test button click rotates photo 90° clockwise (0° → 90° → 180° → 270° → 0°)
- [x] 5.4 Test button is disabled during photo upload processing
- [x] 5.5 Test button touch target on mobile devices (iOS Safari, Android Chrome)
- [x] 5.6 Test keyboard navigation and activation (Tab to button, Enter/Space to activate)
- [x] 5.7 Test screen reader announces ARIA label correctly
- [x] 5.8 Test button contrast on various photo backgrounds (light/dark photos)
- [x] 5.9 Test responsive behavior on small screens (minimum 320px width)

## 6. Code Cleanup (Optional)

- [x] 6.1 Check if `rotateLeft()` is used elsewhere in codebase (grep search)
- [x] 6.2 If `rotateLeft()` is unused, remove from rotation.ts and update tests
- [x] 6.3 Update any documentation or comments referencing dual rotation buttons
