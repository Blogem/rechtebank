## ADDED Requirements

### Requirement: Austere government color palette
The system SHALL use an austere color scheme that reflects Dutch government and court websites.

#### Scenario: Primary background color
- **WHEN** application renders
- **THEN** system uses off-white (#fafafa) as the primary background color

#### Scenario: Header and footer styling
- **WHEN** application renders header and footer
- **THEN** system uses solid charcoal (#2e2e2e) background with white text

#### Scenario: Card background
- **WHEN** content cards are displayed
- **THEN** system uses pure white (#ffffff) background for cards

#### Scenario: No gradient backgrounds
- **WHEN** any component is styled
- **THEN** system MUST NOT use gradient backgrounds anywhere in the application

### Requirement: Formal document border styling
The system SHALL apply formal government-style borders to content sections.

#### Scenario: Card accent borders
- **WHEN** content cards are displayed
- **THEN** system applies a 3px solid slate (#4a4a4a) top border to each card

#### Scenario: Header border accent
- **WHEN** header is displayed
- **THEN** system applies a 3px solid slate (#4a4a4a) bottom border

#### Scenario: Footer border accent
- **WHEN** footer is displayed
- **THEN** system applies a 3px solid slate (#4a4a4a) top border

### Requirement: Minimal border radius for formal appearance
The system SHALL use minimal border radius to create sharp, formal edges.

#### Scenario: Card border radius
- **WHEN** content cards are rendered
- **THEN** system applies 2px border-radius (not 8px or higher)

#### Scenario: Button border radius
- **WHEN** buttons are rendered
- **THEN** system applies minimal border-radius (0-2px) for formal appearance

### Requirement: Horizontal section dividers
The system SHALL display horizontal divider rules between major content sections.

#### Scenario: Section separation
- **WHEN** multiple content sections are displayed
- **THEN** system renders horizontal border lines (2px solid #d1d1d1) between sections

### Requirement: Crisp minimal shadows
The system SHALL use subtle, crisp shadows instead of soft blur effects.

#### Scenario: Card shadows
- **WHEN** content cards are displayed
- **THEN** system applies subtle box-shadow (0 1px 3px rgba(0,0,0,0.1)) with minimal blur

### Requirement: Authoritative typography
The system SHALL increase font weights for formal, authoritative presentation.

#### Scenario: Heading font weight
- **WHEN** headings are displayed
- **THEN** system uses semibold weight (600) for Georgia serif headings

#### Scenario: Formal text styling
- **WHEN** formal content text is displayed
- **THEN** system maintains Georgia serif with increased line-height (1.7) for readability
