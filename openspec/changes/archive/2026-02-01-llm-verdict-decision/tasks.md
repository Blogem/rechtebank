## 1. Backend - Domain Model Updates

- [x] 1.1 Add `VerdictType` field to `VerdictDetails` struct in `backend/internal/core/domain/verdict.go`
- [x] 1.2 Add JSON tag `verdictType` to the new field
- [x] 1.3 Write unit test for VerdictDetails JSON serialization including verdictType

## 2. Backend - Gemini Adapter Schema

- [x] 2.1 Add `VerdictType` field to `VerdictSchema` struct in `backend/internal/adapters/gemini/analyzer.go`
- [x] 2.2 Update Gemini `ResponseSchema` to include `verdictType` property with type `genai.TypeString`
- [x] 2.3 Add enum constraint with values `["vrijspraak", "waarschuwing", "schuldig"]` to verdictType schema
- [x] 2.4 Add `verdictType` to the Required fields array in ResponseSchema
- [x] 2.5 Update response parsing to map LLM verdictType to domain VerdictType field

## 3. Backend - System Prompt Update

- [x] 3.1 Add verdictType field documentation to system prompt in `analyzer.go`
- [x] 3.2 Add criteria for "vrijspraak" (scores 8-10 or exceptional alignment)
- [x] 3.3 Add criteria for "waarschuwing" (scores 6-7 or minor violations)
- [x] 3.4 Add criteria for "schuldig" (scores 1-5 or serious violations)
- [x] 3.5 Add note that LLM can consider context beyond numeric score

## 4. Backend - Integration Tests

- [x] 4.1 Update Gemini integration test to verify verdictType field is present in response
- [x] 4.2 Add test case verifying verdictType enum values are valid
- [x] 4.3 Test that verdictType is included in JSON response from verdict service

## 5. Frontend - TypeScript Types

- [x] 5.1 Add `verdictType` field to `VerdictDetails` interface in `frontend/src/lib/shared/types/Verdict.ts`
- [x] 5.2 Define type as union: `verdictType: "vrijspraak" | "waarschuwing" | "schuldig"`

## 6. Frontend - VerdictDisplay Component Logic

- [x] 6.1 Update `getVerdictClass()` to use `verdict.verdict.verdictType` instead of score comparison
- [x] 6.2 Add case for `verdictType === 'vrijspraak'` returning 'acquittal'
- [x] 6.3 Add case for `verdictType === 'waarschuwing'` returning 'warning'
- [x] 6.4 Add case for `verdictType === 'schuldig'` returning 'guilty'
- [x] 6.5 Remove score-based logic (`verdict.score >= 7`) from getVerdictClass
- [x] 6.6 Update `getVerdictIcon()` to map verdictType to icons (vrijspraak: 🎉, waarschuwing: ⚠️, schuldig: ⚖️)
- [x] 6.7 Update template conditional to use `verdict.verdict.verdictType` for heading display
- [x] 6.8 Add "Waarschuwing" heading case in template

## 7. Frontend - Styling for Warning Category

- [x] 7.1 Add CSS class `.verdict-display.warning` with orange/amber border (e.g., `border-top: 5px solid #ff9800`)
- [x] 7.2 Verify warning styling is visually distinct from acquittal (green) and guilty (red)
- [x] 7.3 Test responsive design with warning state

## 8. Frontend - Unit Tests

- [x] 8.1 Update VerdictDisplay.test.ts to test verdict display with verdictType field
- [x] 8.2 Test vrijspraak verdict displays correctly
- [x] 8.3 Test waarschuwing verdict displays correctly
- [x] 8.4 Test schuldig verdict displays correctly
- [x] 8.5 Test getVerdictClass returns correct class for each verdictType
- [x] 8.6 Test getVerdictIcon returns correct icon for each verdictType

## 9. Integration Testing

- [x] 9.1 Test end-to-end flow: upload photo → receive verdict with verdictType → display correctly
- [x] 9.2 Verify vrijspraak verdict displays with green styling
- [x] 9.3 Verify waarschuwing verdict displays with orange styling
- [x] 9.4 Verify schuldig verdict displays with red styling
- [x] 9.5 Test that score is still displayed but not used for verdict determination

## 10. Documentation and Cleanup

- [x] 10.1 Update README or API documentation to mention three-category verdict system
- [x] 10.2 Remove any comments referencing old score-based verdict logic
- [x] 10.3 Verify conventional commits are used for all changes
