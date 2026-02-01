## ADDED Requirements

### Requirement: Compress JPEG images
The system SHALL compress JPEG images to quality 75 before sending to the LLM API.

#### Scenario: JPEG compression successful
- **WHEN** system receives a JPEG image for LLM analysis
- **THEN** image is re-encoded at quality 75 and compressed size is logged

#### Scenario: JPEG already smaller than threshold
- **WHEN** compressed JPEG is larger than original
- **THEN** system uses the original image and logs the occurrence

### Requirement: Compress PNG images
The system SHALL compress PNG images using BestSpeed compression level before sending to the LLM API.

#### Scenario: PNG compression successful
- **WHEN** system receives a PNG image for LLM analysis
- **THEN** image is re-encoded with BestSpeed compression and compressed size is logged

#### Scenario: PNG already optimized
- **WHEN** compressed PNG is larger than original
- **THEN** system uses the original image and logs the occurrence

### Requirement: Resize oversized images
The system SHALL resize images to a maximum dimension of 1600 pixels while maintaining aspect ratio.

#### Scenario: Image exceeds maximum dimension
- **WHEN** either width or height exceeds 1600 pixels
- **THEN** image is resized proportionally to fit within 1600x1600 boundary

#### Scenario: Image within size limit
- **WHEN** both width and height are 1600 pixels or less
- **THEN** image dimensions remain unchanged

#### Scenario: Aspect ratio preserved
- **WHEN** image is resized
- **THEN** original aspect ratio is maintained

### Requirement: Pass through WebP images
The system SHALL not compress WebP images as they are already efficiently compressed.

#### Scenario: WebP image received
- **WHEN** system receives a WebP image for LLM analysis
- **THEN** image is sent to LLM without compression and original size is logged

### Requirement: Fail gracefully on compression errors
The system SHALL use the original uncompressed image if compression fails.

#### Scenario: Compression fails
- **WHEN** image compression encounters an error
- **THEN** system logs the error, uses the original image, and continues processing

#### Scenario: Invalid image format
- **WHEN** image cannot be decoded for compression
- **THEN** system logs the error and uses the original image data

### Requirement: Log compression metrics
The system SHALL log compression metrics for monitoring and optimization.

#### Scenario: Successful compression
- **WHEN** image is compressed successfully
- **THEN** log includes originalSize, compressedSize, compressionRatio, and imageFormat

#### Scenario: Compression skipped
- **WHEN** compression is skipped or fails
- **THEN** log includes reason for skipping compression
