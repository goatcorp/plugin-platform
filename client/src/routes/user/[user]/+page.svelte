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

<h1>{data.profile.name}</h1>

<a href={`/user/${data.profile.id}/favorites/presets`}>Favorites</a>

<h2>Plugin presets</h2>
<div>
	{#each data.presets as preset}
		<PresetRow {preset} stats={data.presetStats[preset.id]} authorName="" {showSpoilers} />
	{/each}
</div>
