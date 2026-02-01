## ADDED Requirements

### Requirement: Dutch language for all user-facing text
The system SHALL present all user-facing text in Dutch.

#### Scenario: Main route uses Dutch text
- **WHEN** a user visits `/`
- **THEN** all visible labels, headings, buttons, and messages SHALL be in Dutch

#### Scenario: Verdict route uses Dutch text
- **WHEN** a user visits `/verdict/[id]`
- **THEN** all visible labels, headings, buttons, and messages SHALL be in Dutch

### Requirement: Formal tone with “u”
The system SHALL use a formal Dutch tone and address users using “u” rather than informal pronouns.

#### Scenario: Buttons and prompts use formal “u”
- **WHEN** buttons or prompts address the user directly
- **THEN** they SHALL use “u” (formal form) where a pronoun is needed
- **AND** they SHALL avoid informal pronouns such as “je” or “jij”

### Requirement: Legal-sounding section titles
The system SHALL use legal-sounding Dutch titles for key sections related to the court process.

#### Scenario: Case submission titled as “Zaak indienen”
- **WHEN** the case submission (photo capture) section is rendered
- **THEN** it SHALL include a title or label that uses phrasing equivalent to “Zaak indienen” (filing a case)

#### Scenario: Verdict sections use legal headings
- **WHEN** verdict details are rendered
- **THEN** section headings SHALL use legal-sounding terms such as “Feiten”, “Overwegingen”, and “Uitspraak” where applicable

### Requirement: Humor located in content, not chrome
The system SHALL keep the humor primarily in the content of descriptions and verdicts, while the UI chrome remains serious and institutional.

#### Scenario: UI chrome remains serious
- **WHEN** a user views the general layout (header, footer, cards, labels)
- **THEN** the styling and wording SHALL resemble a serious Dutch court site
- **AND** any humorous or absurd elements SHALL be confined to the descriptive text and verdict reasoning, not to the structural labels or navigation
