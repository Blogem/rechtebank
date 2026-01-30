import type { ICameraPort } from '../ports/ICameraPort';

export class CameraAdapter implements ICameraPort {
    private currentStream: MediaStream | null = null;
    private videoElement: HTMLVideoElement | null = null;

    /**
     * Check if camera access is supported (HTTPS or localhost)
     */
    isCameraSupported(): boolean {
        // Check if MediaDevices API is available
        if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
            return false;
        }

        // Camera requires HTTPS or localhost
        const protocol = window.location.protocol;
        const hostname = window.location.hostname;

        return protocol === 'https:' || hostname === 'localhost' || hostname === '127.0.0.1';
    }

    /**
     * Request camera access with optional facing mode preference
     */
    async requestCameraAccess(
        preferredFacingMode: 'user' | 'environment' = 'environment'
    ): Promise<MediaStream | null> {
        if (!this.isCameraSupported()) {
            console.error('Camera access requires HTTPS or localhost');
            return null;
        }

        try {
            // Try with ideal facingMode first (works better on Android)
            const constraints: MediaStreamConstraints = {
                video: {
                    facingMode: { ideal: preferredFacingMode },
                    width: { ideal: 1920 },
                    height: { ideal: 1080 }
                },
                audio: false
            };

            const stream = await navigator.mediaDevices.getUserMedia(constraints);
            this.currentStream = stream;
            return stream;
        } catch (error) {
            // If ideal constraints fail, try with minimal constraints
            console.warn('Failed with ideal constraints, trying basic constraints:', error);
            try {
                const basicConstraints: MediaStreamConstraints = {
                    video: true,
                    audio: false
                };
                
                const stream = await navigator.mediaDevices.getUserMedia(basicConstraints);
                this.currentStream = stream;
                return stream;
            } catch (fallbackError) {
                console.error('Failed to access camera:', fallbackError);
                return null;
            }
        }
    }

    /**
     * Stop the camera stream and release resources
     */
    stopCamera(): void {
        if (this.currentStream) {
            this.currentStream.getTracks().forEach((track) => track.stop());
            this.currentStream = null;
        }

        if (this.videoElement) {
            this.videoElement.srcObject = null;
        }
    }

    /**
     * Capture a photo from the current camera stream
     */
    async capturePhoto(): Promise<Blob | null> {
        if (!this.currentStream) {
            console.error('No active camera stream');
            return null;
        }

        try {
            // Create a video element to capture the frame
            if (!this.videoElement) {
                this.videoElement = document.createElement('video');
                this.videoElement.srcObject = this.currentStream;
                this.videoElement.play();
            }

            // Wait for video to be ready
            await new Promise((resolve) => {
                if (this.videoElement!.readyState >= 2) {
                    resolve(true);
                } else {
                    this.videoElement!.addEventListener('loadeddata', () => resolve(true), { once: true });
                }
            });

            // Create canvas and capture frame
            const canvas = document.createElement('canvas');
            canvas.width = this.videoElement.videoWidth;
            canvas.height = this.videoElement.videoHeight;

            const context = canvas.getContext('2d');
            if (!context) {
                throw new Error('Failed to get canvas context');
            }

            context.drawImage(this.videoElement, 0, 0, canvas.width, canvas.height);

            // Convert to JPEG blob
            return new Promise((resolve) => {
                canvas.toBlob(
                    (blob) => {
                        resolve(blob);
                    },
                    'image/jpeg',
                    0.9
                );
            });
        } catch (error) {
            console.error('Failed to capture photo:', error);
            return null;
        }
    }

    /**
     * Get list of available camera devices
     */
    async getAvailableCameras(): Promise<MediaDeviceInfo[]> {
        try {
            const devices = await navigator.mediaDevices.enumerateDevices();
            return devices.filter((device) => device.kind === 'videoinput');
        } catch (error) {
            console.error('Failed to enumerate devices:', error);
            return [];
        }
    }
}
