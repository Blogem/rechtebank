## Why

The application currently uses "verdict" in the URL path, but the Dutch term "vonnis" is more appropriate for the target audience. We need to support both URL paths to maintain backward compatibility while introducing the preferred Dutch terminology.

## What Changes

- Add a new `/vonnis/[id]` route that mirrors the existing `/verdict/[id]` functionality
- Keep the existing `/verdict/[id]` route working for backward compatibility (old links remain functional)
- Both routes will display the same verdict content using shared components
- Update share functionality to generate `/vonnis/[id]` URLs (not `/verdict/[id]`)
- `/vonnis/[id]` becomes the canonical URL format for new shares
- No breaking changes to existing functionality or URLs

## Capabilities

### New Capabilities
- `route-alias-vonnis`: Support for `/vonnis/[id]` route as an alias to verdict functionality

### Modified Capabilities
- `shareable-verdict-url`: Share functionality must generate `/vonnis/[id]` URLs instead of `/verdict/[id]`

## Impact

- **Frontend routing**: New route directory structure at `frontend/src/routes/vonnis/[id]/`
- **Existing routes**: No changes to existing `/verdict/[id]` route - both will coexist
- **Components**: Shared verdict display components will be used by both routes
- **Share functionality**: URLs generated for sharing will use `/vonnis/[id]` format
- **URLs**: All existing `/verdict/[id]` URLs continue to work unchanged (backward compatibility)
- **User experience**: Users can access verdicts via either `/verdict/[id]` or `/vonnis/[id]`, but newly shared links use the Dutch terminology
