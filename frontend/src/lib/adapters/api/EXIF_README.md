# EXIF Orientation Correction

## Overview

This module handles automatic EXIF orientation correction for uploaded photos. Mobile devices often save photos with their original pixel orientation and store the actual display orientation in EXIF metadata. This ensures images are correctly oriented before being sent to the backend for analysis.

## Features

- **Automatic EXIF Reading**: Reads orientation metadata from JPEG files
- **Orientation Correction**: Applies the correct transformation (rotation/flip) based on EXIF data
- **Browser-Based Processing**: All processing happens client-side before upload
- **Lightweight**: No external dependencies - pure JavaScript implementation

## EXIF Orientation Values

The EXIF orientation tag can have values from 1-8:

- **1**: Normal (0°)
- **2**: Horizontal flip
- **3**: 180° rotation
- **4**: Vertical flip
- **5**: Vertical flip + 90° CW
- **6**: 90° clockwise rotation
- **7**: Horizontal flip + 90° CW
- **8**: 90° counter-clockwise rotation

## Integration

The orientation correction is automatically applied in `ApiAdapter.uploadPhoto()` before converting the image to JPEG format:

```typescript
async uploadPhoto(photo: Blob, metadata: PhotoMetadata): Promise<Verdict> {
    // First, correct image orientation based on EXIF data
    const orientedPhoto = await correctImageOrientation(photo);
    
    // Then convert to JPEG and upload
    const jpegBlob = await this.convertToJPEG(orientedPhoto);
    // ...
}
```

## Usage

### Direct Usage

```typescript
import { correctImageOrientation, getExifOrientation } from './exifOrientation';

// Get EXIF orientation value
const orientation = await getExifOrientation(photoBlob);
console.log('EXIF orientation:', orientation); // 1-8

// Correct image orientation
const correctedBlob = await correctImageOrientation(photoBlob);
```

### Automatic (via ApiAdapter)

No manual intervention needed - orientation correction happens automatically when uploading photos through the API adapter.

## Implementation Details

### EXIF Parsing

- Reads first 64KB of image file (sufficient for EXIF data)
- Parses JPEG markers to find APP1 (EXIF) segment
- Extracts orientation tag (0x0112) from TIFF IFD
- Handles both little-endian and big-endian byte order
- Defaults to orientation 1 (normal) if no EXIF data found

### Canvas Transformation

- Creates HTML5 Canvas element
- Applies appropriate transformation matrix based on orientation
- Draws corrected image
- Returns new Blob with proper orientation

### Performance

- Only processes images that have EXIF orientation data
- Images with orientation 1 (normal) are returned as-is
- Canvas operations are optimized for minimal overhead
- Quality preserved at 95% during re-encoding

## Browser Support

Works in all modern browsers that support:
- Canvas API
- FileReader API
- Blob API

## Testing

Comprehensive test coverage in `exifOrientation.test.ts`:
- EXIF reading for various image formats
- Error handling for malformed data
- Edge cases (empty blobs, non-JPEG files)
- Integration with image correction pipeline

Run tests:
```bash
npm test -- exifOrientation.test.ts
```
