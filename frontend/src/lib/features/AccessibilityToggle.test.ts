import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import AccessibilityToggle from './AccessibilityToggle.svelte';

describe('AccessibilityToggle', () => {
    it('should render toggle with enabled state by default', () => {
        render(AccessibilityToggle, { props: { levelCheckEnabled: true } });

        expect(screen.getByText(/Waterpas controle vereist/i)).toBeInTheDocument();
        const checkbox = screen.getByRole('checkbox') as HTMLInputElement;
        expect(checkbox.checked).toBe(true);
    });

    it('should render toggle with disabled state', () => {
        render(AccessibilityToggle, { props: { levelCheckEnabled: false } });

        const checkbox = screen.getByRole('checkbox') as HTMLInputElement;
        expect(checkbox.checked).toBe(false);
    });

    it('should show help text when level check is disabled', () => {
        render(AccessibilityToggle, { props: { levelCheckEnabled: false } });

        expect(screen.getByText(/Waterpas controle uitgeschakeld/i)).toBeInTheDocument();
        expect(screen.getByText(/toegankelijkheid/i)).toBeInTheDocument();
    });

    it('should not show help text when level check is enabled', () => {
        render(AccessibilityToggle, { props: { levelCheckEnabled: true } });

        expect(screen.queryByText(/Waterpas controle uitgeschakeld/i)).not.toBeInTheDocument();
    });

    it('should emit toggle event when checkbox changed', async () => {
        const user = userEvent.setup();
        const toggleHandler = vi.fn();
        render(AccessibilityToggle, {
            props: {
                levelCheckEnabled: true,
                ontoggle: toggleHandler
            }
        });

        const checkbox = screen.getByRole('checkbox');
        await user.click(checkbox);

        expect(toggleHandler).toHaveBeenCalledTimes(1);
        expect(toggleHandler).toHaveBeenCalledWith(
            expect.objectContaining({
                detail: { enabled: false }
            })
        );
    });

    it('should toggle from disabled to enabled', async () => {
        const user = userEvent.setup();
        const toggleHandler = vi.fn();
        render(AccessibilityToggle, {
            props: {
                levelCheckEnabled: false,
                ontoggle: toggleHandler
            }
        });

        const checkbox = screen.getByRole('checkbox');
        await user.click(checkbox);

        expect(toggleHandler).toHaveBeenCalledWith(
            expect.objectContaining({
                detail: { enabled: true }
            })
        );
    });
});
