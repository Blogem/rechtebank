## ADDED Requirements

### Requirement: Generate shareable verdict URLs
The system SHALL generate shareable URLs that encode the storage location of verdict data and photo.

#### Scenario: Generate URL for stored verdict
- **WHEN** backend receives a request to create a shareable URL for a stored verdict
- **THEN** system generates a base64url-encoded identifier containing the file path and returns the shareable ID

#### Scenario: URL encoding uses base64url
- **WHEN** system generates a shareable URL identifier
- **THEN** system encodes the file path (date directory + filename without extension) using base64url encoding for URL safety

#### Scenario: Decode shareable URL identifier
- **WHEN** backend receives a request with a shareable verdict ID
- **THEN** system decodes the base64url identifier to extract the file path for retrieving verdict and photo files

#### Scenario: Invalid URL identifier
- **WHEN** backend receives a malformed or invalid shareable ID that cannot be decoded
- **THEN** system returns HTTP 400 with error message "Invalid verdict ID"

### Requirement: Generate shareable URL from current verdict
The system SHALL provide an endpoint to generate a shareable URL from verdict data.

#### Scenario: Create shareable link for existing verdict
- **WHEN** client sends POST request to `/v1/verdict/share` with verdict timestamp and request ID
- **THEN** system locates the corresponding stored files and returns a shareable ID

#### Scenario: Verdict files not found
- **WHEN** client requests a shareable URL for verdict files that don't exist
- **THEN** system returns HTTP 404 with error message "Verdict not found"

### Requirement: Retrieve verdict by shareable ID
The system SHALL provide an endpoint to retrieve verdict data and photo using a shareable ID.

#### Scenario: Successful verdict retrieval
- **WHEN** client sends GET request to `/v1/verdict/:id` with valid shareable ID
- **THEN** system returns HTTP 200 with JSON containing verdict data and base64-encoded photo

#### Scenario: Response includes photo data
- **WHEN** verdict is retrieved by shareable ID
- **THEN** response includes the photo as a base64-encoded data URL with format "data:image/jpeg;base64,..."

#### Scenario: Response includes verdict JSON
- **WHEN** verdict is retrieved by shareable ID
- **THEN** response includes the complete verdict object with all original fields (score, admissible, verdict details, timestamp)

#### Scenario: Missing photo file
- **WHEN** verdict JSON exists but photo file is missing from storage
- **THEN** system returns HTTP 404 with error message "Photo file not found"

#### Scenario: Missing verdict JSON file
- **WHEN** photo exists but verdict JSON file is missing from storage
- **THEN** system returns HTTP 404 with error message "Verdict data not found"

#### Scenario: Verdict files cleaned up by retention policy
- **WHEN** verdict files have been removed by the cleanup policy
- **THEN** system returns HTTP 410 with error message "Verdict no longer available"
