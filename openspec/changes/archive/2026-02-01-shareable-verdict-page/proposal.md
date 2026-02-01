## Why

Users want to share their furniture verdicts with friends and family on social media or messaging apps. Currently, there's no way to share a link that displays both the verdict text and the photo together. This reduces engagement and viral potential of the experience.

## What Changes

- Add a shareable URL scheme that encodes the verdict ID or storage reference
- Create a verdict display page accessible via shareable URL that shows both the photo and verdict together
- Modify the verdict display to include the photo alongside the verdict text
- Add a "Share Verdict" (Deel Vonnis) button/feature to generate and copy shareable links
- Reuse existing photo + JSON storage mechanism (no database required)

## Capabilities

### New Capabilities
- `shareable-verdict-url`: Generate and parse shareable URLs that reference stored verdict + photo data
- `shared-verdict-page`: Display verdict and photo together on a standalone page accessible via shareable URL

### Modified Capabilities
- `verdict-display`: Add photo display to the verdict page (currently shows only verdict text)
- `photo-upload-endpoint`: Ensure photos are stored with stable identifiers for retrieval via shareable URLs

## Impact

**Frontend**:
- New route for shared verdict page (`/verdict/[id]` or similar)
- UI component changes to display photo + verdict together
- Share button component with URL generation logic

**Backend**:
- Storage identifier/URL generation logic
- Endpoint to retrieve verdict JSON + photo by identifier (may already exist)

**Storage**:
- Relies on existing photo + JSON storage structure
- Need stable, URL-safe identifiers for stored files
