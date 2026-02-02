<script lang="ts">
	import type { Verdict } from '$lib/shared/types';
	import { ApiAdapter } from '$lib/adapters/api/ApiAdapter';

	export let verdict: Verdict;
	export let imageData: string | undefined = undefined;
	export let onreset: ((event: CustomEvent) => void) | undefined = undefined;

	const api = new ApiAdapter();

	// Toast notification state
	let toastMessage = '';
	let toastVisible = false;
	let toastTimeout: number | undefined;

	function showToast(message: string, duration = 3000) {
		if (toastTimeout) clearTimeout(toastTimeout);
		toastMessage = message;
		toastVisible = true;
		toastTimeout = window.setTimeout(() => {
			toastVisible = false;
		}, duration);
	}

	function resetFlow() {
		onreset?.(new CustomEvent('reset'));
	}

	function generateCaseNumber(timestamp: string): string {
		const date = new Date(timestamp);
		const year = date.getFullYear();
		const timestampMs = date.getTime();
		return `RVM-${year}-${timestampMs}`;
	}

	function formatDutchTimestamp(timestamp: string): string {
		const date = new Date(timestamp);
		const day = date.getDate();
		const monthNames = [
			'januari',
			'februari',
			'maart',
			'april',
			'mei',
			'juni',
			'juli',
			'augustus',
			'september',
			'oktober',
			'november',
			'december'
		];
		const month = monthNames[date.getMonth()];
		const year = date.getFullYear();
		const hours = String(date.getHours()).padStart(2, '0');
		const minutes = String(date.getMinutes()).padStart(2, '0');
		return `Uitspraak d.d.: ${day} ${month} ${year}, ${hours}:${minutes}`;
	}

	async function shareVerdict() {
		try {
			// Call backend to create shareable URL
			const shareResponse = await api.createShareURL({
				timestamp: verdict.timestamp,
				requestId: verdict.requestId
			});

			// Construct full shareable URL using /vonnis/ as canonical route
			// (legacy /verdict/ URLs are still supported via nginx rewrite)
			const baseUrl = window.location.origin;
			const shareUrl = `${baseUrl}/vonnis/${shareResponse.id}`;

			// Get verdict text for sharing
			let verdictText = '';
			if (verdict.admissible) {
				verdictText = `${verdict.verdict.observation}\n\nUitspraak: ${verdict.verdict.sentence}`;
			} else {
				verdictText = `${verdict.verdict.crime}\n\n${verdict.verdict.observation}`;
			}

			// Share data with URL in separate field
			const shareData: ShareData = {
				title: 'Vonnis van de Rechtbank voor Meubilair',
				text: `${verdictText}\n\nScore: ${verdict.score}/10\n\nBekijk het volledige vonnis hier: ${shareUrl}`
			};

			// Try to include the photo if available
			if (imageData) {
				try {
					// Convert base64 data URL to Blob
					const response = await fetch(imageData);
					const blob = await response.blob();
					const file = new File([blob], 'meubelstuk.jpg', { type: blob.type });

					// Check if sharing files is supported
					const dataWithFile = { ...shareData, files: [file] };
					if (navigator.canShare && navigator.canShare(dataWithFile)) {
						shareData.files = [file];
					}
				} catch (err) {
					// If file conversion fails, continue without the image
					console.log('Could not include image in share:', err);
				}
			}

			// Try Web Share API first (mobile)
			if (navigator.share) {
				try {
					await navigator.share(shareData);
					// Success - exit without showing any toast
					return;
				} catch (err) {
					// User cancelled or dismissed the share dialog
					if (err instanceof Error && err.name === 'AbortError') {
						return; // Exit silently
					}
					// For any other error with Web Share API, show error and exit
					// Don't fall back to clipboard on mobile where share API exists
					console.error('Web Share API error:', err);
					showToast('⚠ Delen mislukt. Probeer het opnieuw.');
					return;
				}
			}

			// Fallback: Copy shareable URL to clipboard (only for desktop)
			try {
				await navigator.clipboard.writeText(shareUrl);
				showToast('✓ Link gekopieerd naar klembord!');
			} catch (err) {
				// Clipboard API failed, create a selectable text element
				const textArea = document.createElement('textarea');
				textArea.value = shareUrl;
				textArea.style.position = 'fixed';
				textArea.style.opacity = '0';
				document.body.appendChild(textArea);
				textArea.select();
				try {
					document.execCommand('copy');
					showToast('✓ Link gekopieerd naar klembord!');
				} catch (e) {
					showToast('Kopieer deze link: ' + shareUrl, 8000);
				}
				document.body.removeChild(textArea);
			}
		} catch (error) {
			console.error('Failed to share verdict:', error);
			showToast('⚠ Kon geen deelbare link maken. Probeer het later opnieuw.');
		}
	}

	function getVerdictClass(): string {
		if (!verdict.admissible) return 'dismissed';
		if (verdict.verdict.verdictType === 'vrijspraak') return 'acquittal';
		if (verdict.verdict.verdictType === 'waarschuwing') return 'warning';
		return 'guilty';
	}

	function getVerdictIcon(): string {
		if (!verdict.admissible) return '🚫';
		if (verdict.verdict.verdictType === 'vrijspraak') return '✅';
		if (verdict.verdict.verdictType === 'waarschuwing') return '⚠️';
		return '⚖️';
	}

	function getScoreClass(score: number): string {
		if (score >= 8) return 'excellent';
		if (score >= 5) return 'moderate';
		return 'poor';
	}
