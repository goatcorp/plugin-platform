<script lang="ts">
	import type { Preset } from './preset';

	export let preset: Preset;
	export let authorName: string;
</script>

<div class="result-entry">
	<a href={`/preset/${preset.id}`}><div class="thumbnail" /></a>
	<div class="info">
		<div class="basic">
			<span class="title"><a href={`/preset/${preset.id}`}>{preset.title}</a></span>
			{#if authorName}
				<span>By <a href={`/user/${preset.author}`}>{authorName}</a></span>
			{/if}
		</div>
		<div><small>Last updated: {preset.updated.toLocaleDateString()}</small></div>
	</div>
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	small {
		color: #444;
	}

	.result-entry {
		display: flex;

		margin-bottom: 15px;

		border: 1px solid #ccc;
		border-radius: 15px 10px 10px 15px; // Hides the left corners behind the thumbnail

		.thumbnail {
			height: 100px;
			width: 150px;

			background-color: blue;

			border-radius: 10px 0 0 10px;
		}

		.info {
			display: flex;
			justify-content: space-between;
			width: 100%;

			padding: 10px;

			.basic > * {
				display: block;
			}

			.title {
				margin-bottom: 10px;
			}
		}
	}
</style>
