<script lang="ts">
	import { page } from '$app/stores';
</script>

<svelte:head>
	<title>Fout - Rechtbank voor Meubilair</title>
</svelte:head>

<div class="error-page">
	<div class="error-content">
		<div class="gavel-icon">🔨</div>
		<h1>Fout {$page.status}</h1>
		<p class="error-message">{$page.error?.message || 'Er is een fout opgetreden'}</p>

		{#if $page.status === 404}
			<p class="hint">Dit vonnis bestaat niet of is niet meer beschikbaar.</p>
		{:else if $page.status === 400}
			<p class="hint">De vonnis-link is ongeldig.</p>
		{:else}
			<p class="hint">Probeer het later opnieuw.</p>
		{/if}

		<a href="/" class="back-button">Terug naar Rechtbank</a>
	</div>
</div>

<style>
	.error-page {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		background: #f5f5f5;
		padding: 2rem;
	}

	.error-content {
		max-width: 600px;
		text-align: center;
		background: white;
		padding: 3rem 2rem;
		border-radius: 2px;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
		border-top: 5px solid #dc3545;
	}

	.gavel-icon {
		font-size: 4rem;
		margin-bottom: 1rem;
	}

	h1 {
		color: #2c3e50;
		font-size: 2rem;
		margin: 1rem 0;
		font-family: Georgia, serif;
	}

	.error-message {
		font-size: 1.2rem;
		color: #666;
		margin: 1rem 0;
	}

	.hint {
		color: #999;
		font-style: italic;
		margin: 1.5rem 0;
	}

	.back-button {
		display: inline-block;
		margin-top: 2rem;
		padding: 0.75rem 1.5rem;
		background: #2c3e50;
		color: white;
		text-decoration: none;
		border-radius: 2px;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
		transition: all 0.2s;
	}

	.back-button:hover {
		background: #34495e;
		transform: translateY(-2px);
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.15);
	}
</style>
