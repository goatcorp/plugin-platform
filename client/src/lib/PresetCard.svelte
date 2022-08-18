<script lang="ts">
	import type { Preset } from './preset';

	export let preset: Preset;
</script>

<div class="preset-card">
	<a href={`/presets/${preset.id}`}>
		<div class="thumbnail" />
		<p class="title">{preset.title}</p>
	</a>
	<p class="author">By <a href={`/users/${preset.author.id}`}>{preset.author.name}</a></p>
</div>

<style lang="scss">
	p,
	a {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	.preset-card {
		height: 200px;
		width: 300px;

		.thumbnail {
			height: 60%;
			width: 100%;
			background-color: blue;
		}
	}
</style>
