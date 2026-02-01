## Context

Currently, the Rechtbank voor Meubilair application:
- Stores photos + verdict JSON on disk using timestamp-based filenames (`HHMMSS_requestID.jpg` and `.json`)
- Organizes files in date-based directories (`YYYY-MM-DD`)
- Does not expose stored photos or verdicts via public URLs
- Has a `VerdictDisplay.svelte` component that shows only verdict text (no photo)
- Uses a single-page application flow where users see their verdict immediately after upload

Users want to share their verdicts on social media, but there's no way to generate a link that displays both the photo and verdict together.

**Constraints:**
- No database - must use existing file-based storage
- Backend is Go + Gin, frontend is Svelte + SvelteKit
- Photos are already stored with JSON verdicts using `timestamp_requestID` naming
- Follows hexagonal architecture principles

## Goals / Non-Goals

**Goals:**
- Enable users to share a link that displays their verdict + photo
- Reuse existing file-based storage without database
- Create shareable URLs that encode the storage identifier
- Make the shared page look identical to the initial verdict display
- Support social media sharing with proper metadata (Open Graph tags)

**Non-Goals:**
- Database integration for verdict persistence
- User accounts or authentication
- Editing/deleting verdicts after creation
- Analytics or view tracking on shared verdicts
- Long-term storage guarantees (existing cleanup policy still applies)

## Decisions

### Decision 1: URL Structure for Shareable Verdicts

**Chosen:** Use base64url-encoded path as the identifier: `/verdict/{base64url(dateDir/filename)}`

Example: `/verdict/MjAyNi0wMi0wMS8xNTMwNDVfYWJjMTIz`

**Rationale:**
- Encodes the full file path (date directory + filename without extension) in the URL
- No need to maintain separate ID mapping or database
- Base64url is URL-safe and reversible
- Keeps the backend stateless

**Alternatives considered:**
- Timestamp + RequestID in URL (`/verdict/2026-02-01/153045_abc123`) - Rejected: exposes file structure, longer URLs
- Hash-based IDs - Rejected: requires mapping storage (database or file lookup)
- Sequential integer IDs - Rejected: requires database or persistent counter

### Decision 2: Backend Endpoint for Retrieving Verdicts

**Chosen:** Add `GET /v1/verdict/:id` endpoint that:
1. Decodes base64url ID to get file path
2. Reads both `.jpg` and `.json` files
3. Returns JSON response with verdict data + base64-encoded image

**Response format:**
```json
{
  "verdict": { /* existing verdict structure */ },
  "image": "data:image/jpeg;base64,..." 
}
```

**Rationale:**
- Single API call to get both verdict and photo
- Image as base64 data URL simplifies frontend (no separate image endpoint)
- Reuses existing file storage structure
- No CORS complexity for image serving

**Alternatives considered:**
- Separate endpoints for JSON and image - Rejected: requires 2 HTTP requests, more complex state management
- Stream raw image with Content-Type - Rejected: harder to combine with verdict JSON in single response
- Store images in blob storage - Rejected: violates "no database" constraint

### Decision 3: Frontend Routing Strategy

**Chosen:** SvelteKit dynamic route at `/src/routes/verdict/[id]/+page.svelte` with server-side load function

**Implementation:**
- `+page.server.ts` calls backend API with verdict ID
- Passes verdict + image data to page component
- Reuse existing `VerdictDisplay.svelte` component
- Add photo display above/beside verdict

**Rationale:**
- SvelteKit's SSR enables Open Graph meta tags for social sharing
- Dynamic routes handle arbitrary verdict IDs
- Server-side loading prevents exposing API endpoint to scrapers
- Component reuse maintains consistency

**Alternatives considered:**
- Client-side only route - Rejected: no SSR means poor social media previews
- Duplicate component for shared page - Rejected: violates DRY, hard to maintain consistency
- Embed both pages in single route with state toggle - Rejected: confusing URL structure

### Decision 4: Share Button Implementation

**Chosen:** Add "Deel Vonnis" button to `VerdictDisplay.svelte` that:
1. Calls backend to generate shareable URL (returns base64url ID)
2. Constructs full URL with current origin
3. Uses Web Share API (mobile) or clipboard fallback (desktop)

**Flow:**
1. User clicks "Deel Vonnis"
2. Frontend calls `POST /v1/verdict/share` with current verdict data
3. Backend returns `{id: "base64url..."}`
4. Frontend constructs URL: `${window.location.origin}/verdict/${id}`
5. Share via native API or copy to clipboard

**Rationale:**
- Backend generates ID to ensure consistency
- Frontend handles sharing UX based on device capabilities
- Reuses existing verdict data (already in component props)

**Alternatives considered:**
- Generate ID client-side - Rejected: client doesn't know file path structure
- Always use clipboard - Rejected: native sharing is better UX on mobile
- Store verdict ID in component when first created - Rejected: requires backend to return ID on initial upload (breaking change)

### Decision 5: Photo Display in VerdictDisplay Component

**Chosen:** Add optional `imageData` prop to `VerdictDisplay.svelte`:
- If present, display photo in bordered frame above verdict
- Maintain existing layout when `imageData` is undefined
- Use CSS to ensure responsive image sizing

**Rationale:**
- Backward compatible (existing usage without photo still works)
- Single component for both use cases
- Clear separation of concerns

**Alternatives considered:**
- Create separate component - Rejected: duplicates verdict display logic
- Always require image - Rejected: breaks existing usage
- Use slots - Rejected: over-engineered for simple image display

## Risks / Trade-offs

**[Risk]** URLs break if files are cleaned up by retention policy  
→ **Mitigation:** Document that shared links expire with the cleanup policy (currently not implemented but planned). Consider adding warning text on share button.

**[Risk]** Base64url IDs could be guessable if timestamp pattern is known  
→ **Mitigation:** This is acceptable - no sensitive data, just furniture photos. Users sharing publicly already expect discoverability.

**[Risk]** Large images in base64 increase response size  
→ **Mitigation:** Image size is already limited to 10MB by upload endpoint. Base64 adds ~33% overhead but still acceptable for modern networks. Consider image compression in future if needed.

**[Risk]** Backend needs to read file system on every verdict view  
→ **Mitigation:** File reads are fast for modern SSDs. No caching needed initially - can add HTTP caching headers (ETag, Cache-Control) if performance becomes an issue.

**[Trade-off]** Stateless design means no view counts or analytics  
→ **Acceptable:** Aligns with "no database" constraint and privacy-first approach.

**[Trade-off]** Shared pages require backend to be running  
→ **Acceptable:** This is inherent to the architecture. Static export not feasible without database.

## Migration Plan

**Deployment:**
1. Deploy backend changes first (new GET endpoint)
2. Deploy frontend changes (new route + share button)
3. No data migration needed - existing files are compatible

**Rollback:**
- Backend: Remove new endpoint, no impact on existing functionality
- Frontend: Remove route and share button, no data loss

**Testing:**
- Test share button on both mobile (Web Share API) and desktop (clipboard)
- Verify social media preview with Open Graph validators
- Test with various image sizes and verdict types
- Confirm cleanup policy doesn't break recently created links

## Open Questions

None - design is ready for implementation.
