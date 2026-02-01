## ADDED Requirements

### Requirement: Court branding tokens
The system SHALL define a set of court branding tokens (colors, typography, and emblem usage) that can be applied consistently across all frontend routes and components.

#### Scenario: Branding tokens available as CSS variables
- **WHEN** the frontend application is built
- **THEN** the global styles SHALL expose CSS variables for at least:
  - a primary court color
  - an accent court color
  - a background color
  - a surface (card) color
  - a border/divider color
  - a primary text color

#### Scenario: Branding tokens used in header and footer
- **WHEN** the header and footer are rendered
- **THEN** they SHALL use the court primary color for their background
- **AND** they SHALL use the court surface/text tokens for legible text and icons

### Requirement: Court typography system
The system SHALL define a typography system using self-hosted Cormorant Garamond for headings and key “official” text, and Source Sans 3 for body text and UI elements, with appropriate fallbacks.

#### Scenario: Court fonts are declared and self-hosted
- **WHEN** the application is loaded
- **THEN** `@font-face` declarations SHALL be present for Cormorant Garamond and Source Sans 3
- **AND** the fonts SHALL be served from the application's static assets (not from an external CDN)

#### Scenario: Headings use court serif font
- **WHEN** a page heading for the court (e.g., the main H1 in the masthead) is rendered
- **THEN** it SHALL use the court serif font family as its primary font
- **AND** it SHALL fall back to system serif fonts if the custom font fails to load

#### Scenario: Body text uses court sans-serif font
- **WHEN** main body text and UI labels are rendered
- **THEN** they SHALL use the court sans-serif font family as their primary font
- **AND** they SHALL fall back to system sans-serif fonts if the custom font fails to load

### Requirement: Court seal and emoji usage
The system SHALL define how the court seal asset and the ⚖ emoji are used as part of the brand.

#### Scenario: Masthead displays court seal
- **WHEN** the masthead is rendered
- **THEN** the court seal image SHALL be displayed in the header area
- **AND** it SHALL be visually associated with the court name

#### Scenario: Court symbol usage is consistent
- **WHEN** the court symbol (⚖) is rendered in headings or dividers
- **THEN** it SHALL be used consistently in the masthead and/or as a decorative divider
- **AND** it SHALL not be used in a way that conflicts with the serious judicial visual style"## ADDED Requirements

### Requirement: Court branding tokens
The system SHALL define a set of court branding tokens (colors, typography, and emblem usage) that can be applied consistently across all frontend routes and components.

#### Scenario: Branding tokens available in styles
- **WHEN** the frontend is built
- **THEN** the global styles SHALL expose CSS variables for court primary color, accent color, background color, surface color, border color, and base text color

..."