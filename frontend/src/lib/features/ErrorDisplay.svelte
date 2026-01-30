<script lang="ts">
	export let message: string;
	export let retryable = true;
	export let onretry: ((event: CustomEvent) => void) | undefined = undefined;
	export let onreset: ((event: CustomEvent) => void) | undefined = undefined;

	function retry() {
		onretry?.(new CustomEvent('retry'));
	}

	function reset() {
		onreset?.(new CustomEvent('reset'));
	}
</script>

<div class="error-display">
	<div class="court-header">
		<div class="gavel-icon">🔨</div>
		<h1>Zaak Opgeschort</h1>
	</div>

	<div class="error-content">
		<div class="error-icon">⚠️</div>
		<p class="error-message">{message}</p>

		<div class="legal-language">
			<p><em>De rechtbank kan op dit moment geen uitspraak doen.</em></p>
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
</style>
