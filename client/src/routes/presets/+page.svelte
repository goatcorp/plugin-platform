<script lang="ts">
	import PresetRow from '$lib/PresetRow.svelte';
	import { id } from '$lib/session';
	import { getSettings } from '$lib/settings';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	let showSpoilers = false;

	onMount(async () => {
		const settings = await getSettings($id);
		showSpoilers = settings.show_spoilers;
	});
</script>

<h1>All presets</h1>

<div>
	{#each data.presets as preset}
		<PresetRow
			{preset}
			authorName={data.authors[preset.author]}
			stats={data.stats[preset.id]}
			{showSpoilers}
		/>
	{/each}
</div>

<div>
	{#if data.page > 1}
		<a href={`?page=${data.page - 1}`}>Previous</a>
	{:else}
		<a href={`?page=${data.page - 1}`} disabled tabindex="-1">Previous</a>
	{/if}
	<a href={`?page=${data.page}`} disabled tabindex="-1">{data.page}</a>
	{#if data.page < data.totalPages}
		<a href={`?page=${data.page + 1}`}>Next</a>
	{:else}
		<a href={`?page=${data.page + 1}`} disabled tabindex="-1">Next</a>
	{/if}
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	a[disabled] {
		pointer-events: none;
		text-decoration: none;
		color: inherit;
	}
</style>
