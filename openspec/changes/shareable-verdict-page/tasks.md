## 1. Backend - Base64 URL Encoding Utilities

- [x] 1.1 Write unit test for encoding file paths to base64url IDs
- [x] 1.2 Write unit test for decoding base64url IDs back to file paths
- [x] 1.3 Write unit test for malformed/invalid ID error handling
- [x] 1.4 Create `internal/core/domain/verdict_id.go` with encode function to pass tests
- [x] 1.5 Implement decode function to pass tests
- [x] 1.6 Implement error handling to pass validation tests

## 2. Backend - Verdict Retrieval Endpoint

- [x] 2.1 Write handler test for successful verdict retrieval (valid ID returns 200 with verdict + image)
- [x] 2.2 Write handler test for invalid/malformed ID (returns 400)
- [x] 2.3 Write handler test for missing verdict JSON file (returns 404)
- [x] 2.4 Write handler test for missing photo file (returns 404)
- [x] 2.5 Write handler test for file read errors (returns 500)
- [x] 2.6 Create response struct combining verdict JSON + base64 image
- [x] 2.7 Create `GET /v1/verdict/:id` handler in `internal/adapters/http/handlers/verdict_handler.go`
- [x] 2.8 Implement base64url ID decoding to extract file path
- [x] 2.9 Implement verdict JSON file reading from storage
- [x] 2.10 Implement photo file reading from storage
- [x] 2.11 Implement base64 encoding of photo data to data URL format
- [x] 2.12 Implement error handling for all test scenarios
- [x] 2.13 Register route in `internal/adapters/http/router.go`

## 3. Backend - Share URL Generation Endpoint

- [x] 3.1 Write handler test for successful share URL generation (returns shareable ID)
- [x] 3.2 Write handler test for missing verdict files (returns 404)
- [x] 3.3 Write handler test for invalid request data (returns 400)
- [x] 3.4 Create request struct accepting timestamp and requestID
- [x] 3.5 Create `POST /v1/verdict/share` handler in verdict_handler.go
- [x] 3.6 Implement file path construction from timestamp + requestID
- [x] 3.7 Implement validation that both .jpg and .json files exist
- [x] 3.8 Generate base64url ID from file path
- [x] 3.9 Implement error handling for missing files
- [x] 3.10 Register route in router.go

## 4. Backend - Photo Upload Response Enhancement

- [x] 4.1 Write test verifying response includes timestamp field
- [x] 4.2 Write test verifying response includes requestID field
- [x] 4.3 Update response struct in `internal/adapters/http/handlers/judge_handler.go`
- [x] 4.4 Modify judge handler to populate timestamp and requestID fields
- [x] 4.5 Update integration tests if needed

## 5. Backend - Integration Tests

- [x] 5.1 Write integration test for verdict retrieval flow (upload → share → retrieve)
- [x] 5.2 Test verdict retrieval with various image formats (JPEG, PNG, WebP)
- [x] 5.3 Test error scenarios (invalid ID, missing files, corrupted data)

## 6. Frontend - VerdictDisplay Component Enhancement

- [x] 6.1 Write component test: renders correctly with photo data provided
- [x] 6.2 Write component test: renders correctly without photo (backward compatibility)
- [x] 6.3 Write component test: photo displays with responsive sizing
- [x] 6.4 Add optional `imageData` prop to VerdictDisplay.svelte (type: string | undefined)
- [x] 6.5 Add conditional photo display section using {#if imageData}
- [x] 6.6 Implement responsive image styling with CSS
- [x] 6.7 Add bordered frame styling for photo

## 7. Frontend - Share Button Implementation

- [x] 7.1 Write test: share button appears on verdict display
- [x] 7.2 Write test: clicking share calls /v1/verdict/share endpoint
- [x] 7.3 Write test: successful share constructs correct URL
- [x] 7.4 Write test: Web Share API invoked on supported devices
- [x] 7.5 Write test: clipboard fallback used when Web Share unavailable
- [x] 7.6 Write test: error handling for API failures
- [x] 7.7 Add TypeScript types for share request/response
- [x] 7.8 Add API adapter method for /v1/verdict/share endpoint
- [x] 7.9 Add "Deel Vonnis" button to VerdictDisplay.svelte
- [x] 7.10 Implement shareVerdict() function calling POST /v1/verdict/share
- [x] 7.11 Implement URL construction from share response
- [x] 7.12 Implement Web Share API detection and invocation
- [x] 7.13 Implement clipboard API fallback
- [x] 7.14 Add user feedback (success/error messages)
- [x] 7.15 Style share button to match court theme

## 8. Frontend - Shared Verdict Page Route

- [x] 8.1 Write test: page loads successfully with valid verdict ID
- [x] 8.2 Write test: page displays error for invalid verdict ID
- [x] 8.3 Write test: page displays error for missing/expired verdict
- [x] 8.4 Write test: server load function calls GET /v1/verdict/:id
- [x] 8.5 Write test: VerdictDisplay receives both verdict and image data
- [x] 8.6 Add TypeScript types for verdict retrieval response
- [x] 8.7 Add API adapter method for GET /v1/verdict/:id
- [x] 8.8 Create `/src/routes/verdict/[id]/+page.server.ts` with load function
- [x] 8.9 Implement server load function calling GET /v1/verdict/:id
- [x] 8.10 Create `/src/routes/verdict/[id]/+page.svelte`
- [x] 8.11 Render VerdictDisplay component with verdict and imageData props
- [x] 8.12 Add error page for invalid verdict ID
- [x] 8.13 Add error page for missing/expired verdicts

## 9. Frontend - Social Media Meta Tags

- [x] 9.1 Add Open Graph meta tags to verdict page head (og:title, og:description, og:image, og:url)
- [x] 9.2 Add Twitter Card meta tags (twitter:card, twitter:title, twitter:description, twitter:image)
- [x] 9.3 Use verdict photo as og:image (convert to absolute URL or use base64)
- [x] 9.4 Add dynamic title based on verdict content
- [x] 9.5 Test meta tags with Open Graph debugger (Facebook/LinkedIn)
- [x] 9.6 Test meta tags with Twitter Card validator

## 10. Frontend - Shared Page Styling

- [x] 10.1 Ensure shared verdict page uses same CSS as original verdict display
- [x] 10.2 Add responsive layout for photo + verdict on different screen sizes
- [x] 10.3 Test layout on mobile, tablet, and desktop viewports
- [x] 10.4 Verify court theme styling is consistent

## 11. Documentation

- [x] 11.1 Add comments documenting base64url encoding/decoding logic
- [x] 11.2 Update README if needed with sharing feature description

## 12. Testing & Documentation

- [x] 1End-to-End Testing

- [x] 12.1 Write E2E test: complete share flow (upload → share → view)
- [x] 12.2 Write E2E test: shared verdict displays correctly for admissible case
- [x] 12.3 Write E2E test: shared verdict displays correctly for guilty verdict
- [x] 12.4 Write E2E test: shared verdict displays correctly for acquittal
- [x] 12.5 Write E2E test: share button error handling (network failures)
- [x] 12.6 Write E2E test: social media meta tags render correctly

- [x] 13.1 Deploy backend changes to staging/production
- [x] 13.2 Deploy frontend changes to staging/production
- [x] 13.3 Test shareable URLs work in production environment
- [x] 13.4 Verify social media previews display correctly when shared
- [x] 13.5 Confirm existing functionality still works (no regressions)
