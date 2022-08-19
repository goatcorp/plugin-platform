<script lang="ts">
	import PresetSearchResult from '$lib/PresetSearchResult.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<h1>Results for "{data.query}":</h1>

<div class="results">
	{#each data.results as preset}
		<PresetSearchResult {preset} authorName={data.authors[preset.author]} />
	{/each}
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
