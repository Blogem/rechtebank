import type { IApiPort, PhotoMetadata } from '../ports/IApiPort';
import type { Verdict } from '$lib/shared/types/Verdict';

export class ApiAdapter implements IApiPort {
    private apiBaseUrl: string;
    private maxRetries: number;
    private retryDelay: number;

    constructor(apiBaseUrl?: string, maxRetries = 3, retryDelay = 1000) {
        // Use environment variable or default to same host
        this.apiBaseUrl = apiBaseUrl || import.meta.env.PUBLIC_API_URL || '';
        this.maxRetries = maxRetries;
        this.retryDelay = retryDelay;
    }

    /**
     * Upload a photo to the backend for judgment
     */
    async uploadPhoto(photo: Blob, metadata: PhotoMetadata): Promise<Verdict> {
        // Ensure photo is JPEG format
        const jpegBlob = await this.convertToJPEG(photo);

        // Validate file size (10MB max)
        const maxSize = 10 * 1024 * 1024;
        if (jpegBlob.size > maxSize) {
            throw new Error('Foto is te groot. Maximaal 10MB toegestaan.');
        }

        // Prepare form data
        const formData = new FormData();
        formData.append('photo', jpegBlob, 'furniture.jpg');
        formData.append('userAgent', metadata.userAgent);
        formData.append('timestamp', metadata.timestamp);
        formData.append('captureMethod', metadata.captureMethod);

        // Upload with retry logic
        return this.uploadWithRetry(formData);
    }

    /**
     * Convert image blob to JPEG format
     */
    private async convertToJPEG(blob: Blob): Promise<Blob> {
        // If already JPEG, return as-is
        if (blob.type === 'image/jpeg') {
            return blob;
        }

        // Convert to JPEG using canvas
        return new Promise((resolve, reject) => {
            const img = new Image();
            const url = URL.createObjectURL(blob);

            img.onload = () => {
                const canvas = document.createElement('canvas');
                canvas.width = img.width;
                canvas.height = img.height;

                const ctx = canvas.getContext('2d');
                if (!ctx) {
                    reject(new Error('Failed to get canvas context'));
                    return;
                }

                ctx.drawImage(img, 0, 0);
                URL.revokeObjectURL(url);

                canvas.toBlob(
                    (jpegBlob) => {
                        if (jpegBlob) {
                            resolve(jpegBlob);
                        } else {
                            reject(new Error('Failed to convert image to JPEG'));
                        }
                    },
                    'image/jpeg',
                    0.9
                );
            };

            img.onerror = () => {
                URL.revokeObjectURL(url);
                reject(new Error('Failed to load image for conversion'));
            };

            img.src = url;
        });
    }

    /**
     * Upload with retry logic for network errors
     */
    private async uploadWithRetry(formData: FormData, attempt = 1): Promise<Verdict> {
        try {
            // Create abort controller for timeout
            const controller = new AbortController();
            const timeoutId = setTimeout(() => controller.abort(), 30000); // 30 second timeout

            try {
                const response = await fetch(`${this.apiBaseUrl}/v1/judge`, {
                    method: 'POST',
                    body: formData,
                    signal: controller.signal
                });

                clearTimeout(timeoutId);

                if (!response.ok) {
                    const errorText = await response.text();
                    throw new Error(`Server error (${response.status}): ${errorText}`);
                }

                const verdict: Verdict = await response.json();
                return verdict;
            } catch (fetchError) {
                clearTimeout(timeoutId);

                // Check if it's an abort error (timeout)
                if (fetchError instanceof DOMException && fetchError.name === 'AbortError') {
                    throw new Error('De rechter heeft te lang beraadslaagd. Probeer het opnieuw.');
                }
                if (fetchError instanceof Error && fetchError.name === 'AbortError') {
                    throw new Error('De rechter heeft te lang beraadslaagd. Probeer het opnieuw.');
                }

                throw fetchError;
            }
        } catch (error) {
            // Retry on network errors (but not on timeout errors)
            if (attempt < this.maxRetries && this.isNetworkError(error) &&
                !(error instanceof Error && error.message.includes('te lang beraadslaagd'))) {
                await this.delay(this.retryDelay * attempt);
                return this.uploadWithRetry(formData, attempt + 1);
            }

            // Re-throw error if max retries reached or non-network error
            throw error;
        }
    }

    /**
     * Check if error is a network error (retryable)
     */
    private isNetworkError(error: unknown): boolean {
        if (error instanceof TypeError) {
            // Network errors are often TypeErrors (e.g., "Failed to fetch")
            return true;
        }
        return false;
    }

    /**
     * Delay utility for retry logic
     */
    private delay(ms: number): Promise<void> {
        return new Promise((resolve) => setTimeout(resolve, ms));
    }
}
