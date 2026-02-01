"## ADDED Requirements

### Requirement: Mobile-first single-column layout
The system SHALL present the main content for the court site in a mobile-first, single-column layout that is readable and usable on typical smartphone viewports.

#### Scenario: Main route uses single-column layout on mobile
- **WHEN** the main route `/` is viewed on a mobile viewport
- **THEN** the primary content (intro notice, case submission, status, verdict, error) SHALL be arranged in a single column
- **AND** horizontal scrolling SHALL not be required for normal interaction

#### Scenario: Verdict route uses single-column layout on mobile
- **WHEN** the verdict route `/verdict/[id]` is viewed on a mobile viewport
- **THEN** the verdict content SHALL be arranged in a single column
- **AND** the verdict SHALL be readable without horizontal scrolling

### Requirement: Shared masthead and footer
The system SHALL render a shared masthead and footer on both the main route and the verdict route, reflecting the court brand.

#### Scenario: Masthead present on all primary routes
- **WHEN** a user visits `/` or `/verdict/[id]`
- **THEN** a masthead with the court seal and court name SHALL be visible at the top of the page

#### Scenario: Footer present on all primary routes
- **WHEN** a user visits `/` or `/verdict/[id]`
- **THEN** a footer with formal information and/or disclaimer text SHALL be visible at the bottom of the page

### Requirement: Document-like sections for content
The system SHALL present major content areas (intro notice, case submission, verdict, errors) as document-like sections or cards that visually separate them from the background.

#### Scenario: Intro section styled as official notice
- **WHEN** the introduction content on `/` is rendered
- **THEN** it SHALL be contained in a card or panel with a distinct surface color and border
- **AND** it SHALL include a heading or label that indicates an official notice

#### Scenario: Case submission section styled as official form
- **WHEN** the camera/photo capture section is shown
- **THEN** it SHALL be contained in a card or panel labeled as a case submission (e.g., “Zaak indienen”)
- **AND** the controls SHALL be clearly grouped as part of this form

#### Scenario: Verdict and error sections styled as documents
- **WHEN** a verdict or error state is shown
- **THEN** the content SHALL appear within a card or panel styled like an official document on top of the page background"