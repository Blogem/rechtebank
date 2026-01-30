<script lang="ts">
	import { onMount, onDestroy } from 'svelte';

	export let stream: MediaStream | null = null;
	export let isLevelCheckEnabled = true;
	export let isDeviceLevel = false;
	export let oncapture: ((event: CustomEvent) => void) | undefined = undefined;
	export let onswitchcamera: ((event: CustomEvent) => void) | undefined = undefined;
	let videoElement: HTMLVideoElement;

	onMount(() => {
		if (stream && videoElement) {
			videoElement.srcObject = stream;
			// Ensure video plays on mobile browsers
			videoElement.play().catch((err) => {
				console.warn('Video autoplay failed:', err);
			});
		}
	});

	onDestroy(() => {
		if (videoElement) {
			videoElement.srcObject = null;
		}
	});

	$: if (stream && videoElement) {
		videoElement.srcObject = stream;
		// Ensure video plays when stream changes
		videoElement.play().catch((err) => {
			console.warn('Video play failed:', err);
		});
	}

	function capturePhoto() {
		oncapture?.(new CustomEvent('capture'));
	}

	function switchCamera() {
		onswitchcamera?.(new CustomEvent('switch-camera'));
	}

	$: canCapture = !isLevelCheckEnabled || isDeviceLevel;
</script>

<div class="camera-preview">
	<div class="video-container">
		<!-- svelte-ignore a11y_media_has_caption -->
		<video bind:this={videoElement} autoplay playsinline muted></video>

		{#if isLevelCheckEnabled && !isDeviceLevel}
			<div class="level-warning">
				<p>⚖️ Houd uw apparaat waterpas</p>
			</div>
		{/if}
	</div>

	<div class="controls">
		<button onclick={switchCamera} class="control-button secondary"> Wissel Camera </button>

		<button onclick={capturePhoto} class="control-button primary" disabled={!canCapture}>
			{#if canCapture}
				📸 Neem Foto
			{:else}
				🔒 Waterpas Vereist
			{/if}
		</button>
	</div>

	{#if !canCapture}
		<p class="help-text">Houd uw apparaat waterpas om een foto te maken</p>
	{/if}
</div>

<style>
	.camera-preview {
		max-width: 800px;
		margin: 0 auto;
	}

	.video-container {
		position: relative;
		background: #000;
		border-radius: 8px;
		overflow: hidden;
	}

	video {
		width: 100%;
		height: auto;
		display: block;
	}

	.level-warning {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		background: rgba(220, 53, 69, 0.9);
		color: white;
		padding: 1rem 2rem;
		border-radius: 8px;
		font-weight: bold;
	}

	.controls {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin: 1rem 0;
	}

	.control-button {
		padding: 0.75rem 1.5rem;
		font-size: 1rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		transition: all 0.3s;
	}

	.control-button.primary {
		background: #28a745;
		color: white;
	}

	.control-button.primary:hover:not(:disabled) {
		background: #218838;
	}

	.control-button.primary:disabled {
		background: #6c757d;
		cursor: not-allowed;
		opacity: 0.6;
	}

	.control-button.secondary {
		background: #6c757d;
		color: white;
	}

	.control-button.secondary:hover {
		background: #5a6268;
	}

	.help-text {
		text-align: center;
		color: #dc3545;
		font-weight: 500;
		margin-top: 0.5rem;
	}
</style>
