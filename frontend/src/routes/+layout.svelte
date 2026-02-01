<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.png';
	import courtSeal from '$lib/assets/court-seal.png';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
</svelte:head>

<div class="layout-container">
	<header class="masthead">
		<div class="masthead-content">
			<img src={courtSeal} alt="Zegel van de Rechtbank" class="court-seal" />
			<div class="court-identity">
				<h1 class="court-name">⚖️ Rechtbank voor Meubilair</h1>
				<p class="court-tagline">Rechtvaardigheid in elk meubelstuk</p>
			</div>
		</div>
	</header>

	<main class="main-content">
		{@render children()}
	</main>

	<footer class="site-footer">
		<div class="footer-content">
			<nav class="footer-links">
				<a href="#privacybeleid">Privacybeleid</a>
				<a href="#procesreglement">Procesreglement</a>
				<a href="#over">Over deze site</a>
			</nav>
			<p class="footer-copyright">© 2026 Rechtbank voor Meubilair</p>
		</div>
	</footer>
</div>

<style>
	.layout-container {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
	}

	.masthead {
		background: var(--color-court-primary);
		color: white;
		padding: 1.5rem 1rem;
		box-shadow: var(--shadow-md);
		position: relative;
	}

	.masthead-content {
		max-width: 1200px;
		margin: 0 auto;
		display: flex;
		align-items: center;
		gap: 1.5rem;
	}

	.court-seal {
		width: 80px;
		height: 80px;
		flex-shrink: 0;
		filter: brightness(1.1);
	}

	.court-identity {
		flex: 1;
		text-align: center;
	}

	.court-name {
		font-size: 2.25rem;
		font-weight: 700;
		margin: 0;
		color: white;
		letter-spacing: 0.02em;
	}

	.court-tagline {
		margin: 0.5rem 0 0 0;
		font-size: 1rem;
		font-style: italic;
		opacity: 0.95;
		font-family: var(--font-serif);
	}

	.main-content {
		flex: 1;
		width: 100%;
		max-width: 1200px;
		margin: 0 auto;
		padding: 2rem 1rem;
	}

	.site-footer {
		background: var(--color-court-primary);
		color: white;
		padding: 2rem 1rem;
		margin-top: 3rem;
		border-top: 3px solid var(--color-court-accent);
	}

	.footer-content {
		max-width: 1200px;
		margin: 0 auto;
		text-align: center;
	}

	.footer-disclaimer {
		font-size: 0.9rem;
		opacity: 0.9;
		margin-bottom: 1.5rem;
		font-style: italic;
		max-width: 600px;
		margin-left: auto;
		margin-right: auto;
	}

	.footer-links {
		display: flex;
		justify-content: center;
		gap: 2rem;
		margin-bottom: 1.5rem;
		flex-wrap: wrap;
	}

	.footer-links a {
		color: white;
		text-decoration: none;
		font-size: 0.9rem;
		opacity: 0.9;
		transition: opacity var(--timing-default) var(--ease-out);
	}

	.footer-links a:hover {
		opacity: 1;
		text-decoration: underline;
	}

	.footer-copyright {
		font-size: 0.85rem;
		opacity: 0.8;
		margin: 0;
	}

	@media (max-width: 768px) {
		.masthead-content {
			gap: 1rem;
		}

		.court-seal {
			width: 60px;
			height: 60px;
		}

		.court-name {
			font-size: 1.75rem;
		}

		.court-tagline {
			font-size: 0.9rem;
		}

		.main-content {
			padding: 1.5rem 1rem;
		}

		.footer-links {
			flex-direction: column;
			gap: 0.75rem;
		}
	}

	@media (max-width: 480px) {
		.masthead {
			padding: 1rem 0.75rem;
		}

		.masthead-content {
			gap: 0.75rem;
		}

		.court-seal {
			width: 50px;
			height: 50px;
		}

		.court-name {
			font-size: 1.4rem;
		}

		.court-tagline {
			font-size: 0.85rem;
		}

		.main-content {
			padding: 1rem 0.75rem;
		}
	}
</style>
