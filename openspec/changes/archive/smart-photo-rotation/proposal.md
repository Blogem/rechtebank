## Why

Photos sent to Gemini for furniture analysis are sometimes oriented incorrectly (sideways or upside-down), reducing analysis accuracy. The spirit level feature ensures photos are level but doesn't address rotational orientation, and automatic EXIF-based detection doesn't work for camera captures.

## What Changes

- Add smart initial rotation detection using DeviceOrientationEvent beta reading at capture time
- Provide manual rotation controls (left/right buttons) in photo confirmation screen
- Apply canvas transformation before upload to ensure correctly oriented images reach Gemini
- Reset rotation state when capturing new photos

## Capabilities

### New Capabilities
- `device-orientation-detection`: Detect phone orientation (portrait vs landscape) using spirit level beta reading at moment of photo capture
- `manual-rotation-controls`: Provide UI controls for users to manually rotate photos left or right before submission
- `canvas-rotation-transform`: Apply rotation transformation to image canvas before upload to ensure correct orientation

### Modified Capabilities

## Impact

**Affected Code:**
- Frontend photo confirmation component (add rotation controls and preview)
- API adapter (add rotation transformation before upload)
- Orientation adapter (extend to capture beta reading at moment of capture)
- Photo state management (add rotation angle state)

**User Experience:**
- Most photos will have correct orientation automatically via smart heuristic
- Users see preview with applied rotation before submission
- Manual controls handle edge cases (floor photos, unusual angles)

**Dependencies:**
- No new external dependencies required
- Leverages existing DeviceOrientationEvent infrastructure from spirit level feature
