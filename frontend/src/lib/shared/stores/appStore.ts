import { writable } from 'svelte/store';
import type { Verdict } from '$lib/shared/types';

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

// Upload state
export const uploadProgress = writable<number>(0);
export const uploadError = writable<string | null>(null);

// Verdict state
export const currentVerdict = writable<Verdict | null>(null);

// Track if permissions have been granted
let permissionsGranted = false;

export function setPermissionsGranted(granted: boolean) {
    permissionsGranted = granted;
}

// Reset all state to initial values
export function resetAppState() {
    // If permissions were already granted, go straight to camera-ready
    // Otherwise, go back to requesting permissions
    appState.set(permissionsGranted ? 'camera-ready' : 'requesting-permissions');
    uploadProgress.set(0);
    uploadError.set(null);
    currentVerdict.set(null);
}
