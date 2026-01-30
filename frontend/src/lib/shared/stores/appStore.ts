import { writable, derived } from 'svelte/store';
import type { OrientationData, Verdict, PhotoMetadata } from '$lib/shared/types';

// App flow state machine states
export type AppState =
    | 'requesting-permissions'
    | 'camera-ready'
    | 'photo-captured'
    | 'uploading'
    | 'showing-verdict'
    | 'error';

// Main app state
export const appState = writable<AppState>('requesting-permissions');

// Camera state
export const cameraStream = writable<MediaStream | null>(null);
export const cameraPermissionGranted = writable<boolean>(false);
export const capturedPhoto = writable<Blob | null>(null);

// Orientation state
export const orientationData = writable<OrientationData | null>(null);
export const orientationPermissionGranted = writable<boolean>(false);
export const levelCheckEnabled = writable<boolean>(true); // Accessibility toggle

// Derived store: is device level?
export const isDeviceLevel = derived(
    [orientationData, levelCheckEnabled],
    ([$orientationData, $levelCheckEnabled]) => {
        if (!$levelCheckEnabled) return true; // Always level if check is disabled
        if (!$orientationData) return false;
        return $orientationData.isLevel;
    }
);

// Upload state
export const uploadProgress = writable<number>(0);
export const uploadError = writable<string | null>(null);

// Verdict state
export const currentVerdict = writable<Verdict | null>(null);

// Reset all state to initial values
export function resetAppState() {
    appState.set('requesting-permissions');
    cameraStream.set(null);
    capturedPhoto.set(null);
    orientationData.set(null);
    uploadProgress.set(0);
    uploadError.set(null);
    currentVerdict.set(null);
}
