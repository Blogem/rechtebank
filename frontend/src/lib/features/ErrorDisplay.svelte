<script lang="ts">
	export let message: string;
	export let retryable = true;
	export let onretry: ((event: CustomEvent) => void) | undefined = undefined;
	export let onreset: ((event: CustomEvent) => void) | undefined = undefined;

	let showTechnicalDetails = false;

	function retry() {
		onretry?.(new CustomEvent('retry'));
	}

	function reset() {
		onreset?.(new CustomEvent('reset'));
	}

	function toggleDetails() {
		showTechnicalDetails = !showTechnicalDetails;
	}

	// Extract user-friendly message from technical error
	function getUserFriendlyMessage(errorMessage: string): string {
		// Check for specific error patterns
		if (errorMessage.includes('te lang beraadslaagd')) {
			return 'De rechtbank heeft te lang nodig voor beraadslaging.';
		}
		if (errorMessage.includes('Server error (5')) {
			return 'De rechtbank ondervindt momenteel technische problemen.';
		}
		if (errorMessage.includes('Server error (4')) {
			return 'Het bewijsmateriaal kon niet worden geaccepteerd.';
		}
		if (errorMessage.includes('te groot') || errorMessage.includes('10MB')) {
			return 'De ingediende foto is te groot voor beoordeling.';
		}
		if (
			errorMessage.toLowerCase().includes('network') ||
			errorMessage.includes('Failed to fetch')
		) {
			return 'De verbinding met de rechtbank is verbroken.';
		}

		// Generic fallback
		return 'De rechtbank kan het bewijsmateriaal momenteel niet beoordelen.';
	}

	$: userFriendlyMessage = getUserFriendlyMessage(message);
</script>

<div class="error-document">
	<div class="error-header">
		<div class="error-icon-header">🔨</div>
		<h1 class="error-title">Zaak Opgeschort</h1>
	</div>

	<div class="error-content">
		<div class="error-icon">⚠️</div>
		<p class="error-message">{userFriendlyMessage}</p>

		<div class="legal-notice">
			<p><em>De rechtbank kan op dit moment geen uitspraak doen over uw zaak.</em></p>
		</div>

		<div class="technical-details">
			<button onclick={toggleDetails} class="details-toggle">
				{showTechnicalDetails ? '▼' : '▶'} Technische details
			</button>

			{#if showTechnicalDetails}
				<div class="details-content">
					<pre>{message}</pre>
				</div>
			{/if}
		</div>
	</div>

	<div class="error-actions">
		{#if retryable}
			<button onclick={retry} class="action-button primary">Probeer Opnieuw</button>
		{/if}
		<button onclick={reset} class="action-button secondary">Nieuwe Zaak</button>
	</div>
</div>

<style>
	.error-document {
		max-width: 600px;
		margin: 0 auto;
		background: var(--color-court-surface);
		border-radius: 4px;
		box-shadow: var(--shadow-base);
		border-top: 4px solid #dc3545;
		overflow: hidden;
		opacity: 0;
		transform: translateY(8px);
		animation: revealError var(--timing-reveal) var(--ease-out) forwards;
	}

	@keyframes revealError {
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.error-header {
		background: var(--color-court-primary);
		color: white;
		text-align: center;
		padding: 2rem 1.5rem;
		border-bottom: 3px solid #dc3545;
	}

	.error-icon-header {
		font-size: 2.5rem;
		margin-bottom: 0.75rem;
	}

	.error-title {
		font-size: 1.75rem;
		margin: 0;
		color: white;
	}

	.error-content {
		text-align: center;
		padding: 2rem 1.5rem;
	}

	.error-icon {
		font-size: 3.5rem;
		margin-bottom: 1.25rem;
	}

	.error-message {
		font-size: 1.15rem;
		color: #dc3545;
		margin: 1.5rem 0;
		line-height: 1.6;
		font-family: var(--font-serif);
	}

	.legal-notice {
		font-style: italic;
		color: var(--color-court-text-light);
		margin-top: 1.5rem;
		font-size: 0.95rem;
	}

	.error-actions {
		display: flex;
		gap: 1rem;
		justify-content: center;
		padding: 1.5rem 2rem 2rem;
		border-top: 1px solid var(--color-court-border);
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

	.technical-details {
		margin-top: 2rem;
		padding-top: 1.5rem;
		border-top: 1px solid var(--color-court-border);
	}

	.details-toggle {
		background: transparent;
		border: 1px solid var(--color-court-text-light);
		color: var(--color-court-text-light);
		padding: 0.5rem 1rem;
		border-radius: 4px;
		cursor: pointer;
		font-size: 0.9rem;
		font-family: var(--font-sans);
		transition: all var(--timing-default) var(--ease-out);
	}

	.details-toggle:hover {
		background: var(--color-court-text-light);
		color: white;
	}

	.details-content {
		margin-top: 1rem;
		background: var(--color-court-bg);
		border: 1px solid var(--color-court-border);
		border-radius: 4px;
		padding: 1rem;
		text-align: left;
		max-height: 300px;
		overflow-y: auto;
	}

	.details-content pre {
		margin: 0;
		font-family: 'Courier New', monospace;
		font-size: 0.85rem;
		color: var(--color-court-text);
		white-space: pre-wrap;
		word-wrap: break-word;
	}

	@media (prefers-reduced-motion: reduce) {
		.error-document {
			animation: none;
			opacity: 1;
			transform: none;
		}
	}

	@media (max-width: 768px) {
		.error-header {
			padding: 1.5rem 1rem;
		}

		.error-title {
			font-size: 1.5rem;
		}

		.error-content {
			padding: 1.5rem 1rem;
		}

		.error-actions {
			flex-direction: column;
			padding: 1rem;
		}

		.action-button {
			width: 100%;
		}
	}
</style>
