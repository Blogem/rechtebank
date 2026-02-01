## 1. JPEG Compression (TDD)

- [x] 1.1 Write test: JPEG compression produces smaller output at quality 75
- [x] 1.2 Write test: Fallback to original when compressed JPEG is larger
- [x] 1.3 Implement JPEG compression with quality 75 using `image/jpeg` package
- [x] 1.4 Implement size comparison logic to use original if compressed is larger
- [x] 1.5 Verify tests pass

## 2. PNG Compression (TDD)

- [x] 2.1 Write test: PNG compression with BestSpeed level produces smaller output
- [x] 2.2 Write test: Fallback to original when compressed PNG is larger
- [x] 2.3 Implement PNG compression with BestSpeed level using `image/png` package
- [x] 2.4 Verify tests pass

## 3. WebP Pass-through (TDD)

- [x] 3.1 Write test: WebP images pass through unchanged
- [x] 3.2 Implement WebP pass-through (return original bytes unchanged)
- [x] 3.3 Verify test passes

## 4. Image Resizing (TDD)

- [x] 4.1 Write test: Resize images over 1600px in width or height
- [x] 4.2 Write test: Resize maintains aspect ratio
- [x] 4.3 Write test: Images within size limit remain unchanged
- [x] 4.4 Create `resizeIfNeeded()` function that checks dimensions and resizes to max 1600px
- [x] 4.5 Implement proportional scaling maintaining aspect ratio
- [x] 4.6 Verify tests pass

## 5. Compression Utility Integration (TDD)

- [x] 5.1 Write test: `compressImage()` routes to correct handler based on MIME type
- [x] 5.2 Write test: Error handling when image decode fails
- [x] 5.3 Create `compressImage()` function that takes imageData []byte and returns compressed []byte
- [x] 5.4 Implement MIME type detection to route to correct compression handler
- [x] 5.5 Integrate resizing before compression in the pipeline
- [x] 5.6 Add error handling to fall back to original image on failure
- [x] 5.7 Verify tests pass

## 6. Compression Logging (TDD)

- [x] 6.1 Write test: Compression metrics are logged correctly
- [x] 6.2 Write test: Log reason when compression is skipped or fails
- [x] 6.3 Add structured log fields for originalSize, compressedSize, compressionRatio, and imageFormat
- [x] 6.4 Log compression metrics on successful compression
- [x] 6.5 Log reason when compression is skipped or fails
- [x] 6.6 Verify tests pass

## 7. Gemini Flow Integration (TDD)

- [x] 7.1 Write test: `GenerateContent()` compresses imageData before Gemini API call
- [x] 7.2 Write test: Compressed image still has correct MIME type detection
- [x] 7.3 Modify `GenerateContent()` to compress imageData before sending to Gemini API
- [x] 7.4 Ensure compressed image still has correct MIME type detection
- [x] 7.5 Verify tests pass

## 8. Integration Testing

- [x] 8.1 Test end-to-end with real photos through Gemini API
- [x] 8.2 Verify verdict accuracy is maintained with compressed images
- [x] 8.3 Test all three image formats (JPEG, PNG, WebP) in integration test
- [x] 8.4 Validate compression logs appear correctly in test output

## 9. Documentation

- [x] 9.1 Add code comments explaining compression strategy and quality settings
- [x] 9.2 Document configuration constants (JPEG_QUALITY, MAX_DIMENSION)
