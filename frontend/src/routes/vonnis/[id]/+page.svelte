<script lang="ts">
	// Canonical route for shared verdicts (was /verdict/[id], now /vonnis/[id])
	// Legacy /verdict/[id] URLs are redirected via nginx rewrite for backward compatibility
	import VerdictDisplay from '$lib/features/VerdictDisplay.svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';

	export let data: PageData;

	// Generate dynamic title and description
	$: verdictTitle = data.verdict.admissible
		? `Vonnis: Score ${data.verdict.score}/10 - ${data.verdict.verdict.verdictType}`
		: 'Zaak Niet-Ontvankelijk';

	$: verdictDescription = data.verdict.admissible
		? data.verdict.verdict.observation
		: data.verdict.verdict.crime;

	// Get current URL for og:url
	$: pageUrl = typeof window !== 'undefined' ? window.location.href : '';

	function handleReset() {
		goto('/');
	}
</script>

<svelte:head>
	<title>{verdictTitle} - Rechtbank voor Meubilair</title>
	<meta name="description" content={verdictDescription} />

	<!-- Open Graph meta tags -->
	<meta property="og:title" content="{verdictTitle} - Rechtbank voor Meubilair" />
	<meta property="og:description" content={verdictDescription} />
	<meta property="og:image" content={data.imageData} />
	<meta property="og:url" content={pageUrl} />
	<meta property="og:type" content="website" />

	<!-- Twitter Card meta tags -->
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content="{verdictTitle} - Rechtbank voor Meubilair" />
	<meta name="twitter:description" content={verdictDescription} />
	<meta name="twitter:image" content={data.imageData} />
</svelte:head>

<div class="shared-verdict-page">
	<VerdictDisplay verdict={data.verdict} imageData={data.imageData} onreset={handleReset} />
</div>

<style>
	.shared-verdict-page {
		padding: 2rem 0;
	}

	@media (max-width: 768px) {
		.shared-verdict-page {
			padding: 1rem 0;
		}
	}
</style>
