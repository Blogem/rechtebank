## ADDED Requirements

### Requirement: Display Dutch court seal in header
The system SHALL display the official court seal badge in the application header.

#### Scenario: Seal placement
- **WHEN** header is rendered
- **THEN** system displays court seal image on the left side of the header

#### Scenario: Seal dimensions
- **WHEN** court seal is displayed
- **THEN** system renders the seal at 80px diameter

#### Scenario: Seal image source
- **WHEN** court seal is displayed
- **THEN** system loads the seal from `lib/assets/court-seal.png`

#### Scenario: Seal positioning
- **WHEN** header is rendered
- **THEN** system positions seal using absolute positioning, vertically centered within header

### Requirement: Seal responsive behavior
The system SHALL adapt court seal display for different screen sizes.

#### Scenario: Desktop seal display
- **WHEN** viewport is wider than 768px
- **THEN** system displays court seal at 80px diameter

#### Scenario: Mobile seal display
- **WHEN** viewport is 768px or narrower
- **THEN** system either reduces seal size to 60px or hides seal to maintain layout clarity

### Requirement: Seal visual integration
The system SHALL integrate the seal harmoniously with existing header elements.

#### Scenario: Seal with title layout
- **WHEN** header contains both seal and title
- **THEN** system positions seal on left, title and tagline remain centered or right-aligned

#### Scenario: Seal opacity
- **WHEN** court seal is displayed
- **THEN** system renders seal at 95% opacity for subtle professional appearance
