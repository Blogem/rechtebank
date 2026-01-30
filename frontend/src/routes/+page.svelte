<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import {
		appState,
		cameraStream,
		cameraPermissionGranted,
		capturedPhoto,
		orientationData,
		orientationPermissionGranted,
		levelCheckEnabled,
		isDeviceLevel,
		currentVerdict,
		uploadError,
		resetAppState
	} from '$lib/shared/stores/appStore';

	import { CameraAdapter } from '$lib/adapters/camera/CameraAdapter';
	import { OrientationAdapter } from '$lib/adapters/orientation/OrientationAdapter';
	import { ApiAdapter } from '$lib/adapters/api/ApiAdapter';

	import CameraPermission from '$lib/features/CameraPermission.svelte';
	import CameraPreview from '$lib/features/CameraPreview.svelte';
	import PhotoConfirmation from '$lib/features/PhotoConfirmation.svelte';
	import FileUploadFallback from '$lib/features/FileUploadFallback.svelte';
	import SpiritLevel from '$lib/features/SpiritLevel.svelte';
	import AccessibilityToggle from '$lib/features/AccessibilityToggle.svelte';
	import UploadProgress from '$lib/features/UploadProgress.svelte';
	import VerdictDisplay from '$lib/features/VerdictDisplay.svelte';
	import ErrorDisplay from '$lib/features/ErrorDisplay.svelte';

	const cameraAdapter = new CameraAdapter();
	const orientationAdapter = new OrientationAdapter();
	const apiAdapter = new ApiAdapter();

	let currentFacingMode: 'user' | 'environment' = 'environment';
	let photoObjectUrl: string | null = null;
	let showFileUpload = false;

	onMount(async () => {
		// Check if camera is supported
		if (!cameraAdapter.isCameraSupported()) {
			showFileUpload = true;
			appState.set('camera-ready');
		}

		// On non-iOS devices, start orientation monitoring immediately
		if (!orientationAdapter.requiresPermission()) {
			orientationPermissionGranted.set(true);
			orientationAdapter.startMonitoring((data) => {
				orientationData.set(data);
			});
		}
	});

	onDestroy(() => {
		cameraAdapter.stopCamera();
		orientationAdapter.stopMonitoring();
		if (photoObjectUrl) {
			URL.revokeObjectURL(photoObjectUrl);
		}
	});

	async function handlePermissionRequested() {
		// Request orientation permission first (iOS requires user gesture)
		if (orientationAdapter.requiresPermission() && !$orientationPermissionGranted) {
			const granted = await orientationAdapter.requestOrientationPermission();
			orientationPermissionGranted.set(granted);

			if (granted) {
				orientationAdapter.startMonitoring((data) => {
					orientationData.set(data);
				});
			}
		}

		const stream = await cameraAdapter.requestCameraAccess(currentFacingMode);

		if (stream) {
			cameraStream.set(stream);
			cameraPermissionGranted.set(true);
			appState.set('camera-ready');
		} else {
			showFileUpload = true;
			cameraPermissionGranted.set(false);
		}
	}

	async function handleCapture() {
		const photo = await cameraAdapter.capturePhoto();

		if (photo) {
			capturedPhoto.set(photo);
			photoObjectUrl = URL.createObjectURL(photo);
			appState.set('photo-captured');
		}
	}

	async function handleSwitchCamera() {
		currentFacingMode = currentFacingMode === 'user' ? 'environment' : 'user';
		cameraAdapter.stopCamera();

		const stream = await cameraAdapter.requestCameraAccess(currentFacingMode);
		if (stream) {
			cameraStream.set(stream);
		}
	}

	function handleRetake() {
		if (photoObjectUrl) {
			URL.revokeObjectURL(photoObjectUrl);
			photoObjectUrl = null;
		}
		capturedPhoto.set(null);
		appState.set('camera-ready');
	}

	async function handleConfirm() {
		if (!$capturedPhoto) return;

		appState.set('uploading');
		uploadError.set(null);

		try {
			const metadata = {
				userAgent: navigator.userAgent,
				timestamp: new Date().toISOString(),
				captureMethod: 'camera' as const
			};

			const verdict = await apiAdapter.uploadPhoto($capturedPhoto, metadata);
			currentVerdict.set(verdict);
			appState.set('showing-verdict');
		} catch (error) {
			uploadError.set(error instanceof Error ? error.message : 'Onbekende fout opgetreden');
			appState.set('error');
		}
	}

	async function handleFileSelected(event: CustomEvent<{ file: File }>) {
		const file = event.detail.file;

		appState.set('uploading');
		uploadError.set(null);

		try {
			const metadata = {
				userAgent: navigator.userAgent,
				timestamp: new Date().toISOString(),
				captureMethod: 'file' as const
			};

			const verdict = await apiAdapter.uploadPhoto(file, metadata);
			currentVerdict.set(verdict);
			appState.set('showing-verdict');
		} catch (error) {
			uploadError.set(error instanceof Error ? error.message : 'Onbekende fout opgetreden');
			appState.set('error');
		}
	}

	function handleReset() {
		if (photoObjectUrl) {
			URL.revokeObjectURL(photoObjectUrl);
			photoObjectUrl = null;
		}
		resetAppState();
	}

	function handleToggleLevelCheck(event: CustomEvent<{ enabled: boolean }>) {
		levelCheckEnabled.set(event.detail.enabled);
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
			<CameraPermission
				httpsRequired={!cameraAdapter.isCameraSupported()}
				onpermissionrequested={handlePermissionRequested}
			/>

			{#if !cameraAdapter.isCameraSupported()}
				<div class="or-divider">
					<span>of</span>
				</div>
				<FileUploadFallback onfileselected={handleFileSelected} />
			{/if}
		{:else if $appState === 'camera-ready'}
			{#if showFileUpload || !$cameraPermissionGranted}
				<FileUploadFallback onfileselected={handleFileSelected} />
			{:else if $cameraStream}
				<div class="camera-section">
					<AccessibilityToggle
						bind:levelCheckEnabled={$levelCheckEnabled}
						ontoggle={handleToggleLevelCheck}
					/>

					<SpiritLevel orientationData={$orientationData} enabled={$levelCheckEnabled} />

					<CameraPreview
						stream={$cameraStream}
						isLevelCheckEnabled={$levelCheckEnabled}
						isDeviceLevel={$isDeviceLevel}
						oncapture={handleCapture}
						onswitchcamera={handleSwitchCamera}
					/>
				</div>
			{/if}
		{:else if $appState === 'photo-captured' && photoObjectUrl}
			<PhotoConfirmation
				photoUrl={photoObjectUrl}
				onconfirm={handleConfirm}
				onretake={handleRetake}
			/>
		{:else if $appState === 'uploading'}
			<UploadProgress progress={50} message="De rechter beraadslaagt..." />
		{:else if $appState === 'showing-verdict' && $currentVerdict}
			<VerdictDisplay verdict={$currentVerdict} onreset={handleReset} />
		{:else if $appState === 'error'}
			<ErrorDisplay
				message={$uploadError || 'Er is een fout opgetreden'}
				onretry={handleConfirm}
				onreset={handleReset}
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

	.or-divider {
		text-align: center;
		margin: 2rem 0;
		position: relative;
	}

	.or-divider::before {
		content: '';
		position: absolute;
		top: 50%;
		left: 0;
		right: 0;
		height: 1px;
		background: rgba(255, 255, 255, 0.3);
	}

	.or-divider span {
		background: rgba(0, 0, 0, 0.3);
		color: white;
		padding: 0.5rem 1rem;
		border-radius: 20px;
		position: relative;
		font-weight: 500;
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
