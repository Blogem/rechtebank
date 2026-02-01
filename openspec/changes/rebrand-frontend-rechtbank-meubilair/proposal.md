## Why

The current frontend of the “Rechtbank voor Meubilair” visually feels more like a generic web app than an official Dutch court. The core joke—a very serious court that passes judgment on furniture—only fully works when the visual presentation looks highly official and judicial, while the content itself turns out to be humorous once you actually read it.

Right now, the styling is inconsistent, not explicitly designed mobile‑first, and there is no coherent brand and typography system. This change introduces a clear Dutch court identity (color, typography, copy and component styling) with a strong mobile focus, so that the experience reads as an official judicial site at first glance.

## What Changes

- Introduce a consistent, mobile‑first branding and theme system for the frontend:
  - Self‑hosted Google Fonts:
    - Serif: Cormorant Garamond for headings, court name, and verdict text.
    - Sans-serif: Source Sans 3 for body text, UI labels, and buttons.
  - Color palette aligned with a Dutch court aesthetic:
    - Deep blue for primary elements (header, footer, primary buttons).
    - Warm neutral “paper‑like” background.
    - Subtle gold/bronze accents for labels, rules, and emphasis.
- Redesign the general layout to feel like the official site of the “Rechtbank voor Meubilair”:
  - Masthead/header area with seal, court name, ⚖ emoji, and a formal Dutch tagline.
  - Footer with formal information, disclaimer, and consistent styling.
  - Mobile‑first single‑column layout with clear, document‑like sections.
- Visual and textual rework of the main route (`/`):
  - Introduction section styled as an “official notice” in a formal tone.
  - Camera/PhotoCapture section presented as "Zaak indienen" (filing a case) with clear headings and explanation.
  - Upload, error, and verdict states presented as separate “documents”/notices with formal labels and sections.
- Visual alignment of the verdict route (`/verdict/[id]`) with the new court branding:
  - Same masthead/footer elements and background.
  - Verdict view styled as an official judgment (Feiten, Overwegingen, Uitspraak), with humorous content.
- Uniform Dutch copy for all user‑facing text:
  - Formal tone with consistent use of the polite “u”.
  - Section titles, buttons, and error messages in legal‑sounding Dutch.
  - Humor mainly in the furniture descriptions and reasoning of the verdict, not in the UI chrome.
- Introduction of a reusable “court branding system” for Svelte components:
  - Guidelines/patterns for cards (intro, case submission, verdict, error).
  - Consistent typography and spacing scales.
  - Applied to existing feature components (`PhotoCapture`, `VerdictDisplay`, `ErrorDisplay`, `UploadProgress`, `CameraPermission`) without changing their logic.
- No changes to backend APIs or functional logic; the change is purely visual, copy, and layout oriented on the frontend (no **BREAKING** API changes).

## Capabilities

### New Capabilities

- `court-branding-system`: Defines a coherent court brand for the frontend, including color palette, typography system (Cormorant Garamond + Source Sans 3), use of the seal, use of the ⚖ emoji, and guidelines for how these elements are applied across Svelte components and routes.
- `mobile-first-judicial-layout`: Describes the mobile‑first layout principles for the court site, including header/footer structure, card‑like sections for introduction, case submission, status/progress, verdicts, and errors, with an emphasis on single‑column, highly readable presentation on phones.
- `dutch-legal-microcopy`: Specifies requirements for Dutch, formal UI copy with consistent use of “u”, legal‑sounding section titles (e.g., Zaak indienen, Feiten, Overwegingen, Uitspraak), and a style where the humor lives in the content while the UI chrome remains serious and institutional.
- `verdict-view-consistent-branding`: Describes how the verdict view, both on the main route (inline verdict) and on `/verdict/[id]`, must visually and textually align with the court brand, including document‑like structure, section layout, and use of seal, typography, and colors.- `polished-interaction-system`: Defines the micro-interaction and visual depth system to avoid the flat, dated feel of the previous design, including layered shadow tokens, button hover/active/focus states, verdict reveal animations, page transitions, and consistent timing values inspired by modern Dutch government sites like hogeraad.nl.
### Modified Capabilities

- *None* — there are currently no existing frontend specs for branding or layout that define formal behavioral requirements. This change introduces new capabilities and does not alter any existing functional requirements.

## Impact

- **Frontend codebase**
  - Svelte routes:
    - `src/routes/+page.svelte` (home page: introduction, camera/case submission, inline verdict/error display).
    - `src/routes/verdict/[id]/+page.svelte` (verdict detail page).
    - Any layout files such as `src/routes/+layout.svelte` for global header/footer structure.
  - Svelte components:
    - `src/lib/features/PhotoCapture.svelte`
    - `src/lib/features/VerdictDisplay.svelte`
    - `src/lib/features/ErrorDisplay.svelte`
    - `src/lib/features/UploadProgress.svelte`
    - `src/lib/features/CameraPermission.svelte`
    - Any shared layout/utility components introduced later to centralize branding.
  - Stylesheets / styling structure:
    - Global styles (e.g., `app.html`, global CSS/SCSS) for body background, fonts, and base typography.
    - Local component styles in the above Svelte files to align them with the new brand.
- **Assets**
  - Add and self‑host font files (Cormorant Garamond, Source Sans 3) in the frontend static assets.
  - Possible reuse/minor adjustments of `court-seal.png` and consistent use as brand icon and favicon (favicon already uses the seal and remains in use).
- **User experience**
  - On mobile, the site will immediately present itself as an official Dutch court site, with the joke only becoming apparent when reading the text and verdicts.
  - No changes to functionality or API contracts; all changes are visual, copy, and layout focused.
- **Technical constraints**
  - No new external backend dependencies.
  - Additional static assets (fonts) and adjusted styling must remain compatible with existing tooling (Svelte, Vite) and test strategy (TDD), and do not introduce new architectural layers.