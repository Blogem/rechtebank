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

<div class="error-display">
	<div class="court-header">
		<div class="gavel-icon">🔨</div>
		<h1>Zaak Opgeschort</h1>
	</div>

	<div class="error-content">
		<div class="error-icon">⚠️</div>
		<p class="error-message">{userFriendlyMessage}</p>

		<div class="legal-language">
			<p><em>De rechtbank kan op dit moment geen uitspraak doen.</em></p>
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
			<button onclick={retry} class="action-button primary"> 🔄 Probeer Opnieuw </button>
		{/if}

		<button onclick={reset} class="action-button secondary"> ← Terug naar Begin </button>
	</div>
</div>

<style>
	.error-display {
		max-width: 600px;
		margin: 2rem auto;
		padding: 2rem;
		background: #fff;
		border-radius: 8px;
		border-top: 5px solid #dc3545;
		box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
		font-family: Georgia, serif;
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
	}

	.error-content {
		text-align: center;
		margin: 2rem 0;
	}

	.error-icon {
		font-size: 4rem;
		margin-bottom: 1rem;
	}

	.error-message {
		font-size: 1.2rem;
		color: #dc3545;
		margin: 1.5rem 0;
		line-height: 1.6;
	}

	.legal-language {
		font-style: italic;
		color: #666;
		margin-top: 1.5rem;
	}

	.error-actions {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin-top: 2rem;
		padding-top: 2rem;
		border-top: 1px solid #dee2e6;
		flex-wrap: wrap;
	}

	.action-button {
		padding: 1rem 2rem;
		font-size: 1rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		transition: all 0.3s;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
		font-weight: 500;
	}

	.action-button.primary {
		background: #2c3e50;
		color: white;
	}

	.action-button.primary:hover {
		background: #34495e;
	}

	.action-button.secondary {
		background: #6c757d;
		color: white;
	}

	.action-button.secondary:hover {
		background: #5a6268;
	}

	.technical-details {
		margin-top: 2rem;
		padding-top: 1.5rem;
		border-top: 1px solid #dee2e6;
	}

	.details-toggle {
		background: transparent;
		border: 1px solid #6c757d;
		color: #6c757d;
		padding: 0.5rem 1rem;
		border-radius: 4px;
		cursor: pointer;
		font-size: 0.9rem;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
		transition: all 0.3s;
	}

	.details-toggle:hover {
		background: #6c757d;
		color: white;
	}

	.details-content {
		margin-top: 1rem;
		background: #f8f9fa;
		border: 1px solid #dee2e6;
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
		color: #495057;
		white-space: pre-wrap;
		word-wrap: break-word;
	}
</style>
