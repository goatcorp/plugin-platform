<script lang="ts">
	import type { Preset } from '../lib/preset';
	import PresetCard from '../lib/PresetCard.svelte';

	const popularPresets: Preset[] = new Array(6).fill(0).map((_, i) => ({
		id: `cool-preset-${i}`,
		thumbnail: null,
		title: `Cool plugin preset ${i}`,
		author: {
			id: 'karashiiro',
			name: 'karashiiro'
		}
	}));

	const trendingPresets: Preset[] = [
		{
			id: 'cool-preset-1',
			thumbnail: null,
			title: 'Cool plugin preset 1',
			author: {
				id: 'karashiiro',
				name: 'karashiiro'
			}
		}
	];

	const newPresets: Preset[] = [];
</script>

<main>
	<div>
		<h1>Dalamud Plugin Presets</h1>
		<!--Plugins themselves will come later-->
		<p>We have cool presets</p>
	</div>

	<div>
		<h1>Search</h1>
		<a href="/search">Search</a>
	</div>

	<div>
		<h1>Plugin presets</h1>
		<a href="/presets">Go to all presets</a>
		<h2>Popular</h2>
		<!--Presets with a high absolute number of downloads-->
		<div class="gallery">
			{#each popularPresets as preset}
				<PresetCard {preset} />
			{/each}
		</div>
		<h2>Trending</h2>
		<!--Presets with a high rate of growth over the past 24 hours-->
		<div class="gallery">
			{#each trendingPresets as preset}
				<PresetCard {preset} />
			{/each}
		</div>
		<h2>New</h2>
		<div class="gallery">
			{#each newPresets as preset}
				<PresetCard {preset} />
			{/each}
		</div>
	</div>
</main>

<style lang="scss">
	main {
		width: 70%;
		margin: auto;
	}

	h1,
	h2,
	p,
	a {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	.gallery {
		display: flex;
		min-height: 200px;
		flex-wrap: wrap;
		justify-content: space-around;
	}
</style>
