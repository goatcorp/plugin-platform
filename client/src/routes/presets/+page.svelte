<script lang="ts">
	import PresetSearchResult from '$lib/PresetRow.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<h1>All presets</h1>

<div>
	{#each data.presets as preset}
		<PresetSearchResult
			{preset}
			authorName={data.authors[preset.author]}
			stats={data.stats[preset.id]}
		/>
	{/each}
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
