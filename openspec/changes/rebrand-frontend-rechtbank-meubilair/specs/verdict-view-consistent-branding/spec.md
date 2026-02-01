## ADDED Requirements

### Requirement: Consistent verdict layout across routes
The system SHALL present verdicts using a consistent layout and branding on both the main route and the verdict route.

#### Scenario: Inline verdict matches verdict route layout
- **WHEN** a verdict is shown inline on `/`
- **THEN** its structure (headings, sections, visual style) SHALL match the verdict shown on `/verdict/[id]` for the same case, apart from route-specific navigation elements

#### Scenario: Verdict sections follow court brand
- **WHEN** a verdict is displayed
- **THEN** it SHALL use the court typography system for headings and body text
- **AND** it SHALL use the court color and card styles defined by the branding tokens

### Requirement: Verdict structure with sections
The system SHALL structure verdicts into clearly identified sections that resemble an official judgment.

#### Scenario: Verdict includes facts, considerations, and decision
- **WHEN** a verdict is rendered
- **THEN** the content SHALL be organized into sections representing at least:
  - the facts of the case (e.g., “Feiten”)
  - the court’s reasoning (e.g., “Overwegingen”)
  - the decision or outcome (e.g., “Uitspraak”)

### Requirement: Evidence/photo presentation
The system SHALL present the furniture photo as evidence within the verdict view.

#### Scenario: Verdict displays evidence photo with caption
- **WHEN** a verdict is displayed for a submitted photo
- **THEN** the corresponding furniture image SHALL be shown as part of the verdict
- **AND** it SHALL include a caption or label indicating that it is evidence (e.g., “Bewijsmateriaal”)

### Requirement: Shareable verdict appearance
The system SHALL ensure that verdict views appear as self-contained official documents that can be shared while retaining the court branding.

#### Scenario: Verdict route is visually complete for sharing
- **WHEN** a user opens `/verdict/[id]`
- **THEN** the page SHALL include the masthead, the verdict document, and the footer
- **AND** the page SHALL appear visually complete and branded without requiring context from other routes


