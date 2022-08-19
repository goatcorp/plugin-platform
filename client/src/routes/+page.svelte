<script lang="ts">
	import type { PageData } from './$types';
	import PresetCard from '$lib/PresetCard.svelte';

	export let data: PageData;
</script>

<div>
	<h1>Dalamud Plugin Presets</h1>
	<!--Plugins themselves will come later-->
	<p>We have cool presets</p>
</div>

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
	<div class="gallery">
		{#each data.popular as preset}
			<div class="wrapper"><PresetCard {preset} stats={{ views: preset.views }} /></div>
		{/each}
	</div>
	<h2>Trending</h2>
	<!--Presets with a high rate of growth over the past 24 hours-->
	<div class="gallery" />
	<h2>New</h2>
	<div class="gallery">
		{#each data.presets as preset}
			<div class="wrapper"><PresetCard {preset} stats={data.presetStats[preset.id]} /></div>
		{/each}
	</div>
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	.gallery {
		display: flex;
		min-height: 50px;
		flex-wrap: wrap;
		justify-content: space-around;

		.wrapper {
			margin-bottom: 10px;
		}
	}
</style>
