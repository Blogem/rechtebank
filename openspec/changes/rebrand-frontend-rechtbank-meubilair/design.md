## Context

The frontend is a Svelte-based single-page application that currently presents the "Rechtbank voor Meubilair" as a simple, mostly generic web app. The main flow lives on `/` (camera, photo upload, verdict display, error handling) with an additional verdict route at `/verdict/[id]`. Styling is mostly local to Svelte components, uses system fonts with some ad-hoc serif usage (e.g., Georgia), and a simple grey-on-white palette. There is a court seal asset and some legal-flavoured copy, but no consistent branding system.

The goal of this change is to introduce a coherent, mobile-first Dutch court brand for the frontend, making the site look like an official judicial website at first glance, while keeping the humorous content in the details of the text and verdicts. The backend APIs, domain logic, and state management are not changing; the work is purely on the visual and textual presentation layer.

Constraints and context:
- Tech stack: Svelte, TypeScript, Vite; tests exist and TDD is used.
- No changes to Go backend or API contracts.
- No new external runtime dependencies; fonts should be self-hosted instead of pulled from a third-party CDN.
- The UI must remain fully functional on mobile devices and prioritize a mobile-first, single-column layout.

Key files in scope include:
- `src/routes/+page.svelte` (main experience)
- `src/routes/verdict/[id]/+page.svelte` (verdict details)
- `src/routes/+layout.svelte` (layout shell)
- Feature components: `PhotoCapture`, `VerdictDisplay`, `ErrorDisplay`, `UploadProgress`, `CameraPermission` in `src/lib/features/`
- Global styling entry points (`app.html`, any global CSS) and static assets (fonts, court seal).

## Goals / Non-Goals

**Goals:**
- Establish a reusable court branding system for the frontend:
  - Self-hosted type system using Cormorant Garamond (serif) and Source Sans 3 (sans-serif).
  - A Dutch court-inspired color palette (deep blue primary, warm neutral background, subtle gold/bronze accents).
  - Consistent usage guidelines for the court seal and the ⚖ emoji.
- Redesign the layout and styling to be clearly mobile-first and to read as an official Dutch court site:
  - Single-column, document-like sections on mobile.
  - Clear header (masthead) and footer structure consistent across routes.
  - Card/document patterns for intro, case submission, status, verdict, and error views.
- Align `/` and `/verdict/[id]` visually:
  - Shared header/footer and background.
  - Verdict views styled as official judgments with sections such as Feiten, Overwegingen, Uitspraak.
- Enforce consistent Dutch, formal UI copy conventions:
  - All user-facing text in Dutch.
  - Formal tone with consistent use of "u".
  - Legal-sounding labels and section titles while keeping the humor in the content of verdicts and descriptions.
- Implement the new branding and layout without changing application behavior or API contracts.

**Non-Goals:**
- No changes to backend services, data models, or API endpoints.
- No introduction of new routes, navigation structures, or flow changes beyond what is necessary to apply the new styling.
- No feature-level behavior changes in `PhotoCapture`, `VerdictDisplay`, `ErrorDisplay`, `UploadProgress`, or `CameraPermission` beyond what is required to present them in the new style.
- No configuration or internationalization system for multiple languages; the focus is on Dutch only.
- No re-architecture of the frontend state management or component hierarchy beyond what is minimally necessary to support consistent branding.

## Decisions

1. **Font selection and self-hosting strategy**
   - **Decision:** Use Cormorant Garamond for headings and key "official" text (e.g., court name, section titles, verdict sections) and Source Sans 3 for body text, UI labels, and buttons. Self-host both fonts from the frontend static assets.
   - **Rationale:**
     - Cormorant Garamond provides a distinctly European, classical look appropriate for a court context.
     - Source Sans 3 is neutral, readable on mobile, and fits institutional UI patterns without feeling like a startup brand.
     - Self-hosting avoids external requests to Google Fonts, which is better for privacy and reduces external dependencies.
   - **Alternatives considered:**
     - *System fonts only:* Good performance and simplicity, but lack the distinct, designed identity we want. Rejected because the brand would remain generic.
     - *Google Fonts via CDN:* Easy to integrate, but introduces privacy concerns and a runtime dependency on Google.

