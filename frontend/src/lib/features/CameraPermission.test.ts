import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import CameraPermission from './CameraPermission.svelte';

describe('CameraPermission', () => {
    it('should render permission request UI', () => {
        render(CameraPermission, { props: { httpsRequired: false } });

        expect(screen.getByRole('heading', { name: /Camera Toestemming/i })).toBeInTheDocument();
        expect(screen.getByText(/Sta Camera Toe/i)).toBeInTheDocument();
    });

    it('should show HTTPS warning when required', () => {
        render(CameraPermission, { props: { httpsRequired: true } });

        expect(screen.getByText(/HTTPS Vereist/i)).toBeInTheDocument();
        expect(screen.getByText(/Camera toegang werkt alleen via HTTPS/i)).toBeInTheDocument();
    });

    it('should not show permission button when HTTPS required', () => {
        render(CameraPermission, { props: { httpsRequired: true } });

        expect(screen.queryByText(/Sta Camera Toe/i)).not.toBeInTheDocument();
    });

    it('should emit permission-requested event when button clicked', async () => {
        const user = userEvent.setup();
        const permissionHandler = vi.fn();
        render(CameraPermission, {
            props: {
                httpsRequired: false,
                onpermissionrequested: permissionHandler
            }
        });

        const button = screen.getByText(/Sta Camera Toe/i);
        await user.click(button);

        expect(permissionHandler).toHaveBeenCalledTimes(1);
    });

    it('should show help text about file upload fallback', () => {
        render(CameraPermission, { props: { httpsRequired: false } });

        expect(
            screen.getByText(/U kunt ook een foto uploaden als u geen camera toestemming wilt verlenen/i)
        ).toBeInTheDocument();
    });
});
