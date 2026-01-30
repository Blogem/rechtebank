<script lang="ts">
	export let levelCheckEnabled = true;
	export let ontoggle: ((event: CustomEvent<{ enabled: boolean }>) => void) | undefined = undefined;

	function toggle() {
		levelCheckEnabled = !levelCheckEnabled;
		ontoggle?.(new CustomEvent('toggle', { detail: { enabled: levelCheckEnabled } }));
	}
</script>

<div class="accessibility-toggle">
	<label>
		<input type="checkbox" checked={levelCheckEnabled} onchange={toggle} />
		<span>Waterpas controle vereist</span>
	</label>

	{#if !levelCheckEnabled}
		<p class="help-text">
			<small
				>♿ Waterpas controle uitgeschakeld voor toegankelijkheid. U kunt nu op elk moment een foto
				maken.</small
			>
		</p>
	{/if}
</div>

<style>
	.accessibility-toggle {
		margin: 1rem 0;
		padding: 1rem;
		background: #f8f9fa;
		border-radius: 4px;
	}

	label {
		display: flex;
		align-items: center;
		cursor: pointer;
		font-weight: 500;
	}

	input[type='checkbox'] {
		margin-right: 0.5rem;
		width: 20px;
		height: 20px;
		cursor: pointer;
	}

	.help-text {
		margin: 0.5rem 0 0 0;
		color: #666;
	}
</style>
