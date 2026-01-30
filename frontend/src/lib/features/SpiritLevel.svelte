<script lang="ts">
	import type { OrientationData } from '$lib/shared/types';

	export let orientationData: OrientationData | null = null;
	export let enabled = true;

	$: beta = orientationData?.beta ?? 0;
	$: isLevel = orientationData?.isLevel ?? false;
	$: tiltDegrees = Math.round(beta);

	// Calculate bubble position based on tilt (-5 to +5 degree range mapped to -100px to +100px)
	$: bubbleOffset = Math.max(-100, Math.min(100, beta * 20));

	function getTiltDirection(): string {
		if (Math.abs(beta) <= 5) return '';
		return beta > 0 ? 'naar voren' : 'naar achteren';
	}
</script>

{#if enabled}
	<div class="spirit-level" class:level={isLevel} class:off-level={!isLevel}>
		<div class="level-container">
			<div class="level-tube">
				<div class="level-marks">
					<span class="mark center">|</span>
					<span class="mark left">|</span>
					<span class="mark right">|</span>
				</div>
				<div class="bubble" style="transform: translateX({bubbleOffset}px)">
					<div class="bubble-inner"></div>
				</div>
			</div>

			<div class="level-indicator">
				{#if isLevel}
					<span class="status-icon">✅</span>
					<span class="status-text">Waterpas</span>
				{:else}
					<span class="status-icon">⚠️</span>
					<span class="status-text">Kantel {tiltDegrees}° {getTiltDirection()}</span>
				{/if}
			</div>
		</div>
	</div>
{/if}

<style>
	.spirit-level {
		padding: 1rem;
		margin: 1rem 0;
		border-radius: 8px;
		transition: all 0.3s;
	}

	.spirit-level.level {
		background: rgba(40, 167, 69, 0.1);
		border: 2px solid #28a745;
	}

	.spirit-level.off-level {
		background: rgba(220, 53, 69, 0.1);
		border: 2px solid #dc3545;
	}

	.level-container {
		max-width: 400px;
		margin: 0 auto;
	}

	.level-tube {
		position: relative;
		height: 60px;
		background: linear-gradient(to bottom, #555, #333, #555);
		border-radius: 30px;
		box-shadow: inset 0 2px 10px rgba(0, 0, 0, 0.5);
		overflow: visible;
		margin: 0 auto;
		width: 100%;
		max-width: 300px;
	}

	.level-marks {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		display: flex;
		justify-content: center;
		align-items: center;
		color: rgba(255, 255, 255, 0.5);
		font-size: 2rem;
	}

	.mark {
		position: absolute;
	}

	.mark.center {
		left: 50%;
		transform: translateX(-50%);
	}

	.mark.left {
		left: 20%;
	}

	.mark.right {
		right: 20%;
	}

	.bubble {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		transition: transform 0.1s ease-out;
		will-change: transform;
	}

	.bubble-inner {
		width: 40px;
		height: 40px;
		background: radial-gradient(circle at 30% 30%, #ffeb3b, #fbc02d);
		border-radius: 50%;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
	}

	.level-indicator {
		margin-top: 1rem;
		text-align: center;
		font-weight: 500;
	}

	.status-icon {
		font-size: 1.5rem;
		margin-right: 0.5rem;
	}

	.status-text {
		font-size: 1.1rem;
	}

	.level .status-text {
		color: #28a745;
	}

	.off-level .status-text {
		color: #dc3545;
	}
</style>
