import { describe, it, expect } from 'vitest';
import { getExifOrientation } from './exifOrientation';

describe('EXIF Orientation', () => {
    describe('getExifOrientation', () => {
        it('should return 1 for non-JPEG blobs', async () => {
            const blob = new Blob(['not a jpeg'], { type: 'text/plain' });
            const orientation = await getExifOrientation(blob);
            expect(orientation).toBe(1);
        });

        it('should return 1 for JPEG without EXIF data', async () => {
            // Minimal JPEG header without EXIF
            const jpegData = new Uint8Array([
                0xFF, 0xD8, // JPEG SOI marker
                0xFF, 0xE0, // APP0 marker
                0x00, 0x10, // Segment length (16 bytes)
                0x4A, 0x46, 0x49, 0x46, 0x00, // "JFIF"
                0x01, 0x01, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00,
                0xFF, 0xDB, // DQT marker
                0x00, 0x09, // Length
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0xFF, 0xD9  // JPEG EOI marker
            ]);
            const blob = new Blob([jpegData], { type: 'image/jpeg' });
            const orientation = await getExifOrientation(blob);
            expect(orientation).toBe(1);
        });

        it('should return 1 for empty blob', async () => {
            const blob = new Blob([], { type: 'image/jpeg' });
            const orientation = await getExifOrientation(blob);
            expect(orientation).toBe(1);
        });

        it('should handle small blobs gracefully', async () => {
            const blob = new Blob([new Uint8Array([0xFF, 0xD8])], { type: 'image/jpeg' });
            const orientation = await getExifOrientation(blob);
            expect(orientation).toBe(1);
        });
    });
});
