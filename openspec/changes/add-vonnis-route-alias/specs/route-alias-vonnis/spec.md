## ADDED Requirements

### Requirement: Support /vonnis/[id] route
The system SHALL support accessing verdicts via the `/vonnis/[id]` URL path.

#### Scenario: Access verdict via vonnis URL
- **WHEN** user navigates to `/vonnis/[id]` with a valid verdict ID
- **THEN** system displays the verdict page with all verdict details and photo

#### Scenario: Invalid verdict ID via vonnis route
- **WHEN** user navigates to `/vonnis/[id]` with an invalid or malformed ID
- **THEN** system displays HTTP 400 error page with message "Ongeldige vonnis-ID"

#### Scenario: Non-existent verdict via vonnis route
- **WHEN** user navigates to `/vonnis/[id]` with a valid ID format but non-existent verdict
- **THEN** system displays HTTP 404 error page with message "Vonnis niet gevonden"

### Requirement: Maintain /verdict/[id] route compatibility
The system SHALL continue to support the existing `/verdict/[id]` URL path for backward compatibility.

#### Scenario: Access verdict via legacy verdict URL
- **WHEN** user navigates to `/verdict/[id]` with a valid verdict ID
- **THEN** system displays the verdict page identically to the `/vonnis/[id]` route

#### Scenario: Legacy URL shows identical content
- **WHEN** same verdict is accessed via both `/verdict/[id]` and `/vonnis/[id]`
- **THEN** both routes display identical verdict data, photo, and metadata

### Requirement: Shared verdict display logic
The system SHALL use the same verdict display component for both route paths.

#### Scenario: Component reuse across routes
- **WHEN** verdict is rendered on either `/vonnis/[id]` or `/verdict/[id]`
- **THEN** system uses the VerdictDisplay component to render the verdict

#### Scenario: Identical meta tags on both routes
- **WHEN** verdict page is rendered on either route
- **THEN** Open Graph and Twitter Card meta tags are identical regardless of URL path

### Requirement: URL path reflects current route
The system SHALL preserve the URL path that the user navigated to without redirects.

#### Scenario: No redirect from verdict to vonnis
- **WHEN** user navigates to `/verdict/[id]`
- **THEN** system does not redirect to `/vonnis/[id]` and keeps the original URL

#### Scenario: No redirect from vonnis to verdict
- **WHEN** user navigates to `/vonnis/[id]`
- **THEN** system does not redirect to `/verdict/[id]` and keeps the original URL
