## Context

The current system sends full-resolution photos directly to the Gemini API for furniture analysis. Each photo is encoded as base64 and sent in the API request, consuming tokens proportional to the image size. Gemini's token cost for vision models is based on image dimensions, so larger images directly increase API costs.

The backend architecture follows hexagonal/layered principles with the `GeminiAnalyzer` in `internal/adapters/gemini/analyzer.go` handling all Gemini API interactions. Photos arrive as byte arrays and are sent with MIME type detection but no preprocessing.

## Goals / Non-Goals

**Goals:**
- Reduce token consumption per Gemini API request by 50-70% through image compression
- Maintain furniture detection accuracy (no degradation in verdict quality)
- Implement compression transparently without changing external APIs or frontend behavior
- Add observability metrics to track compression effectiveness (original size, compressed size, ratio)
- Support JPEG, PNG, and WebP formats

**Non-Goals:**
- Client-side compression (frontend already sends photos, changing this would require API changes)
- Changing photo storage format (compression only applies to LLM API calls)
- Adaptive quality based on furniture detection confidence (future enhancement)
- Caching compressed versions (stateless compression per request)

## Decisions

### 1. Compression Location: Backend Before API Call
**Decision:** Compress photos in `analyzer.go` immediately before sending to Gemini API.

**Rationale:** 
- Keeps compression logic isolated in the Gemini adapter (hexagonal architecture)
- Frontend continues to send original photos (no breaking changes)
- Compression happens only for LLM calls, not for storage or other uses
- Easy to A/B test and roll back if needed

**Alternatives Considered:**
- Frontend compression: Rejected due to API breaking change and inconsistent device capabilities
- Middleware compression: Rejected as it would require all consumers to use compressed photos

### 2. Compression Library: Go Standard Library `image/jpeg` and `image/png`
**Decision:** Use Go's built-in `image` packages for compression.

**Rationale:**
- No external dependencies (simpler deployment, smaller attack surface)
- JPEG quality parameter provides fine-grained control (0-100)
- PNG supports compression levels (BestSpeed to BestCompression)
- Already battle-tested and optimized

**Alternatives Considered:**
- Third-party libraries (e.g., `bimg`, `imagick`): Rejected due to C dependencies and deployment complexity
- WebP conversion: Deferred as it requires external libraries; can add later if needed

### 3. Target Quality: JPEG Quality 75, PNG CompressionLevel 5
**Decision:** 
- JPEG images: Compress to quality 75
- PNG images: Use `png.BestSpeed` (level 5)
- WebP images: Pass through without compression (already efficient)

**Rationale:**
- Quality 75 balances file size reduction (~50-60%) with visual quality retention
- Furniture edges and shapes remain clear for LLM analysis
- Can be tuned via configuration if needed
- Standard recommendation for web images

**Alternatives Considered:**
- Lower quality (50-60): Rejected due to potential degradation of fine details
- Higher quality (85-90): Rejected as compression ratio too low
- Resize dimensions: Deferred to avoid over-optimization without data

### 4. Resize Strategy: Max Dimension 1600px
**Decision:** If either dimension exceeds 1600px, resize proportionally to fit within 1600x1600.

**Rationale:**
- Gemini API likely downsamples large images internally anyway
- Furniture can be recognized clearly at 1600px
- Reduces tokens further without quality loss
- Typical phone photos are 3000-4000px, so this saves significant bandwidth

**Alternatives Considered:**
- No resizing: Rejected as it leaves easy optimization on the table
- Smaller max (1024px): Rejected as it may affect detail recognition
- Larger max (2048px): Provides minimal additional detail for furniture

### 5. Error Handling: Fail Gracefully to Original Image
**Decision:** If compression fails, log the error and send the original uncompressed image.

**Rationale:**
- Compression errors shouldn't break the user experience
- Better to pay for full tokens than return an error
- Errors will be logged for monitoring and debugging

**Alternatives Considered:**
- Fail the request: Rejected as too disruptive
- Retry with different settings: Rejected as adds latency

### 6. Metrics: Log Original and Compressed Sizes
**Decision:** Add structured logging with `originalSize`, `compressedSize`, and `compressionRatio` fields.

**Rationale:**
- Track actual compression effectiveness in production
- Identify images that don't compress well
- Tune quality settings based on real data
- No infrastructure changes needed (use existing logging)

**Alternatives Considered:**
- Prometheus metrics: Deferred until we see if logging is sufficient
- Per-format breakdown: Can add if needed based on logs

## Risks / Trade-offs

**[Risk] Compression degrades furniture detection accuracy** → Mitigation:
- Start with conservative quality (75) known to preserve detail
- Monitor verdict distribution before/after deployment
- Can A/B test by compressing only 50% of requests initially
- Easy rollback by disabling compression

**[Risk] Compression adds latency** → Mitigation:
- Compression is CPU-bound and very fast (~10-50ms for typical photos)
- Much faster than network transmission time savings
- Measure actual latency impact in logs

**[Risk] Memory usage increases during compression** → Mitigation:
- Compression uses temporary buffers but Go's GC handles it well
- Typical photos (5MB) decode to ~30MB in memory temporarily
- Not a concern for backend with adequate RAM

**[Risk] Different formats compress differently** → Mitigation:
- PNG may not compress as much (already compressed)
- JPEG typically compresses well
- Log per-format metrics to track this
- WebP passes through (already efficient)
