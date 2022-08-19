<script lang="ts">
	import type { PageData } from './$types';
	import PresetRow from '$lib/PresetRow.svelte';

	export let data: PageData;
</script>

<h1>{data.profile.name}</h1>

<h2>Plugin presets</h2>
<div>
	{#each data.presets as preset}
		<PresetRow {preset} authorName="" />
	{/each}
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
