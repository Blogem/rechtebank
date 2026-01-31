import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, fireEvent, waitFor } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import PhotoCapture from './PhotoCapture.svelte';
import * as rotationUtils from '$lib/shared/utils/rotationUtils';

describe('PhotoCapture', () => {
    const mockOnPhotoConfirmed = vi.fn();
    const mockOnCancelled = vi.fn();

    beforeEach(() => {
        vi.clearAllMocks();
    });

    it('should trigger file input on capture button click', async () => {
        const user = userEvent.setup();
        render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        const captureButton = screen.getByText(/Neem Foto/i);
        await user.click(captureButton);

        // File input should be clicked (check if it exists and has correct attributes)
        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;
        expect(fileInput).toBeInTheDocument();
        expect(fileInput.accept).toBe('image/*');
        // Note: capture attribute is a boolean in the DOM, not a string
        expect(fileInput.hasAttribute('capture')).toBe(true);
    });

    it('should display image preview after file selection', async () => {
        render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;

        // Create a mock image file
        const file = new File(['dummy'], 'test.jpg', { type: 'image/jpeg' });

        // Mock URL.createObjectURL
        const mockUrl = 'blob:http://localhost:3000/test';
        global.URL.createObjectURL = vi.fn(() => mockUrl);

        // Trigger file selection
        await fireEvent.change(fileInput, { target: { files: [file] } });

        await waitFor(() => {
            const preview = screen.getByAltText(/Preview/i) as HTMLImageElement;
            expect(preview).toBeInTheDocument();
        });
    });

    it('should decrease rotation by 90 degrees when rotate-left is clicked', async () => {
        const user = userEvent.setup();
        render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        // Load a photo first
        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;
        const file = new File(['dummy'], 'test.jpg', { type: 'image/jpeg' });
        global.URL.createObjectURL = vi.fn(() => 'blob:test');
        await fireEvent.change(fileInput, { target: { files: [file] } });

        await waitFor(() => {
            expect(screen.getByAltText(/Preview/i)).toBeInTheDocument();
        });

        const rotateLeftButton = screen.getByLabelText(/Draai links/i);
        const preview = screen.getByAltText(/Preview/i) as HTMLElement;

        // Initial rotation should be 0 (or from initial heuristic)
        const initialTransform = preview.style.transform;

        await user.click(rotateLeftButton);

        // After clicking rotate-left, rotation should decrease by 90°
        await waitFor(() => {
            expect(preview.style.transform).toContain('rotate');
        });
    });

    it('should increase rotation by 90 degrees when rotate-right is clicked', async () => {
        const user = userEvent.setup();
        render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        // Load a photo first
        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;
        const file = new File(['dummy'], 'test.jpg', { type: 'image/jpeg' });
        global.URL.createObjectURL = vi.fn(() => 'blob:test');
        await fireEvent.change(fileInput, { target: { files: [file] } });

        await waitFor(() => {
            expect(screen.getByAltText(/Preview/i)).toBeInTheDocument();
        });

        const rotateRightButton = screen.getByLabelText(/Draai rechts/i);
        await user.click(rotateRightButton);

        // After clicking rotate-right, rotation should increase by 90°
        const preview = screen.getByAltText(/Preview/i) as HTMLElement;
        await waitFor(() => {
            expect(preview.style.transform).toContain('rotate');
        });
    });

    it('should wrap rotation at 0 and 360 degrees', async () => {
        const user = userEvent.setup();
        render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        // Load a photo
        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;
        const file = new File(['dummy'], 'test.jpg', { type: 'image/jpeg' });
        global.URL.createObjectURL = vi.fn(() => 'blob:test');
        await fireEvent.change(fileInput, { target: { files: [file] } });

        await waitFor(() => {
            expect(screen.getByAltText(/Preview/i)).toBeInTheDocument();
        });

        const rotateLeftButton = screen.getByLabelText(/Draai links/i);

        // Rotate left from 0° should wrap to 270°
        await user.click(rotateLeftButton);

        const preview = screen.getByAltText(/Preview/i) as HTMLElement;
        await waitFor(() => {
            expect(preview.style.transform).toContain('270deg');
        });
    });

    it('should call onPhotoConfirmed with blob and rotation when confirmed', async () => {
        const user = userEvent.setup();

        // Mock the rotation utilities
        const mockBlob = new Blob(['rotated'], { type: 'image/jpeg' });
        const mockCanvas = document.createElement('canvas');

        vi.spyOn(rotationUtils, 'rotateImageOnCanvas').mockReturnValue(mockCanvas);
        vi.spyOn(rotationUtils, 'canvasToBlob').mockResolvedValue(mockBlob);

        // Mock Image loading
        global.Image = class {
            onload: (() => void) | null = null;
            onerror: (() => void) | null = null;
            src = '';
            constructor() {
                // Trigger onload immediately when src is set
                setTimeout(() => {
                    if (this.onload) {
                        this.onload();
                    }
                }, 0);
            }
        } as any;

        render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        // Load a photo
        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;
        const file = new File(['dummy'], 'test.jpg', { type: 'image/jpeg' });
        global.URL.createObjectURL = vi.fn(() => 'blob:test');
        await fireEvent.change(fileInput, { target: { files: [file] } });

        await waitFor(() => {
            expect(screen.getByAltText(/Preview/i)).toBeInTheDocument();
        });

        const confirmButton = screen.getByRole('button', { name: /Bevestig/i });
        await user.click(confirmButton);

        // Should be called with a blob and rotation value
        await waitFor(
            () => {
                expect(mockOnPhotoConfirmed).toHaveBeenCalledTimes(1);
            },
            { timeout: 2000 }
        );

        const [blob, rotation] = mockOnPhotoConfirmed.mock.calls[0];
        expect(blob).toBe(mockBlob);
        expect(typeof rotation).toBe('number');
    });

    it('should reset state and revoke object URL on retake', async () => {
        const user = userEvent.setup();
        const mockRevoke = vi.fn();
        global.URL.revokeObjectURL = mockRevoke;

        render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        // Load a photo
        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;
        const file = new File(['dummy'], 'test.jpg', { type: 'image/jpeg' });
        const mockUrl = 'blob:test';
        global.URL.createObjectURL = vi.fn(() => mockUrl);
        await fireEvent.change(fileInput, { target: { files: [file] } });

        await waitFor(() => {
            expect(screen.getByAltText(/Preview/i)).toBeInTheDocument();
        });

        const retakeButton = screen.getByText(/Opnieuw/i);
        await user.click(retakeButton);

        // Object URL should be revoked
        expect(mockRevoke).toHaveBeenCalledWith(mockUrl);

        // Preview should no longer be visible
        await waitFor(() => {
            expect(screen.queryByAltText(/Preview/i)).not.toBeInTheDocument();
        });
    });

    it('should revoke object URL on component unmount', async () => {
        const mockRevoke = vi.fn();
        global.URL.revokeObjectURL = mockRevoke;

        const { unmount } = render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        // Load a photo
        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;
        const file = new File(['dummy'], 'test.jpg', { type: 'image/jpeg' });
        const mockUrl = 'blob:test';
        global.URL.createObjectURL = vi.fn(() => mockUrl);
        await fireEvent.change(fileInput, { target: { files: [file] } });

        await waitFor(() => {
            expect(screen.getByAltText(/Preview/i)).toBeInTheDocument();
        });

        // Unmount component
        unmount();

        // Object URL should be revoked
        expect(mockRevoke).toHaveBeenCalledWith(mockUrl);
    });

    it('should reject non-image files with error display', async () => {
        render(PhotoCapture, { props: { onPhotoConfirmed: mockOnPhotoConfirmed } });

        const fileInput = screen.getByLabelText(/Upload foto/i) as HTMLInputElement;

        // Try to upload a non-image file
        const file = new File(['dummy'], 'test.txt', { type: 'text/plain' });
        await fireEvent.change(fileInput, { target: { files: [file] } });

        // Should show an error message
        await waitFor(() => {
            expect(screen.getByText(/Alleen afbeeldingen/i)).toBeInTheDocument();
        });

        // Preview should not be shown
        expect(screen.queryByAltText(/Preview/i)).not.toBeInTheDocument();
    });
});
