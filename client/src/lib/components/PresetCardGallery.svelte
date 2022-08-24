<script lang="ts">
	import type { Preset, PresetStats } from '$lib/preset';

	import PresetCard from './PresetCard.svelte';

	export let presets: Preset[];
	export let stats: Record<string, PresetStats>;
	export let authors: Record<string, string>;
	export let showSpoilers: boolean;
</script>

<div class="gallery">
	{#each presets as preset}
		<div class="wrapper">
			<PresetCard
				{preset}
				stats={stats[preset.id]}
				authorName={authors[preset.author]}
				{showSpoilers}
			/>
		</div>
	{/each}
</div>

<style lang="scss">
	.gallery {
		display: flex;
		min-height: 50px;
		flex-wrap: wrap;
		justify-content: space-around;

		.wrapper {
			margin-bottom: 10px;
		}
	}
</style>
