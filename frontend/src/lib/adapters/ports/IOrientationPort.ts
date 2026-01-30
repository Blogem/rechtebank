// Port interface for device orientation access
export interface IOrientationPort {
    /**
     * Request permission to access device orientation (iOS 13+ requirement)
     * @returns true if permission granted, false otherwise
     */
    requestOrientationPermission(): Promise<boolean>;

    /**
     * Start monitoring device orientation
     * @param callback - Function to call with orientation data updates
     */
    startMonitoring(callback: (data: OrientationData) => void): void;

    /**
     * Stop monitoring device orientation
     */
    stopMonitoring(): void;

    /**
     * Check if device orientation is supported
     */
    isOrientationSupported(): boolean;

    /**
     * Check if iOS 13+ permission request is required
     */
    requiresPermission(): boolean;
}

export interface OrientationData {
    /** Front-to-back tilt in degrees (-180 to 180) */
    beta: number;
    /** Left-to-right tilt in degrees (-90 to 90) */
    gamma: number;
    /** Compass direction in degrees (0 to 360) */
    alpha: number;
    /** Whether the device is within ±5° of level */
    isLevel: boolean;
}
