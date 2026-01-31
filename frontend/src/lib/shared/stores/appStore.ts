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

// Reset all state to initial values
export function resetAppState() {
    appState.set('requesting-permissions');
    uploadProgress.set(0);
    uploadError.set(null);
    currentVerdict.set(null);
}
