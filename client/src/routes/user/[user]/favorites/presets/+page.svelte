<script lang="ts">
	import type { PageData } from './$types';
	import PresetRow from '$lib/components/PresetRow.svelte';
	import { onMount } from 'svelte';
	import { getSettings } from '$lib/settings';
	import { connectBackend } from '$lib/backend';

	export let data: PageData;

	let showSpoilers = false;

	onMount(async () => {
		const backend = connectBackend();
		const user = backend.getCurrentUser();
		const settings = await getSettings(user ? user.id : undefined);
		showSpoilers = settings.show_spoilers;
	});
</script>

<h1>Favorite presets</h1>

<span>Back to <a href={`/user/${data.profile.id}`}>{data.profile.name}</a></span>
<div class="gap" />

<div>
	{#each data.presets as preset}
		<PresetRow
			{preset}
			authorName={data.presetAuthors[preset.author]}
			stats={data.presetStats[preset.id]}
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

	.gap {
		height: 15px;
	}

	a[disabled] {
		pointer-events: none;
		text-decoration: none;
		color: inherit;
	}
</style>
