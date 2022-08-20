<script lang="ts">
	import type { PageData } from './$types';
	import PresetRow from '$lib/PresetRow.svelte';
	import { onMount } from 'svelte';
	import { id } from '$lib/session';
	import { getSettings } from '$lib/settings';

	export let data: PageData;

	let showSpoilers = false;

	onMount(async () => {
		const settings = await getSettings($id);
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
