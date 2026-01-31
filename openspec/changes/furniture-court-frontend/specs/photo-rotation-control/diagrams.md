# Photo Rotation Flow Diagrams

## How Spirit Level Enables Smart Rotation Detection

```
┌─────────────────────────────────────────────────────────────────┐
│      SPIRIT LEVEL (DeviceOrientationEvent) FOR ORIENTATION       │
└─────────────────────────────────────────────────────────────────┘

Scenario 1: Phone VERTICAL (Portrait), camera pointing forward
───────────────────────────────────────────────────────────────
     ↑ Top
  ┌─────┐
  │ 📷  │ Phone held vertically
  │     │ Camera points forward (horizontal)
  │     │ Normal portrait photo stance
  └─────┘
  
  Gravity vector: ↓ (down)
  Phone orientation: VERTICAL

  DeviceOrientationEvent:
  • beta ≈ 90°  ← Phone is vertical!
  • gamma ≈ 0°
  • alpha = compass

  → HEURISTIC: beta > 45° = Portrait photo
  → Initial rotation = 0°


Scenario 2: Phone HORIZONTAL (Landscape), camera pointing forward
──────────────────────────────────────────────────────────────────
         ↑ Top
  ┌──────────────┐
  │ 📷           │ Phone held horizontally
  └──────────────┘ Camera points forward (horizontal)
                   Normal landscape photo stance
  
  Gravity vector: ↓ (down)
  Phone orientation: HORIZONTAL

  DeviceOrientationEvent:
  • beta ≈ 0°   ← Phone is horizontal!
  • gamma ≈ 0°
  • alpha = compass

  → HEURISTIC: beta ≤ 45° = Landscape photo
  → Initial rotation = 90° (assume needs correction)


Scenario 3: Phone VERTICAL, camera pointing DOWN (floor photo)
───────────────────────────────────────────────────────────────
  ┌─────┐
  │     │ Phone held vertically
  │ 📷  │ Camera points DOWN at floor
  │     │ Unusual but possible
  └─────┘
     ↓ Camera direction
  
  Gravity vector: ↓ (down, same direction as camera)
  Phone orientation: VERTICAL (but camera down)

  DeviceOrientationEvent:
  • beta ≈ 0°   ← Looks like landscape!
  • gamma ≈ 0°
  • alpha = compass

  → HEURISTIC: beta ≤ 45° = Landscape photo
  → Initial rotation = 90° (WRONG for floor photo)
  → USER ROTATES MANUALLY to fix


SMART HEURISTIC:
────────────────
if (beta > 45°) {
  rotation = 0°;  // Portrait photo (most common)
} else {
  rotation = 90°; // Landscape OR floor photo
}

User can manually adjust if initial guess is wrong.
```

## User Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                    PHOTO ROTATION USER FLOW                      │
└─────────────────────────────────────────────────────────────────┘

1. User Captures/Selects Photo
   │
   ├─── Camera: Takes photo, reads beta angle
   │             if beta > 45°: rotation = 0° (portrait)
   │             if beta ≤ 45°: rotation = 90° (landscape/floor)
   │
   └─── File: Selects existing photo from gallery
                rotation = 0° (default, no sensor data)
   │
   ▼

2. Confirmation Screen
   ┌──────────────────────────────────────┐
   │  📸 [Photo Preview]                   │
   │      transform: rotate(0deg)         │
   │                                      │
   │  Staat de foto niet goed?            │
   │  Roteer hem eerst:                   │
   │                                      │
   │  [↶ Links]    [↷ Rechts]            │
   │                                      │
   │  [🔄 Opnieuw]  [✅ Indienen]         │
   └──────────────────────────────────────┘
   │
   ├─── Photo looks correct?
   │    └─→ Click "Indienen" → Go to step 3
   │
   └─── Photo is sideways?
        ├─→ Click "↷ Rechts"
        │   rotation: 0° → 90° → 180° → 270° → 0° ...
        │   Visual: CSS transform updates instantly
        │
        └─→ When correct, click "Indienen" → Go to step 3
   │
   ▼

3. Upload with Rotation
   │
   ├─── rotation = 0°
   │    └─→ Send original image (no transform)
   │
   └─── rotation ≠ 0°
        └─→ Apply canvas transformation
            ├─ Create rotated canvas
            ├─ Draw image with rotation
            ├─ Export as JPEG (no EXIF)
            └─→ Send to backend
   │
   ▼

4. Backend → Gemini
   └─→ Correctly oriented image analyzed ✅
```

## Technical Flow

```
┌─────────────────────────────────────────────────────────────────┐
│              ROTATION TRANSFORMATION FLOW                        │
└─────────────────────────────────────────────────────────────────┘

Photo Blob (original)
   │
   ▼
┌─────────────────────────────────┐
│  ApiAdapter.uploadPhoto()       │
│  - Receives: photo, metadata,   │
│              rotation (0-270°)  │
└─────────────────────────────────┘
   │
   ├─── rotation = 0°?
   │    └─→ convertToJPEG() → Upload
   │        (no transformation)
   │
   └─── rotation ≠ 0°?
        └─→ applyRotation()
   │
   ▼
