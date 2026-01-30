// Port interface for camera access
export interface ICameraPort {
    /**
     * Request camera permission and access the camera stream
     * @param preferredFacingMode - 'user' for front camera, 'environment' for rear camera
     * @returns MediaStream if successful, null if denied/error
     */
    requestCameraAccess(preferredFacingMode?: 'user' | 'environment'): Promise<MediaStream | null>;

    /**
     * Stop the camera stream and release resources
     */
    stopCamera(): void;

    /**
     * Capture a photo from the current camera stream
     * @returns Blob containing the captured image in JPEG format
     */
    capturePhoto(): Promise<Blob | null>;

    /**
     * Check if camera access is supported in the current environment
     * @returns true if HTTPS or localhost, false otherwise
     */
    isCameraSupported(): boolean;

    /**
     * Get list of available camera devices
     */
    getAvailableCameras(): Promise<MediaDeviceInfo[]>;
}
