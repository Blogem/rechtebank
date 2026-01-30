import { describe, it, expect, vi, beforeEach } from 'vitest';
import { OrientationAdapter } from './OrientationAdapter';

describe('OrientationAdapter', () => {
    let adapter: OrientationAdapter;

    beforeEach(() => {
        adapter = new OrientationAdapter();
    });

    describe('isOrientationSupported', () => {
        it('should return true when DeviceOrientationEvent is available', () => {
            expect(adapter.isOrientationSupported()).toBe(true);
        });
    });

    describe('startMonitoring', () => {
        it('should call callback with orientation data', () => {
            const callback = vi.fn();
            adapter.startMonitoring(callback);

            const event = new DeviceOrientationEvent('deviceorientation', {
                beta: 3,
                gamma: 2,
                alpha: 90
            });

            window.dispatchEvent(event);

            expect(callback).toHaveBeenCalledWith({
                beta: 3,
                gamma: 2,
                alpha: 90,
                isLevel: true // Within ±5° threshold
            });

            adapter.stopMonitoring();
        });

        it('should detect off-level state when beta > 5°', () => {
            const callback = vi.fn();
            adapter.startMonitoring(callback);

            const event = new DeviceOrientationEvent('deviceorientation', {
                beta: 10,
                gamma: 0,
                alpha: 0
            });

            window.dispatchEvent(event);

            expect(callback).toHaveBeenCalledWith({
                beta: 10,
                gamma: 0,
                alpha: 0,
                isLevel: false // Exceeds ±5° threshold
            });

            adapter.stopMonitoring();
        });
    });

    describe('stopMonitoring', () => {
        it('should remove event listener', () => {
            const callback = vi.fn();
            adapter.startMonitoring(callback);
            adapter.stopMonitoring();

            const event = new DeviceOrientationEvent('deviceorientation', {
                beta: 0,
                gamma: 0,
                alpha: 0
            });

            window.dispatchEvent(event);

            expect(callback).not.toHaveBeenCalled();
        });
    });
});
