import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import SpiritLevel from './SpiritLevel.svelte';
import type { OrientationData } from '$lib/shared/types';

describe('SpiritLevel', () => {
    const levelOrientation: OrientationData = {
        beta: 0,
        gamma: 0,
        alpha: 0,
        isLevel: true
    };

    const tiltedOrientation: OrientationData = {
        beta: 10,
        gamma: 0,
        alpha: 0,
        isLevel: false
    };

    it('should render spirit level when enabled', () => {
        const { container } = render(SpiritLevel, {
            props: { orientationData: levelOrientation, enabled: true }
        });

        expect(container.querySelector('.spirit-level')).toBeInTheDocument();
    });

    it('should not render when disabled', () => {
        const { container } = render(SpiritLevel, {
            props: { orientationData: levelOrientation, enabled: false }
        });

        expect(container.querySelector('.spirit-level')).not.toBeInTheDocument();
    });

    it('should show green state when device is level', () => {
        const { container } = render(SpiritLevel, {
            props: { orientationData: levelOrientation, enabled: true }
        });

        const levelElement = container.querySelector('.spirit-level');
        expect(levelElement).toHaveClass('level');
        expect(screen.getByText(/Waterpas/i)).toBeInTheDocument();
    });

    it('should show red state when device is tilted', () => {
        const { container } = render(SpiritLevel, {
            props: { orientationData: tiltedOrientation, enabled: true }
        });

        const levelElement = container.querySelector('.spirit-level');
        expect(levelElement).toHaveClass('off-level');
    });

    it('should display tilt angle and direction', () => {
        render(SpiritLevel, {
            props: { orientationData: tiltedOrientation, enabled: true }
        });

        expect(screen.getByText(/10°/)).toBeInTheDocument();
        expect(screen.getByText(/naar voren/i)).toBeInTheDocument();
    });

    it('should show level indicator with bubble', () => {
        const { container } = render(SpiritLevel, {
            props: { orientationData: levelOrientation, enabled: true }
        });

        expect(container.querySelector('.level-tube')).toBeInTheDocument();
        expect(container.querySelector('.bubble')).toBeInTheDocument();
    });

    it('should handle null orientation data gracefully', () => {
        const { container } = render(SpiritLevel, {
            props: { orientationData: null, enabled: true }
        });

        expect(container.querySelector('.spirit-level')).toBeInTheDocument();
    });
});
