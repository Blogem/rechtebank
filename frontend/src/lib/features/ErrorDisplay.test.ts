import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import ErrorDisplay from './ErrorDisplay.svelte';

describe('ErrorDisplay', () => {
    it('should display error message', () => {
        render(ErrorDisplay, { props: { message: 'Upload failed' } });

        expect(screen.getByText(/Upload failed/i)).toBeInTheDocument();
    });

    it('should show legal-styled header', () => {
        render(ErrorDisplay, { props: { message: 'Test error' } });

        expect(screen.getByText(/Zaak Opgeschort/i)).toBeInTheDocument();
    });

    it('should show retry button when retryable is true', () => {
        render(ErrorDisplay, { props: { message: 'Network error', retryable: true } });

        expect(screen.getByText(/Probeer Opnieuw/i)).toBeInTheDocument();
    });

    it('should not show retry button when retryable is false', () => {
        render(ErrorDisplay, { props: { message: 'Fatal error', retryable: false } });

        expect(screen.queryByText(/Probeer Opnieuw/i)).not.toBeInTheDocument();
    });

    it('should always show reset button', () => {
        render(ErrorDisplay, { props: { message: 'Test error' } });

        expect(screen.getByText(/Terug naar Begin/i)).toBeInTheDocument();
    });

    it('should emit retry event when retry button clicked', async () => {
        const user = userEvent.setup();
        const retryHandler = vi.fn();
        render(ErrorDisplay, {
            props: {
                message: 'Test error',
                retryable: true,
                onretry: retryHandler
            }
        });

        const retryButton = screen.getByText(/Probeer Opnieuw/i);
        await user.click(retryButton);

        expect(retryHandler).toHaveBeenCalledTimes(1);
    });

    it('should emit reset event when reset button clicked', async () => {
        const user = userEvent.setup();
        const resetHandler = vi.fn();
        render(ErrorDisplay, {
            props: {
                message: 'Test error',
                onreset: resetHandler
            }
        });

        const resetButton = screen.getByText(/Terug naar Begin/i);
        await user.click(resetButton);

        expect(resetHandler).toHaveBeenCalledTimes(1);
    });

    it('should show gavel icon', () => {
        const { container } = render(ErrorDisplay, { props: { message: 'Test error' } });

        expect(container.querySelector('.gavel-icon')).toBeInTheDocument();
    });

    it('should show legal language text', () => {
        render(ErrorDisplay, { props: { message: 'Test error' } });

        expect(screen.getByText(/rechtbank kan op dit moment geen uitspraak doen/i)).toBeInTheDocument();
    });
});
