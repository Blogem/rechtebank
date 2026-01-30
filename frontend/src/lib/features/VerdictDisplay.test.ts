import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import VerdictDisplay from './VerdictDisplay.svelte';
import type { Verdict } from '$lib/shared/types';

describe('VerdictDisplay', () => {
    const mockGuiltyVerdict: Verdict = {
        type: 'guilty',
        score: 5,
        verdictText: 'Dit meubelstuk is schuldig bevonden aan scheefheid.',
        sentence: 'Veroordeling tot rechtzetting',
        angleDeviation: 15,
        isFurniture: true
    };

    const mockAcquittalVerdict: Verdict = {
        type: 'acquittal',
        score: 10,
        verdictText: 'Dit meubelstuk is volmaakt recht.',
        isFurniture: true
    };

    const mockDismissedVerdict: Verdict = {
        type: 'niet-ontvankelijk',
        score: 0,
        verdictText: 'Dit is geen meubelstuk.',
        isFurniture: false
    };

    it('should render guilty verdict with score', () => {
        render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        expect(screen.getByText('5')).toBeInTheDocument();
        expect(screen.getByText('van 10')).toBeInTheDocument();
        expect(screen.getByRole('heading', { name: /schuldig bevonden/i })).toBeInTheDocument();
    });

    it('should display angle deviation when present', () => {
        render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        expect(screen.getByText(/15°/)).toBeInTheDocument();
    });

    it('should display sentence when present', () => {
        render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        expect(screen.getByText(/Veroordeling tot rechtzetting/)).toBeInTheDocument();
    });

    it('should render acquittal verdict', () => {
        render(VerdictDisplay, { props: { verdict: mockAcquittalVerdict } });

        expect(screen.getByText('10')).toBeInTheDocument();
        expect(screen.getByText(/Vrijspraak/i)).toBeInTheDocument();
    });

    it('should render dismissed verdict for non-furniture', () => {
        render(VerdictDisplay, { props: { verdict: mockDismissedVerdict } });

        expect(screen.getByText(/Niet-Ontvankelijk/i)).toBeInTheDocument();
        expect(screen.getByText(/Dit is geen meubelstuk/i)).toBeInTheDocument();
    });

    it('should apply correct CSS class for guilty verdict', () => {
        const { container } = render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        const verdictDisplay = container.querySelector('.verdict-display');
        expect(verdictDisplay).toHaveClass('guilty');
    });

    it('should apply correct CSS class for acquittal verdict', () => {
        const { container } = render(VerdictDisplay, { props: { verdict: mockAcquittalVerdict } });

        const verdictDisplay = container.querySelector('.verdict-display');
        expect(verdictDisplay).toHaveClass('acquittal');
    });

    it('should apply correct score class for excellent score', () => {
        const { container } = render(VerdictDisplay, { props: { verdict: mockAcquittalVerdict } });

        const scoreDisplay = container.querySelector('.score-display');
        expect(scoreDisplay).toHaveClass('excellent');
    });

    it('should emit reset event when reset button clicked', async () => {
        const user = userEvent.setup();
        const resetHandler = vi.fn();
        render(VerdictDisplay, {
            props: {
                verdict: mockGuiltyVerdict,
                onreset: resetHandler
            }
        });

        const resetButton = screen.getByText(/Dien Ander Meubelstuk In/i);
        await user.click(resetButton);

        expect(resetHandler).toHaveBeenCalledTimes(1);
    });

    it('should have share button', () => {
        render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        expect(screen.getByText(/Deel Vonnis/i)).toBeInTheDocument();
    });

    it('should display current date in court seal', () => {
        render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        const today = new Date().toLocaleDateString('nl-NL');
        expect(screen.getByText(today)).toBeInTheDocument();
    });
});
