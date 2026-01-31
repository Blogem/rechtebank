## ADDED Requirements

### Requirement: Return structured JSON verdict
The API SHALL return a JSON response with structured verdict data including admissibility, score, and verdict components.

#### Scenario: Admissible furniture case
- **WHEN** photo contains valid furniture
- **THEN** response includes JSON with `{"admissible": true, "score": <1-10>, "verdict": {"crime": "<offense>", "sentence": "<punishment>", "reasoning": "<justification>"}}`

#### Scenario: Non-admissible case
- **WHEN** photo does not contain furniture
- **THEN** response includes JSON with `{"admissible": false, "score": 0, "verdict": {"crime": "Geen meubilair gedetecteerd", "sentence": "Zaak niet-ontvankelijk", "reasoning": "Alleen meubilair kan worden berecht"}}`

### Requirement: Include alignment score
The verdict SHALL include a numeric score from 1 to 10 representing furniture alignment quality.

#### Scenario: Perfect alignment
- **WHEN** furniture is perfectly straight (180 degrees)
- **THEN** score is 10

#### Scenario: Slight misalignment
- **WHEN** furniture has minor angle deviation
- **THEN** score is between 6-9 based on deviation severity

#### Scenario: Severe misalignment
- **WHEN** furniture has significant angle deviation
- **THEN** score is between 1-5 based on deviation severity

#### Scenario: Non-admissible case score
- **WHEN** case is non-admissible
- **THEN** score is 0

### Requirement: Provide humorous legal verdict components
The verdict SHALL include three structured components in Dutch legal jargon: crime (offense description), sentence (punishment), and reasoning (legal justification).

#### Scenario: Crime component
- **WHEN** verdict is generated
- **THEN** `verdict.crime` field describes the furniture offense (e.g., "Rugleuning-afwijking van 5 graden", "Ernstige horizontale schending")

#### Scenario: Sentence component
- **WHEN** verdict is generated
- **THEN** `verdict.sentence` field states the punishment in humorous legal terms (e.g., "Veroordeeld tot de brandstapel", "Vrijgesproken wegens voorbeeldige rechtheid")

#### Scenario: Reasoning component
- **WHEN** verdict is generated
- **THEN** `verdict.reasoning` field provides legal justification (e.g., "Artikel 42 van de Meubilair-wet verbiedt afwijkingen groter dan 3 graden")

#### Scenario: High score verdict
- **WHEN** score is 8-10
- **THEN** sentence includes praise or acquittal in legal jargon

#### Scenario: Medium score verdict
- **WHEN** score is 5-7
- **THEN** sentence includes mild criticism or minor penalty

#### Scenario: Low score verdict
- **WHEN** score is 1-4
- **THEN** sentence includes harsh judgment or severe penalty

### Requirement: Set correct content type
The API SHALL set Content-Type header to `application/json` for all verdict responses.

#### Scenario: Successful response
- **WHEN** API returns verdict
- **THEN** Content-Type header is `application/json`

### Requirement: Enable CORS for frontend
The API SHALL include CORS headers to allow requests from the frontend domain.

#### Scenario: Preflight request
- **WHEN** browser sends OPTIONS preflight request
- **THEN** API returns appropriate CORS headers including Access-Control-Allow-Origin

#### Scenario: POST request from frontend
- **WHEN** frontend sends POST request to `/v1/judge`
- **THEN** response includes Access-Control-Allow-Origin header matching frontend domain

### Requirement: Include request metadata
The verdict response SHALL include metadata about the analysis request.

#### Scenario: Successful analysis
- **WHEN** API returns verdict
- **THEN** response includes `timestamp` field with ISO 8601 formatted analysis time

#### Scenario: Response includes request ID
- **WHEN** API processes request
- **THEN** response includes unique `requestId` for tracking and debugging
