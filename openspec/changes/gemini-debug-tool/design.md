## Context

The existing Gemini analyzer (`backend/internal/adapters/gemini/analyzer.go`) wraps API interactions, making it difficult to debug issues with:
- System prompt effectiveness
- Response schema validation
- API error handling
- Timeout and retry behavior

The analyzer currently logs minimal information and doesn't expose the raw request/response cycle. Developers need a standalone tool to test prompt changes and troubleshoot API behavior without running the full application.

## Goals / Non-Goals

**Goals:**
- Create a simple CLI tool that runs independently of the web server
- Show complete Gemini API interaction (system prompt, user prompt, image metadata, raw response)
- Allow quick iteration on prompt engineering by testing different images
- Reuse existing analyzer code to ensure debug behavior matches production

**Non-Goals:**
- Modifying the existing HTTP API or verdict service
- Adding permanent verbose logging to production code (debug tool only)
- Creating a new Gemini client implementation (reuse existing)
- Building a GUI or web interface for debugging

## Decisions

### Decision 1: Standalone CLI vs Analyzer Modification
**Choice:** Create a new standalone CLI in `backend/cmd/debug-gemini/`

**Rationale:**
- Keeps production code clean (no debug-only flags or logging clutter)
- Allows independent execution without starting the server
- Easy to add to `.gitignore` or documentation for developer use only
- Can be run with different environment variables for testing

**Alternatives considered:**
- Add `--debug` flag to main server → rejected, adds complexity to production code
- Modify analyzer to always log verbosely → rejected, pollutes production logs

### Decision 2: Code Reuse Strategy
**Choice:** Import and use the existing `GeminiAnalyzer` directly, with wrapper functions to expose internals

**Rationale:**
- Ensures debug tool behavior matches production exactly
- No code duplication or drift between debug and production paths
- May need to add debug-friendly methods to analyzer (e.g., `DebugAnalyzePhoto`)

**Alternatives considered:**
- Copy-paste analyzer code into debug tool → rejected, creates maintenance burden
- Mock the Gemini client → rejected, defeats purpose of testing real API

### Decision 3: Output Format
**Choice:** Structured output with clear sections for prompt, metadata, and response

**Format:**
```
=== GEMINI DEBUG TOOL ===
Image: /path/to/image.jpg
Size: 245.3 KB
MIME Type: jpeg
Dimensions: 1920x1080

=== SYSTEM PROMPT ===
Je bent de Eerwaarde Rechter...
[full system prompt]

=== USER PROMPT ===
Analyseer dit meubelstuk en spreek je vonnis uit.

=== REQUEST METADATA ===
Model: gemini-2.5-flash-lite
Timeout: 10s
Max Retries: 3

=== RAW RESPONSE ===
{
  "observation": "...",
  "admissible": true,
  ...
}

=== PARSED VERDICT ===
Admissible: Yes
Score: 7/10
Crime: Lichte scheefstand van 3 graden
Sentence: Veroordeeld tot heroriëntatie
...
```

**Rationale:**
- Clear visual separation between sections
- Easy to copy-paste sections for documentation or bug reports
- Human-readable format (not just JSON dump)

**Alternatives considered:**
- Pure JSON output → rejected, hard to read for quick debugging
- Minimal output → rejected, defeats purpose of "seeing everything"

### Decision 4: Error Handling
**Choice:** Show all errors with full stack traces and don't retry automatically

**Rationale:**
- Debug tool should expose failures, not hide them
- Developers need to see timeout, rate limit, and parsing errors clearly
- Automatic retry in debug mode obscures underlying issues

## Risks / Trade-offs

**[Risk]** Tool exposes sensitive API keys in output → **Mitigation:** Documentation warns to never commit tool output; consider redacting API key in logs

**[Risk]** Debug tool diverges from production behavior → **Mitigation:** Import analyzer directly, add integration tests that compare debug vs production output

**[Trade-off]** Adding debug methods to analyzer clutters production code → **Mitigation:** Keep debug surface minimal; use build tags if needed to exclude from production builds

**[Risk]** Large images may cause output to be unwieldy → **Mitigation:** Show image metadata (size, dimensions) but not raw bytes; limit response output to reasonable length
