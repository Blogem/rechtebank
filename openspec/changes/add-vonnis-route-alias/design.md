## Context

The application currently uses `/verdict/[id]` as the route for displaying shared verdicts. We want to introduce `/vonnis/[id]` (Dutch for "verdict") as the canonical route while maintaining backward compatibility with existing `/verdict/[id]` URLs.

**Current State:**
- Route exists at `frontend/src/routes/verdict/[id]/`
- Share URL generation in `VerdictDisplay.svelte` line 70: `${baseUrl}/verdict/${shareResponse.id}`
- Both +page.svelte and +page.ts files handle verdict display logic
- The route uses SvelteKit's file-based routing system

**Constraints:**
- Must not break existing `/verdict/[id]` URLs (backward compatibility)
- Share functionality should generate new format `/vonnis/[id]`
- No backend API changes required (backend doesn't care about frontend routes)
- Follow Svelte/SvelteKit conventions

## Goals / Non-Goals

**Goals:**
- Create `/vonnis/[id]` route that displays verdicts identically to `/verdict/[id]`
- Maintain `/verdict/[id]` route for backward compatibility
- Update share URL generation to use `/vonnis/[id]` format
- Both routes share the same display logic (no code duplication)

**Non-Goals:**
- Not changing backend API endpoints (remain at `/v1/verdict/*`)
- Not redirecting `/verdict/[id]` to `/vonnis/[id]` (both coexist)
- Not updating existing shared URLs (they continue to work)
- Not changing the "verdict" terminology in code/components (only user-facing URLs)

## Decisions

### Decision 1: Route Structure - Rename and Rewrite via nginx

**Chosen:** Rename `/verdict/[id]` to `/vonnis/[id]` and use nginx rewrite rule to map legacy `/verdict/[id]` URLs to `/vonnis/[id]`

**Rationale:**
- `/vonnis/` becomes the canonical route (single source of truth)
- `/verdict/` is treated as the legacy alias
- nginx rewrite handles URL mapping at infrastructure level
- No file duplication - single set of route files
- Simpler than SvelteKit hooks - already using nginx for static serving
- URL in browser changes to `/vonnis/` (internal redirect)
- Clearer intent: vonnis is primary, verdict is for backward compatibility

**Alternatives Considered:**
- Copy files: Would create duplication and maintenance burden
- SvelteKit reroute hook: Would work but adds code complexity when nginx already handles routing
- Client-side redirect: Would cause flash of content before navigation

### Decision 2: nginx Rewrite Implementation

**Chosen:** Use nginx rewrite rule in `nginx.conf`

**Rationale:**
- nginx already serves the static SvelteKit build
- Simple 3-line addition to existing nginx config
- Handles URL mapping at infrastructure level (before SvelteKit)
- `last` flag performs internal redirect (URL changes in browser to `/vonnis/`)
- Works for all `/verdict/*` paths automatically
- No SvelteKit code changes needed

**Implementation:**
```nginx
# Rewrite legacy /verdict/ URLs to /vonnis/
location /verdict/ {
    rewrite ^/verdict/(.*)$ /vonnis/$1 last;
}
```

**Alternatives Considered:**
- SvelteKit reroute hook: Would work but unnecessary when nginx handles routing
- nginx proxy_pass: Overkill for simple path rewrite
- Client-side redirect: Creates flash of content before navigation

**Chosen:** Update VerdictDisplay.svelte to generate `/vonnis/[id]` URLs

**Rationale:**
- VerdictDisplay.svelte contains the share logic (line 70)
- Single point of change for share URL format
- No need to detect which route the user is on (always generate `/vonnis/`)
- Simple string replacement: `/verdict/` → `/vonnis/`

**Alternatives Considered:**
- Route-aware URL generation: Would require passing route context, adds complexity
- Environment-based toggle: Unnecessary for this simple change
- Keep generating `/verdict/` URLs: Defeats the purpose of the change

### Decision 3: Share URL Generation Strategy

**Chosen:** Update VerdictDisplay.svelte to generate `/vonnis/[id]` URLs

**Rationale:**
- VerdictDisplay.svelte contains the share logic (line 70)
- Single point of change for share URL format
- Simple string replacement: `/verdict/` → `/vonnis/`
- Aligns with making `/vonnis/` the canonical route

**Alternatives Considered:**
- Route-aware URL generation: Would require passing route context, adds complexity
- Environment-based toggle: Unnecessary for this simple change
- Keep generating `/verdict/` URLs: Defeats the purpose of the change

### Decision 4: Internal Redirect Behavior

**Chosen:** nginx rewrite with `last` flag performs internal redirect (URL changes to `/vonnis/`)

**Rationale:**
- Users see the canonical `/vonnis/` URL after accessing `/verdict/`
- Helps users discover the new preferred terminology
- Search engines will eventually index `/vonnis/` as canonical
- Internal redirect is fast (no HTTP round-trip)
- Old links don't break - they just redirect to new URL

**Alternatives Considered:**
- Preserve original URL: Would require nginx proxy instead of rewrite, adds complexity
- 301 permanent redirect: nginx internal redirect is effectively the same but simpler
- No redirect: Would defeat the purpose of migrating to `/vonnis/` terminology

## Risks / Trade-offs

**Risk:** nginx rewrite won't work in pure dev mode (`npm run dev` without nginx)  
→ **Mitigation:** Document that `/verdict/` → `/vonnis/` rewrite only works when running through nginx (docker-compose dev/prod). For pure dev, access via `/vonnis/` directly.

**Risk:** URL change in browser might confuse users  
→ **Mitigation:** This is actually a feature - helps users discover the canonical `/vonnis/` URL. Internal redirect is fast and transparent.

**Risk:** Confusion about which route is "canonical"  
→ **Mitigation:** File structure makes it clear (`/vonnis/[id]/` files exist, `/verdict/[id]/` does not). nginx config shows the rewrite rule. Share functionality always generates `/vonnis/` URLs.

**Trade-off:** Infrastructure dependency on nginx  
→ **Acceptable:** nginx is already required for production deployment. The rewrite is simple and well-documented in nginx.conf.
