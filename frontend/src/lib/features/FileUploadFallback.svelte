<script lang="ts">
	export let onfileselected: ((event: CustomEvent<{ file: File }>) => void) | undefined = undefined;
	const MAX_FILE_SIZE = 10 * 1024 * 1024; // 10MB
	const ACCEPTED_FORMATS = ['image/jpeg', 'image/png', 'image/webp', 'image/gif'];

	let fileInput: HTMLInputElement;
	let selectedFile: File | null = null;
	let previewUrl: string | null = null;
	let error: string | null = null;

	function handleFileSelect(event: Event) {
		const target = event.target as HTMLInputElement;
		const file = target.files?.[0];

		if (!file) return;

		error = null;

		// Validate file type
		if (!ACCEPTED_FORMATS.includes(file.type)) {
			error = 'Ongeldig bestandsformaat. Gebruik JPEG, PNG, WEBP of GIF.';
			return;
		}

		// Validate file size
		if (file.size > MAX_FILE_SIZE) {
			error = 'Bestand is te groot. Maximaal 10MB toegestaan.';
			return;
		}

		selectedFile = file;
		previewUrl = URL.createObjectURL(file);
	}

	function uploadFile() {
		if (selectedFile) {
			onfileselected?.(new CustomEvent('file-selected', { detail: { file: selectedFile } }));
		}
	}

	function clearSelection() {
		selectedFile = null;
		if (previewUrl) {
			URL.revokeObjectURL(previewUrl);
			previewUrl = null;
		}
		if (fileInput) {
			fileInput.value = '';
		}
		error = null;
	}

	function triggerFileInput() {
		fileInput?.click();
	}
</script>

<div class="file-upload-fallback">
	<h2>📎 Upload Meubelfoto</h2>

	<p>Upload een foto van uw meubelstuk voor oordeel van de rechtbank.</p>

	<input
		bind:this={fileInput}
		type="file"
		accept="image/jpeg,image/png,image/webp,image/gif"
		onchange={handleFileSelect}
		style="display: none;"
	/>

	{#if error}
		<div class="error-message">
			<p>⚠️ {error}</p>
		</div>
	{/if}

	{#if previewUrl}
		<div class="preview">
			<img src={previewUrl} alt="Selected file preview" />
		</div>

		<div class="controls">
			<button onclick={clearSelection} class="button secondary"> Annuleren </button>

			<button onclick={uploadFile} class="button primary"> Indienen voor Vonnis </button>
		</div>
	{:else}
		<button onclick={triggerFileInput} class="upload-button"> 📁 Kies Bestand </button>

		<p class="help-text">
			<small>Maximaal 10MB | JPEG, PNG, WEBP, GIF</small>
		</p>
	{/if}
</div>

<style>
	.file-upload-fallback {
		max-width: 600px;
		margin: 2rem auto;
		padding: 2rem;
		text-align: center;
		background: #f5f5f5;
		border-radius: 8px;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
	}

	.upload-button {
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

	.upload-button:hover {
		background: #34495e;
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
		max-height: 500px;
		object-fit: contain;
	}

	.error-message {
		background: #f8d7da;
		border: 1px solid #f5c6cb;
		color: #721c24;
		padding: 0.75rem;
		border-radius: 4px;
		margin: 1rem 0;
	}

	.controls {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin-top: 1.5rem;
	}

	.button {
		padding: 0.75rem 1.5rem;
		font-size: 1rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		transition: all 0.3s;
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

	.help-text {
		color: #666;
		margin-top: 0.5rem;
	}
</style>