2. **Color palette and theming structure**
   - **Decision:** Introduce a small set of design tokens (implemented via CSS variables or a base CSS layer) capturing the Dutch court brand palette:
     - `--color-court-primary`: deep blue for header, footer, and primary actions.
     - `--color-court-accent`: muted gold/bronze for labels, borders, and small highlights.
     - `--color-court-bg`: warm neutral background for the page body.
     - `--color-court-surface`: white for cards/documents.
     - `--color-court-border`: soft neutral for card borders and dividers.
     - `--color-court-text`: dark text color for primary text.
   - **Rationale:**
     - Tokens give a single source of truth for brand colors while keeping implementation simple (plain CSS with variables, no theming framework required).
     - A restrained palette reinforces the serious, institutional look and reduces visual noise.
   - **Alternatives considered:**
     - *Ad-hoc colors in each component:* Easier to start with, but quickly leads to inconsistency and harder future changes. Rejected in favor of tokens.

3. **Layout structure and header/footer design**
   - **Decision:**
     - Implement a masthead in `+layout.svelte` or a shared header component that is present on both `/` and `/verdict/[id]`.
     - Masthead content:
       - Court seal (existing `court-seal.png`) prominently displayed.
       - Court name "Rechtbank voor Meubilair" as an H1 in the serif font.
       - ⚖ emoji as part of the visual identity (in the title or as a small inline mark).
       - Formal Dutch tagline beneath the title.
     - Footer content:
       - Copyright line.
       - Links or placeholders for items like "Privacybeleid", "Procesreglement", "Over deze site".
       - Short disclaimer that the court is fictional.
     - Use a mobile-first, single-column layout for main content on both routes.
   - **Rationale:**
     - Moving the header/footer into the layout ensures branding consistency across routes.
     - Single-column layout is natural on mobile and reinforces the "reading a document" experience.
   - **Alternatives considered:**
     - *Route-specific headers/footers:* More flexibility but risks inconsistent branding and duplication.

4. **Card/document patterns for content sections**
   - **Decision:** Standardize on document-like cards for major sections:
     - Introductory "official notice" card on `/`.
     - "Zaak indienen" (case submission) card wrapping `PhotoCapture`.
     - Status/Upload progress message styled as a narrow card or banner.
     - Verdict and error views styled as separate document cards.
   - **Rationale:**
     - Card/document patterns provide a strong visual metaphor (official documents) that fits the court theme.
     - They are easy to implement in Svelte and work well on mobile.
   - **Alternatives considered:**
     - *Full-bleed sections without cards:* Simpler visually, but less evocative of official paperwork and more like a generic app.

5. **Copy placement and language handling**
   - **Decision:**
     - Keep copy inline in Svelte templates for now, but ensure all user-facing strings are Dutch and formal.
     - Consolidate core labels and section titles where it improves maintainability (e.g., constants or small helper modules) only if needed; avoid building a full i18n system.
   - **Rationale:**
     - The app is Dutch-only; a full internationalization system is unnecessary overhead.
     - Keeping copy near the components that use it maintains clarity while still allowing refactoring later if needed.
   - **Alternatives considered:**
     - *Centralized translation files (JSON/i18n):* Overkill for a single-language, small app. Rejected to keep the change focused on branding.

6. **Verdict view structure on `/` and `/verdict/[id]`**
   - **Decision:**
     - Align the structural layout of verdict content on both routes:
       - Header block: label (e.g., "Uitspraak"), case context, date.
       - Evidence/photo section with caption (e.g., "Bewijsmateriaal A").
       - Text sections: Feiten, Overwegingen, Uitspraak.
     - Use shared styles or a shared component if practical (e.g., a base verdict layout component used by both views, with route-specific data wiring).
   - **Rationale:**
     - Consistency reinforces the brand and makes verdicts easy to share and recognize.
     - A shared layout component avoids duplication of styling logic.
   - **Alternatives considered:**
     - *Keeping the verdict detail route visually distinct:* Would dilute the overall court identity and create more styling work.

