import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import VerdictDisplay from './VerdictDisplay.svelte';
import type { Verdict } from '$lib/shared/types';

describe('VerdictDisplay', () => {
    const mockGuiltyVerdict: Verdict = {
        admissible: true,
        score: 5,
        verdict: {
            crime: 'Ernstige scheefstand',
            sentence: 'Veroordeling tot rechtzetting',
            reasoning: 'Artikel 42 van de Meubilair-wet',
            observation: 'Dit meubelstuk is schuldig bevonden aan scheefheid.',
            verdictType: 'schuldig'
        },
        requestId: 'test-request-123',
        timestamp: new Date().toISOString()
    };

    const mockAcquittalVerdict: Verdict = {
        admissible: true,
        score: 10,
        verdict: {
            crime: 'Geen afwijkingen geconstateerd',
            sentence: 'Volledige vrijspraak',
            reasoning: 'Het meubelstuk voldoet aan alle eisen',
            observation: 'Dit meubelstuk is volmaakt recht.',
            verdictType: 'vrijspraak'
        },
        requestId: 'test-request-124',
        timestamp: new Date().toISOString()
    };

    const mockWarningVerdict: Verdict = {
        admissible: true,
        score: 6,
        verdict: {
            crime: 'Lichte afwijking',
            sentence: 'Waarschuwing',
            reasoning: 'Artikel 15 van de Meubilair-wet',
            observation: 'Dit meubelstuk heeft een kleine afwijking.',
            verdictType: 'waarschuwing'
        },
        requestId: 'test-request-125',
        timestamp: new Date().toISOString()
    };

    const mockDismissedVerdict: Verdict = {
        admissible: false,
        score: 0,
        verdict: {
            crime: 'Geen meubilair gedetecteerd',
            sentence: 'Zaak niet-ontvankelijk verklaard',
            reasoning: 'Dit Hof oordeelt alleen over meubilair',
            observation: 'Dit is geen meubelstuk.',
            verdictType: 'schuldig'
        },
        requestId: 'test-request-126',
        timestamp: new Date().toISOString()
    };

    it('should render guilty verdict with score', () => {
        render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        expect(screen.getByText('5')).toBeInTheDocument();
        expect(screen.getByText('/10')).toBeInTheDocument();
        expect(screen.getByRole('heading', { name: /schuldig bevonden/i })).toBeInTheDocument();
    });

    it('should render acquittal verdict', () => {
        render(VerdictDisplay, { props: { verdict: mockAcquittalVerdict } });

        expect(screen.getByText('10')).toBeInTheDocument();
        expect(screen.getByRole('heading', { name: /Vrijspraak/i })).toBeInTheDocument();
    });

    it('should render warning verdict', () => {
        render(VerdictDisplay, { props: { verdict: mockWarningVerdict } });

        expect(screen.getByText('6')).toBeInTheDocument();
        expect(screen.getByRole('heading', { name: /Waarschuwing/i })).toBeInTheDocument();
    });

    it('should render dismissed verdict for non-furniture', () => {
        render(VerdictDisplay, { props: { verdict: mockDismissedVerdict } });

        expect(screen.getByText(/Niet-Ontvankelijk/i)).toBeInTheDocument();
        expect(screen.getByText(/Dit is geen meubelstuk/i)).toBeInTheDocument();
    });

    it('should apply correct CSS class for guilty verdict based on verdictType', () => {
        const { container } = render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        const verdictDisplay = container.querySelector('.verdict-document');
        expect(verdictDisplay).toHaveClass('guilty');
    });

    it('should apply correct CSS class for acquittal verdict based on verdictType', () => {
        const { container } = render(VerdictDisplay, { props: { verdict: mockAcquittalVerdict } });

        const verdictDisplay = container.querySelector('.verdict-document');
        expect(verdictDisplay).toHaveClass('acquittal');
    });

    it('should apply correct CSS class for warning verdict based on verdictType', () => {
        const { container } = render(VerdictDisplay, { props: { verdict: mockWarningVerdict } });

        const verdictDisplay = container.querySelector('.verdict-document');
        expect(verdictDisplay).toHaveClass('warning');
    });

    it('should display correct icon for vrijspraak verdict', () => {
        render(VerdictDisplay, { props: { verdict: mockAcquittalVerdict } });

        expect(screen.getByText('✅')).toBeInTheDocument();
    });

    it('should display correct icon for waarschuwing verdict', () => {
        render(VerdictDisplay, { props: { verdict: mockWarningVerdict } });

        expect(screen.getByText('⚠️')).toBeInTheDocument();
    });

    it('should display correct icon for schuldig verdict', () => {
        const { container } = render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        const verdictIcon = container.querySelector('.verdict-icon');
        expect(verdictIcon).toBeInTheDocument();
        expect(verdictIcon?.textContent).toBe('⚖️');
    });

    it('should apply correct score class for excellent score', () => {
        const { container } = render(VerdictDisplay, { props: { verdict: mockAcquittalVerdict } });

        const scoreDisplay = container.querySelector('.score-badge');
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

        const resetButton = screen.getByText(/Nieuwe Zaak/i);
        await user.click(resetButton);

        expect(resetHandler).toHaveBeenCalledTimes(1);
    });

    it('should have share button', () => {
        render(VerdictDisplay, { props: { verdict: mockGuiltyVerdict } });

        expect(screen.getByText(/Deel Vonnis/i)).toBeInTheDocument();
    });
});
