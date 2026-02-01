## Context

The current frontend uses a vibrant purple gradient background (`linear-gradient(135deg, #667eea 0%, #764ba2 100%)`) with rounded cards and modern styling. This aesthetic is playful but undermines the satirical impact of judging furniture in a mock courtroom.

The redesign applies Dutch government/court visual language to create deadpan presentation. The court seal (`court-seal.png`) has already been generated and placed in `frontend/lib/assets/`.

Current styling is primarily in:
- `frontend/src/routes/+page.svelte` - Global styles via `:global(body)` and component classes
- Individual feature components with scoped styles

All changes are CSS-only with no functional modifications to components.

## Goals / Non-Goals

**Goals:**
- Replace all gradient backgrounds with solid government colors (charcoal #2e2e2e, off-white #fafafa)
- Integrate court seal badge in header
- Apply formal document styling (3px top borders, minimal radius, crisp shadows)
- Add case number (RVM-YYYY-timestamp) and formatted timestamp to verdict display
- Create consistent austere visual language across all components
- Maintain responsive design and accessibility

**Non-Goals:**
- Changing component functionality or behavior
- Modifying API contracts or data structures
- Adding new interactive features
- Changing routing or state management
- Backend modifications

## Decisions

### Decision 1: CSS-only approach with scoped component updates
**Chosen:** Update component-scoped `<style>` blocks individually rather than global CSS file

**Rationale:** 
- Svelte's scoped styling keeps changes isolated and maintainable
- Easier to review component-by-component
- No risk of cascading side effects
- Aligns with existing architecture

**Alternatives considered:**
- Global CSS file: Would require new file structure and potential specificity conflicts
- CSS variables: Over-engineering for one-time color replacement

### Decision 2: Court seal placement in header
**Chosen:** Absolute positioning on left side of header, 80px diameter

**Rationale:**
- Mirrors typical government website branding (logo left, title center/right)
- Maintains visual hierarchy with ⚖️ emoji and title
- Scales well on mobile (can reduce size or hide via media query)

**Alternatives considered:**
- Center above title: Too prominent, competes with main heading
- Replace emoji: Emoji works well, seal is supplementary branding

### Decision 3: Case number format
**Chosen:** `RVM-{YEAR}-{TIMESTAMP}` where timestamp is Unix milliseconds

**Rationale:**
- RVM = Rechtbank voor Meubilair abbreviation
- Year provides human-readable context
- Timestamp ensures uniqueness and is already available via `verdict.timestamp`
- Format mimics Dutch court case numbers (e.g., "ECLI:NL:RBDHA:2026:1234")

**Alternatives considered:**
- Sequential numbering: Requires backend state management, out of scope
- UUID: Too technical, breaks immersion
- Timestamp only: Less readable

### Decision 4: Color palette implementation
**Chosen:** 
- Primary: #2e2e2e (charcoal) for header/footer
- Background: #fafafa (off-white)
- Cards: #ffffff (white)
- Accent borders: #4a4a4a (slate)
- Light borders: #d1d1d1

**Rationale:**
- High contrast for accessibility (WCAG AA compliant)
- Monochrome government aesthetic
- Distinct from any existing brand colors
- Professional without being sterile

**Alternatives considered:**
- Navy blue (#1a3a52): Traditional government, but less austere
- Pure grays (#f0f0f0): Too neutral, lacks warmth
- Black/white only: Too harsh, reduces visual hierarchy

### Decision 5: Border radius reduction
**Chosen:** 2px minimal radius instead of 0px

**Rationale:**
- 2px provides subtle softness while maintaining formal appearance
- Prevents sharp pixel artifacts on high-DPI displays
- Compromise between complete formality (0px) and modern feel (8px)

**Alternatives considered:**
- 0px (completely square): Too harsh, can look pixelated
- 4px: Still too rounded for government aesthetic

### Decision 6: Typography weight adjustments
**Chosen:** Keep Georgia serif for headings, increase weight to 600 (semibold) where appropriate

**Rationale:**
- Georgia already chosen and works well
- Increased weight adds authority without changing font family
- Maintains readability while appearing more formal
- No new font loading required

**Alternatives considered:**
- New serif font (Garamond, Baskerville): Requires font loading, minimal benefit
- Sans-serif everywhere: Loses the legal/formal character

## Risks / Trade-offs

**[Risk]** Court seal image quality at different sizes → **Mitigation:** Use PNG at 400x400px minimum, test at 80px display size and mobile breakpoints. If pixelated, replace with SVG version.

**[Risk]** High contrast (dark header, light background) may feel jarring after gradient → **Mitigation:** This is intentional for formal aesthetic. User feedback will validate if it enhances humor as intended.

**[Risk]** Case number display adds visual clutter to verdict → **Mitigation:** Place in header area above verdict text, use smaller font size, separate with horizontal rule. Treat as metadata, not primary content.

**[Trade-off]** Removing rounded corners makes UI feel less modern → **Accepted:** Formal appearance is the goal; modern feel is antithetical to government aesthetic.

**[Trade-off]** More complex styling rules (borders, dividers) increase CSS maintenance → **Accepted:** Changes are scoped to components; increased lines of CSS are minimal and well-documented.

## Migration Plan

**Deployment:**
1. Update is frontend-only, no backend coordination needed
2. Changes are purely visual, no feature flags required
3. Deploy via standard frontend build process
4. Verify court seal loads correctly (check network tab for 404s)

**Rollback:**
- Git revert of styling changes
- No database migrations or API changes to revert
- Court seal asset can remain even if unused (622KB, negligible)

**Testing:**
- Visual regression testing on Chrome, Firefox, Safari
- Mobile responsive testing (iOS Safari, Chrome Android)
- Accessibility testing (contrast ratios, screen reader compatibility)
- Verify all components render with new color scheme

## Open Questions

None - design is straightforward CSS update with no architectural ambiguity.
