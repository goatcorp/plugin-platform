<script lang="ts">
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { getSettings } from '$lib/settings';
	import { connectBackend } from '$lib/backend';
	import PresetCardGallery from '$lib/components/PresetCardGallery.svelte';
	import { SearchIcon } from 'svelte-feather-icons';

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
	<form method="get" action="/search" class="search">
		<label class="input-group">
			<input
				type="search"
				placeholder="Search for presets..."
				name="q"
				class="input input-bordered input-primary w-full max-w-xs"
			/>
			<button class="btn"><SearchIcon /></button>
		</label>
	</form>
</div>

<div>
	<h1>Plugin presets</h1>
	<a href="/presets" class="link link-primary">Go to all presets</a>
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
