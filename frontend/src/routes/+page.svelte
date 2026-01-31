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
	import courtSeal from '$lib/assets/court-seal.png';

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

	function handleRetry() {
		// Go back to camera to retry the photo submission
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
		<img src={courtSeal} alt="Court Seal" class="court-seal" />
		<h1>⚖️ Rechtbank voor Meubilair</h1>
		<p class="tagline">Uw meubels aan het oordeel onderworpen</p>
	</header>

	<main class="app-main">
		{#if $appState === 'requesting-permissions'}
			<CameraPermission httpsRequired={false} onpermissionrequested={handlePermissionRequested} />
		{:else if $appState === 'camera-ready'}
			<div class="introduction">
				<p class="welcome-text">
					Welkom bij de Rechtbank voor Meubilair, waar elk meubelstuk rekenschap moet afleggen van
					zijn houding en constructie. Dien uw bewijsmateriaal in en verneem het onverbiddelijke
					oordeel van de Eerwaarde Rechter.
				</p>
			</div>
			<div class="section-divider"></div>
			<div class="camera-section">
				<PhotoCapture onPhotoConfirmed={handlePhotoConfirmed} />
			</div>
		{:else if $appState === 'uploading'}
			<UploadProgress progress={50} message="De rechter beraadslaagt..." />
		{:else if $appState === 'showing-verdict' && $currentVerdict}
			<VerdictDisplay verdict={$currentVerdict} onreset={handleReset} />
		{:else if $appState === 'error'}
			<ErrorDisplay
				message={$uploadError || 'Er is een fout opgetreden'}
				onreset={handleReset}
				onretry={handleRetry}
			/>
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
		background: #fafafa;
		min-height: 100vh;
	}

	.app-container {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
	}

	.app-header {
		background: #2e2e2e;
		color: white;
		text-align: center;
		padding: 2rem 1rem;
		box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
		border-bottom: 3px solid #4a4a4a;
		position: relative;
	}

	.app-header h1 {
		margin: 0;
		font-size: 2.5rem;
		font-family: Georgia, serif;
		font-weight: 600;
	}

	.court-seal {
		position: absolute;
		left: 2rem;
		top: 50%;
		transform: translateY(-50%);
		width: 80px;
		height: 80px;
		opacity: 0.95;
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

	.introduction {
		background: white;
		border-radius: 2px;
		padding: 2rem;
		margin-bottom: 1.5rem;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
		border-top: 3px solid #4a4a4a;
	}

	.welcome-text {
		font-size: 1.1rem;
		line-height: 1.7;
		color: #333;
		margin: 0;
		text-align: center;
		font-family: Georgia, serif;
	}

	.camera-section {
		background: white;
		border-radius: 2px;
		padding: 1.5rem;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
		border-top: 3px solid #4a4a4a;
	}

	.section-divider {
		height: 2px;
		background: #d1d1d1;
		margin: 1.5rem 0;
	}

	.app-footer {
		background: #2e2e2e;
		color: white;
		text-align: center;
		padding: 1.5rem 1rem;
		margin-top: 2rem;
		border-top: 3px solid #4a4a4a;
	}

	.app-footer p {
		margin: 0;
		opacity: 0.8;
	}

	@media (max-width: 768px) {
		.app-header h1 {
			font-size: 1.8rem;
		}

		.court-seal {
			width: 50px;
			height: 50px;
			left: 0.5rem;
		}

		.app-main {
			padding: 1rem 0.5rem;
		}
	}

	@media (max-width: 480px) {
		.app-header h1 {
			font-size: 1.4rem;
		}

		.court-seal {
			display: none;
		}
	}
</style>
