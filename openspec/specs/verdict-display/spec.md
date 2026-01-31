## ADDED Requirements

### Requirement: Display court verdict
The system SHALL display the court verdict returned from the backend API.

#### Scenario: Successful verdict retrieval
- **WHEN** backend returns a verdict response
- **THEN** system displays the verdict text in a legal/courtroom themed interface

#### Scenario: Verdict with score display
- **WHEN** verdict includes a straightness score (1-10)
- **THEN** system prominently displays the numerical score alongside the verdict text

#### Scenario: Verdict formatting
- **WHEN** verdict is displayed
- **THEN** system formats the text with appropriate typography and styling for comedic legal effect

### Requirement: Court case status indicators
The system SHALL display different visual styles based on the verdict outcome.

#### Scenario: Non-admissible case
- **WHEN** verdict declares the case "niet-ontvankelijk" (not admissible - not a couch)
- **THEN** system displays verdict with dismissal styling and gavel icon

#### Scenario: Guilty verdict
- **WHEN** verdict declares furniture guilty of misalignment
- **THEN** system displays verdict with stern/condemning styling and appropriate sentencing

#### Scenario: Acquittal verdict
- **WHEN** verdict declares furniture perfectly straight (score 9-10)
- **THEN** system displays verdict with celebratory styling and honorable mention

### Requirement: Verdict comedic presentation
The system SHALL present verdicts with appropriate legal humor and Dutch language styling.

#### Scenario: Legal terminology usage
- **WHEN** verdict is displayed
- **THEN** system presents text using Dutch legal jargon and formal court language

#### Scenario: Humorous sentencing display
- **WHEN** verdict includes a sentence (e.g., "brandstapel" for severe misalignment)
- **THEN** system displays the sentence with dramatic styling and appropriate emoji/icons

#### Scenario: Degree of deviation display
- **WHEN** verdict mentions angle deviation
- **THEN** system highlights the specific degree measurement with visual emphasis

### Requirement: Verdict result actions
The system SHALL provide actions after verdict is displayed.

#### Scenario: Try another judgment
- **WHEN** verdict is displayed
- **THEN** system provides a button to submit another furniture photo for judgment

#### Scenario: Share verdict option
- **WHEN** verdict is displayed
- **THEN** system provides an option to share the verdict (copy link or download image)

### Requirement: Loading state during analysis
The system SHALL display appropriate loading state while waiting for backend verdict.

#### Scenario: Analysis in progress
- **WHEN** photo is uploaded and backend is processing
- **THEN** system displays a themed loading animation (e.g., "De rechter beraadslaagt...")

#### Scenario: Loading timeout
- **WHEN** backend takes longer than 30 seconds to respond
- **THEN** system displays a timeout message with option to retry

### Requirement: Error handling for verdict failures
The system SHALL gracefully handle verdict retrieval errors.

#### Scenario: Backend error response
- **WHEN** backend returns an error instead of a verdict
- **THEN** system displays the error in a legal-styled "case dismissed" format

#### Scenario: Network timeout during verdict
- **WHEN** network request times out while fetching verdict
- **THEN** system displays an error message with retry option
