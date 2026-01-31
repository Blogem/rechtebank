<script lang="ts">
	export let progress = 0; // 0-100
	export let message = 'De rechter beraadslaagt...';
</script>

<div class="upload-progress">
	<div class="court-icon">⚖️</div>

	<h2>{message}</h2>

	<div class="progress-bar">
		<div class="progress-fill" style="width: {progress}%"></div>
	</div>

	<p class="legal-text">
		<em>Uw meubel wordt beoordeeld op rechtheid en deugdelijkheid...</em>
	</p>
</div>

<style>
	.upload-progress {
		max-width: 600px;
		margin: 3rem auto;
		padding: 3rem 2rem;
		text-align: center;
		background: #f8f9fa;
		border-radius: 2px;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
		border-top: 3px solid #4a4a4a;
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

	h2 {
		color: #2c3e50;
		margin-bottom: 1.5rem;
		font-family: Georgia, serif;
	}

	.progress-bar {
		width: 100%;
		height: 30px;
		background: #e9ecef;
		border-radius: 2px;
		overflow: hidden;
		margin: 1.5rem 0;
		box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.1);
	}

	.progress-fill {
		height: 100%;
		background: #2c3e50;
		transition: width 0.3s ease-out;
		border-radius: 2px;
	}

	.legal-text {
		color: #666;
		font-style: italic;
		margin-top: 1rem;
	}
</style>
