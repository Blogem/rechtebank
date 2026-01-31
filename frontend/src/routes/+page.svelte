<script lang="ts">
	import { onMount } from 'svelte';
	import {
		appState,
		currentVerdict,
		uploadError,
		resetAppState
	} from '$lib/shared/stores/appStore';

	import { ApiAdapter } from '$lib/adapters/api/ApiAdapter';

	import PhotoCapture from '$lib/features/PhotoCapture.svelte';
	import CameraPermission from '$lib/features/CameraPermission.svelte';
	import UploadProgress from '$lib/features/UploadProgress.svelte';
	import VerdictDisplay from '$lib/features/VerdictDisplay.svelte';
	import ErrorDisplay from '$lib/features/ErrorDisplay.svelte';

	const apiAdapter = new ApiAdapter();

	onMount(() => {
		appState.set('camera-ready');
	});

	function handlePermissionRequested() {
		appState.set('camera-ready');
	}

	async function handlePhotoConfirmed(blob: Blob, rotation: number) {
		// PhotoCapture component handles rotation internally and returns rotated blob
		// Pass rotation=0 since blob is already rotated (rotation value kept for metadata/logging)
		appState.set('uploading');
		uploadError.set(null);

		try {
			const metadata = {
				userAgent: navigator.userAgent,
				timestamp: new Date().toISOString(),
				captureMethod: 'camera' as const
			};

			// Blob is already rotated by PhotoCapture, so pass rotation=0
			const verdict = await apiAdapter.uploadPhoto(blob, metadata, 0);
			currentVerdict.set(verdict);
			appState.set('showing-verdict');
		} catch (error) {
			uploadError.set(error instanceof Error ? error.message : 'Onbekende fout opgetreden');
			appState.set('error');
		}
	}

	function handleRetake() {
		// PhotoCapture component handles its own cleanup
		appState.set('camera-ready');
	}

	function handleReset() {
		resetAppState();
	}
</script>

<svelte:head>
	<title>Rechtbank voor Meubilair</title>
	<meta name="description" content="Submit your furniture for legal judgment" />
</svelte:head>

<div class="app-container">
	<header class="app-header">
		<h1>⚖️ Rechtbank voor Meubilair</h1>
		<p class="tagline">Uw meubels aan het oordeel onderworpen</p>
	</header>

	<main class="app-main">
		{#if $appState === 'requesting-permissions'}
			<CameraPermission httpsRequired={false} onpermissionrequested={handlePermissionRequested} />
		{:else if $appState === 'camera-ready'}
			<div class="camera-section">
				<PhotoCapture onPhotoConfirmed={handlePhotoConfirmed} />
			</div>
		{:else if $appState === 'uploading'}
			<UploadProgress progress={50} message="De rechter beraadslaagt..." />
		{:else if $appState === 'showing-verdict' && $currentVerdict}
			<VerdictDisplay verdict={$currentVerdict} onreset={handleReset} />
		{:else if $appState === 'error'}
			<ErrorDisplay message={$uploadError || 'Er is een fout opgetreden'} onreset={handleReset} />
		{/if}
	</main>

	<footer class="app-footer">
		<p>
			<small
				>© 2026 Rechtbank voor Meubilair | Een satirisch project | Geen echte juridische uitspraken</small
			>
		</p>
	</footer>
</div>

<style>
	:global(body) {
		margin: 0;
		padding: 0;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
		background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
		min-height: 100vh;
	}

	.app-container {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
	}

	.app-header {
		background: rgba(0, 0, 0, 0.3);
		color: white;
		text-align: center;
		padding: 2rem 1rem;
		box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
	}

	.app-header h1 {
		margin: 0;
		font-size: 2.5rem;
		font-family: Georgia, serif;
	}

	.tagline {
		margin: 0.5rem 0 0 0;
		font-style: italic;
		opacity: 0.9;
	}

	.app-main {
		flex: 1;
		padding: 2rem 1rem;
		max-width: 1200px;
		margin: 0 auto;
		width: 100%;
	}

	.camera-section {
		background: white;
		border-radius: 8px;
		padding: 1.5rem;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
	}

	.app-footer {
		background: rgba(0, 0, 0, 0.3);
		color: white;
		text-align: center;
		padding: 1.5rem 1rem;
		margin-top: 2rem;
	}

	.app-footer p {
		margin: 0;
		opacity: 0.8;
	}

	@media (max-width: 768px) {
		.app-header h1 {
			font-size: 1.8rem;
		}

		.app-main {
			padding: 1rem 0.5rem;
		}
	}
</style>
