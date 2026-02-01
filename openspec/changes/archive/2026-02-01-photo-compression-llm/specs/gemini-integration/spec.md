## MODIFIED Requirements

### Requirement: Send photo with legal system prompt
The system SHALL compress and optionally resize the uploaded photo before sending to Gemini API with a Dutch legal-themed system instruction prompt that requests structured verdict components.

#### Scenario: Photo compressed before API call
- **WHEN** system prepares photo for Gemini API
- **THEN** photo is compressed according to photo-compression capability requirements before being sent

#### Scenario: Photo analysis request with structured output
- **WHEN** system sends compressed photo to Gemini API
- **THEN** request includes system prompt instructing Gemini to analyze furniture and return structured JSON with crime, sentence, and reasoning fields in Dutch legal jargon

#### Scenario: JSON schema definition
- **WHEN** system configures Gemini request
- **THEN** request includes JSON schema defining required fields: admissible (boolean), score (integer 0-10), verdict object with crime, sentence, and reasoning (strings)

#### Scenario: Compression metrics logged
- **WHEN** photo is sent to Gemini API
- **THEN** compression metrics including original size, compressed size, and ratio are logged
