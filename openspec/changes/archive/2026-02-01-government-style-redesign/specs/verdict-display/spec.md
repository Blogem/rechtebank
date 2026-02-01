## MODIFIED Requirements

### Requirement: Display court verdict
The system SHALL display the court verdict returned from the backend API.

#### Scenario: Successful verdict retrieval
- **WHEN** backend returns a verdict response
- **THEN** system displays the verdict text in a government-styled legal interface with austere colors and formal borders

#### Scenario: Verdict with score display
- **WHEN** verdict includes a straightness score (1-10)
- **THEN** system prominently displays the numerical score alongside the verdict text

#### Scenario: Verdict formatting
- **WHEN** verdict is displayed
- **THEN** system formats the text with government-style typography (Georgia serif, semibold weights) and minimal border-radius for formal presentation

### Requirement: Court case status indicators
The system SHALL display different visual styles based on the verdict outcome.

#### Scenario: Non-admissible case
- **WHEN** verdict declares the case "niet-ontvankelijk" (not admissible - not a couch)
- **THEN** system displays verdict with dismissal styling using government color palette and gavel icon

#### Scenario: Guilty verdict
- **WHEN** verdict declares furniture guilty of misalignment
- **THEN** system displays verdict with stern/condemning styling using government color palette and appropriate sentencing

#### Scenario: Acquittal verdict
- **WHEN** verdict declares furniture perfectly straight (score 9-10)
- **THEN** system displays verdict with celebratory styling using government color palette and honorable mention
