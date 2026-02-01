<!--
@component
PhotoCapture - Manual photo capture component with rotation control

A self-contained component for capturing photos using native file input with camera access.
Provides a manual rotation control overlay and "bakes" rotation into the final image using canvas.

**Features:**
- Native file input with camera capture (mobile-friendly)
- Manual rotation control (rotate clockwise in 90° increments via overlay button)
- Initial rotation heuristic using screen.orientation API
- Canvas-based rotation processing (no EXIF dependency)
- Automatic memory management (URL cleanup)
- Error handling for non-image files

**Props:**
- `onPhotoConfirmed: (blob: Blob, rotation: number) => void` - Called when user confirms photo (required)
- `onPreviewStateChange?: (isPreviewing: boolean) => void` - Called when preview state changes (optional)

**Usage:**
```svelte
<PhotoCapture 
  onPhotoConfirmed={(blob, rotation) => {
    // Handle confirmed photo
    console.log('Photo confirmed with rotation:', rotation);
  }}
/>
```
-->
<script lang="ts">
	import { onDestroy } from 'svelte';
	import { rotateRight } from '$lib/shared/utils/rotation';
	import {
		getInitialRotation,
		rotateImageOnCanvas,
		canvasToBlob
	} from '$lib/shared/utils/rotationUtils';

	// Props
	export let onPhotoConfirmed: (blob: Blob, rotation: number) => void;
	export let onPreviewStateChange: ((isPreviewing: boolean) => void) | undefined = undefined;

	// State
	let fileInput: HTMLInputElement;
	let currentFile: File | null = null;
	let photoUrl: string | null = null;
	let rotation = 0;
	let error: string | null = null;
	let isProcessing = false;

	// Lifecycle cleanup
	onDestroy(() => {
		cleanupObjectUrl();
	});

	function cleanupObjectUrl() {
		if (photoUrl) {
			URL.revokeObjectURL(photoUrl);
			photoUrl = null;
		}
	}

	function triggerFileInput() {
		fileInput?.click();
	}

	async function handleFileSelection(event: Event) {
		const input = event.target as HTMLInputElement;
		const file = input.files?.[0];

		if (!file) {
			return;
		}

		// Validate MIME type
		if (!file.type.startsWith('image/')) {
			error = 'Alleen afbeeldingen zijn toegestaan';
			return;
		}

		// Clean up previous photo
		cleanupObjectUrl();
		error = null;

		// Create object URL for preview
		currentFile = file;
		photoUrl = URL.createObjectURL(file);

		// Set initial rotation based on screen orientation
		rotation = getInitialRotation();

		// Notify parent that we're now in preview mode
		onPreviewStateChange?.(true);
	}

	function handleRotateRight() {
		rotation = rotateRight(rotation);
	}

	async function handleConfirm() {
		if (!photoUrl || !currentFile) {
			return;
		}

		isProcessing = true;
		error = null;

		try {
			// Load image
			const img = new Image();
			img.src = photoUrl;

			await new Promise((resolve, reject) => {
				img.onload = resolve;
				img.onerror = () => reject(new Error('Failed to load image'));
			});

			// Rotate image on canvas
			const canvas = rotateImageOnCanvas(img, rotation);

			// Convert to blob
			const blob = await canvasToBlob(canvas);

			// Call callback with rotated blob and rotation value
			onPhotoConfirmed(blob, rotation);

			// Clean up
			cleanupObjectUrl();
			currentFile = null;
			rotation = 0;
		} catch (err) {
			error = 'Failed to process image';
			console.error('Image processing error:', err);
		} finally {
			isProcessing = false;
		}
	}

	function handleRetake() {
		cleanupObjectUrl();
		currentFile = null;
		rotation = 0;
		error = null;

		// Notify parent that we're leaving preview mode
		onPreviewStateChange?.(false);

		// Trigger file input again
		triggerFileInput();
	}
</script>

