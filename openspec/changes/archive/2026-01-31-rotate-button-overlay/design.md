## Context

The PhotoCapture component currently displays rotation controls as two separate buttons ("Links" and "Rechts") positioned below the photo preview, with hint text explaining their purpose. This UI pattern creates visual separation between the photo and its controls, requiring users to shift focus away from the image.

**Current Implementation:**
- Two buttons below photo preview (`.rotation-controls` div)
- Hint text section (`.rotation-hint`)
- Both buttons call separate handlers: `handleRotateLeft()` and `handleRotateRight()`
- Utility functions in `rotation.ts`: `rotateLeft()` and `rotateRight()`

**Constraints:**
- Frontend-only change (no backend API modifications)
- Must maintain existing rotation logic (0° → 90° → 180° → 270° → 0° cycle)
- Must preserve canvas-based rotation processing on upload
- Must work on mobile (touch) and desktop (mouse)
- Must meet accessibility standards (ARIA labels, keyboard support)

## Goals / Non-Goals

**Goals:**
- Replace dual button UI with single overlay button on photo
- Position button as semi-transparent overlay in bottom-right corner of photo
- Remove counter-clockwise rotation (clockwise-only rotation)
- Remove rotation hint text
- Maintain smooth CSS rotation transitions
- Ensure touch-friendly button size (minimum 44×44px)
- Keep existing canvas transformation logic unchanged

**Non-Goals:**
- Backend changes (rotation processing remains client-side)
- Changes to rotation angles or cycle logic (still 0°/90°/180°/270°)
- Auto-rotation or orientation detection improvements
- Gesture-based rotation (pinch/rotate)
- Changes to other components (UploadProgress, VerdictDisplay, etc.)

## Decisions

### Decision 1: Overlay button positioning approach
**Choice:** Use absolute positioning within photo preview container with `position: relative` parent.

**Rationale:**
- Simple CSS-only solution, no complex layout calculations
- Button remains in consistent position regardless of photo rotation
- Easy to implement and maintain
- Works across all viewports and photo dimensions

**Alternatives considered:**
- Fixed positioning relative to viewport: Rejected - would not stay with photo during scroll
- Transform-based positioning: Rejected - overly complex, issues with nested transforms

**Implementation:**
```css
.preview {
  position: relative; /* Parent container */
}

.rotation-button-overlay {
  position: absolute;
  bottom: 16px;
  right: 16px;
  z-index: 10;
}
```

### Decision 2: Remove `rotateLeft()` function vs keep for potential reuse
**Choice:** Keep `rotateLeft()` function in `rotation.ts` but remove its import/usage in PhotoCapture.

**Rationale:**
- Minimal code change (only remove import and one button)
- Future-proofs for potential reuse in other components
- No breaking changes to utility module
- Easy to clean up later if truly unused (IDE/linter will flag it)

**Alternatives considered:**
- Delete function entirely: Rejected - could be used elsewhere, unnecessary risk
- Keep both buttons but hide one: Rejected - adds code complexity for no benefit

### Decision 3: Button icon/label
**Choice:** Use Unicode circular arrow symbol (↻) with ARIA label "Roteer foto 90 graden".

**Rationale:**
- No additional icon dependencies or SVG complexity
- Unicode arrows render consistently across browsers
- ARIA label provides clear accessibility text
- Matches existing button style (emoji/unicode usage in "📷 Neem Foto")

**Alternatives considered:**
- SVG icon: Rejected - adds complexity, file dependency
- Text label "Roteer": Rejected - takes more space, less elegant
- No label/icon: Rejected - fails accessibility requirements

### Decision 4: Semi-transparent background style
**Choice:** `background: rgba(0, 0, 0, 0.6)` with white icon color.

**Rationale:**
- High contrast ratio for white icon on dark semi-transparent background
- Opacity 0.6 balances visibility vs photo obstruction
- Black background works with most photo colors
- Matches modern overlay UI patterns (video controls, image galleries)

