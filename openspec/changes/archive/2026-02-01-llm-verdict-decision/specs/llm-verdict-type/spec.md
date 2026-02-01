## ADDED Requirements

### Requirement: LLM provides verdict type classification

The Gemini analyzer SHALL include a `verdictType` field in its structured response that explicitly classifies the verdict into one of three categories: "vrijspraak" (acquittal), "waarschuwing" (warning), or "schuldig" (found guilty).

#### Scenario: LLM returns vrijspraak for well-aligned furniture
- **WHEN** the LLM analyzes furniture with excellent alignment (typically score 8-10)
- **THEN** the response SHALL include `verdictType: "vrijspraak"`

#### Scenario: LLM returns waarschuwing for borderline cases
- **WHEN** the LLM analyzes furniture with minor alignment issues (typically score 6-7)
- **THEN** the response SHALL include `verdictType: "waarschuwing"`

#### Scenario: LLM returns schuldig for poorly aligned furniture
- **WHEN** the LLM analyzes furniture with serious alignment violations (typically score 1-5)
- **THEN** the response SHALL include `verdictType: "schuldig"`

#### Scenario: VerdictType is always present in response
- **WHEN** the Gemini analyzer returns a verdict for admissible furniture
- **THEN** the response MUST include the `verdictType` field

#### Scenario: VerdictType uses exact Dutch terminology
- **WHEN** the Gemini analyzer returns a verdict type
- **THEN** the value SHALL be exactly one of: "vrijspraak", "waarschuwing", or "schuldig" (lowercase)

### Requirement: Backend domain model includes verdict type

The `VerdictDetails` domain struct SHALL include a `VerdictType` field to store the LLM's verdict classification.

#### Scenario: VerdictDetails contains VerdictType field
- **WHEN** a VerdictResponse is created with verdict details
- **THEN** the VerdictDetails SHALL contain a `VerdictType` string field

#### Scenario: VerdictType is serialized in JSON response
- **WHEN** the API returns a verdict to the frontend
- **THEN** the JSON response SHALL include `verdictType` in the `verdict` object

### Requirement: Gemini schema enforces valid verdict types

The Gemini ResponseSchema SHALL define `verdictType` as an enumerated string field with exactly three allowed values.

#### Scenario: Schema allows only valid verdict types
- **WHEN** the Gemini model is configured with ResponseSchema
- **THEN** the `verdictType` field SHALL have an enum constraint with values ["vrijspraak", "waarschuwing", "schuldig"]

#### Scenario: VerdictType is a required field
- **WHEN** the Gemini model returns structured output
- **THEN** the `verdictType` field MUST be included in the required fields list

#### Scenario: Invalid verdict types are rejected
- **WHEN** the LLM attempts to return a verdict type not in the enum
- **THEN** the response SHALL be rejected or the LLM SHALL be forced to use one of the valid values

### Requirement: System prompt instructs verdict decision-making

The Gemini system prompt SHALL provide explicit instructions for determining the verdict type based on the furniture analysis.

#### Scenario: Prompt explains vrijspraak criteria
- **WHEN** the system prompt is provided to the LLM
- **THEN** it SHALL define when to use "vrijspraak" (scores 8-10 or exceptional alignment)

#### Scenario: Prompt explains waarschuwing criteria
- **WHEN** the system prompt is provided to the LLM
- **THEN** it SHALL define when to use "waarschuwing" (scores 5-7 or minor violations)

#### Scenario: Prompt explains schuldig criteria
- **WHEN** the system prompt is provided to the LLM
- **THEN** it SHALL define when to use "schuldig" (scores 1-4 or serious violations)

#### Scenario: Prompt allows context-based decisions
- **WHEN** the system prompt instructs verdict decision-making
- **THEN** it SHALL indicate that the LLM can consider context beyond just the numeric score

### Requirement: Frontend displays LLM-provided verdict type

The VerdictDisplay component SHALL use the `verdictType` field from the API response to determine how to display the verdict, removing score-based logic.

#### Scenario: Frontend reads verdictType from response
- **WHEN** the VerdictDisplay component receives a verdict
- **THEN** it SHALL access `verdict.verdict.verdictType` to determine the verdict classification

#### Scenario: Vrijspraak displays acquittal UI
- **WHEN** `verdictType` is "vrijspraak"
- **THEN** the component SHALL display "Vrijspraak" heading and acquittal styling (green)

#### Scenario: Waarschuwing displays warning UI
- **WHEN** `verdictType` is "waarschuwing"
- **THEN** the component SHALL display "Waarschuwing" heading and warning styling (orange/amber)

#### Scenario: Schuldig displays guilty UI
- **WHEN** `verdictType` is "schuldig"
- **THEN** the component SHALL display "Schuldig Bevonden" heading and guilty styling (red)

#### Scenario: Score-based verdict logic is removed
- **WHEN** determining verdict display class
- **THEN** the component SHALL NOT use `verdict.score >= 7` or similar score comparisons

### Requirement: Frontend TypeScript types include verdict type

The `VerdictDetails` TypeScript interface SHALL include a `verdictType` field matching the backend schema.

#### Scenario: VerdictDetails type includes verdictType
- **WHEN** the VerdictDetails interface is defined
- **THEN** it SHALL include a `verdictType` field of type string

#### Scenario: VerdictType has type-safe values
- **WHEN** defining the verdictType field type
- **THEN** it SHOULD use a union type: `"vrijspraak" | "waarschuwing" | "schuldig"`

### Requirement: Frontend styling supports three verdict categories

The VerdictDisplay component SHALL include CSS styling for all three verdict types.

#### Scenario: Warning verdict has distinct styling
- **WHEN** the verdict type is "waarschuwing"
- **THEN** the component SHALL apply a `.warning` or `.verdict-display.warning` CSS class

#### Scenario: Warning style is visually distinct
- **WHEN** the warning CSS class is applied
- **THEN** it SHALL use a color scheme distinct from acquittal (green) and guilty (red), such as orange or amber

#### Scenario: Warning icon is displayed
- **WHEN** the verdict type is "waarschuwing"
- **THEN** the `getVerdictIcon()` function SHALL return an appropriate warning icon (e.g., ⚠️)
