/**
 * Canvas-based image rotation utilities for manual photo orientation control.
 * Provides functions for detecting initial rotation and rotating images using HTML Canvas API.
 */

/**
 * Get initial rotation angle using screen orientation API.
 * Attempts to detect device orientation with fallbacks:
 * 1. screen.orientation.angle (modern API)
 * 2. window.orientation (legacy iOS)
 * 3. Default to 0 if unavailable
 * 
 * @returns Rotation angle in degrees (0, 90, 180, or 270)
 */
export function getInitialRotation(): number {
    // Try modern screen.orientation API
    if (typeof screen !== 'undefined' && screen.orientation?.angle !== undefined) {
        return screen.orientation.angle;
    }

    // Fallback to legacy window.orientation (iOS)
    if (typeof window !== 'undefined' && (window as any).orientation !== undefined) {
        const angle = (window as any).orientation;
        // Normalize negative angles (iOS uses -90, 0, 90, 180)
        return angle < 0 ? 360 + angle : angle;
    }

    // Default: no rotation
    return 0;
}

/**
 * Rotate an image on canvas and return the rotated canvas.
 * Properly handles dimension swapping for portrait rotations (90°, 270°)
 * and applies coordinate transformations to keep the image centered.
 * 
 * @param img - HTMLImageElement to rotate
 * @param rotation - Rotation angle in degrees (0, 90, 180, or 270)
 * @returns HTMLCanvasElement with rotated image
 */
export function rotateImageOnCanvas(img: HTMLImageElement, rotation: number): HTMLCanvasElement {
    const canvas = document.createElement('canvas');
    const ctx = canvas.getContext('2d');

    if (!ctx) {
        throw new Error('Failed to get canvas 2D context');
    }

    // Determine if we need to swap dimensions (portrait rotations)
    const isPortrait = rotation === 90 || rotation === 270;

    // Set canvas dimensions (swap for portrait)
    canvas.width = isPortrait ? img.height : img.width;
    canvas.height = isPortrait ? img.width : img.height;

    // Move origin to center of canvas
    ctx.translate(canvas.width / 2, canvas.height / 2);

    // Rotate around center
    ctx.rotate((rotation * Math.PI) / 180);

    // Draw image centered at origin
    ctx.drawImage(img, -img.width / 2, -img.height / 2);

    return canvas;
}

/**
 * Convert canvas to Blob with JPEG encoding.
 * Uses quality parameter of 0.9 for good quality/size balance.
 * 
 * @param canvas - HTMLCanvasElement to convert
 * @returns Promise<Blob> - JPEG blob
 */
export function canvasToBlob(canvas: HTMLCanvasElement): Promise<Blob> {
    return new Promise((resolve, reject) => {
        canvas.toBlob(
            (blob) => {
                if (blob) {
                    resolve(blob);
                } else {
                    reject(new Error('Failed to convert canvas to blob'));
                }
            },
            'image/jpeg',
            0.9
        );
    });
}
