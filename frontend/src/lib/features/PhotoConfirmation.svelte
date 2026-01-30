<script lang="ts">
	export let photoUrl: string;
	export let onconfirm: ((event: CustomEvent) => void) | undefined = undefined;
	export let onretake: ((event: CustomEvent) => void) | undefined = undefined;

	function confirmPhoto() {
		onconfirm?.(new CustomEvent('confirm'));
	}

	function retakePhoto() {
		onretake?.(new CustomEvent('retake'));
	}
</script>

<div class="photo-confirmation">
	<h2>Bevestig Foto</h2>

	<div class="preview">
		<img src={photoUrl} alt="Captured furniture" />
	</div>

	<div class="controls">
		<button onclick={retakePhoto} class="button secondary"> 🔄 Opnieuw </button>

		<button onclick={confirmPhoto} class="button primary"> ✅ Indienen voor Vonnis </button>
	</div>
</div>

<style>
	.photo-confirmation {
		max-width: 800px;
		margin: 2rem auto;
		text-align: center;
	}

	.preview {
		margin: 1rem 0;
		border-radius: 8px;
		overflow: hidden;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
	}

	.preview img {
		width: 100%;
		height: auto;
		display: block;
	}

	.controls {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin-top: 1.5rem;
	}

	.button {
		padding: 1rem 2rem;
		font-size: 1.1rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		transition: all 0.3s;
		font-weight: 500;
	}

	.button.primary {
		background: #2c3e50;
		color: white;
	}

	.button.primary:hover {
		background: #34495e;
	}

	.button.secondary {
		background: #6c757d;
		color: white;
	}

	.button.secondary:hover {
		background: #5a6268;
	}
</style>
