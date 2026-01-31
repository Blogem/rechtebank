# Photo Rotation Control

Manual photo rotation controls to ensure images are correctly oriented for Gemini AI analysis.

## The Problem

Photos captured via camera or uploaded from files can appear sideways or upside-down to Gemini, reducing analysis accuracy. This happens because:

- **Camera captures** have no EXIF metadata, orientation depends on device sensor
- **File uploads** lose EXIF metadata when converted to JPEG via canvas
- **Auto-detection** is unreliable (varies by device, browser, platform)

## The Solution

Instead of trying to auto-detect orientation (which fails across devices), we:

1. Show the photo to the user as Gemini will see it
2. Provide rotation buttons if it's wrong (↶ Links / ↷ Rechts)
3. Apply canvas transformation before upload

Simple, reliable, works for all cases.

## Documents

- **[spec.md](spec.md)** - Complete capability specification with requirements and scenarios
- **[decision-record.md](decision-record.md)** - Analysis of the problem, options considered, and why we chose this approach
- **[implementation-summary.md](implementation-summary.md)** - Code changes needed, technical details, implementation checklist

## Key Points

- Default rotation: 0° (no transformation)
- Rotation increments: 90° (0°, 90°, 180°, 270°)
- Visual feedback: CSS transform (instant)
- Upload: Canvas transformation (pixel-level)
- Reset: On new photo capture or retake

## Status

- ✅ Spec documented
- ✅ Design decision recorded
- ⏸️ Implementation started (stores, API interface)
- ⏸️ UI pending (PhotoConfirmation + page handlers)
- ⏸️ Testing pending

See `implementation-summary.md` for completion checklist.
