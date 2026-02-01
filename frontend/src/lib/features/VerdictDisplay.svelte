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

			// Construct full shareable URL
			const baseUrl = window.location.origin;
			const shareUrl = `${baseUrl}/verdict/${shareResponse.id}`;

			// Get verdict text for sharing
			const verdictText = verdict.admissible ? verdict.verdict.observation : verdict.verdict.crime;

			// Include URL in text for apps that don't support separate url field (like WhatsApp)
			const shareData = {
				title: 'Vonnis van de Rechtbank voor Meubilair',
				text: `${verdictText}\n\nScore: ${verdict.score}/10\n\n${shareUrl}`,
				url: shareUrl
			};

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

<div class="verdict-display {getVerdictClass()}">
	<div class="court-header">
		<div class="gavel-icon">{getVerdictIcon()}</div>
		<h1>Vonnis van de Rechtbank voor Meubilair</h1>
	</div>

	{#if imageData}
		<div class="photo-display">
			<img src={imageData} alt="Ingediend meubelstuk" class="verdict-photo" />
		</div>
	{/if}

	<div class="case-metadata">
		<p class="case-number">Zaaknummer: {generateCaseNumber(verdict.timestamp)}</p>
		<p class="case-date">{formatDutchTimestamp(verdict.timestamp)}</p>
	</div>
	<hr class="metadata-separator" />

	<div class="verdict-content">
		{#if !verdict.admissible}
			<div class="case-dismissed">
				<h2>Zaak Niet-Ontvankelijk</h2>
				<p class="verdict-text">{verdict.verdict.crime}</p>
				<p class="verdict-text">{verdict.verdict.observation}</p>
				<p class="legal-note">
					<em>Dit object is geen meubelstuk en valt buiten de jurisdictie van deze rechtbank.</em>
				</p>
			</div>
		{:else}
			<div class="score-display {getScoreClass(verdict.score)}">
				<div class="score-number">{verdict.score}</div>
				<div class="score-label">van 10</div>
			</div>

			<div class="verdict-type">
				<h2>
					{#if verdict.verdict.verdictType === 'vrijspraak'}
						Vrijspraak
					{:else if verdict.verdict.verdictType === 'waarschuwing'}
						Waarschuwing
					{:else}
						Schuldig Bevonden
					{/if}
				</h2>
			</div>

			<div class="verdict-body">
				<p class="verdict-text">{verdict.verdict.observation}</p>

				<div class="crime-section">
					<h3>Overtreding:</h3>
					<p>{verdict.verdict.crime}</p>
				</div>

				<div class="reasoning-section">
					<h3>Juridische Overwegingen:</h3>
					<p>{verdict.verdict.reasoning}</p>
				</div>

				<div class="sentence">
					<h3>Uitspraak:</h3>
					<p>{verdict.verdict.sentence}</p>
				</div>
			</div>
		{/if}
	</div>

	<div class="verdict-actions">
		<button onclick={shareVerdict} class="action-button secondary">Deel Vonnis</button>

		<button onclick={resetFlow} class="action-button primary">Dien Ander Meubelstuk In</button>
	</div>

	<div class="court-seal">
		<p><em>Uitgesproken in het openbaar</em></p>
		<p><small>{new Date(verdict.timestamp).toLocaleDateString('nl-NL')}</small></p>
	</div>
</div>

{#if toastVisible}
	<div class="toast" class:visible={toastVisible}>
		{toastMessage}
	</div>
{/if}

<style>
	.verdict-display {
		max-width: 800px;
		margin: 2rem auto;
		padding: 2rem;
		background: #fff;
		border-radius: 2px;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
		font-family: Georgia, serif;
		border-top: 3px solid #4a4a4a;
	}

	.court-header {
		text-align: center;
		border-bottom: 3px double #2c3e50;
		padding-bottom: 1.5rem;
		margin-bottom: 2rem;
	}

	.gavel-icon {
		font-size: 3rem;
		margin-bottom: 0.5rem;
	}

	.court-header h1 {
		color: #2c3e50;
		font-size: 1.8rem;
		margin: 0;
		font-weight: 600;
	}

	.photo-display {
		display: flex;
		justify-content: center;
		margin: 2rem 0;
		padding: 1rem;
		background: #f8f9fa;
		border: 2px solid #2c3e50;
		border-radius: 2px;
	}

	.verdict-photo {
		max-width: 100%;
		max-height: 500px;
		width: auto;
		height: auto;
		display: block;
		border-radius: 2px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
	}

	.case-metadata {
		text-align: center;
		margin: 1.5rem 0;
		font-size: 0.9rem;
		color: #666;
	}

	.case-number,
	.case-date {
		margin: 0.5rem 0;
	}

	.metadata-separator {
		border: none;
		border-top: 2px solid #d1d1d1;
		margin: 1.5rem 0;
	}

	.verdict-content {
		margin: 2rem 0;
	}

	.score-display {
		text-align: center;
		margin: 2rem 0;
		padding: 2rem;
		border-radius: 2px;
	}

	.score-number {
		font-size: 5rem;
		font-weight: bold;
		line-height: 1;
	}

	.score-label {
		font-size: 1.2rem;
		opacity: 0.8;
		margin-top: 0.5rem;
	}

	.score-display.excellent {
		background: linear-gradient(135deg, #28a745, #20c997);
		color: white;
	}

	.score-display.good {
		background: linear-gradient(135deg, #17a2b8, #3498db);
		color: white;
	}

	.score-display.moderate {
		background: linear-gradient(135deg, #ffc107, #fd7e14);
		color: white;
	}

	.score-display.poor {
		background: linear-gradient(135deg, #dc3545, #c82333);
		color: white;
	}

	.verdict-type {
		text-align: center;
		margin: 1.5rem 0;
	}

	.verdict-type h2 {
		color: #2c3e50;
		font-size: 1.8rem;
		font-weight: 600;
	}

	.verdict-body {
		line-height: 1.8;
		color: #333;
	}

	.verdict-text {
		font-size: 1.1rem;
		margin: 1.5rem 0;
		text-align: justify;
	}

	.angle-deviation {
		background: #f8f9fa;
		padding: 1rem;
		border-left: 4px solid #2c3e50;
		margin: 1.5rem 0;
	}

	.sentence {
		background: #fff3cd;
		border: 2px solid #ffc107;
		padding: 1.5rem;
		border-radius: 4px;
		margin: 2rem 0;
	}

	.sentence h3 {
		color: #856404;
		margin-top: 0;
	}

	.crime-section,
	.reasoning-section {
		background: #f8f9fa;
		padding: 1.5rem;
		border-left: 4px solid #2c3e50;
		margin: 1.5rem 0;
	}

	.crime-section h3,
	.reasoning-section h3 {
		color: #2c3e50;
		margin-top: 0;
		font-weight: 600;
	}

	.case-dismissed {
		text-align: center;
		padding: 2rem;
	}

	.case-dismissed h2 {
		color: #dc3545;
		font-size: 2rem;
		font-weight: 600;
	}

	.legal-note {
		font-style: italic;
		color: #666;
		margin-top: 1rem;
	}

	.verdict-actions {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin: 2rem 0;
		padding-top: 2rem;
		border-top: 1px solid #dee2e6;
	}

	.action-button {
		padding: 0.75rem 1.5rem;
		font-size: 1rem;
		border: none;
		border-radius: 2px;
		cursor: pointer;
		transition: all 0.2s;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
		font-weight: 500;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
	}

	.action-button:hover {
		transform: translateY(-2px);
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.15);
	}

	.action-button.primary {
		background: #2c3e50;
		color: white;
	}

	.action-button.primary:hover {
		background: #34495e;
	}

	.action-button.secondary {
		background: #757575;
		color: white;
	}

	.action-button.secondary:hover {
		background: #616161;
	}

	.action-button.secondary:hover {
		background: #5a6268;
	}

	.court-seal {
		text-align: center;
		margin-top: 2rem;
		padding-top: 1.5rem;
		border-top: 1px solid #dee2e6;
		color: #666;
		font-style: italic;
	}

	.verdict-display.guilty {
		border-top: 5px solid #dc3545;
	}

	.verdict-display.acquittal {
		border-top: 5px solid #28a745;
	}

	.verdict-display.warning {
		border-top: 5px solid #ff9800;
	}

	.verdict-display.dismissed {
		border-top: 5px solid #6c757d;
	}

	.toast {
		position: fixed;
		bottom: 2rem;
		left: 50%;
		transform: translateX(-50%) translateY(100px);
		background: #2c3e50;
		color: white;
		padding: 1rem 1.5rem;
		border-radius: 8px;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
		font-size: 0.95rem;
		z-index: 1000;
		max-width: 90%;
		word-wrap: break-word;
		opacity: 0;
		transition: all 0.3s ease-in-out;
	}

	.toast.visible {
		opacity: 1;
		transform: translateX(-50%) translateY(0);
	}

	@media (max-width: 768px) {
		.verdict-display {
			padding: 1rem;
			margin: 1rem;
		}

		.court-header h1 {
			font-size: 1.4rem;
		}

		.score-number {
			font-size: 3.5rem;
		}

		.verdict-actions {
			flex-direction: column;
		}

		.action-button {
			width: 100%;
		}

		.verdict-photo {
			max-height: 300px;
		}
	}
</style>