</script>

<div class="verdict-document {getVerdictClass()}">
	<div class="verdict-header">
		<div class="verdict-icon">{getVerdictIcon()}</div>
		<h1 class="verdict-title">Vonnis</h1>
		<p class="case-number">Zaaknummer: {generateCaseNumber(verdict.timestamp)}</p>
		<p class="case-date">{formatDutchTimestamp(verdict.timestamp)}</p>
	</div>

	{#if imageData}
		<section class="evidence-section">
			<h2 class="section-label">Bewijsmateriaal A</h2>
			<div class="photo-evidence">
				<img src={imageData} alt="Ingediend meubelstuk" class="evidence-photo" />
			</div>
		</section>
	{/if}

	{#if !verdict.admissible}
		<section class="verdict-section">
			<h2 class="section-heading">Zaak Niet-Ontvankelijk</h2>
			<div class="section-content">
				<p class="legal-text">{verdict.verdict.crime}</p>
				<p class="legal-text">{verdict.verdict.observation}</p>
				<p class="legal-notice">
					<em>Dit object is geen meubelstuk en valt buiten de jurisdictie van deze rechtbank.</em>
				</p>
			</div>
		</section>
	{:else}
		<section class="verdict-section">
			<h2 class="section-heading">Feiten</h2>
			<div class="section-content">
				<p class="legal-text">{verdict.verdict.observation}</p>
				<div class="score-badge {getScoreClass(verdict.score)}">
					<span class="score-number">{verdict.score}</span>
					<span class="score-label">/10</span>
				</div>
			</div>
		</section>

		<section class="verdict-section">
			<h2 class="section-heading">Overwegingen</h2>
			<div class="section-content">
				<div class="consideration-item">
					<h3 class="consideration-label">Overtreding:</h3>
					<p class="legal-text">{verdict.verdict.crime}</p>
				</div>
				<div class="consideration-item">
					<h3 class="consideration-label">Juridische Grondslag:</h3>
					<p class="legal-text">{verdict.verdict.reasoning}</p>
				</div>
			</div>
		</section>

		<section class="verdict-section uitspraak">
			<h2 class="section-heading">
				{#if verdict.verdict.verdictType === 'vrijspraak'}
					Vrijspraak
				{:else if verdict.verdict.verdictType === 'waarschuwing'}
					Waarschuwing
				{:else}
					Schuldig Bevonden
				{/if}
			</h2>
			<div class="section-content">
				<p class="legal-text">{verdict.verdict.sentence}</p>
			</div>
		</section>
	{/if}

	<div class="verdict-actions">
		<button onclick={shareVerdict} class="action-button secondary">Deel Vonnis</button>
		<button onclick={resetFlow} class="action-button primary">Nieuwe Zaak</button>
	</div>

	<footer class="verdict-footer">
		<p class="proclamation">Uitgesproken in het openbaar</p>
		<p class="verdict-seal">⚖️</p>
	</footer>
</div>

{#if toastVisible}
	<div class="toast" class:visible={toastVisible}>
		{toastMessage}
	</div>
{/if}

<style>
	.verdict-document {
		max-width: 800px;
		margin: 0 auto;
		background: var(--color-court-surface);
		border-radius: 4px;
		box-shadow: var(--shadow-base);
		overflow: hidden;
		opacity: 0;
		transform: translateY(8px);
		animation: revealVerdict var(--timing-reveal) var(--ease-out) forwards;
	}

	@keyframes revealVerdict {
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.verdict-header {
		background: var(--color-court-primary);
		color: white;
		text-align: center;
		padding: 2rem 1.5rem;
		border-bottom: 3px solid var(--color-court-accent);
	}

	.verdict-icon {
		font-size: 2.5rem;
		margin-bottom: 0.75rem;
		opacity: 0;
		animation: fadeIn var(--timing-reveal) var(--ease-out) 0.1s forwards;
	}

	.verdict-title {
		font-size: 2rem;
		margin: 0 0 1rem 0;
		color: white;
		letter-spacing: 0.02em;
	}

	.case-number {
		font-size: 0.9rem;
		opacity: 0.9;
		margin: 0.25rem 0;
		font-family: var(--font-sans);
	}

	.case-date {
		font-size: 0.85rem;
		opacity: 0.85;
		margin: 0.25rem 0;
		font-family: var(--font-sans);
		font-style: italic;
	}

	.evidence-section {
		padding: 2rem;
		background: var(--color-court-surface);
		border-bottom: 1px solid var(--color-court-border);
		opacity: 0;
		animation: fadeInSection var(--timing-reveal) var(--ease-out) 0.2s forwards;
	}

	.section-label {
		font-size: 0.85rem;
		font-weight: 600;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		color: var(--color-court-accent);
		margin: 0 0 1rem 0;
		font-family: var(--font-sans);
		text-align: center;
	}

	.photo-evidence {
		display: flex;
		justify-content: center;
		background: #f8f9fa;
		padding: 1rem;
		border: 2px solid var(--color-court-border);
		border-radius: 2px;
	}

	.evidence-photo {
		max-width: 100%;
		max-height: 400px;
		width: auto;
		height: auto;
		display: block;
		box-shadow: var(--shadow-sm);
	}

	.verdict-section {
		padding: 2rem;
		border-bottom: 1px solid var(--color-court-border);
		opacity: 0;
		animation: fadeInSection var(--timing-reveal) var(--ease-out) 0.3s forwards;
	}

	.verdict-section:nth-of-type(3) {
		animation-delay: 0.4s;
	}

	.verdict-section:nth-of-type(4) {
		animation-delay: 0.5s;
	}

	.verdict-section.uitspraak {
		background: #fffef8;
		border-left: 4px solid var(--color-court-accent);
		animation-delay: 0.6s;
	}

	@keyframes fadeInSection {
		to {
			opacity: 1;
		}
	}

	@keyframes fadeIn {
		to {
			opacity: 1;
		}
	}

	.section-heading {
		font-size: 1.5rem;
		color: var(--color-court-primary);
		margin: 0 0 1.25rem 0;
		padding-bottom: 0.75rem;
		border-bottom: 2px solid var(--color-court-border);
	}

	.section-content {
		line-height: 1.75;
		color: var(--color-court-text);
	}

	.legal-text {
		font-size: 1.05rem;
		margin: 1rem 0;
		text-align: justify;
		font-family: var(--font-serif);
	}

	.legal-notice {
		font-style: italic;
		color: var(--color-court-text-light);
		margin-top: 1.5rem;
		text-align: center;
		font-size: 0.95rem;
	}

	.score-badge {
		display: inline-flex;
		align-items: baseline;
		gap: 0.25rem;
		background: var(--color-court-primary);
		color: white;
		padding: 0.75rem 1.5rem;
		border-radius: 4px;
		margin-top: 1.5rem;
		box-shadow: var(--shadow-sm);
	}

	.score-badge.excellent {
		background: #28a745;
	}

	.score-badge.moderate {
		background: #ffc107;
		color: var(--color-court-text);
	}

	.score-badge.poor {
		background: #dc3545;
	}

	.score-number {
		font-size: 2rem;
		font-weight: 700;
		font-family: var(--font-serif);
	}

	.score-label {
		font-size: 1rem;
		opacity: 0.9;
	}

	.consideration-item {
		margin: 1.5rem 0;
	}

	.consideration-item:first-child {
		margin-top: 0;
	}

	.consideration-label {
		font-size: 1rem;
		font-weight: 600;
		color: var(--color-court-accent);
		margin: 0 0 0.5rem 0;
		font-family: var(--font-sans);
		text-transform: uppercase;
		letter-spacing: 0.02em;
		font-size: 0.85rem;
	}

	.verdict-footer {
		text-align: center;
		padding: 2rem 1.5rem;
		background: var(--color-court-surface);
		color: var(--color-court-text-light);
	}

	.proclamation {
		font-style: italic;
		margin: 0 0 1rem 0;
		font-family: var(--font-serif);
	}

	.verdict-seal {
		font-size: 2rem;
		margin: 0;
		opacity: 0;
		animation: sealAppear 600ms var(--ease-out) 0.8s forwards;
	}

	@keyframes sealAppear {
		from {
			opacity: 0;
			transform: scale(0.9);
		}
		to {
			opacity: 1;
			transform: scale(1);
		}
	}

	.verdict-actions {
		display: flex;
		gap: 1rem;
		padding: 1.5rem 2rem 2rem;
		justify-content: center;
		flex-wrap: wrap;
	}

	.action-button {
		padding: 0.875rem 2rem;
		font-size: 1rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		font-family: var(--font-sans);
		font-weight: 600;
		box-shadow: var(--shadow-base);
		transition: all var(--timing-interactive) var(--ease-out);
		position: relative;
	}

	.action-button:hover {
		transform: translateY(-3px);
		box-shadow: var(--shadow-md);
	}

	.action-button:active {
		transform: translateY(-1px) scale(0.98);
		box-shadow: var(--shadow-sm);
		transition: all var(--timing-quick) var(--ease-out);
	}

	.action-button:focus-visible {
		outline: none;
		box-shadow:
			var(--shadow-md),
			0 0 0 3px var(--color-court-accent);
	}

	.action-button.primary {
		background: var(--color-court-primary);
		color: white;
	}

	.action-button.primary:hover {
		background: #244a66;
	}

	.action-button.secondary {
		background: var(--color-court-text-light);
		color: white;
	}

	.action-button.secondary:hover {
		background: #555555;
	}

	.verdict-document.guilty {
		border-top: 4px solid #dc3545;
	}

	.verdict-document.acquittal {
		border-top: 4px solid #28a745;
	}

	.verdict-document.warning {
		border-top: 4px solid #ffc107;
	}

	.verdict-document.dismissed {
		border-top: 4px solid #6c757d;
	}

	.toast {
		position: fixed;
		bottom: 2rem;
		left: 50%;
		transform: translateX(-50%) translateY(100px);
		background: var(--color-court-primary);
		color: white;
		padding: 1rem 1.5rem;
		border-radius: 8px;
		box-shadow: var(--shadow-lg);
		font-family: var(--font-sans);
		font-size: 0.95rem;
		z-index: 1000;
		max-width: 90%;
		word-wrap: break-word;
		opacity: 0;
		transition: all var(--timing-interactive) var(--ease-out);
	}

	.toast.visible {
		opacity: 1;
		transform: translateX(-50%) translateY(0);
	}

	@media (prefers-reduced-motion: reduce) {
		.verdict-document,
		.verdict-section,
		.evidence-section,
		.verdict-icon,
		.verdict-seal {
			animation: none;
			opacity: 1;
			transform: none;
		}
	}

	@media (max-width: 768px) {
		.verdict-header {
			padding: 1.5rem 1rem;
		}

		.verdict-title {
			font-size: 1.5rem;
		}

		.verdict-section,
		.evidence-section {
			padding: 1.5rem 1rem;
		}

		.section-heading {
			font-size: 1.3rem;
		}

		.legal-text {
			font-size: 1rem;
		}

		.evidence-photo {
			max-height: 300px;
		}

		.verdict-actions {
			flex-direction: column;
			padding: 1rem;
		}

		.action-button {
			width: 100%;
		}
	}

	@media (max-width: 480px) {
		.verdict-header {
			padding: 1rem 0.75rem;
		}

		.verdict-title {
			font-size: 1.25rem;
		}

		.verdict-section,
		.evidence-section {
			padding: 1rem 0.75rem;
		}

		.section-heading {
			font-size: 1.2rem;
		}

		.legal-text {
			font-size: 0.95rem;
			text-align: left;
		}
	}
</style>
