<script lang="ts">
	import type { PageData } from './$types';
	import PresetRow from '$lib/PresetRow.svelte';

	export let data: PageData;
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
		/>
	{/each}
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	.gap {
		height: 15px;
	}
</style>
