<script lang="ts">
	import PocketBase from 'pocketbase';
	import type { PageData } from './$types';
	import PresetRow from '$lib/PresetRow.svelte';
	import { onMount } from 'svelte';
	import { getSettings } from '$lib/settings';

	export let data: PageData;

	let showSpoilers = false;

	const connect = () => {
		return new PocketBase('http://127.0.0.1:8090');
	};

	onMount(async () => {
		const client = connect();
		const user = client.authStore.model;
		const settings = await getSettings('id' in user ? user.id : undefined);
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
