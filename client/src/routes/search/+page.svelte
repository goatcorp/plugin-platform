<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
	console.log(data.results);
</script>

<div>
	{#each data.results as preset}
		<div>
			<a href={`/preset/${preset.id}`}>{preset.title}</a>
		</div>
	{/each}
</div>

<style lang="scss">
	a {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
