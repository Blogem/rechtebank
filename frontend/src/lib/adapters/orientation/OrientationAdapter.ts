import type { IOrientationPort, OrientationData } from '../ports/IOrientationPort';

const LEVEL_THRESHOLD = 5; // ±5 degrees

export class OrientationAdapter implements IOrientationPort {
    private callback: ((data: OrientationData) => void) | null = null;
    private orientationHandler: ((event: DeviceOrientationEvent) => void) | null = null;

    /**
     * Check if device orientation API is supported
     */
    isOrientationSupported(): boolean {
        return 'DeviceOrientationEvent' in window;
    }

    /**
     * Check if iOS 13+ permission request is required
     */
    requiresPermission(): boolean {
        // iOS 13+ requires explicit permission for DeviceOrientationEvent
        return typeof (DeviceOrientationEvent as any).requestPermission === 'function';
    }

    /**
     * Request orientation permission (iOS 13+ only)
     */
    async requestOrientationPermission(): Promise<boolean> {
        if (!this.requiresPermission()) {
            return true; // Permission not required on non-iOS devices
        }

        try {
            const permission = await (DeviceOrientationEvent as any).requestPermission();
            return permission === 'granted';
        } catch (error) {
            console.error('Failed to request orientation permission:', error);
            return false;
        }
    }

    /**
     * Start monitoring device orientation
     */
    startMonitoring(callback: (data: OrientationData) => void): void {
        if (!this.isOrientationSupported()) {
            console.error('Device orientation not supported');
            return;
        }

        this.callback = callback;

        this.orientationHandler = (event: DeviceOrientationEvent) => {
            const beta = event.beta ?? 0; // Front-to-back tilt (-180 to 180)
            const gamma = event.gamma ?? 0; // Left-to-right tilt (-90 to 90)
            const alpha = event.alpha ?? 0; // Compass direction (0 to 360)

            const isLevel = Math.abs(beta) <= LEVEL_THRESHOLD;

            const data: OrientationData = {
                beta,
                gamma,
                alpha,
                isLevel
            };

            if (this.callback) {
                this.callback(data);
            }
        };

        window.addEventListener('deviceorientation', this.orientationHandler);
    }

    /**
     * Stop monitoring device orientation
     */
    stopMonitoring(): void {
        if (this.orientationHandler) {
            window.removeEventListener('deviceorientation', this.orientationHandler);
            this.orientationHandler = null;
            this.callback = null;
        }
    }
}