7. **Micro-interactions and visual depth system**
   - **Decision:**
     - Introduce a balanced micro-interaction language to avoid the flat, dated feel of the current UI:
       - **Button interactions:** Smooth hover states with subtle lift (2-4px translateY), shadow increase, and color transitions (200-300ms ease-out). Active states with slight scale-down feedback.
       - **Card depth:** Use layered shadows (ambient + direct light) with 2-4px blur for resting state, increasing to 6-8px on hover where appropriate. Smooth transitions on shadow changes.
       - **Verdict reveal:** Staggered fade-in animation for verdict sections (Feiten → Overwegingen → Uitspraak) with 100-200ms delay between sections. Optional subtle court seal animation on verdict load.
       - **Page transitions:** Fade-in content with slight upward slide (translateY(8px) → 0) on route load and state changes.
       - **Focus states:** Designed focus rings using court accent color with offset shadow rather than default browser outline.
     - Define shadow tokens:
       - `--shadow-sm`: 0 1px 2px rgba(0,0,0,0.05) (subtle elements)
       - `--shadow-base`: 0 2px 4px rgba(0,0,0,0.08) (cards, resting)
       - `--shadow-md`: 0 4px 8px rgba(0,0,0,0.12) (elevated cards, hover)
       - `--shadow-lg`: 0 8px 16px rgba(0,0,0,0.15) (modals, prominent elements)
     - Define standard transition timing:
       - Default: `200ms ease-out` for colors and small movements
       - Interactive: `300ms ease-out` for transforms and shadows
       - Reveal: `400ms ease-out` for content appearing
   - **Rationale:**
     - Modern Dutch government sites like hogeraad.nl use subtle depth and smooth interactions to feel polished rather than flat.
     - Balanced animations (not minimal, not expressive) provide feedback without feeling gimmicky or distracting from the serious court tone.
     - Staggered verdict reveals add moment and gravitas to the judgment without being playful.
     - Layered shadows create hierarchy and tactile quality that makes the UI feel current rather than dated.
   - **Alternatives considered:**
     - *Minimal depth with single shadows:* Easier to implement but perpetuates the flat feel. Rejected in favor of richer, more modern depth system.
     - *Expressive animations with spring physics:* Too playful for the formal court context. Balanced easing curves are more appropriate.
     - *No animations:* Would keep the dated, static feel that the current UI suffers from.

## Risks / Trade-offs

- **Risk:** Increased CSS complexity and potential for regressions.
  - **Mitigation:**
    - Introduce a small, well-documented set of CSS variables for brand colors and typography.
    - Prefer incremental styling refactors within existing components rather than wholesale rewrites.
    - Use existing tests and add basic UI-level checks where appropriate (e.g., snapshot or visual tests if available).

- **Risk:** Fonts increase bundle size and may slow initial load on slow mobile connections.
  - **Mitigation:**
    - Limit the number of font weights and styles (e.g., regular + semibold only for each family).
    - Use `font-display: swap` to avoid blank text during font loading.
    - Ensure system font fallbacks are well-defined if custom fonts fail to load.

- **Risk:** Visual changes might make existing screenshots or documentation outdated.
  - **Mitigation:**
    - Update any documented screenshots or README snippets that rely on the old look.
    - Clearly note in commit messages and change logs that this is a branding/layout update.

- **Risk:** Over-stylization could harm readability or usability on very small screens.
  - **Mitigation:**
    - Prioritize readable font sizes (at least 14–16px for body text) and sufficient contrast.
    - Test on representative mobile viewports and adjust spacing/line-heights for clarity.

- **Risk:** Alignment and spacing changes may affect how camera and photo capture elements fit on smaller devices.
  - **Mitigation:**
    - Keep functional components like `PhotoCapture` largely intact; wrap them in styled containers rather than deeply altering their layouts.
    - Test the complete flow (permission → capture → verdict) on mobile-sized viewports after styling changes.
