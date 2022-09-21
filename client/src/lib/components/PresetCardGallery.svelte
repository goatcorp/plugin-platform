<script lang="ts">
	import type { Preset, PresetStats } from '$lib/preset';

	import PresetCard from './PresetCard.svelte';

	export let presets: Preset[];
	export let stats: Record<string, PresetStats>;
	export let authors: Record<string, string>;
	export let showSpoilers: boolean;
</script>

<div class="flex flex-wrap justify-around min-h-50">
	{#each presets as preset}
		<div class="mb-6 px-2">
			<PresetCard
				{preset}
				stats={stats[preset.id]}
				authorName={authors[preset.author]}
				{showSpoilers}
			/>
		</div>
	{/each}
</div>
