<script lang="ts">
	export let httpsRequired = false;
	export let onpermissionrequested: ((event: CustomEvent) => void) | undefined = undefined;

	function requestPermission() {
		onpermissionrequested?.(new CustomEvent('permission-requested'));
	}
</script>

<div class="camera-permission">
	<div class="permission-icon">📸</div>
	<h2 class="permission-heading">Camera Toestemming Vereist</h2>

	{#if httpsRequired}
		<div class="warning-notice">
			<p class="warning-title"><strong>⚠️ HTTPS Vereist</strong></p>
			<p>
				Camera toegang werkt alleen via HTTPS. Voor lokale ontwikkeling gebruik localhost of
				configureer HTTPS met mkcert.
			</p>
		</div>
	{:else}
		<p class="permission-text">
			De Rechtbank voor Meubilair vereist toegang tot uw camera om bewijsmateriaal te verzamelen.
		</p>

		<button onclick={requestPermission} class="permission-button">Sta Camera Toe</button>

		<p class="help-text">
			<small>
				U kunt ook een foto uploaden als u geen camera toestemming wilt verlenen of als uw browser
				dit niet ondersteunt.
			</small>
		</p>
	{/if}
</div>

<style>
	.camera-permission {
		max-width: 500px;
		margin: 0 auto;
		padding: 2rem;
		text-align: center;
		background: var(--color-court-surface);
		border-radius: 4px;
		box-shadow: var(--shadow-base);
		border-top: 3px solid var(--color-court-primary);
	}

	.permission-icon {
		font-size: 3rem;
		margin-bottom: 1rem;
	}

	.permission-heading {
		color: var(--color-court-primary);
		margin-bottom: 1.5rem;
		font-size: 1.5rem;
	}

	.permission-text {
		color: var(--color-court-text);
		line-height: 1.6;
		margin: 1.5rem 0;
	}

	.warning-notice {
		background: #fff3cd;
		border: 2px solid #ffc107;
		border-radius: 4px;
		padding: 1rem;
		margin: 1rem 0;
	}

	.warning-title {
		margin: 0 0 0.5rem 0;
	}

	.permission-button {
		background: var(--color-court-primary);
		color: white;
		border: none;
		padding: 0.875rem 2rem;
		font-size: 1rem;
		border-radius: 4px;
		cursor: pointer;
		margin: 1.5rem 0;
		font-family: var(--font-sans);
		font-weight: 600;
		box-shadow: var(--shadow-base);
		transition: all var(--timing-interactive) var(--ease-out);
	}

	.permission-button:hover {
		background: #244a66;
		transform: translateY(-3px);
		box-shadow: var(--shadow-md);
	}

	.permission-button:active {
		transform: translateY(-1px) scale(0.98);
		box-shadow: var(--shadow-sm);
		transition: all var(--timing-quick) var(--ease-out);
	}

	.permission-button:focus-visible {
		outline: none;
		box-shadow:
			var(--shadow-md),
			0 0 0 3px var(--color-court-accent);
	}

	.help-text {
		color: var(--color-court-text-light);
		margin-top: 1rem;
		font-size: 0.9rem;
	}

	@media (max-width: 768px) {
		.camera-permission {
			padding: 1.5rem;
		}

		.permission-heading {
			font-size: 1.3rem;
		}

		.permission-button {
			width: 100%;
		}
	}
</style>
