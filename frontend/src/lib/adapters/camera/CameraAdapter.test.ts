import { describe, it, expect, vi, beforeEach } from 'vitest';
import { CameraAdapter } from './CameraAdapter.js';

describe('CameraAdapter', () => {
    let adapter: CameraAdapter;
    let mockGetUserMedia: ReturnType<typeof vi.fn>;
    let mockEnumerateDevices: ReturnType<typeof vi.fn>;

    beforeEach(() => {
        mockGetUserMedia = vi.fn();
        mockEnumerateDevices = vi.fn();

        // Mock navigator.mediaDevices
        Object.defineProperty(globalThis.navigator, 'mediaDevices', {
            value: {
                getUserMedia: mockGetUserMedia,
                enumerateDevices: mockEnumerateDevices
            },
            writable: true,
            configurable: true
        });

        // Reset adapter after mocks are set up
        adapter = new CameraAdapter();
    });

    describe('isCameraSupported', () => {
        it('should return true for HTTPS connections', () => {
            Object.defineProperty(window, 'location', {
                value: { protocol: 'https:' },
                writable: true
            });

            expect(adapter.isCameraSupported()).toBe(true);
        });

        it('should return true for localhost HTTP', () => {
            Object.defineProperty(window, 'location', {
                value: { protocol: 'http:', hostname: 'localhost' },
                writable: true
            });

            expect(adapter.isCameraSupported()).toBe(true);
        });

        it('should return false for non-localhost HTTP', () => {
            Object.defineProperty(window, 'location', {
                value: { protocol: 'http:', hostname: 'example.com' },
                writable: true
            });

            expect(adapter.isCameraSupported()).toBe(false);
        });
    });

    describe('requestCameraAccess', () => {
        it('should request camera permission with rear camera preference', async () => {
            const mockStream = { id: 'test-stream' } as MediaStream;
            mockGetUserMedia.mockResolvedValue(mockStream);

            // Mock HTTPS location
            Object.defineProperty(window, 'location', {
                value: {
                    protocol: 'https:',
                    hostname: 'example.com'
                },
                writable: true,
                configurable: true
            });

            const stream = await adapter.requestCameraAccess('environment');

            expect(mockGetUserMedia).toHaveBeenCalledWith({
                video: {
                    facingMode: { ideal: 'environment' },
                    width: { ideal: 1920 },
                    height: { ideal: 1080 }
                },
                audio: false
            });
            expect(stream).toBe(mockStream);
        });

        it('should request front camera when specified', async () => {
            const mockStream = { id: 'test-stream' } as MediaStream;
            mockGetUserMedia.mockResolvedValue(mockStream);

            // Mock HTTPS location
            Object.defineProperty(window, 'location', {
                value: {
                    protocol: 'https:',
                    hostname: 'example.com'
                },
                writable: true,
                configurable: true
            });

            await adapter.requestCameraAccess('user');

            expect(mockGetUserMedia).toHaveBeenCalledWith({
                video: {
                    facingMode: { ideal: 'user' },
                    width: { ideal: 1920 },
                    height: { ideal: 1080 }
                },
                audio: false
            });
        });

        it('should return null on permission denial', async () => {
            mockGetUserMedia.mockRejectedValue(new Error('Permission denied'));

            const stream = await adapter.requestCameraAccess();

            expect(stream).toBeNull();
        });
    });

    describe('getAvailableCameras', () => {
        it('should return list of video input devices', async () => {
            const mockDevices = [
                { kind: 'videoinput', deviceId: '1', label: 'Front Camera' },
                { kind: 'videoinput', deviceId: '2', label: 'Rear Camera' },
                { kind: 'audioinput', deviceId: '3', label: 'Microphone' }
            ] as MediaDeviceInfo[];

            mockEnumerateDevices.mockResolvedValue(mockDevices);

            const cameras = await adapter.getAvailableCameras();

            expect(cameras).toHaveLength(2);
            expect(cameras[0].kind).toBe('videoinput');
        });
    });
});
