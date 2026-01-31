/**
 * EXIF Orientation utilities
 * Handles EXIF orientation metadata to correctly orient images
 */

/**
 * Read EXIF orientation from image blob
 * Returns orientation value (1-8) or 1 (default/no rotation)
 */
export async function getExifOrientation(blob: Blob): Promise<number> {
    return new Promise((resolve) => {
        const reader = new FileReader();

        reader.onload = (e) => {
            const buffer = e.target?.result as ArrayBuffer;

            if (!buffer || buffer.byteLength < 2) {
                resolve(1);
                return;
            }

            const view = new DataView(buffer);

            if (view.getUint16(0, false) !== 0xFFD8) {
                // Not a JPEG
                resolve(1);
                return;
            }

            const length = view.byteLength;
            let offset = 2;

            while (offset < length - 4) {
                const marker = view.getUint16(offset, false);

                if ((marker & 0xFF00) !== 0xFF00) {
                    // Not a valid marker
                    break;
                }

                offset += 2;

                if (marker === 0xFFD9 || marker === 0xFFD8) {
                    // EOI or SOI marker (no length field)
                    break;
                }

                if (offset + 2 > length) {
                    break;
                }

                const segmentLength = view.getUint16(offset, false);

                if (segmentLength < 2 || offset + segmentLength > length) {
                    // Invalid segment size
                    break;
                }

                if (marker === 0xFFE1 && segmentLength > 10) {
                    // APP1 marker - check for EXIF
                    if (view.getUint32(offset + 2, false) === 0x45786966) {
                        // "Exif" identifier found
                        const exifStart = offset + 6; // Skip length + "Exif\0\0"

                        if (exifStart + 8 > length) {
                            break;
                        }

                        const little = view.getUint16(exifStart, false) === 0x4949;
                        const ifdOffset = exifStart + view.getUint32(exifStart + 4, little);

                        if (ifdOffset + 2 > length) {
                            break;
                        }

                        const tags = view.getUint16(ifdOffset, little);

                        for (let i = 0; i < tags; i++) {
                            const tagOffset = ifdOffset + 2 + (i * 12);
                            if (tagOffset + 12 > length) {
                                break;
                            }

                            const tag = view.getUint16(tagOffset, little);
                            if (tag === 0x0112) {
                                // Orientation tag found
                                const orientation = view.getUint16(tagOffset + 8, little);
                                resolve(orientation);
                                return;
                            }
                        }
                    }
                }

                // Move to next segment
                offset += segmentLength;
            }

            resolve(1); // Default orientation
        };

        reader.onerror = () => resolve(1);
        reader.readAsArrayBuffer(blob.slice(0, 64 * 1024)); // Read first 64KB
    });
}

/**
 * Apply EXIF orientation correction to an image blob
 * Returns a new blob with the image correctly oriented
 * 
 * Note: Modern browsers automatically apply EXIF orientation when displaying images.
 * This function strips EXIF data and renders the image in its corrected orientation
 * so that the backend receives a properly oriented image without EXIF metadata.
 */
export async function correctImageOrientation(blob: Blob): Promise<Blob> {
    return new Promise((resolve, reject) => {
        const img = new Image();
        const url = URL.createObjectURL(blob);

        img.onload = () => {
            const canvas = document.createElement('canvas');
            const ctx = canvas.getContext('2d');

            if (!ctx) {
                URL.revokeObjectURL(url);
                reject(new Error('Failed to get canvas context'));
                return;
            }

            // Modern browsers auto-correct EXIF orientation when loading images
            // The img.width and img.height already reflect the corrected dimensions
            // So we just use the natural dimensions without any transformation
            canvas.width = img.naturalWidth || img.width;
            canvas.height = img.naturalHeight || img.height;

            // Draw the image directly - browser has already applied EXIF orientation
            ctx.drawImage(img, 0, 0);
            URL.revokeObjectURL(url);

            canvas.toBlob(
                (correctedBlob) => {
                    if (correctedBlob) {
                        resolve(correctedBlob);
                    } else {
                        reject(new Error('Failed to create blob from corrected image'));
                    }
                },
                blob.type || 'image/jpeg',
                0.95
            );
        };

        img.onerror = () => {
            URL.revokeObjectURL(url);
            reject(new Error('Failed to load image for orientation correction'));
        };

        img.src = url;
    });
}
