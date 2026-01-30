<script lang="ts">
	export let httpsRequired = false;
	export let onpermissionrequested: ((event: CustomEvent) => void) | undefined = undefined;

	function requestPermission() {
		onpermissionrequested?.(new CustomEvent('permission-requested'));
	}
</script>

<div class="camera-permission">
	<h2>📸 Camera Toestemming</h2>

	{#if httpsRequired}
		<div class="warning">
			<p><strong>⚠️ HTTPS Vereist</strong></p>
			<p>
				Camera toegang werkt alleen via HTTPS. Voor lokale ontwikkeling gebruik localhost of
				configureer HTTPS met mkcert.
			</p>
		</div>
	{:else}
		<p>De Rechtbank voor Meubilair vereist toegang tot uw camera om meubels te fotograferen.</p>

		<button onclick={requestPermission} class="permission-button"> Sta Camera Toe </button>

		<p class="help-text">
			<small
				>U kunt ook een foto uploaden als u geen camera toestemming wilt geven of als uw browser dit
				niet ondersteunt.</small
			>
		</p>
	{/if}
</div>

<style>
	.camera-permission {
		max-width: 500px;
		margin: 2rem auto;
		padding: 2rem;
		text-align: center;
		background: #f5f5f5;
		border-radius: 8px;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
	}

	.warning {
		background: #fff3cd;
		border: 1px solid #ffc107;
		border-radius: 4px;
		padding: 1rem;
		margin: 1rem 0;
	}

	.permission-button {
		background: #2c3e50;
		color: white;
		border: none;
		padding: 1rem 2rem;
		font-size: 1.1rem;
		border-radius: 4px;
		cursor: pointer;
		margin: 1rem 0;
		transition: background 0.3s;
	}

	.permission-button:hover {
		background: #34495e;
	}

	.help-text {
		color: #666;
		margin-top: 1rem;
	}
</style>