<div class="photo-capture">
	{#if !photoUrl}
		<!-- Capture button -->
		<div class="capture-section">
			<input
				bind:this={fileInput}
				type="file"
				accept="image/*"
				capture="environment"
				onchange={handleFileSelection}
				aria-label="Upload foto"
				class="hidden-input"
			/>
			<button onclick={triggerFileInput} class="button primary capture-button"> Neem Foto </button>

			{#if error}
				<div class="error-message" role="alert">
					{error}
				</div>
			{/if}
		</div>
	{:else}
		<!-- Preview and controls -->
		<div class="preview-section">
			<h2>Bevestig Foto</h2>

			<div class="preview">
				<img src={photoUrl} alt="Preview" style="transform: rotate({rotation}deg);" />
				<button
					onclick={handleRotateRight}
					class="rotation-button-overlay"
					aria-label="Roteer foto 90 graden"
					disabled={isProcessing}
				>
					↻
				</button>
			</div>

			<div class="controls">
				<button onclick={handleConfirm} class="button primary" disabled={isProcessing}>
					{isProcessing ? 'Verwerken...' : 'Bevestig'}
				</button>

				<button onclick={handleRetake} class="button secondary" disabled={isProcessing}>
					Opnieuw
				</button>
			</div>

			{#if error}
				<div class="error-message" role="alert">
					{error}
				</div>
			{/if}
		</div>
	{/if}
</div>

<style>
	.photo-capture {
		max-width: 800px;
		margin: 2rem auto;
		text-align: center;
	}

	.hidden-input {
		position: absolute;
		width: 1px;
		height: 1px;
		padding: 0;
		margin: -1px;
		overflow: hidden;
		clip: rect(0, 0, 0, 0);
		white-space: nowrap;
		border-width: 0;
	}

	.capture-section {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1rem;
	}

	.capture-button {
		font-size: 1.25rem;
		padding: 1.5rem 3rem;
	}

	.preview-section {
		width: 100%;
	}

	.preview {
		position: relative;
		margin: 1rem 0;
		border-radius: 2px;
		overflow: hidden;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
		max-height: 70vh;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.preview img {
		max-width: 100%;
		max-height: 70vh;
		height: auto;
		display: block;
		transition: transform 0.3s ease;
	}

	.rotation-button-overlay {
		position: absolute;
		bottom: 16px;
		right: 16px;
		z-index: 10;
		background: rgba(0, 0, 0, 0.6);
		color: white;
		min-width: 48px;
		min-height: 48px;
		width: 48px;
		height: 48px;
		border: none;
		border-radius: 50%;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
		font-size: 1.5rem;
		cursor: pointer;
		transition: all 0.2s ease;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.rotation-button-overlay:hover:not(:disabled) {
		background: rgba(0, 0, 0, 0.8);
		transform: scale(1.05);
	}

	.rotation-button-overlay:disabled {
		opacity: 0.4;
		cursor: not-allowed;
	}

	.controls {
		display: flex;
		flex-direction: column;
		gap: 1rem;
		justify-content: center;
		align-items: center;
		margin: 1.5rem 0;
	}

	.button {
		padding: 0.75rem 1.5rem;
		border: none;
		border-radius: 2px;
		font-size: 1rem;
		cursor: pointer;
		transition: all 0.2s;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
	}

	.button:hover:not(:disabled) {
		transform: translateY(-2px);
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.15);
	}

	.button:disabled {
		opacity: 0.6;
		cursor: not-allowed;
	}

	.button.primary {
		background: #2c3e50;
		color: white;
	}

	.button.primary:hover:not(:disabled) {
		background: #34495e;
	}

	.button.secondary {
		background: #757575;
		color: white;
	}

	.button.secondary:hover:not(:disabled) {
		background: #616161;
	}

	.error-message {
		color: #d32f2f;
		background: #ffebee;
		padding: 0.75rem;
		border-radius: 2px;
		margin-top: 1rem;
		font-size: 0.9rem;
	}

	@media (max-width: 600px) {
		.photo-capture {
			margin: 1rem;
		}

		.controls {
			flex-direction: column;
		}

		.button {
			width: 100%;
		}
	}
</style>
