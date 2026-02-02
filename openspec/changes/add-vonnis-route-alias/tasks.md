## 1. Rename Route Directory

- [x] 1.1 Rename directory `frontend/src/routes/verdict/` to `frontend/src/routes/vonnis/`
- [x] 1.2 Verify all files moved correctly (+page.svelte, +page.ts, +error.svelte in [id] subdirectory)

## 2. Add nginx Rewrite for Legacy /verdict/ URLs

- [x] 2.1 Open `frontend/nginx.conf`
- [x] 2.2 Add rewrite rule before the `location /` block (around line 53)
- [x] 2.3 Add comment documenting the legacy /verdict → /vonnis rewrite
- [x] 2.4 Use `rewrite ^/verdict/(.*)$ /vonnis/$1 last;` for internal redirect

## 3. Update Share URL Generation

- [x] 3.1 Locate share URL construction in `VerdictDisplay.svelte` (line ~70)
- [x] 3.2 Change `/verdict/${shareResponse.id}` to `/vonnis/${shareResponse.id}`

## 4. Testing

- [x] 4.1 Test accessing verdict via `/vonnis/[id]` with valid ID displays correctly
- [x] 4.2 Test accessing verdict via `/verdict/[id]` redirects to `/vonnis/[id]` (nginx rewrite)
- [x] 4.3 Test invalid ID on `/vonnis/[id]` shows 400 error with "Ongeldige vonnis-ID"
- [x] 4.4 Test non-existent verdict on `/vonnis/[id]` shows 404 error with "Vonnis niet gevonden"
- [x] 4.5 Test `/verdict/[id]` URL changes to `/vonnis/[id]` in browser address bar
- [x] 4.6 Test share functionality generates `/vonnis/[id]` URL (not `/verdict/[id]`)
- [x] 4.7 Test Web Share API receives `/vonnis/[id]` URL
- [x] 4.8 Test Open Graph meta tags work correctly on vonnis route
- [x] 4.9 Rebuild frontend Docker image and test rewrite works in container
- [x] 4.10 Test nginx rewrite only works when running through nginx (not pure dev mode)

## 5. Documentation

- [x] 5.1 Add code comment in `vonnis/[id]/+page.svelte` documenting that it's the canonical route
- [x] 5.2 Add comment in nginx.conf explaining `/verdict/` → `/vonnis/` rewrite for backward compatibility
- [x] 5.3 Add comment in `VerdictDisplay.svelte` explaining vonnis URL generation
- [x] 5.4 Note in README or docs that `/verdict/` rewrite only works when running through nginx
