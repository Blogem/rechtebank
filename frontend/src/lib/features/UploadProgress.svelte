<script lang="ts">
	export let progress = 0; // 0-100
	export let message = 'De rechter beraadslaagt...';
</script>

<div class="upload-progress">
	<div class="court-icon">⚖️</div>

	<h2 class="progress-heading">{message}</h2>

	<div class="progress-bar">
		<div class="progress-fill" style="width: {progress}%"></div>
	</div>

	<p class="legal-notice">
		<em>Uw meubelstuk wordt beoordeeld op rechtheid en deugdelijkheid</em>
	</p>
</div>

<style>
	.upload-progress {
		max-width: 600px;
		margin: 0 auto;
		padding: 3rem 2rem;
		text-align: center;
		background: var(--color-court-surface);
		border-radius: 4px;
		box-shadow: var(--shadow-base);
		border-top: 3px solid var(--color-court-accent);
	}

	.court-icon {
		font-size: 4rem;
		margin-bottom: 1rem;
		animation: swing 2s ease-in-out infinite;
	}

	@keyframes swing {
		0%,
		100% {
			transform: rotate(-5deg);
		}
		50% {
			transform: rotate(5deg);
		}
	}

	.progress-heading {
		color: var(--color-court-primary);
		margin-bottom: 1.5rem;
		font-size: 1.5rem;
	}

	.progress-bar {
		width: 100%;
		height: 30px;
		background: var(--color-court-bg);
		border-radius: 4px;
		overflow: hidden;
		margin: 1.5rem 0;
		box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.1);
		border: 1px solid var(--color-court-border);
	}

	.progress-fill {
		height: 100%;
		background: var(--color-court-primary);
		transition: width var(--timing-interactive) var(--ease-out);
	}

	.legal-notice {
		color: var(--color-court-text-light);
		font-style: italic;
		margin-top: 1rem;
		font-size: 0.95rem;
	}

	@media (prefers-reduced-motion: reduce) {
		.court-icon {
			animation: none;
		}
	}

	@media (max-width: 768px) {
		.upload-progress {
			padding: 2rem 1.5rem;
		}

		.court-icon {
			font-size: 3rem;
		}

		.progress-heading {
			font-size: 1.3rem;
		}
	}
</style>
