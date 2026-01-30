import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import PhotoConfirmation from './PhotoConfirmation.svelte';

describe('PhotoConfirmation', () => {
    const mockPhotoUrl = 'data:image/jpeg;base64,/9j/4AAQSkZJRg==';

    it('should render photo preview', () => {
        render(PhotoConfirmation, { props: { photoUrl: mockPhotoUrl } });

        const img = screen.getByAltText(/Captured furniture/i) as HTMLImageElement;
        expect(img).toBeInTheDocument();
        expect(img.src).toBe(mockPhotoUrl);
    });

    it('should show confirmation header', () => {
        render(PhotoConfirmation, { props: { photoUrl: mockPhotoUrl } });

        expect(screen.getByText(/Bevestig Foto/i)).toBeInTheDocument();
    });

    it('should show retake button', () => {
        render(PhotoConfirmation, { props: { photoUrl: mockPhotoUrl } });

        expect(screen.getByText(/Opnieuw/i)).toBeInTheDocument();
    });

    it('should show confirm button', () => {
        render(PhotoConfirmation, { props: { photoUrl: mockPhotoUrl } });

        expect(screen.getByText(/Indienen voor Vonnis/i)).toBeInTheDocument();
    });

    it('should emit retake event when retake button clicked', async () => {
        const user = userEvent.setup();
        const retakeHandler = vi.fn();
        render(PhotoConfirmation, {
            props: {
                photoUrl: mockPhotoUrl,
                onretake: retakeHandler
            }
        });

        const retakeButton = screen.getByText(/Opnieuw/i);
        await user.click(retakeButton);

        expect(retakeHandler).toHaveBeenCalledTimes(1);
    });

    it('should emit confirm event when confirm button clicked', async () => {
        const user = userEvent.setup();
        const confirmHandler = vi.fn();
        render(PhotoConfirmation, {
            props: {
                photoUrl: mockPhotoUrl,
                onconfirm: confirmHandler
            }
        });

        const confirmButton = screen.getByText(/Indienen voor Vonnis/i);
        await user.click(confirmButton);

        expect(confirmHandler).toHaveBeenCalledTimes(1);
    });
});