**Alternatives considered:**
- Light background (white/gray): Rejected - poor contrast on light photos
- Fully opaque background: Rejected - too obstructive
- Blur backdrop (backdrop-filter): Rejected - browser support issues, performance cost

### Decision 5: Disable button during processing
**Choice:** Reuse existing `isProcessing` state to disable overlay button during upload.

**Rationale:**
- Prevents accidental rotation changes during upload
- Consistent with existing "Bevestig" and "Opnieuw" button behavior
- No additional state needed
- Clear visual feedback (reduced opacity via `:disabled` style)

**Implementation:**
```svelte
<button 
  onclick={handleRotateRight} 
  disabled={isProcessing}
  class="rotation-button-overlay"
>
```

### Decision 6: Component structure changes
**Choice:** Remove `.rotation-hint` and `.rotation-controls` sections entirely, move button into `.preview` container.

**Rationale:**
- Simplifies component structure (fewer nested divs)
- Button naturally positioned relative to photo
- Cleaner DOM hierarchy
- Reduces CSS complexity

**Before:**
```svelte
<div class="preview">...</div>
<div class="rotation-hint">...</div>
<div class="rotation-controls">
  <button>Links</button>
  <button>Rechts</button>
</div>
```

**After:**
```svelte
<div class="preview">
  <img ... />
  <button class="rotation-button-overlay">↻</button>
</div>
```

## Risks / Trade-offs

### Risk: Button overlaps important photo content
**Mitigation:** 
- Position in bottom-right corner (typically least important area)
- Use semi-transparent background (content partially visible through button)
- Keep button size modest (48×48px minimum, not larger than necessary)
- If issue persists, users can rotate photo to move content away from button area

### Risk: Poor contrast on dark photos
**Mitigation:**
- White icon provides contrast against dark semi-transparent background
- Semi-transparent background (rgba 0.6 alpha) ensures button outline visible against any photo color
- Add subtle border or shadow if needed during testing

### Risk: Accidental clicks on mobile
**Mitigation:**
- Button positioned in corner (away from typical tap areas)
- Rotation is non-destructive (users can rotate back or continue clicking)
- `isProcessing` state prevents clicks during upload
- 48×48px minimum size reduces mis-taps (not too small)

### Risk: Users expect counter-clockwise rotation
**Trade-off accepted:**
- Simpler UI worth the cost of 3 clicks for 270° rotation
- Most users need 90° or 180° rotation (1-2 clicks)
- Counter-clockwise was rarely used in practice (most users click right repeatedly)
- Modern mobile apps typically use single rotation button (Instagram, Photos apps)

### Risk: Keyboard navigation to overlay button
**Mitigation:**
- Button is focusable (standard `<button>` element)
- ARIA label provides clear keyboard navigation context
- Tab order: photo preview area includes overlay button naturally
- Support Enter/Space keys for activation (browser default)

### Risk: Button position breaks on very small screens
**Mitigation:**
- Use responsive margins (maintain 16px spacing at all viewports)
- Button size remains touch-friendly (48×48px) even on small screens
- Existing mobile styles (`.photo-capture` media query) handle small viewports
- Test on 320px width (iPhone SE) as minimum

## Migration Plan

**Deployment:**
1. Deploy frontend changes (no backend coordination needed)
2. No feature flag required (UI-only change, low risk)
3. No database migrations or API versioning needed

**Rollback:**
- Revert frontend commit if issues arise
- No data migration concerns (purely presentational change)

**Testing before deployment:**
- Unit tests: PhotoCapture component with single rotation button
- Visual regression: Compare screenshots of photo preview with overlay button
- Manual testing: iOS Safari, Android Chrome, desktop browsers
- Touch testing: Verify 48×48px button tappable on mobile devices
- Accessibility testing: Keyboard navigation, screen reader announces ARIA label

**Monitoring post-deployment:**
- User feedback on rotation UX
- Analytics: Track if users rotate multiple times (suggests missing counter-clockwise)
- Error logs: Check for any rotation-related errors (unlikely)

## Open Questions

None - implementation approach is straightforward UI refactoring with clear requirements.
