## Why

The current purple gradient background feels modern and playful, resembling AI-generated designs rather than an authoritative legal website. To maximize the satirical impact of "Rechtbank voor Meubilair," the visual design should be completely serious and austere, creating a stark contrast between the formal government aesthetic and the absurd premise of judging furniture. This enhances the humor through deadpan presentation.

## What Changes

- Replace purple gradient background with solid off-white (#fafafa) government-style background
- Replace semi-transparent header/footer with solid charcoal (#2e2e2e) bars
- Add Dutch court seal badge to header for official government branding
- Add 3px slate border accents to cards for formal document styling
- Add horizontal divider rules between major sections
- Change card border-radius from 8px to 2px for minimal, formal appearance
- Add case number display (format: RVM-YYYY-timestamp) to verdict pages
- Add timestamp formatting for verdict "uitspraak datum"
- Update typography weights to medium/semibold for authoritative presentation
- Update box-shadows to crisp, minimal government-site style
- Remove all gradient styling throughout the application

## Capabilities

### New Capabilities
- `government-visual-theme`: Austere color palette, formal borders, and official document styling for maximum legal authenticity
- `court-seal-branding`: Integration of Dutch court seal badge in header for government authority
- `verdict-case-number`: Display formal case number and timestamp on verdict documents

### Modified Capabilities
- `verdict-display`: Visual presentation changes to match government aesthetic (requirements unchanged)

## Impact

**Frontend Components:**
- `frontend/src/routes/+page.svelte` - Global background, header, footer, card styling
- `frontend/src/lib/features/VerdictDisplay.svelte` - Case number display, formal typography
- `frontend/src/lib/features/ErrorDisplay.svelte` - Styling updates for consistency
- `frontend/src/lib/features/PhotoCapture.svelte` - Card styling updates
- `frontend/src/lib/features/CameraPermission.svelte` - Card styling updates
- `frontend/src/lib/features/UploadProgress.svelte` - Styling updates

**Assets:**
- `frontend/lib/assets/court-seal.png` - Already added, needs integration

**No breaking changes** - Pure visual redesign with no API or functional changes
