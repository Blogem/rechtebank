## Why

Currently, the frontend determines whether a verdict is "Vrijspraak" (acquittal) or "Schuldig Bevonden" (found guilty) based solely on the numeric score (score >= 7 means acquittal). This binary classification limits the LLM's creative freedom and creates a disconnect where the LLM generates legal reasoning but doesn't make the actual verdict decision. Moving this decision to the LLM allows for more nuanced verdicts with three categories: "vrijspraak" (full acquittal), "waarschuwing" (warning - not a formal sentence but not full acquittal), and "schuldig" (found guilty).

## What Changes

- **Backend**: Add `verdictType` field to the Gemini analyzer's structured output schema and domain models
- **Backend**: Update system prompt to instruct the LLM to provide an explicit verdict ("vrijspraak", "waarschuwing", or "schuldig")
- **Frontend**: Remove score-based verdict logic from VerdictDisplay component
- **Frontend**: Display the LLM-provided verdict type with support for three categories instead of calculating it from the score
- **Frontend**: Update TypeScript types to include the new verdict field
- **Frontend**: Add styling and display logic for the new "waarschuwing" category

## Capabilities

### New Capabilities
- `llm-verdict-type`: LLM explicitly provides the verdict type (vrijspraak/waarschuwing/schuldig) as part of its structured response with three-category classification

### Modified Capabilities
<!-- No existing specs to modify -->

## Impact

**Affected Code:**
- `backend/internal/adapters/gemini/analyzer.go` - VerdictSchema struct and system prompt
- `backend/internal/core/domain/verdict.go` - VerdictDetails struct
- `frontend/src/lib/shared/types/Verdict.ts` - VerdictDetails interface
- `frontend/src/lib/features/VerdictDisplay.svelte` - verdict display logic and getVerdictClass/getVerdictIcon functions

**Benefits:**
- More creative, context-aware verdicts from the LLM
- Three-category system allows for nuanced verdicts (warning for borderline cases)
- Removes arbitrary binary score threshold (>= 7) from frontend
- Better separation of concerns (LLM decides, frontend displays)
- Allows for edge cases where score might not directly correlate with verdict

**Breaking Changes:**
None - the score field remains for display purposes, we're adding a new field
