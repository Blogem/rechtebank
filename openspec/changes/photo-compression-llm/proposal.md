## Why

Photos sent to the Gemini LLM API consume significant tokens, directly increasing API costs. By compressing photos before transmission while maintaining sufficient quality for furniture analysis, we can reduce token usage and lower operational costs without sacrificing verdict accuracy.

## What Changes

- Implement photo compression in the backend before sending to Gemini API
- Configure compression quality settings optimized for furniture detection (balancing file size vs. recognition accuracy)
- Add compression size/quality metrics to logging for monitoring effectiveness
- Validate that compressed photos still meet Gemini API image format requirements

## Capabilities

### New Capabilities
- `photo-compression`: Compress uploaded photos to reduce file size before LLM API calls while maintaining sufficient quality for furniture analysis

### Modified Capabilities
- `gemini-integration`: Add photo preprocessing step to compress images before sending to Gemini API

## Impact

- **Backend**: `internal/adapters/gemini/analyzer.go` - Add compression logic before API call
- **Storage**: Photo compression utilities/libraries (e.g., Go image encoding packages)
- **API Costs**: Reduced token consumption per request to Gemini API
- **Performance**: Slight processing overhead for compression (negligible compared to network/API latency)
- **Quality**: Need to validate compression doesn't negatively affect furniture detection accuracy
