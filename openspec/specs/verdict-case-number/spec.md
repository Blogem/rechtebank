## ADDED Requirements

### Requirement: Display unique case number for each verdict
The system SHALL generate and display a unique case number for each verdict.

#### Scenario: Case number format
- **WHEN** verdict is displayed
- **THEN** system generates case number in format "RVM-{YEAR}-{TIMESTAMP}"

#### Scenario: Case number uniqueness
- **WHEN** verdict is displayed
- **THEN** system uses Unix timestamp in milliseconds to ensure unique case number

#### Scenario: Case number derivation
- **WHEN** generating case number
- **THEN** system extracts year from verdict timestamp and uses full timestamp for unique identifier

### Requirement: Display verdict timestamp
The system SHALL display the formatted date and time when the verdict was issued.

#### Scenario: Timestamp formatting
- **WHEN** verdict is displayed
- **THEN** system formats timestamp as "Uitspraak d.d.: {day} {month} {year}, {HH:mm}"

#### Scenario: Dutch date formatting
- **WHEN** formatting verdict date
- **THEN** system uses Dutch locale for month names (e.g., "31 januari 2026")

### Requirement: Case metadata presentation
The system SHALL present case number and timestamp in formal document header style.

#### Scenario: Metadata placement
- **WHEN** verdict is displayed
- **THEN** system displays case number and timestamp above verdict content, below verdict title

#### Scenario: Metadata styling
- **WHEN** case metadata is displayed
- **THEN** system uses smaller font size and muted color for metadata text

#### Scenario: Metadata separation
- **WHEN** case metadata is displayed
- **THEN** system separates metadata from verdict content with horizontal rule (2px solid #d1d1d1)

### Requirement: Case number label formatting
The system SHALL format case metadata with proper labels.

#### Scenario: Case number label
- **WHEN** case number is displayed
- **THEN** system prefixes number with "Zaaknummer:" label

#### Scenario: Timestamp label
- **WHEN** verdict timestamp is displayed
- **THEN** system prefixes date with "Uitspraak d.d.:" label
