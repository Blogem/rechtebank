import type { Verdict } from '$lib/shared/types/Verdict';

// Port interface for backend API communication
export interface IApiPort {
    /**
     * Upload a photo to the backend for judgment
     * @param photo - Image file blob
     * @param metadata - Additional metadata to send with the photo
     * @returns Verdict response from the backend
     */
    uploadPhoto(photo: Blob, metadata: PhotoMetadata): Promise<Verdict>;
}

export interface PhotoMetadata {
    /** Browser user agent string */
    userAgent: string;
    /** Timestamp when photo was captured/selected */
    timestamp: string;
    /** Whether photo was captured via camera or file upload */
    captureMethod: 'camera' | 'file';
}
