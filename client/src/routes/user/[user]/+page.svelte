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

<h1>{data.profile.name}</h1>

<a href={`/user/${data.profile.id}/favorites/presets`}>Favorites</a>

<h2>Plugin presets</h2>
<div>
	{#each data.presets as preset}
		<PresetRow {preset} stats={data.presetStats[preset.id]} authorName="" {showSpoilers} />
	{/each}
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
