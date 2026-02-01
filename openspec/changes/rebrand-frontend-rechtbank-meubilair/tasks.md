## 1. Fonts and branding tokens setup

- [x] 1.1 Add and self-host Cormorant Garamond and Source Sans 3 font files in the frontend static assets
- [x] 1.2 Define @font-face rules and global font-family variables (serif and sans-serif) in the global styles
- [x] 1.3 Introduce CSS variables for court branding colors (primary, accent, background, surface, border, text)
- [x] 1.4 Introduce CSS variables for shadow tokens (--shadow-sm, --shadow-base, --shadow-md, --shadow-lg) with layered ambient + directional values
- [x] 1.5 Introduce CSS variables for transition timing (--timing-default: 200ms, --timing-interactive: 300ms, --timing-reveal: 400ms)
- [x] 1.6 Apply branding tokens to the global body styles (background color, base text color, default font)

## 2. Layout shell: masthead and footer

- [x] 2.1 Update `src/routes/+layout.svelte` (or introduce shared components) to render a masthead on all pages
- [x] 2.2 Integrate the court seal, court name, ⚖️ symbol, and formal Dutch tagline into the masthead using the court typography system
- [x] 2.3 Style the masthead with the primary court color and ensure it is mobile-first and single-column friendly
- [x] 2.4 Implement a shared footer with copyright, disclaimer, and placeholder links (e.g., Privacybeleid, Procesreglement, Over deze site), styled with court branding tokens

## 3. Main route (`/`) branding and layout

- [x] 3.1 Refactor `src/routes/+page.svelte` to remove ad-hoc colors and fonts in favor of the court branding tokens and typography system
- [x] 3.2 Style the introduction section as an official notice card (distinct surface, border, heading/label) using document-like card patterns
- [x] 3.3 Wrap the `PhotoCapture` area in a "Zaak indienen" card with appropriate heading, description, and spacing for mobile
- [x] 3.4 Ensure all main route content uses a mobile-first single-column layout without horizontal scrolling on typical smartphone widths

## 4. Verdict and error presentation

- [x] 4.1 Update `src/lib/features/VerdictDisplay.svelte` to use the court typography system and card/document styling
- [x] 4.2 Structure verdict content into sections corresponding to Feiten, Overwegingen, and Uitspraak where applicable
- [x] 4.3 Ensure the submitted furniture image is shown as evidence with a caption/label (e.g., Bewijsmateriaal)
- [x] 4.4 Update `src/lib/features/ErrorDisplay.svelte` to present errors as formal procedure notices with document-like styling
- [x] 4.5 Update `src/lib/features/UploadProgress.svelte` to use court branding tokens and present upload status as a formal notice
- [x] 4.6 Update `src/lib/features/CameraPermission.svelte` to use court typography and formal Dutch copy for permission request messages

## 5. Verdict route (`/verdict/[id]`) alignment

- [x] 5.1 Apply the shared masthead and footer to `src/routes/verdict/[id]/+page.svelte` so it matches the main route branding
- [x] 5.2 Align the verdict layout on `/verdict/[id]` with the inline verdict layout (headings, sections, card styling)
- [x] 5.3 Verify that the verdict route appears as a self-contained, shareable document (masthead, verdict card, footer) on mobile

## 6. Dutch legal microcopy and tone

- [x] 6.1 Audit all user-facing text on `/` and `/verdict/[id]` to ensure it is in Dutch
- [x] 6.2 Replace any informal pronouns or phrasing with formal Dutch using "u" where the user is addressed
- [x] 6.3 Ensure key section titles and labels use legal-sounding Dutch terms (e.g., Zaak indienen, Feiten, Overwegingen, Uitspraak)
- [x] 6.4 Review copy to keep humor in the content of descriptions and verdict reasoning, while keeping UI chrome serious and institutional

## 7. Polished interaction system

- [x] 7.1 Implement layered shadow system on all cards (intro, case submission, verdict, error) using --shadow-base in resting state
- [x] 7.2 Add button hover states with lift effect (translateY -2 to -4px), shadow increase (base → md), and color transition using --timing-interactive
- [x] 7.3 Add button active states with scale-down effect (scale 0.98) and shadow reduction with quick timing (< 100ms)
- [x] 7.4 Implement custom focus rings for all interactive elements using court accent color with offset shadow (not default outline)
- [x] 7.5 Implement verdict reveal animation with staggered section fade-in (photo → Feiten → Overwegingen → Uitspraak) with 100-200ms delays
- [x] 7.6 Add subtle court seal animation on verdict load (fade-in or small rotation, < 600ms, restrained)
- [x] 7.7 Implement page transition fade-in with upward slide (translateY 8px → 0) for state changes using --timing-reveal
- [x] 7.8 Add prefers-reduced-motion media query to disable/reduce all non-essential animations
- [x] 7.9 Verify all animations use CSS transforms (not layout properties) for performance

## 8. Testing and polish

- [x] 8.1 Verify on mobile viewports that all primary flows (permission, capture, upload, verdict, error, verdict route) are readable and usable without horizontal scrolling
- [x] 8.2 Confirm that court branding tokens (fonts, colors, seal, emoji) are applied consistently across header, footer, cards, and verdict views
- [x] 8.3 Test all button interactions (hover, active, focus) across different devices and browsers
- [x] 8.4 Verify verdict reveal animation timing feels natural and adds gravitas without being distracting
- [x] 8.5 Test with prefers-reduced-motion enabled to ensure accessibility
- [x] 8.6 Update or add tests as needed (e.g., snapshot or basic rendering tests) to cover new layout/branding assumptions
- [x] 8.7 Update README or other documentation screenshots/snippets to reflect the new court branding where relevant
