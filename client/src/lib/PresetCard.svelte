<script lang="ts">
	import type { Preset, PresetStats } from './preset';

	export let preset: Preset;
	export let stats: PresetStats;
	export let authorName: string;

	const url = `/preset/${preset.id}`;
</script>

<div class="preset-card">
	<div class="preview">
		<a href={url}>
			<div class="thumbnail" />
			<div class="views-wrapper">
				<span class="views">{stats.views} views</span>
			</div>
		</a>
	</div>
	<div class="info">
		<a href={url}><p class="title">{preset.title}</p></a>
		<p class="author">By <a href={`/user/${preset.author}`}>{authorName}</a></p>
	</div>
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	.preset-card {
		height: 200px;
		width: 300px;

		border: 1px solid #ccc;
		border-radius: 10px 10px 15px 15px; // Hides the top corners behind the thumbnail

		.info {
			padding-left: 10px;
			padding-right: 10px;
		}

		.preview {
			position: relative;
			height: 60%;
			width: 100%;

			background-color: blue;
			border-radius: 10px 10px 0 0;

			.thumbnail {
				height: 100%;
			}
		}

		.views-wrapper {
			position: absolute;
			bottom: 10px;
			right: 10px;
			padding: 4px;

			background-color: rgba(0.7, 0.7, 0.7, 0.6);
			border-radius: 15px;
		}

		.views {
			font-size: smaller;
			color: white;
		}
	}
</style>
