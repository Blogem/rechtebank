<script lang="ts">
	import { onMount } from 'svelte';
	import {
		appState,
		currentVerdict,
		uploadError,
		resetAppState,
		setPermissionsGranted
	} from '$lib/shared/stores/appStore';

	import { ApiAdapter } from '$lib/adapters/api/ApiAdapter';

	import PhotoCapture from '$lib/features/PhotoCapture.svelte';
	import CameraPermission from '$lib/features/CameraPermission.svelte';
	import UploadProgress from '$lib/features/UploadProgress.svelte';
	import VerdictDisplay from '$lib/features/VerdictDisplay.svelte';
	import ErrorDisplay from '$lib/features/ErrorDisplay.svelte';

	const apiAdapter = new ApiAdapter();
	let currentPhotoData: string | undefined = undefined;
	let isPreviewingPhoto = false;

	onMount(() => {
		appState.set('camera-ready');
		setPermissionsGranted(true);
	});

	function handlePermissionRequested() {
		appState.set('camera-ready');
		setPermissionsGranted(true);
	}

	async function handlePhotoConfirmed(blob: Blob, rotation: number) {
		// PhotoCapture component handles rotation internally and returns rotated blob
		// Pass rotation=0 since blob is already rotated (rotation value kept for metadata/logging)
		appState.set('uploading');
		uploadError.set(null);

		try {
			// Convert blob to data URL for display
			const reader = new FileReader();
			reader.readAsDataURL(blob);
			reader.onloadend = () => {
				currentPhotoData = reader.result as string;
			};

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
		currentPhotoData = undefined;
		isPreviewingPhoto = false;
		resetAppState();
	}

	function handlePreviewStateChange(isPreviewing: boolean) {
		isPreviewingPhoto = isPreviewing;
	}
</script>

<svelte:head>
	<title>Rechtbank voor Meubilair - Zaak indienen</title>
	<meta name="description" content="Dien uw meubelstuk in voor officieel rechterlijk oordeel" />
</svelte:head>

{#if $appState === 'requesting-permissions'}
	<CameraPermission httpsRequired={false} onpermissionrequested={handlePermissionRequested} />
{:else if $appState === 'camera-ready'}
	{#if !isPreviewingPhoto}
		<section class="introduction-card">
			<h2 class="card-label">Officiële mededeling</h2>
			<p class="introduction-text">
				Welkom bij de Rechtbank voor Meubilair, waar elk meubelstuk rekenschap moet afleggen van
				zijn verticale integriteit. Dien uw bewijsmateriaal in en ontdek of uw stoel, tafel of kast
				recht genoeg staat volgens de Eerwaarde Rechter.
			</p>
		</section>
	{/if}

	<section class="case-submission-card">
		{#if !isPreviewingPhoto}
			<h2 class="section-heading">Zaak indienen</h2>
			<p class="section-description">
				Maak een foto van het meubelstuk dat u ter beoordeling wilt voorleggen. Zorg ervoor dat het
				meubelstuk goed zichtbaar is in het beeld.
			</p>
		{/if}
		<PhotoCapture
			onPhotoConfirmed={handlePhotoConfirmed}
			onPreviewStateChange={handlePreviewStateChange}
		/>
	</section>
{:else if $appState === 'uploading'}
	<UploadProgress progress={50} message="De rechter beraadslaagt..." />
{:else if $appState === 'showing-verdict' && $currentVerdict}
	<VerdictDisplay verdict={$currentVerdict} imageData={currentPhotoData} onreset={handleReset} />
{:else if $appState === 'error'}
	<ErrorDisplay
		message={$uploadError || 'Er is een fout opgetreden'}
		onreset={handleReset}
		onretry={handleRetry}
	/>
{/if}

<style>
	.introduction-card {
		background: var(--color-court-surface);
		border-radius: 4px;
		padding: 2rem;
		margin-bottom: 2rem;
		box-shadow: var(--shadow-base);
		border-top: 3px solid var(--color-court-accent);
		transition: box-shadow var(--timing-interactive) var(--ease-out);
	}

	.card-label {
		font-size: 0.9rem;
		font-weight: 600;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		color: var(--color-court-accent);
		margin: 0 0 1rem 0;
		font-family: var(--font-sans);
	}

	.introduction-text {
		font-size: 1.1rem;
		line-height: 1.7;
		color: var(--color-court-text);
		margin: 0;
		text-align: center;
		font-family: var(--font-serif);
	}

	.case-submission-card {
		background: var(--color-court-surface);
		border-radius: 4px;
		padding: 1.5rem;
		box-shadow: var(--shadow-base);
		border-top: 3px solid var(--color-court-primary);
		transition: box-shadow var(--timing-interactive) var(--ease-out);
	}

	.section-heading {
		font-size: 1.75rem;
		color: var(--color-court-primary);
		margin: 0 0 0.75rem 0;
		text-align: center;
	}

	.section-description {
		font-size: 1rem;
		line-height: 1.6;
		color: var(--color-court-text-light);
		margin: 0 0 1.5rem 0;
		text-align: center;
		max-width: 600px;
		margin-left: auto;
		margin-right: auto;
	}

	@media (max-width: 768px) {
		.introduction-card,
		.case-submission-card {
			padding: 1.25rem;
		}

		.introduction-text {
			font-size: 1rem;
		}

		.section-heading {
			font-size: 1.5rem;
		}
	}

	@media (max-width: 480px) {
		.introduction-card,
		.case-submission-card {
			padding: 1rem;
			border-radius: 2px;
		}

		.introduction-text {
			font-size: 0.95rem;
		}

		.section-heading {
			font-size: 1.3rem;
		}

		.section-description {
			font-size: 0.9rem;
		}
	}
</style>
