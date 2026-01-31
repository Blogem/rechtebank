import { describe, it, expect } from 'vitest';
import { get } from 'svelte/store';
import { appState, resetAppState } from './appStore';

describe('appStore', () => {
    it('should initialize appState to requesting-permissions', () => {
        expect(get(appState)).toBe('requesting-permissions');
    });

    it('should reset appState when resetAppState is called', () => {
        appState.set('showing-verdict');
        expect(get(appState)).toBe('showing-verdict');

        resetAppState();
        expect(get(appState)).toBe('requesting-permissions');
    });
});
