<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<h1>{data.preset.title}</h1>
<div class="info">
	<span>By <a href={`/user/${data.preset.author}`}>{data.authorName}</a></span>
	<span>Created {data.preset.created.toLocaleDateString()}</span>
	<span>Last updated {data.preset.updated.toLocaleDateString()}</span>
</div>

{#each data.presetData as presetDataEntry, i}
	<h2>{presetDataEntry.title || `File ${i + 1}`}</h2>
	<div class="preset-data">
		<code>{@html presetDataEntry.data.replace(/[\r\n]+/g, '<br />')}</code>
	</div>
{/each}

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	code {
		font-family: 'Courier New', Courier, monospace;
	}

	.info > * {
		display: block;
	}

	.preset-data {
		padding: 10px;
		background-color: #eee;
	}
</style>
