import { describe, it, expect, beforeEach, afterEach, vi } from 'vitest';
import { getInitialRotation, rotateImageOnCanvas, canvasToBlob } from './rotationUtils';

describe('rotationUtils', () => {
    describe('getInitialRotation', () => {
        let originalScreen: typeof globalThis.screen;
        let originalOrientation: number | undefined;

        beforeEach(() => {
            originalScreen = globalThis.screen;
            originalOrientation = (globalThis.window as any).orientation;
        });

        afterEach(() => {
            globalThis.screen = originalScreen;
            if (originalOrientation !== undefined) {
                (globalThis.window as any).orientation = originalOrientation;
            } else {
                delete (globalThis.window as any).orientation;
            }
        });

        it('should return screen.orientation.angle when available', () => {
            // Mock screen.orientation.angle
            Object.defineProperty(globalThis.screen, 'orientation', {
                value: { angle: 90 },
                configurable: true
            });

            const rotation = getInitialRotation();
            expect(rotation).toBe(90);
        });

        it('should fallback to window.orientation when screen.orientation.angle is unavailable', () => {
            // Remove screen.orientation
            Object.defineProperty(globalThis.screen, 'orientation', {
                value: undefined,
                configurable: true
            });

            // Set window.orientation
            (globalThis.window as any).orientation = 90;

            const rotation = getInitialRotation();
            expect(rotation).toBe(90);
        });

        it('should normalize negative window.orientation values', () => {
            // Remove screen.orientation
            Object.defineProperty(globalThis.screen, 'orientation', {
                value: undefined,
                configurable: true
            });

            // Set window.orientation to -90 (iOS convention)
            (globalThis.window as any).orientation = -90;

            const rotation = getInitialRotation();
            expect(rotation).toBe(270); // -90 + 360 = 270
        });

        it('should default to 0 when no orientation API is available', () => {
            // Remove screen.orientation
            Object.defineProperty(globalThis.screen, 'orientation', {
                value: undefined,
                configurable: true
            });

            // Remove window.orientation
            delete (globalThis.window as any).orientation;

            const rotation = getInitialRotation();
            expect(rotation).toBe(0);
        });

        it('should normalize rotation angles to 0, 90, 180, 270', () => {
            // Test various angles that should normalize
            Object.defineProperty(globalThis.screen, 'orientation', {
                value: { angle: 0 },
                configurable: true
            });
            expect(getInitialRotation()).toBe(0);

            Object.defineProperty(globalThis.screen, 'orientation', {
                value: { angle: 90 },
                configurable: true
            });
            expect(getInitialRotation()).toBe(90);

            Object.defineProperty(globalThis.screen, 'orientation', {
                value: { angle: 180 },
                configurable: true
            });
            expect(getInitialRotation()).toBe(180);

            Object.defineProperty(globalThis.screen, 'orientation', {
                value: { angle: 270 },
                configurable: true
            });
            expect(getInitialRotation()).toBe(270);
        });
    });

    describe('rotateImageOnCanvas', () => {
        let mockImage: HTMLImageElement;
        let mockCanvas: HTMLCanvasElement;
        let mockContext: CanvasRenderingContext2D;

        beforeEach(() => {
            // Create a mock image
            mockImage = new Image();
            mockImage.width = 800;
            mockImage.height = 600;

            // Create mock canvas and context
            mockContext = {
                translate: vi.fn(),
                rotate: vi.fn(),
                drawImage: vi.fn()
            } as any;

            mockCanvas = {
                width: 0,
                height: 0,
                getContext: vi.fn(() => mockContext)
            } as any;

            // Mock document.createElement to return our mock canvas
            vi.spyOn(document, 'createElement').mockImplementation((tagName) => {
                if (tagName === 'canvas') {
                    return mockCanvas;
                }
                return document.createElement(tagName);
            });
        });

        afterEach(() => {
            vi.restoreAllMocks();
        });

        it('should not swap dimensions at 0° rotation', () => {
            const result = rotateImageOnCanvas(mockImage, 0);

            expect(result.width).toBe(800);
            expect(result.height).toBe(600);
        });

        it('should swap dimensions at 90° rotation', () => {
            const result = rotateImageOnCanvas(mockImage, 90);

            expect(result.width).toBe(600); // swapped
            expect(result.height).toBe(800); // swapped
        });

        it('should not swap dimensions at 180° rotation', () => {
            const result = rotateImageOnCanvas(mockImage, 180);

            expect(result.width).toBe(800);
            expect(result.height).toBe(600);
        });

        it('should swap dimensions at 270° rotation', () => {
            const result = rotateImageOnCanvas(mockImage, 270);

            expect(result.width).toBe(600); // swapped
            expect(result.height).toBe(800); // swapped
        });

        it('should apply proper canvas transformations', () => {
            const result = rotateImageOnCanvas(mockImage, 90);

            // Canvas should be created with swapped dimensions
            expect(result.width).toBe(600);
            expect(result.height).toBe(800);

            // Verify transformations were applied
            expect(mockContext.translate).toHaveBeenCalledWith(300, 400); // width/2, height/2
            expect(mockContext.rotate).toHaveBeenCalledWith((90 * Math.PI) / 180);
            expect(mockContext.drawImage).toHaveBeenCalledWith(mockImage, -400, -300); // -img.width/2, -img.height/2
        });
    });

    describe('canvasToBlob', () => {
        let canvas: HTMLCanvasElement;

        beforeEach(() => {
            canvas = document.createElement('canvas');
            canvas.width = 100;
            canvas.height = 100;
        });

        it('should convert canvas to JPEG blob with quality parameter', async () => {
            // Mock toBlob to call callback immediately
            const mockBlob = new Blob(['fake'], { type: 'image/jpeg' });
            vi.spyOn(canvas, 'toBlob').mockImplementation((callback) => {
                callback?.(mockBlob);
                return undefined;
            });

            const blob = await canvasToBlob(canvas);

            expect(blob).toBe(mockBlob);
            expect(canvas.toBlob).toHaveBeenCalledWith(expect.any(Function), 'image/jpeg', 0.9);
        });

        it('should use JPEG quality of 0.9', async () => {
            const mockBlob = new Blob(['fake'], { type: 'image/jpeg' });
            vi.spyOn(canvas, 'toBlob').mockImplementation((callback) => {
                callback?.(mockBlob);
                return undefined;
            });

            await canvasToBlob(canvas);

            expect(canvas.toBlob).toHaveBeenCalledWith(expect.any(Function), 'image/jpeg', 0.9);
        });

        it('should reject if toBlob returns null', async () => {
            vi.spyOn(canvas, 'toBlob').mockImplementation((callback) => {
                callback?.(null);
                return undefined;
            });

            await expect(canvasToBlob(canvas)).rejects.toThrow('Failed to convert canvas to blob');
        });
    });
});