┌─────────────────────────────────┐
│  applyRotation(blob, rotation)  │
│                                 │
│  1. Load image from blob        │
│  2. Create canvas               │
│     needsSwap = (90° or 270°)   │
│     width = swap ? H : W        │
│     height = swap ? W : H       │
│  3. Apply transform             │
│     translate(cx, cy)           │
│     rotate(radians)             │
│     drawImage(centered)         │
│  4. Export as JPEG              │
└─────────────────────────────────┘
   │
   ▼
Rotated JPEG Blob
   │
   ▼
FormData → Backend → Gemini ✅
```

## Canvas Transformation Details

```
┌─────────────────────────────────────────────────────────────────┐
│           CANVAS TRANSFORMATION BY ROTATION ANGLE                │
└─────────────────────────────────────────────────────────────────┘

Original Image: 1080 (W) × 1920 (H) pixels


Rotation: 0° (No Change)
─────────────────────────
Canvas: 1080 × 1920
Transform: None
Result:
  ┌──────┐
  │      │
  │  ↑   │  Original orientation
  │      │
  └──────┘


Rotation: 90° (Clockwise)
─────────────────────────
Canvas: 1920 × 1080  ← Swapped!
Transform: rotate(90°)
Result:
  ┌───────────┐
  │     →     │  Rotated right
  └───────────┘


Rotation: 180° (Upside Down)
────────────────────────────
Canvas: 1080 × 1920
Transform: rotate(180°)
Result:
  ┌──────┐
  │      │
  │  ↓   │  Flipped
  │      │
  └──────┘


Rotation: 270° (Counter-Clockwise)
───────────────────────────────────
Canvas: 1920 × 1080  ← Swapped!
Transform: rotate(270°)
Result:
  ┌───────────┐
  │     ←     │  Rotated left
  └───────────┘
```

## State Management

```
┌─────────────────────────────────────────────────────────────────┐
│                    ROTATION STATE LIFECYCLE                      │
└─────────────────────────────────────────────────────────────────┘

photoRotation Store (Svelte writable)
───────────────────────────────────

Initial: 0°

Events that change rotation:
  ├─ handleRotate({ direction: 'right' })
  │  └─→ rotation = (current + 90) % 360
  │
  └─ handleRotate({ direction: 'left' })
     └─→ rotation = (current - 90 + 360) % 360

Events that reset rotation:
  ├─ handleCapture() - New photo captured
  ├─ handleRetake() - User clicked "Opnieuw"
  └─ resetAppState() - Full app reset

Read by:
  ├─ PhotoConfirmation.svelte
  │  └─→ <img style="transform: rotate({rotation}deg)" />
  │
  └─ +page.svelte
     └─→ apiAdapter.uploadPhoto(..., $photoRotation)
```

## Component Interaction

```
┌─────────────────────────────────────────────────────────────────┐
│                  COMPONENT COMMUNICATION                         │
└─────────────────────────────────────────────────────────────────┘

+page.svelte (Parent)
  │
  ├─ State: $photoRotation (from store)
  │
  ├─ Handlers:
  │   ├─ handleRotate(event)
  │   │   └─→ Updates photoRotation store
  │   │
  │   ├─ handleCapture()
  │   │   └─→ Resets photoRotation = 0°
  │   │
  │   └─ handleConfirm()
  │       └─→ Passes $photoRotation to API
  │
  └─ Renders:
      │
      ▼
    PhotoConfirmation.svelte (Child)
      │
      ├─ Props:
      │   ├─ photoUrl: string
      │   ├─ rotation: number  ← from $photoRotation
      │   ├─ onconfirm: handler
      │   ├─ onretake: handler
      │   └─ onrotate: handler  ← NEW
      │
      ├─ Emits:
      │   ├─ confirm → handleConfirm()
      │   ├─ retake → handleRetake()
      │   └─ rotate → handleRotate()  ← NEW
      │
      └─ Renders:
          ├─ <img style="transform: rotate({rotation}deg)" />
          ├─ Button: "↶ Links" → rotateLeft()
          └─ Button: "↷ Rechts" → rotateRight()
```

## Data Flow: Camera Capture Example

```
User taps "Capture Photo" button
  │
  ▼
CameraAdapter.capturePhoto()
  └─→ Returns Blob (JPEG, no EXIF)
  │
  ▼
+page.svelte.handleCapture()
  ├─ capturedPhoto.set(blob)
  ├─ photoRotation.set(0)  ← Reset
  └─ appState.set('photo-captured')
  │
  ▼
PhotoConfirmation renders
  └─ <img src={objectURL} style="transform: rotate(0deg)" />
  │
  ▼
User sees photo is sideways
  │
  ▼
User clicks "↷ Rechts"
  │
  ▼
rotateRight() → emits rotate event
  │
  ▼
handleRotate({ direction: 'right' })
  └─→ photoRotation: 0° → 90°
  │
  ▼
PhotoConfirmation re-renders
  └─ <img style="transform: rotate(90deg)" />  ← Visual feedback
  │
  ▼
User clicks "✅ Indienen"
  │
  ▼
handleConfirm()
  └─→ apiAdapter.uploadPhoto(blob, metadata, 90)
  │
  ▼
ApiAdapter.applyRotation(blob, 90)
  ├─ Canvas: 1920×1080 (swapped dimensions)
  ├─ Transform: rotate(90° * π/180)
  └─→ Returns rotated JPEG blob
  │
  ▼
FormData → POST /v1/judge
  │
  ▼
Backend → Gemini
  └─→ Receives correctly oriented image ✅
```
