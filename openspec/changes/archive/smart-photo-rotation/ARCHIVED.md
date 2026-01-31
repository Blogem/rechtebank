# ARCHIVED

**Date:** January 31, 2026  
**Reason:** Superseded by `manual-photo-rotation-control` change

## Why Archived

This change implemented automatic photo rotation using DeviceOrientationEvent sensors (beta/gamma readings). The approach proved:

- **Complex**: Required intricate sensor math with multiple edge cases
- **Unreliable**: Inconsistent across iOS Safari and Android Chrome
- **Problematic**: Sensor permissions, inverted logic bugs, landscape ambiguity

## What Replaced It

The `manual-photo-rotation-control` change took a simpler approach:
- Native file input with camera capture (`<input type="file" capture="environment">`)
- Simple initial rotation heuristic using `screen.orientation.angle`
- Manual rotation controls for user correction
- Canvas-based rotation "baking"

**Result:** All orientation sensor code was eventually removed, including:
- OrientationAdapter (completely deleted)
- Spirit level feature (removed - didn't work with native camera)
- CameraAdapter (replaced by file input)
- All beta/gamma sensor logic

## Historical Value

This change documents the evolution of thinking:
- Initial belief that automatic detection was essential
- Deep investigation into sensor APIs and coordinate systems
- Discovery of fundamental issues with sensor-based approach
- Learning that simpler is better

The 235 completed tasks represent real work that informed the final design, even though the code was later deleted.
