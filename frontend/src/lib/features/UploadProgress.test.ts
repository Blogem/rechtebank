import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import UploadProgress from './UploadProgress.svelte';

describe('UploadProgress', () => {
	it('should render default message', () => {
		render(UploadProgress, { props: { progress: 50 } });

		expect(screen.getByText(/De rechter beraadslaagt/i)).toBeInTheDocument();
	});

	it('should render custom message', () => {
		render(UploadProgress, { props: { progress: 75, message: 'Uploading foto...' } });

		expect(screen.getByText(/Uploading foto/i)).toBeInTheDocument();
	});

	it('should display progress bar', () => {
		const { container } = render(UploadProgress, { props: { progress: 60 } });

		const progressFill = container.querySelector('.progress-fill') as HTMLElement;
		expect(progressFill).toBeInTheDocument();
		expect(progressFill.style.width).toBe('60%');
	});

	it('should show court icon', () => {
		const { container } = render(UploadProgress, { props: { progress: 50 } });

		expect(container.querySelector('.court-icon')).toBeInTheDocument();
	});

	it('should show legal text', () => {
		render(UploadProgress, { props: { progress: 50 } });

		expect(screen.getByText(/Uw meubel wordt beoordeeld/i)).toBeInTheDocument();
	});

	it('should handle 0% progress', () => {
		const { container } = render(UploadProgress, { props: { progress: 0 } });

		const progressFill = container.querySelector('.progress-fill') as HTMLElement;
		expect(progressFill.style.width).toBe('0%');
	});

	it('should handle 100% progress', () => {
		const { container } = render(UploadProgress, { props: { progress: 100 } });

		const progressFill = container.querySelector('.progress-fill') as HTMLElement;
		expect(progressFill.style.width).toBe('100%');
	});
});
