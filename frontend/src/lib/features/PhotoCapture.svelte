<!--
@component
PhotoCapture - Manual photo capture component with rotation controls

A self-contained component for capturing photos using native file input with camera access.
Provides manual rotation controls and "bakes" rotation into the final image using canvas.

**Features:**
- Native file input with camera capture (mobile-friendly)
- Manual rotation controls (rotate left/right in 90° increments)
- Initial rotation heuristic using screen.orientation API
- Canvas-based rotation processing (no EXIF dependency)
- Automatic memory management (URL cleanup)
- Error handling for non-image files

**Props:**
- `onPhotoConfirmed: (blob: Blob, rotation: number) => void` - Called when user confirms photo (required)
- `onCancelled?: () => void` - Called when user cancels (optional)

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
	import { rotateLeft, rotateRight } from '$lib/shared/utils/rotation';
	import {
		getInitialRotation,
		rotateImageOnCanvas,
		canvasToBlob
	} from '$lib/shared/utils/rotationUtils';

	// Props
	export let onPhotoConfirmed: (blob: Blob, rotation: number) => void;
	export let onCancelled: (() => void) | undefined = undefined;

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
	}

	function handleRotateLeft() {
		rotation = rotateLeft(rotation);
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
				capture="camera"
				onchange={handleFileSelection}
				aria-label="Upload foto"
				class="hidden-input"
			/>
			<button onclick={triggerFileInput} class="button primary capture-button">
				📷 Neem Foto
			</button>

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
			</div>

			<div class="rotation-hint">
				<p>Staat de foto niet goed? Roteer hem eerst:</p>
			</div>

			<div class="rotation-controls">
				<button
					onclick={handleRotateLeft}
					class="button button-rotation"
					aria-label="Draai links"
					disabled={isProcessing}
				>
					↶ Links
				</button>
				<button
					onclick={handleRotateRight}
					class="button button-rotation"
					aria-label="Draai rechts"
					disabled={isProcessing}
				>
					↷ Rechts
				</button>
			</div>

			<div class="controls">
				<button onclick={handleRetake} class="button secondary" disabled={isProcessing}>
					🔄 Opnieuw
				</button>

				<button onclick={handleConfirm} class="button primary" disabled={isProcessing}>
					{isProcessing ? 'Verwerken...' : '✅ Bevestig'}
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
		margin: 1rem 0;
		border-radius: 8px;
		overflow: hidden;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
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

	.rotation-hint {
		margin: 1rem 0 0.5rem 0;
		text-align: center;
	}

	.rotation-hint p {
		margin: 0;
		font-size: 0.9rem;
		color: #666;
	}

	.rotation-controls {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin: 1rem 0;
	}

	.button-rotation {
		padding: 0.75rem 1.5rem;
		font-size: 1.1rem;
	}

	.controls {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin: 1.5rem 0;
	}

	.button {
		padding: 0.75rem 1.5rem;
		border: none;
		border-radius: 6px;
		font-size: 1rem;
		cursor: pointer;
		transition: all 0.2s;
	}

	.button:hover:not(:disabled) {
		transform: translateY(-2px);
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
	}

	.button:disabled {
		opacity: 0.6;
		cursor: not-allowed;
	}

	.button.primary {
		background: #4caf50;
		color: white;
	}

	.button.primary:hover:not(:disabled) {
		background: #45a049;
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
		border-radius: 4px;
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

		.rotation-controls {
			flex-direction: column;
		}

		.button {
			width: 100%;
		}
	}
</style>
