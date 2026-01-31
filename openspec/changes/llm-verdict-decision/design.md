## Context

Currently, the application uses Gemini to analyze furniture photos and return a structured verdict (observation, crime, reasoning, sentence, score). However, the final verdict type ("Vrijspraak" vs "Schuldig Bevonden") is determined by frontend logic using a hardcoded threshold (score >= 7). This binary classification creates a disconnect where the LLM provides all the reasoning but doesn't make the actual verdict decision. We need a three-category system: "vrijspraak" (full acquittal), "waarschuwing" (warning for borderline cases), and "schuldig" (found guilty).

The system follows hexagonal architecture with:
- **Domain layer**: Core business entities (VerdictResponse, VerdictDetails)
- **Adapter layer**: Gemini integration using structured JSON output
- **Frontend**: Svelte components consuming the verdict API

## Goals / Non-Goals

**Goals:**
- Add `verdictType` field to structured LLM output (values: "vrijspraak", "waarschuwing", or "schuldig")
- Update LLM system prompt to instruct explicit verdict decision-making with three categories
- Remove score-based verdict logic from frontend
- Add UI support for displaying the "waarschuwing" category distinctly
- Maintain backward compatibility (score field remains for display)

**Non-Goals:**
- Changing the score calculation or display
- Modifying admissibility logic or non-furniture handling
- Changing any other verdict fields (crime, sentence, reasoning, observation)
- UI redesign beyond using the new verdict field

## Decisions

### Decision 1: Field Name and Location
**Choice**: Add `verdictType: string` to `VerdictDetails` struct/interface

**Rationale**: 
- Places the verdict alongside other judicial components (crime, sentence, reasoning)
- Keeps `VerdictResponse` top-level fields minimal (admissible, score, verdict, requestId, timestamp)
- Follows existing pattern where detailed verdict information lives in `VerdictDetails`

**Alternatives considered**:
- Top-level field in `VerdictResponse` → Rejected: Would clutter the top-level structure
- Separate `verdict` and `verdictType` at top level → Rejected: Creates confusion about what "verdict" means

### Decision 2: Allowed Values
**Choice**: String with three allowed values: `"vrijspraak"` | `"waarschuwing"` | `"schuldig"`

**Rationale**: 
- Matches Dutch legal terminology used throughout the app
- Three-category system allows nuanced verdicts for borderline cases
- "Waarschuwing" represents furniture that has issues but doesn't warrant full conviction
- LLM can easily understand and generate these exact strings
- Easy to validate and display

**Alternatives considered**:
- Two categories only → Rejected: Too binary, doesn't capture nuance of borderline scores (5-7)
- English values ("acquittal"/"warning"/"guilty") → Rejected: Inconsistent with Dutch theme
- Enum with numeric values → Rejected: Less readable, adds complexity
- More than three categories → Rejected: Over-complicates; three levels provide sufficient nuance
### Decision 3: LLM Prompt Strategy
**Choice**: Add explicit instruction in system prompt requiring `verdictType` field and explaining when to use each value

**Rationale**:
- Gemini responds well to clear, structured instructions
- System prompt already contains detailed guidance on scoring
- Can provide nuanced criteria beyond just score threshold

**Prompt addition**:
```
- verdictType: "vrijspraak" (acquittal), "waarschuwing" (warning), or "schuldig" (found guilty)
  - Use "vrijspraak" for scores 8-10 OR when the furniture shows exceptional alignment and character
  - Use "waarschuwing" for scores 6-7 OR when minor violations exist but don't warrant conviction
  - Use "schuldig" for scores 1-5 OR when clear, serious violations of furniture alignment laws are present
```

**Alternatives considered**:
- Let LLM decide without guidance → Rejected: Inconsistent results
- Strict score-based rule in prompt → Rejected: Defeats the purpose of LLM decision-making

### Decision 4: Frontend Refactoring
**Choice**: Replace conditional logic in `getVerdictClass()` and `getVerdictIcon()` with direct mapping from `verdict.verdict.verdictType`

**Before**:
```typescript
function getVerdictClass(): string {
  if (!verdict.admissible) return 'dismissed';
  return verdict.score >= 7 ? 'acquittal' : 'guilty';
}
```

**After**:
```typescript
function getVerdictClass(): string {
  if (!verdict.admissible) return 'dismissed';
  if (verdict.verdict.verdictType === 'vrijspraak') return 'acquittal';
  if (verdict.verdict.verdictType === 'waarschuwing') return 'warning';
  return 'guilty';
}
```

**Additional changes needed**:
- Add CSS class `.verdict-display.warning` with appropriate styling (e.g., orange/amber border)
- Update `getVerdictIcon()` to return a warning icon (e.g., ⚠️) for "waarschuwing"
- Update display text in template to show "Waarschuwing" heading for this category

**Rationale**:
- Minimal change, preserves existing CSS classes and styling
- Single source of truth (LLM verdict)
- Easy to test and verify

### Decision 5: Gemini Schema Update
**Choice**: Add `verdictType` as a required field in the ResponseSchema

**Implementation**:
```go
model.ResponseSchema = &genai.Schema{
  Type: genai.TypeObject,
  Properties: map[string]*genai.Schema{
    // ... existing fields
    "verdictType": {
      Type: genai.TypeString,
      Enum: []string{"vrijspraak", "waarschuwing", "schuldig"},
    },
  },
  Required: []string{"admissible", "score", "crime", "sentence", "reasoning", "observation", "verdictType"},
}
```

**Rationale**:
- Enforces valid values at API level
- Structured output ensures field is always present
- Enum constraint prevents invalid values

## Risks / Trade-offs

**[Risk]** LLM generates verdict that contradicts the score  
→ **Mitigation**: This is actually acceptable and desired behavior. The LLM can consider context beyond just the angle. We trust the LLM's judgment. If this becomes problematic, we can add validation logging.

**[Risk]** Breaking changes for existing API consumers  
→ **Mitigation**: Adding a required field could break clients. However, since this is an internal app with coupled frontend/backend, we can deploy atomically. If decoupling is needed later, make `verdictType` optional with backend fallback.

**[Risk]** Increased LLM token usage from longer prompt  
→ **Mitigation**: Adding ~2 lines to system prompt has negligible cost impact. System prompts are not counted per-request.

**[Trade-off]** Giving up deterministic verdict calculation  
→ **Acceptance**: This is the intended change. We're trading predictability for creativity and context-awareness.

**[Trade-off]** Score display becomes less meaningful  
→ **Acceptance**: Score remains for user feedback but verdict becomes primary signal. This is clearer UX.
