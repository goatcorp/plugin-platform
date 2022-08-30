<script lang="ts">
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { getSettings } from '$lib/settings';
	import { connectBackend } from '$lib/backend';
	import PresetCardGallery from '$lib/components/PresetCardGallery.svelte';

	export let data: PageData;

	let showSpoilers = false;

	onMount(async () => {
		const backend = connectBackend();
		const user = backend.getCurrentUser();
		const settings = await getSettings(user ? user.id : undefined);
		showSpoilers = settings.show_spoilers;
	});
</script>

<div>
	<h1>Search</h1>
	<div>
		<form method="get" action="/search">
			<input type="search" placeholder="Search for presets..." name="q" />
			<button>Search</button>
		</form>
	</div>
</div>

<div>
	<h1>Plugin presets</h1>
	<a href="/presets">Go to all presets</a>
	<h2>Popular</h2>
	<!--Presets with a high absolute number of views-->
	<PresetCardGallery
		presets={data.popular}
		stats={data.stats}
		authors={data.authors}
		{showSpoilers}
	/>
	<h2>Trending</h2>
	<!--Presets with a high rate of growth over the past 24 hours-->
	<PresetCardGallery
		presets={data.trending}
		stats={data.stats}
		authors={data.authors}
		{showSpoilers}
	/>
	<h2>New</h2>
	<PresetCardGallery
		presets={data.presets}
		stats={data.stats}
		authors={data.authors}
		{showSpoilers}
	/>
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
