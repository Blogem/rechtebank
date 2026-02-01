## ADDED Requirements

### Requirement: Return stable file identifiers
The API SHALL return stable identifiers for stored photos and verdicts to enable retrieval via shareable URLs.

#### Scenario: Include storage metadata in response
- **WHEN** photo upload and verdict generation completes successfully
- **THEN** response includes storage metadata (timestamp, request ID) that can be used to generate shareable URLs

#### Scenario: Consistent file naming
- **WHEN** system stores photo and verdict JSON files
- **THEN** both files use the same naming pattern (timestamp_requestID) for reliable pairing

#### Scenario: URL-safe identifiers
- **WHEN** system generates timestamps and request IDs
- **THEN** identifiers contain only URL-safe characters (alphanumeric, hyphens, underscores)
